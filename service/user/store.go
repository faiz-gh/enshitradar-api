package user

import (
	"database/sql"

	"github.com/faiz-gh/enshitradar-api/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) AddUser() (*types.User, error) {
	var user types.User
	err := s.db.QueryRow(
		"INSERT INTO users (id) VALUES (uuid_generate_v4()) RETURNING id",
	).Scan(&user.ID)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *Store) GetUserByID(id string) (*types.User, error) {
	rows, err := s.db.Query("SELECT * FROM users WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, nil
	}

	user, err := scanRowsIntoUsers(rows)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func scanRowsIntoUsers(rows *sql.Rows) (*types.User, error) {
	user := new(types.User)

	err := rows.Scan(
		&user.ID,
	)
	if err != nil {
		return nil, err
	}

	return user, nil
}
