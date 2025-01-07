package user

import (
	"database/sql"
	"fmt"
	"to-do-api/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetUserByEmail(email string) (*types.Users, error) {
	rows, err := s.db.Query("SELECT * FROM users WHERE email=?", email)
	if err != nil {
		return nil, err
	}
	user := new(types.Users)
	for rows.Next() {
		user, err = scanRowsIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}
	if user.Id == 0 {
		return nil, fmt.Errorf("user not found")
	}
	return user, nil
}

func scanRowsIntoUser(rows *sql.Rows) (*types.Users, error) {
	user := new(types.Users)
	err := rows.Scan(
		&user.Id,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return user, nil
}
