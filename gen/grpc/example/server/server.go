// Code generated by goa v3.0.7, DO NOT EDIT.
//
// example gRPC server
//
// Command:
// $ goa gen github.com/harrytucker/hello-world-goa/design

package server

import (
	"context"

	example "github.com/harrytucker/hello-world-goa/gen/example"
	examplepb "github.com/harrytucker/hello-world-goa/gen/grpc/example/pb"
	goagrpc "goa.design/goa/v3/grpc"
	goa "goa.design/goa/v3/pkg"
)

// Server implements the examplepb.ExampleServer interface.
type Server struct {
	HelloH goagrpc.UnaryHandler
}

// ErrorNamer is an interface implemented by generated error structs that
// exposes the name of the error as defined in the expr.
type ErrorNamer interface {
	ErrorName() string
}

// New instantiates the server struct with the example service endpoints.
func New(e *example.Endpoints, uh goagrpc.UnaryHandler) *Server {
	return &Server{
		HelloH: NewHelloHandler(e.Hello, uh),
	}
}

// NewHelloHandler creates a gRPC handler which serves the "example" service
// "hello" endpoint.
func NewHelloHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, nil, EncodeHelloResponse)
	}
	return h
}

// Hello implements the "Hello" method in examplepb.ExampleServer interface.
func (s *Server) Hello(ctx context.Context, message *examplepb.HelloRequest) (*examplepb.HelloResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "hello")
	ctx = context.WithValue(ctx, goa.ServiceKey, "example")
	resp, err := s.HelloH.Handle(ctx, message)
	if err != nil {
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*examplepb.HelloResponse), nil
}