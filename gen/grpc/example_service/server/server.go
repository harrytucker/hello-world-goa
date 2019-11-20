// Code generated by goa v3.0.7, DO NOT EDIT.
//
// Example Service gRPC server
//
// Command:
// $ goa gen github.com/harrytucker/hello-world-goa/design

package server

import (
	"context"

	exampleservice "github.com/harrytucker/hello-world-goa/gen/example_service"
	example_servicepb "github.com/harrytucker/hello-world-goa/gen/grpc/example_service/pb"
	goagrpc "goa.design/goa/v3/grpc"
	goa "goa.design/goa/v3/pkg"
)

// Server implements the example_servicepb.ExampleServiceServer interface.
type Server struct {
	SayHelloH goagrpc.UnaryHandler
}

// ErrorNamer is an interface implemented by generated error structs that
// exposes the name of the error as defined in the expr.
type ErrorNamer interface {
	ErrorName() string
}

// New instantiates the server struct with the Example Service service
// endpoints.
func New(e *exampleservice.Endpoints, uh goagrpc.UnaryHandler) *Server {
	return &Server{
		SayHelloH: NewSayHelloHandler(e.SayHello, uh),
	}
}

// NewSayHelloHandler creates a gRPC handler which serves the "Example Service"
// service "Say Hello" endpoint.
func NewSayHelloHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, nil, EncodeSayHelloResponse)
	}
	return h
}

// SayHello implements the "SayHello" method in
// example_servicepb.ExampleServiceServer interface.
func (s *Server) SayHello(ctx context.Context, message *example_servicepb.SayHelloRequest) (*example_servicepb.SayHelloResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "Say Hello")
	ctx = context.WithValue(ctx, goa.ServiceKey, "Example Service")
	resp, err := s.SayHelloH.Handle(ctx, message)
	if err != nil {
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*example_servicepb.SayHelloResponse), nil
}
