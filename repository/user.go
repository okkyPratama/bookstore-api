package repository

import (
	"database/sql"

	"github.com/okkyPratama/bookstore-api/structs"
)

type UserRepository struct {
	DB *sql.DB
}

func (r *UserRepository) CreateUser(user *structs.User) error {
	query := `INSERT INTO users (username, password, created_by) VALUES ($1, $2, $3) RETURNING id`
	return r.DB.QueryRow(query, user.Username, user.Password, user.CreatedBy).Scan(&user.ID)
}

func (r *UserRepository) GetUserByUsername(username string) (*structs.User, error) {
	user := &structs.User{}
	query := `SELECT id, username, password, created_at, created_by, modified_at, modified_by FROM users WHERE username = $1`
	err := r.DB.QueryRow(query, username).Scan(&user.ID, &user.Username, &user.Password, &user.CreatedAt, &user.CreatedBy, &user.ModifiedAt, &user.ModifiedBy)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) UpdateUser(user *structs.User) error {
	query := `UPDATE users SET password = $1, modified_at = CURRENT_TIMESTAMP, modified_by = $2 WHERE id = $3`
	_, err := r.DB.Exec(query, user.Password, user.ModifiedBy, user.ID)
	return err
}

func (r *UserRepository) DeleteUser(id int) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := r.DB.Exec(query, id)
	return err
}
