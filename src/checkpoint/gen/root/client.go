// Code generated by goa v3.7.2, DO NOT EDIT.
//
// root client
//
// Command:
// $ goa gen checkpoint/design

package root

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "root" service client.
type Client struct {
	DefaultEndpoint goa.Endpoint
}

// NewClient initializes a "root" service client given the endpoints.
func NewClient(default_ goa.Endpoint) *Client {
	return &Client{
		DefaultEndpoint: default_,
	}
}

// Default calls the "default" endpoint of the "root" service.
func (c *Client) Default(ctx context.Context) (res *DefaultResult, err error) {
	var ires interface{}
	ires, err = c.DefaultEndpoint(ctx, nil)
	if err != nil {
		return
	}
	return ires.(*DefaultResult), nil
}
