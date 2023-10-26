package films_usecase

import (
	"2023_2_Holi/domain"
	logs "2023_2_Holi/logger"
)

type filmsUsecase struct {
	filmRepo domain.FilmsRepository
}

func NewFilmsUsecase(fr domain.FilmsRepository) domain.FilmsUsecase {
	return &filmsUsecase{
		filmRepo: fr,
	}
}

func (u *filmsUsecase) GetFilmsByGenre(genre string) ([]domain.Film, error) {
	films, err := u.filmRepo.GetFilmsByGenre(genre)
	if err != nil {
		logs.LogError(logs.Logger, "films_usecase", "GetFilmsByGenre", err, err.Error())
		return nil, err
	}
	logs.Logger.Debug("Usecase GetFilmsByGenre:", films)
	return films, nil
}

func (u *filmsUsecase) GetFilmData(id int) (domain.Film, []domain.Artist, error) {
	film, err := u.filmRepo.GetFilmData(id)
	if err != nil {
		logs.LogError(logs.Logger, "films_usecase", "GetFilmData", err, err.Error())
		return domain.Film{}, nil, err
	}
	artists, err := u.filmRepo.GetFilmArtists(id)
	if err != nil {
		return domain.Film{}, nil, err
	}

	return film, artists, nil
}
