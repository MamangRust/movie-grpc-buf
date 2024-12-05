package service

import (
	"context"
	moviesv1 "movie_buf/api/movies/v1"
	"movie_buf/internal/domain"
	"testing"

	"github.com/stretchr/testify/suite"
	"go.opentelemetry.io/otel/metric/noop"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestMoviesServiceSuite(t *testing.T) {
	suite.Run(t, new(MoviesServiceTestSuite))
}

type MoviesServiceTestSuite struct {
	suite.Suite
	db     *gorm.DB
	writer moviesv1.MoviesWriterServiceServer
	reader moviesv1.MoviesReaderServiceServer
}

func (suite *MoviesServiceTestSuite) SetupSuite() {

}

func (suite *MoviesServiceTestSuite) SetupTest() {
	var err error

	suite.db, err = gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	suite.Require().NoError(err)
	suite.Require().NoError(suite.db.Migrator().AutoMigrate(&domain.Movie{}))

	suite.writer = NewMoviesWriter(suite.db, zap.NewNop(), noop.NewMeterProvider().Meter(""))
	suite.reader = NewMoviesReader(suite.db, zap.NewNop(), noop.NewMeterProvider().Meter(""))
}

func (suite *MoviesServiceTestSuite) TearDownTest() {
	db, err := suite.db.DB()

	suite.Require().NoError(err)

	suite.Assert().NoError(db.Close())
}

func (suite *MoviesServiceTestSuite) TearDownSuite() {

}

func (suite *MoviesServiceTestSuite) TestCreate_Success() {
	var before int64

	suite.Require().NoError(suite.db.Model(&domain.Movie{}).Count(&before).Error)

	const title = "movie"

	res, err := suite.writer.CreateMovie(context.Background(), &moviesv1.CreateMovieRequest{
		Movie: &moviesv1.Movie{
			Title: title,
		},
	})

	suite.Assert().NoError(err)
	suite.Assert().NotNil(res)
	suite.Assert().Equal(title, res.GetTitle())

	var after int64

	suite.Require().NoError(suite.db.Model(&domain.Movie{}).Count(&after).Error)
	suite.NotEqual(before, after)
	suite.Equal(before+1, after)

}

func (suite *MoviesServiceTestSuite) TestGet_Success() {
	ctx := context.Background()

	expected, err := suite.writer.CreateMovie(ctx, &moviesv1.CreateMovieRequest{
		Movie: &moviesv1.Movie{
			Title: "A Test",
		},
	})

	suite.Require().NoError(err)

	response, err := suite.reader.GetMovie(ctx, &moviesv1.GetMovieRequest{Id: expected.GetId()})

	suite.Assert().NoError(err)

	suite.Assert().Equal(expected.GetTitle(), response.GetTitle())
}

func (suite *MoviesServiceTestSuite) TestGet_NotFound() {
	ctx := context.Background()

	_, err := suite.reader.GetMovie(ctx, &moviesv1.GetMovieRequest{Id: 199452})
	suite.Assert().Error(err)

	suite.Assert().ErrorIs(err, status.Error(codes.NotFound, "movie not found"))
}

func (suite *MoviesServiceTestSuite) TestList_Empty() {
	ctx := context.Background()

	response, err := suite.reader.ListMovies(ctx, &moviesv1.ListMoviesRequest{
		PageSize:  0,
		PageToken: "",
	})

	suite.Assert().NoError(err)
	suite.Assert().Empty(response.GetMovies())

}

func (suite *MoviesServiceTestSuite) TestList_Success() {
	ctx := context.Background()

	for i := 0; i < 10; i++ {
		_, err := suite.writer.CreateMovie(ctx, &moviesv1.CreateMovieRequest{
			Movie: &moviesv1.Movie{
				Title: "A Test",
			},
		})

		suite.Require().NoError(err)
	}

	var expected int64

	suite.Require().NoError(suite.db.Model(&domain.Movie{}).Count(&expected).Error)

	response, err := suite.reader.ListMovies(ctx, &moviesv1.ListMoviesRequest{
		PageSize:  0,
		PageToken: "",
	})

	suite.Assert().NoError(err)
	suite.Assert().NotEmpty(response.GetMovies())
	suite.Assert().Len(response.GetMovies(), int(expected))
}

func (suite *MoviesServiceTestSuite) TestDelete_NotFound() {
	ctx := context.Background()

	_, err := suite.writer.DeleteMovie(ctx, &moviesv1.DeleteMovieRequest{Id: 116644725})
	suite.Assert().Error(err)
	suite.Assert().ErrorIs(err, status.Error(codes.NotFound, "movie not found"))
}

func (suite *MoviesServiceTestSuite) TestDelete_Success() {
	ctx := context.Background()

	expected, err := suite.writer.CreateMovie(ctx, &moviesv1.CreateMovieRequest{
		Movie: &moviesv1.Movie{
			Title: "A movie",
		},
	})

	suite.Require().NoError(err)

	response, err := suite.writer.DeleteMovie(ctx, &moviesv1.DeleteMovieRequest{Id: expected.GetId()})
	suite.Assert().NoError(err)
	suite.Assert().Equal(expected.GetTitle(), response.GetTitle())
}

func (suite *MoviesServiceTestSuite) TestUndelete_Success() {
	ctx := context.Background()

	expected, err := suite.writer.CreateMovie(ctx, &moviesv1.CreateMovieRequest{
		Movie: &moviesv1.Movie{
			Title: "A test",
		},
	})
	suite.Require().NoError(err)

	res, err := suite.reader.ListMovies(ctx, &moviesv1.ListMoviesRequest{})
	suite.Require().NoError(err)
	before := len(res.GetMovies())

	response, err := suite.writer.DeleteMovie(ctx, &moviesv1.DeleteMovieRequest{Id: expected.GetId()})
	suite.Require().NoError(err)

	res, err = suite.reader.ListMovies(ctx, &moviesv1.ListMoviesRequest{})
	suite.Require().NoError(err)
	after := len(res.GetMovies())
	suite.Require().NotEqual(before, after)

	task, err := suite.writer.UndeleteMovie(ctx, &moviesv1.UndeleteMovieRequest{Id: response.GetId()})
	suite.Assert().NoError(err)
	suite.Assert().NotNil(task)

	res, err = suite.reader.ListMovies(ctx, &moviesv1.ListMoviesRequest{})
	suite.Require().NoError(err)
	suite.Require().Equal(before, len(res.GetMovies()))
}

func (suite *MoviesServiceTestSuite) TestUpdate_Success() {
	ctx := context.Background()

	before, err := suite.writer.CreateMovie(ctx, &moviesv1.CreateMovieRequest{
		Movie: &moviesv1.Movie{
			Title: "A test",
		},
	})
	suite.Require().NoError(err)

	after, err := suite.writer.UpdateMovie(ctx, &moviesv1.UpdateMovieRequest{
		Movie: &moviesv1.Movie{
			Id:    before.GetId(),
			Title: "An updated title",
		},
		UpdateMask: &fieldmaskpb.FieldMask{Paths: []string{"title"}},
	})
	suite.Assert().NoError(err)

	suite.Assert().NotEqual(before.GetTitle(), after.GetTitle())
	suite.Assert().Equal(before.GetId(), after.GetId())
	suite.Assert().Equal(before.GetDescription(), after.GetDescription())

	final, err := suite.reader.GetMovie(ctx, &moviesv1.GetMovieRequest{Id: before.GetId()})
	suite.Require().NoError(err)

	suite.Assert().Equal(after, final)
}
