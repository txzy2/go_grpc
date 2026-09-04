package repository

import (
	"database/sql"
	"errors"
	"example/first/src"
)

type ChatRepository struct {
	db *sql.DB
}

func NewChatRepository(db *sql.DB) *ChatRepository {
	return &ChatRepository{db: db}
}

func (repo *ChatRepository) Create(user *src.User, title string) (string, error) {
	if user == nil {
		return "", errors.New("user must not be nil")
	}

	var extId string
	query := `INSERT INTO chats (title, created_by) VALUES ($1, $2) RETURNING ext_id`
	err := repo.db.QueryRow(query, title, user.GetId()).Scan(&extId)
	if err != nil {
		return "", err
	}

	return extId, nil
}
