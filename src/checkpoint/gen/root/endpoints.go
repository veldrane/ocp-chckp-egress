// Code generated by goa v3.7.2, DO NOT EDIT.
//
// root endpoints
//
// Command:
// $ goa gen checkpoint/design

package root

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Endpoints wraps the "root" service endpoints.
type Endpoints struct {
	Default goa.Endpoint
}

// NewEndpoints wraps the methods of the "root" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		Default: NewDefaultEndpoint(s),
	}
}

// Use applies the given middleware to all the "root" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.Default = m(e.Default)
}

// NewDefaultEndpoint returns an endpoint function that calls the method
// "default" of service "root".
func NewDefaultEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.Default(ctx)
	}
}
