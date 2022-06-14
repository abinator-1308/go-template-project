package objects

import "time"

type User struct {
	ID        string    `dynamodbav:"id"`
	Property  string    `dynamodbav:"property"`
	Name      string    `dynamodbav:"name"`
	CreatedAt time.Time `dynamodbav:"created_at"`
	UpdatedAt time.Time `dynamodbav:"updated_at"`
}
