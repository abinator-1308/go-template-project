package clients

import (
	"context"
	"github.com/harishb2k/go-template-project/pkg/clients/jsonplaceholder"
)

type JsonPlaceholderWriterClient interface {
	Publish(post *jsonplaceholder.Post) error
}

type JsonPlaceholderClient interface {
	FetchPosts(ctx context.Context) ([]*jsonplaceholder.Post, error)
	FetchPost(ctx context.Context, id string) (*jsonplaceholder.Post, error)
}
