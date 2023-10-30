package domain

import "github.com/jackc/pgx/v5/pgtype"

type Film struct {
	ID               int             `json:"id"`
	Name             string          `json:"name"`
	Description      string          `json:"-"`
	PreviewPath      string          `json:"previewPath"`
	MediaPath        string          `json:"-"`
	ReleaseYear      int             `json:"-"`
	Rating           float64         `json:"rating"`
	AgeRestriction   int             `json:"-"`
	Duration         pgtype.Interval `json:"-"`
	PreviewVideoPath string          `json:"previewVideoPath"`
}

type FilmsRepository interface {
	GetFilmsByGenre(genre string) ([]Film, error)
	GetFilmData(id int) (Film, error)
	GetFilmArtists(filmId int) ([]Artist, error)
}

type FilmsUsecase interface {
	GetFilmsByGenre(genre string) ([]Film, error)
	GetFilmData(id int) (Film, []Artist, error)
}
