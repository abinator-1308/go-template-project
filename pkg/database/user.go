package database

import (
	"context"
	"time"
)

type User struct {
	ID        string    `dynamodbav:"id"`
	Key       string    `dynamodbav:"key"`
	Name      string    `dynamodbav:"name"`
	CreatedAt time.Time `dynamodbav:"created_at"`
	UpdatedAt time.Time `dynamodbav:"updated_at"`
}

type UserDao interface {
	Get(ctx context.Context, user *User) error
	Persist(ctx context.Context, user *User) error
	UpdateName(ctx context.Context, user *User) error
}
