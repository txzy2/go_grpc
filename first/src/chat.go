package src

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Chat struct {
	id    int       `json:"id"`
	extId uuid.UUID `json:"ext_id"`

	title       string `json:"title"`
	lastMessage string `json:"last_message"`

	createdBy User      `json:"created_by"`
	createdAt time.Time `json:"created_at"`
	updatedAt time.Time `json:"updated_at"`
}

func NewChat(title string, createdBy *User) (*Chat, error) {
	if title == "" {
		return nil, errors.New("title must not be empty")
	}
	if createdBy == nil {
		return nil, errors.New("created by user must not be nil")
	}

	now := time.Now()
	return &Chat{
		extId:     uuid.New(),
		title:     title,
		createdBy: *createdBy,
		createdAt: now,
		updatedAt: now,
	}, nil
}

func (c *Chat) Id() int {
	return c.id
}

func (c *Chat) ExtId() uuid.UUID {
	return c.extId
}

func (c *Chat) Title() string {
	return c.title
}

func (c *Chat) LastMessage() string {
	return c.lastMessage
}

func (c *Chat) CreatedBy() User {
	return c.createdBy
}

func (c *Chat) CreatedAt() time.Time {
	return c.createdAt
}

func (c *Chat) UpdatedAt() time.Time {
	return c.updatedAt
}

func (c *Chat) SetId(id int) {
	c.id = id
}

func (c *Chat) SetExtId(extId uuid.UUID) {
	c.extId = extId
}

func (c *Chat) SetTitle(title string) {
	c.title = title
}

func (c *Chat) SetLastMessage(lastMessage string) {
	c.lastMessage = lastMessage
}

func (c *Chat) SetCreatedBy(createdBy User) {
	c.createdBy = createdBy
}

func (c *Chat) SetCreatedAt(createdAt time.Time) {
	c.createdAt = createdAt
}

func (c *Chat) SetUpdatedAt(updatedAt time.Time) {
	c.updatedAt = updatedAt
}
