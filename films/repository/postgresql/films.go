package films_postgres

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"

	"2023_2_Holi/domain"
	logs "2023_2_Holi/logger"
)

const getFilmsByGenreQuery = `
SELECT video.id, e.name, e.preview_path, video.rating
	FROM video
         JOIN video_cast AS vc ON video_id = vc.video_id
         JOIN "cast" AS c ON vc.cast_id = c.id
         JOIN episode AS e ON e.video_id = video.id
	WHERE c.name = $1
`

const getFilmDataQuery = `
	SELECT e.name, e.description, e.duration,
		e.preview_path, e.media_path, preview_video_path, release_year, rating, age_restriction
	FROM video
		JOIN episode AS e ON video.id = video_id
	WHERE video.id = $1
`

const getFilmCastQuery = `
	SELECT name
	FROM cast
		JOIN video_film AS vf ON id = cast_id
	WHERE vf.video_id = $1
`
const getCastPageQuery = `
	SELECT video.id, e.name, e.preview_path, video.rating
	FROM video
			 JOIN video_cast AS vc ON video_id = vc.video_id
			 JOIN "cast" AS c ON vc.cast_id = c.id
			 JOIN episode AS e ON e.video_id = video.id
	WHERE c.id = $1
`

const getCastNameQuery = `
		SELECT "cast".name
		FROM "cast" 
		WHERE "cast".id = $1
`

type filmsPostgresqlRepository struct {
	db  domain.PgxPoolIface
	ctx context.Context
}

func NewFilmsPostgresqlRepository(pool domain.PgxPoolIface, ctx context.Context) domain.FilmsRepository {
	return &filmsPostgresqlRepository{
		db:  pool,
		ctx: ctx,
	}
}

func (r *filmsPostgresqlRepository) GetFilmsByGenre(genre string) ([]domain.Film, error) {
	rows, err := r.db.Query(r.ctx, getFilmsByGenreQuery, genre)

	if err != nil {
		logs.LogError(logs.Logger, "films_postgresql", "GetFilmsByGenre", err, err.Error())
		return nil, err
	}
	defer rows.Close()
	logs.Logger.Debug("GetFilmsByGenre query result:", rows)

	var films []domain.Film

	for rows.Next() {
		var film domain.Film
		err = rows.Scan(
			&film.ID,
			&film.Name,
			&film.PreviewPath,
			&film.Rating,
		)

		if errors.Is(err, pgx.ErrNoRows) {
			logs.LogError(logs.Logger, "films_postgresql", "GetFilmsByGenre", err, err.Error())
			return nil, domain.ErrNotFound
		}
		if err != nil {
			logs.LogError(logs.Logger, "films_postgresql", "GetFilmsByGenre", err, err.Error())
			return nil, err
		}

		films = append(films, film)
	}

	return films, nil
}

func (r *filmsPostgresqlRepository) GetFilmData(id int) (domain.Film, error) {
	row := r.db.QueryRow(r.ctx, getFilmDataQuery, id)

	logs.Logger.Debug("GetFilmData query result:", row)

	film := new(domain.Film)
	err := row.Scan(
		&film.Name,
		&film.Description,
		&film.Duration,
		&film.PreviewPath,
		&film.MediaPath,
		&film.PreviewVideoPath,
		&film.ReleaseYear,
		&film.Rating,
		&film.AgeRestriction,
	)

	if err == pgx.ErrNoRows {
		logs.LogError(logs.Logger, "films_postgresql", "GetFilmData", err, "No rows")
		return domain.Film{}, domain.ErrNotFound
	}
	if err != nil {
		logs.LogError(logs.Logger, "films_postgresql", "GetFilmData", err, err.Error())
		return domain.Film{}, err
	}

	return domain.Film{}, nil
}

func (r *filmsPostgresqlRepository) GetFilmCast(FilmId int) ([]domain.Cast, error) {
	rows, err := r.db.Query(r.ctx, getFilmCastQuery, FilmId)
	if err == pgx.ErrNoRows {
		logs.LogError(logs.Logger, "films_postgresql", "GetFilmCast", err, err.Error())
		return nil, domain.ErrNotFound
	}
	if err != nil {
		logs.LogError(logs.Logger, "films_postgresql", "GetFilmCast", err, err.Error())
		return nil, err
	}
	defer rows.Close()
	logs.Logger.Debug("GetFilmCast query result:", rows)

	var artists []domain.Cast
	for rows.Next() {
		var artist domain.Cast
		err = rows.Scan(
			&artist.Name,
		)
		if err != nil {
			logs.LogError(logs.Logger, "films_postgresql", "GetFilmCast", err, err.Error())
			return nil, err
		}
		artists = append(artists, artist)
	}

	return artists, nil
}

func (r *filmsPostgresqlRepository) GetCastPage(id int) ([]domain.Film, error) {
	rows, err := r.db.Query(r.ctx, getCastPageQuery, id)
	if err == pgx.ErrNoRows {
		logs.LogError(logs.Logger, "cast_postgres", "GetCastPage", err, err.Error())
		return nil, domain.ErrNotFound
	}
	if err != nil {
		logs.LogError(logs.Logger, "cast_postgres", "GetCastPage", err, err.Error())
		return nil, err
	}
	defer rows.Close()
	logs.Logger.Debug("GetCastPage query result:", rows)

	var films []domain.Film
	for rows.Next() {
		var film domain.Film
		err = rows.Scan(
			&film.ID,
			&film.Name,
			&film.PreviewPath,
			&film.Rating,
		)

		if err != nil {
			return nil, err
		}

		films = append(films, film)
	}

	return films, nil
}

func (r *filmsPostgresqlRepository) GetCastName(FilmId int) (domain.Cast, error) {
	rows := r.db.QueryRow(r.ctx, getCastNameQuery, FilmId)

	logs.Logger.Debug("GetCastName query result:", rows)

	var cast domain.Cast
	err := rows.Scan(
		&cast.Name,
	)

	if err == pgx.ErrNoRows {
		logs.LogError(logs.Logger, "films_postgresql", "GetCastName", err, err.Error())
		return domain.Cast{}, domain.ErrNotFound
	}
	if err != nil {
		logs.LogError(logs.Logger, "films_postgresql", "GetCastName", err, err.Error())
		return domain.Cast{}, err
	}

	return cast, nil
}
