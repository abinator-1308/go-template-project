package database

import (
	"time"
)

type User struct {
	ID        string    `dynamodbav:"id"`
	Key       string    `dynamodbav:"key"`
	Name      string    `dynamodbav:"name"`
	CreatedAt time.Time `dynamodbav:"created_at"`
	UpdatedAt time.Time `dynamodbav:"updated_at"`
}
