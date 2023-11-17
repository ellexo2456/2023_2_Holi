package postgres

import (
	"2023_2_Holi/domain"
	logs "2023_2_Holi/logger"
	"context"
)

const getSuitableFilmQuery = `
	SELECT id, name, preview_path
	FROM "video"
	WHERE name ILIKE $1
`

const getSuitableCastQuery = `
	SELECT id, name
	FROM "cast"
	WHERE name ILIKE $1
`

type searchPostgresqlRepository struct {
	db  domain.PgxPoolIface
	ctx context.Context
}

func NewSearchPostgresqlRepository(pool domain.PgxPoolIface, ctx context.Context) domain.SearchRepository {
	return &searchPostgresqlRepository{
		db:  pool,
		ctx: ctx,
	}
}

func (r *searchPostgresqlRepository) GetSuitableFilms(searchStr string) ([]domain.Film, error) {
	searchStr = "%" + searchStr + "%"
	rows, err := r.db.Query(r.ctx, getSuitableFilmQuery, searchStr)
	if err != nil {
		logs.LogError(logs.Logger, "search_postgresql", "GetSuitableFilms", err, err.Error())
		return nil, err
	}
	defer rows.Close()
	logs.Logger.Debug("GetSuitableFilms query result:", rows)

	var films []domain.Film
	for rows.Next() {
		var film domain.Film
		err = rows.Scan(
			&film.ID,
			&film.Name,
			&film.PreviewPath,
		)
		if err != nil {
			logs.LogError(logs.Logger, "search_postgresql", "GetSuitableFilms", err, err.Error())
			return nil, err
		}
		films = append(films, film)
	}
	if len(films) == 0 {
		return nil, domain.ErrNotFound
	}

	return films, nil
}

func (r *searchPostgresqlRepository) GetSuitableCast(searchStr string) ([]domain.Cast, error) {
	searchStr = "%" + searchStr + "%"
	rows, err := r.db.Query(r.ctx, getSuitableCastQuery, searchStr)
	if err != nil {
		logs.LogError(logs.Logger, "search_postgresql", "GetSuitableCast", err, err.Error())
		return nil, err
	}
	defer rows.Close()
	logs.Logger.Debug("GetSuitableCast query result:", rows)

	var cast []domain.Cast
	for rows.Next() {
		var person domain.Cast
		err = rows.Scan(
			&person.ID,
			&person.Name,
		)
		if err != nil {
			logs.LogError(logs.Logger, "search_postgresql", "GetSuitableCast", err, err.Error())
			return nil, err
		}
		cast = append(cast, person)
	}
	if len(cast) == 0 {
		return nil, domain.ErrNotFound
	}

	return cast, nil
}
