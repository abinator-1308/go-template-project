package jsonplaceholder

import (
	"context"
	"github.com/devlibx/gox-base"
	"github.com/devlibx/gox-base/errors"
	"github.com/devlibx/gox-base/serialization"
	goxHttpApi "github.com/devlibx/gox-http/api"
	"github.com/devlibx/gox-http/command"
)

type Client interface {
	FetchPosts(ctx context.Context) ([]*Post, error)
	FetchPost(ctx context.Context, id string) (*Post, error)
}

type clientImpl struct {
	goxHttpCtx goxHttpApi.GoxHttpContext
	gox.CrossFunction
}

func NewJsonPlaceHolderClient(cf gox.CrossFunction, goxHttpCtx goxHttpApi.GoxHttpContext) (Client, error) {
	c := clientImpl{
		goxHttpCtx:    goxHttpCtx,
		CrossFunction: cf,
	}
	return c, nil
}

func (c clientImpl) FetchPosts(ctx context.Context) ([]*Post, error) {
	return nil, errors.New("Not implemented")
}

func (c clientImpl) FetchPost(ctx context.Context, id string) (*Post, error) {
	request := command.NewGoxRequestBuilder("getPosts").
		WithContentTypeJson().
		WithPathParam("id", id).
		Build()
	if response, err := c.goxHttpCtx.Execute(ctx, "getPost", request); err != nil {
		return nil, err
	} else {
		post := &Post{}
		if err := serialization.JsonBytesToObject(response.Body, post); err != nil {
			return nil, err
		} else {
			return post, nil
		}
	}
}
