// Code generated by goa v3.0.7, DO NOT EDIT.
//
// example gRPC client types
//
// Command:
// $ goa gen github.com/harrytucker/hello-world-goa/design

package client

import (
	examplepb "github.com/harrytucker/hello-world-goa/gen/grpc/example/pb"
)

// NewHelloRequest builds the gRPC request type from the payload of the "hello"
// endpoint of the "example" service.
func NewHelloRequest() *examplepb.HelloRequest {
	message := &examplepb.HelloRequest{}
	return message
}

// NewHelloResult builds the result type of the "hello" endpoint of the
// "example" service from the gRPC response type.
func NewHelloResult(message *examplepb.HelloResponse) string {
	result := message.Field
	return result
}