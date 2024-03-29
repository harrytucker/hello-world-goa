// Code generated by goa v3.0.7, DO NOT EDIT.
//
// example endpoints
//
// Command:
// $ goa gen github.com/harrytucker/hello-world-goa/design

package example

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Endpoints wraps the "example" service endpoints.
type Endpoints struct {
	Hello goa.Endpoint
}

// NewEndpoints wraps the methods of the "example" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		Hello: NewHelloEndpoint(s),
	}
}

// Use applies the given middleware to all the "example" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.Hello = m(e.Hello)
}

// NewHelloEndpoint returns an endpoint function that calls the method "hello"
// of service "example".
func NewHelloEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.Hello(ctx)
	}
}
