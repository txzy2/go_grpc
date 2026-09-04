package repository

import (
	"database/sql"
	"errors"
	"example/first/src"
	"log"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user *src.User) error {
	if user == nil {
		return errors.New("user is nil")
	}

	if user.GetName() == "" {
		return errors.New("name must be not empty")
	}

	var id int
	query := `INSERT INTO users (name) VALUES ($1) RETURNING id`
	err := r.db.QueryRow(query, user.GetName()).Scan(&id)
	if err != nil {
		return err
	}

	user.SetId(id)
	return nil
}

func (r *UserRepository) GetById(id int32) (*src.User, error) {
	if id <= 0 {
		return nil, errors.New("id must be greater than 0")
	}

	user := &src.User{}
	var userID int
	var userName string

	query := `SELECT id, name FROM users WHERE id = $1`
	err := r.db.QueryRow(query, id).Scan(&userID, &userName)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	user.SetId(userID)
	user.SetName(userName)

	log.Printf("USER: %s, ID: %d", user.GetName(), user.GetId())

	return user, nil
}
