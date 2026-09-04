package src

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type ChatMessage struct {
	id       uuid.UUID `json:"id"`
	extId    string    `json:"ext_id"`
	chatId   int       `json:"chat_id"`
	msg      string    `json:"msg"`
	from     *User     `json:"from"`
	to       *User     `json:"to"`
	createAt time.Time `json:"create_at"`
	updateAt time.Time `json:"update_at"`
}

func NewChatMessage(msg string, from *User, to *User, chatId int) (*ChatMessage, error) {
	if msg == "" {
		return nil, errors.New("message must not be empty")
	}
	if from == nil {
		return nil, errors.New("from user must not be nil")
	}
	if to == nil {
		return nil, errors.New("to user must not be nil")
	}
	if chatId <= 0 {
		return nil, errors.New("chat id must be greater than 0")
	}

	return &ChatMessage{
		id:       uuid.New(),
		msg:      msg,
		from:     from,
		to:       to,
		chatId:   chatId,
		createAt: time.Now(),
		updateAt: time.Now(),
	}, nil
}

func (c *ChatMessage) Id() uuid.UUID {
	return c.id
}

func (c *ChatMessage) ExtId() string {
	return c.extId
}

func (c *ChatMessage) ChatId() int {
	return c.chatId
}

func (c *ChatMessage) Msg() string {
	return c.msg
}

func (c *ChatMessage) From() *User {
	return c.from
}

func (c *ChatMessage) To() *User {
	return c.to
}

func (c *ChatMessage) CreateAt() time.Time {
	return c.createAt
}

func (c *ChatMessage) UpdateAt() time.Time {
	return c.updateAt
}

func (c *ChatMessage) SetId(id uuid.UUID) {
	c.id = id
}

func (c *ChatMessage) SetExtId(extId string) {
	c.extId = extId
}

func (c *ChatMessage) SetChatId(chatId int) {
	c.chatId = chatId
}

func (c *ChatMessage) SetMsg(msg string) {
	c.msg = msg
}

func (c *ChatMessage) SetFrom(from *User) {
	c.from = from
}

func (c *ChatMessage) SetTo(to *User) {
	c.to = to
}

func (c *ChatMessage) SetCreateAt(createAt time.Time) {
	c.createAt = createAt
}

func (c *ChatMessage) SetUpdateAt(updateAt time.Time) {
	c.updateAt = updateAt
}
