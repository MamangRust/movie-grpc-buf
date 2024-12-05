package service

import (
	"context"
	"errors"
	moviesv1 "movie_buf/api/movies/v1"
	"movie_buf/internal/domain"

	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type movies struct {
	moviesv1.UnimplementedMoviesWriterServiceServer
	moviesv1.UnimplementedMoviesReaderServiceServer
	db     *gorm.DB
	logger *zap.Logger
	meter  metric.Meter
}

func (svc *movies) GetMovie(ctx context.Context, request *moviesv1.GetMovieRequest) (*moviesv1.Movie, error) {
	svc.logger.Debug("Getting movie by ID", zap.Int64("movie.id", request.GetId()))

	span := trace.SpanFromContext(ctx)

	defer span.End()

	var movie domain.Movie

	span.AddEvent("Getting movie from the movie")

	err := svc.db.Model(&domain.Movie{}).WithContext(ctx).First(&movie, request.GetId()).Error

	if err != nil {
		svc.logger.Error("Failed to get movie", zap.Error(err))

		span.RecordError(err)

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Error(codes.NotFound, "movie not found")
		}

		return nil, status.Errorf(codes.Unavailable, "faield to movie: %v", err)
	}

	svc.logger.Debug("Returning movie", zap.Int64("movie.id", request.GetId()), zap.String("movie.title", movie.Title))

	return movie.API(), nil
}

func (svc *movies) ListMovies(ctx context.Context, request *moviesv1.ListMoviesRequest) (*moviesv1.ListMoviesResponse, error) {
	svc.logger.Debug("Getting movie list", zap.Int32("page_size", request.GetPageSize()), zap.String("page_token", request.GetPageToken()))

	span := trace.SpanFromContext(ctx)

	defer span.End()

	span.AddEvent("Getting movies from the database")

	var out []domain.Movie

	err := svc.db.Model(&domain.Movie{}).WithContext(ctx).Find(&out).Error

	if err != nil {
		svc.logger.Error("Failed to lists movies", zap.Error(err))
		span.RecordError(err)

		return nil, status.Errorf(codes.Unavailable, "failed to get movie: %v", err)
	}

	svc.logger.Debug("Returning task list", zap.Int32("page_size", request.GetPageSize()), zap.String("page_token", request.GetPageToken()), zap.Int("count", len(out)))

	res := moviesv1.ListMoviesResponse{
		Movies:        make([]*moviesv1.Movie, len(out)),
		NextPageToken: "",
	}

	for i, movie := range out {
		res.Movies[i] = movie.API()
	}

	return &res, nil
}

func (svc *movies) CreateMovie(ctx context.Context, request *moviesv1.CreateMovieRequest) (*moviesv1.Movie, error) {
	svc.logger.Debug("Creating movie", zap.String("movie.title", request.GetMovie().GetTitle()))

	span := trace.SpanFromContext(ctx)

	defer span.End()

	var movie domain.Movie

	svc.logger.Debug("Filling out movie information")

	span.AddEvent("Parsing movie from API request")

	movie.FromAPI(request.GetMovie())

	span.AddEvent("Persisting movie in the database")

	svc.logger.Debug("Persisting movie in the database", zap.String("movie.title", request.GetMovie().GetTitle()))

	err := svc.db.Model(&domain.Movie{}).WithContext(ctx).Create(&movie).Error

	if err != nil {
		svc.logger.Error("Failed to create movie", zap.Error(err))

		span.RecordError(err)

		return nil, status.Error(codes.Unavailable, "failed to create movie")
	}

	svc.logger.Debug("Returning created movie", zap.String("movie.title", request.GetMovie().GetTitle()))

	return movie.API(), nil

}

func (svc *movies) DeleteMovie(ctx context.Context, request *moviesv1.DeleteMovieRequest) (*moviesv1.Movie, error) {
	svc.logger.Debug("Deleting movie", zap.Int64("movie.id", request.GetId()))

	span := trace.SpanFromContext(ctx)

	defer span.End()

	movie, err := svc.GetMovie(ctx, &moviesv1.GetMovieRequest{Id: request.GetId()})

	if err != nil {
		return nil, err
	}

	span.AddEvent("Deleting movie from the database")

	err = svc.db.Model(&domain.Movie{}).Delete(&domain.Movie{}, movie.GetId()).Error

	if err != nil {
		svc.logger.Error("Failed to delete movie", zap.Error(err))
		span.RecordError(err)

		return nil, status.Error(codes.Unavailable, "failed to delete movie")
	}

	svc.logger.Debug("Movie was deleted", zap.Int64("movie.id", request.GetId()))

	return movie, nil

}

func (svc *movies) UndeleteMovie(ctx context.Context, request *moviesv1.UndeleteMovieRequest) (*moviesv1.Movie, error) {
	svc.logger.Debug("Undeleting movie", zap.Int64("movie.id", request.GetId()))

	span := trace.SpanFromContext(ctx)

	defer span.End()

	err := svc.db.Model(&domain.Movie{}).Unscoped().Where("id = ?", request.GetId()).Update("deleted_at", nil).Error

	if err != nil {
		svc.logger.Error("Failed to undelete movie", zap.Error(err))
		span.RecordError(err)

		return nil, status.Error(codes.Unavailable, "failed to undelete movie")
	}

	movie, err := svc.GetMovie(ctx, &moviesv1.GetMovieRequest{Id: request.GetId()})

	if err != nil {
		return nil, err
	}

	svc.logger.Debug("Movie was undeleted", zap.Int64("movie.id", request.GetId()))
	return movie, nil
}

func (svc *movies) UpdateMovie(ctx context.Context, request *moviesv1.UpdateMovieRequest) (*moviesv1.Movie, error) {
	svc.logger.Debug("Updating movie", zap.Int64("movie.id", request.GetMovie().GetId()))

	span := trace.SpanFromContext(ctx)

	defer span.End()

	_, err := svc.GetMovie(ctx, &moviesv1.GetMovieRequest{Id: request.GetMovie().GetId()})

	if err != nil {
		return nil, err
	}

	var movie domain.Movie

	movie.FromAPI(request.GetMovie())

	m, err := movie.ApplyMask(request.GetUpdateMask())

	if err != nil {
		svc.logger.Error("Failed to apply update mask", zap.Error(err))

		span.RecordError(err)

		return nil, status.Error(codes.Internal, "failed to  update movie")
	}

	err = svc.db.Model(&domain.Movie{}).Where("id = ?", request.GetMovie().GetId()).Updates(m).Error
	if err != nil {
		svc.logger.Error("Failed to update task", zap.Error(err))
		span.RecordError(err)
		return nil, status.Error(codes.Internal, "failed to update task")
	}

	svc.logger.Debug("Movie was updated", zap.Int64("movie.id", request.GetMovie().GetId()))

	return svc.GetMovie(ctx, &moviesv1.GetMovieRequest{Id: request.GetMovie().GetId()})

}

func NewMoviesWriter(db *gorm.DB, logger *zap.Logger, meter metric.Meter) moviesv1.MoviesWriterServiceServer {
	moviesLogger := logger.Named("service.movies.writer")

	return &movies{
		db:     db,
		logger: moviesLogger,
		meter:  meter,
	}
}

func NewMoviesReader(db *gorm.DB, logger *zap.Logger, meter metric.Meter) moviesv1.MoviesReaderServiceServer {
	moviesLogger := logger.Named("service.movies.reader")

	return &movies{
		db:     db,
		logger: moviesLogger,
		meter:  meter,
	}
}
