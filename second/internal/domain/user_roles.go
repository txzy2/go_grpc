package domain

import (
	"second/pkg/enums"
	"time"
)

type UserRoles struct {
	Id   int    `json:"id"`
	Name string `json:"name"`

	Status enums.StatusState `json:"status"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
