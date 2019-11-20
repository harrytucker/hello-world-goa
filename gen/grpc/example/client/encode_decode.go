// Code generated by goa v3.0.7, DO NOT EDIT.
//
// example gRPC client encoders and decoders
//
// Command:
// $ goa gen github.com/harrytucker/hello-world-goa/design

package client

import (
	"context"

	examplepb "github.com/harrytucker/hello-world-goa/gen/grpc/example/pb"
	goagrpc "goa.design/goa/v3/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// BuildHelloFunc builds the remote method to invoke for "example" service
// "hello" endpoint.
func BuildHelloFunc(grpccli examplepb.ExampleClient, cliopts ...grpc.CallOption) goagrpc.RemoteFunc {
	return func(ctx context.Context, reqpb interface{}, opts ...grpc.CallOption) (interface{}, error) {
		for _, opt := range cliopts {
			opts = append(opts, opt)
		}
		if reqpb != nil {
			return grpccli.Hello(ctx, reqpb.(*examplepb.HelloRequest), opts...)
		}
		return grpccli.Hello(ctx, &examplepb.HelloRequest{}, opts...)
	}
}

// DecodeHelloResponse decodes responses from the example hello endpoint.
func DecodeHelloResponse(ctx context.Context, v interface{}, hdr, trlr metadata.MD) (interface{}, error) {
	message, ok := v.(*examplepb.HelloResponse)
	if !ok {
		return nil, goagrpc.ErrInvalidType("example", "hello", "*examplepb.HelloResponse", v)
	}
	res := NewHelloResult(message)
	return res, nil
}