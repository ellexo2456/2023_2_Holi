package postgresql

import (
	"database/sql"
	"fmt"

	"2023_2_Holi/domain"
)

type authPostgresqlRepository struct {
	db *sql.DB
}

func NewAuthPostgresqlRepository(conn *sql.DB) domain.AuthRepository {
	return &authPostgresqlRepository{conn}
}

func (r *authPostgresqlRepository) GetByEmail(email string) (domain.User, error) {
	result, err := r.db.Query(`SELECT id, email, password FROM "user" WHERE email = ?`, email)
	if err != nil {
		return domain.User{}, err
	}
	defer func(result *sql.Rows) {
		err := result.Close()
		if err != nil {
			fmt.Println("Error: ", err)
		}
	}(result)

	var user domain.User
	for result.Next() {
		err = result.Scan(
			&user.ID,
			&user.Email,
			&user.Password,
		)

		if err != nil {
			return domain.User{}, err
		}
	}

	return user, nil
}

func (r *authPostgresqlRepository) AddUser(user domain.User) (int, error) {
	if user.Email == "" || user.Password == "" {
		return 0, domain.ErrBadRequest
	}

	result := r.db.QueryRow(
		`INSERT INTO "user" (password, name, email, date_joined, image_path) 
		VALUES (?, ?, ?, ?, ?) RETURNING id`,
		user.Password,
		user.Name,
		user.Email,
		user.DateJoined,
		user.ImagePath)

	var id int
	if err := result.Scan(&id); err != nil {
		return 0, domain.ErrInternalServerError
	}
	return id, nil
}

func (r *authPostgresqlRepository) UserExists(email string) (bool, error) {
	result := r.db.QueryRow(`SELECT EXISTS(SELECT 1 FROM "user" WHERE email = ?)`, email)

	var exist bool
	if err := result.Scan(&exist); err != nil {
		return false, err
	}

	return exist, nil
}

type sessionPostgresqlRepository struct {
	db *sql.DB
}

func NewSessionPostgresqlRepository(conn *sql.DB) domain.SessionRepository {
	return &sessionPostgresqlRepository{conn}
}

func (s *sessionPostgresqlRepository) Add(session domain.Session) error {
	_, err := s.db.Exec(`INSERT INTO "session" VALUES (?, ?, ?)`, session.Token, session.ExpiresAt, session.UserID)
	if err != nil {
		return err
	}
	return nil
}

func (s *sessionPostgresqlRepository) DeleteByToken(token string) error {
	_, err := s.db.Exec(`DELETE FROM "session" WHERE token = ?`, token)
	if err != nil {
		return err
	}
	return nil
}

func (s *sessionPostgresqlRepository) SessionExists(token string) (bool, error) {
	result := s.db.QueryRow(`SELECT EXISTS(SELECT 1 FROM "session" WHERE token = ?)`, token)

	var exist bool
	if err := result.Scan(&exist); err != nil {
		return false, err
	}

	return exist, nil
}
