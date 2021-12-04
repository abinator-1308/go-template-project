package database

import (
	"context"
	"time"
)

type User struct {
	ID        string    `dynamodbav:"id"`
	Key       string    `dynamodbav:"key"`
	CreatedAt time.Time `dynamodbav:"created_at"`
	UpdatedAt time.Time `dynamodbav:"updated_at"`
}

type UserDao interface {
	Persist(ctx context.Context, user *User) error
}
