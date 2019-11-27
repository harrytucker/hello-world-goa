// Code generated by goa v3.0.7, DO NOT EDIT.
//
// ipaddr endpoints
//
// Command:
// $ goa gen github.com/harrytucker/hello-world-goa/design

package ipaddr

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Endpoints wraps the "ipaddr" service endpoints.
type Endpoints struct {
	IP goa.Endpoint
}

// NewEndpoints wraps the methods of the "ipaddr" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		IP: NewIPEndpoint(s),
	}
}

// Use applies the given middleware to all the "ipaddr" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.IP = m(e.IP)
}

// NewIPEndpoint returns an endpoint function that calls the method "ip" of
// service "ipaddr".
func NewIPEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.IP(ctx)
	}
}