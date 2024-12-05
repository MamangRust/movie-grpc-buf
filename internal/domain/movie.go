package domain

import (
	"encoding/json"
	"errors"
	moviesv1 "movie_buf/api/movies/v1"
	"movie_buf/internal/serializer"
	"time"

	"github.com/gojaguar/jaguar/strings"
	fieldmask_utils "github.com/mennanov/fieldmask-utils"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gopkg.in/yaml.v3"
	"gorm.io/gorm"
)

var _ serializer.JSON = (*Movie)(nil)
var _ serializer.YAML = (*Movie)(nil)
var _ serializer.API[*moviesv1.Movie] = (*Movie)(nil)

type Movie struct {
	gorm.Model
	Title       string    `gorm:"column:title" json:"title"`
	Genre       string    `gorm:"column:genre" json:"genre"`
	Description string    `gorm:"column:description" json:"description"`
	ReleaseDate time.Time `gorm:"column:release_date" json:"release_date"`
}

func (m *Movie) API() *moviesv1.Movie {
	return &moviesv1.Movie{
		Id:          int64(m.ID),
		Title:       m.Title,
		Genre:       m.Genre,
		Description: m.Description,
		ReleaseDate: timestamppb.New(m.ReleaseDate),
		CreateTime:  timestamppb.New(m.CreatedAt),
		UpdateTime:  timestamppb.New(m.UpdatedAt),
	}
}

func (m *Movie) FromAPI(in *moviesv1.Movie) {
	m.ID = uint(in.GetId())
	m.Title = in.Title
	m.Genre = in.Genre
	m.Description = in.Description
	if in.ReleaseDate != nil {
		m.ReleaseDate = in.ReleaseDate.AsTime()
	}
	if in.CreateTime != nil {
		m.CreatedAt = in.CreateTime.AsTime()
	}
	if in.UpdateTime != nil {
		m.UpdatedAt = in.UpdateTime.AsTime()
	}
}

func (m *Movie) JSON() ([]byte, error) {
	return json.Marshal(m)
}

func (m *Movie) FromJSON(data []byte) error {
	return json.Unmarshal(data, m)
}

func (m *Movie) YAML() ([]byte, error) {
	return yaml.Marshal(m)
}

func (m *Movie) FromYAML(data []byte) error {
	return yaml.Unmarshal(data, m)
}

func (m *Movie) ApplyMask(mask *fieldmaskpb.FieldMask) (map[string]any, error) {
	mask.Normalize()

	if !mask.IsValid(m.API()) {
		return nil, errors.New("invalid mask")
	}

	protoMask, err := fieldmask_utils.MaskFromProtoFieldMask(mask, strings.PascalCase)
	if err != nil {
		return nil, err
	}

	result := make(map[string]any)

	if err = fieldmask_utils.StructToMap(protoMask, m.API(), result); err != nil {
		return nil, err
	}

	return result, nil
}
