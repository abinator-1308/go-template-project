package clients

import (
	"context"
	"github.com/harishb2k/go-template-project/pkg/clients/jsonplaceholder"
)

// JsonPlaceholderWriterClient gives a client to publish a new post
type JsonPlaceholderWriterClient interface {
	Publish(post *jsonplaceholder.Post) error
}

// JsonPlaceholderClient gives a client to get posts
type JsonPlaceholderClient interface {
	FetchPosts(ctx context.Context) ([]*jsonplaceholder.Post, error)
	FetchPost(ctx context.Context, id string) (*jsonplaceholder.Post, error)
}
