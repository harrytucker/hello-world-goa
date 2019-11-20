// Code generated by goa v3.0.7, DO NOT EDIT.
//
// Example Service gRPC client
//
// Command:
// $ goa gen github.com/harrytucker/hello-world-goa/design

package client

import (
	"context"

	example_servicepb "github.com/harrytucker/hello-world-goa/gen/grpc/example_service/pb"
	goagrpc "goa.design/goa/v3/grpc"
	goa "goa.design/goa/v3/pkg"
	"google.golang.org/grpc"
)

// Client lists the service endpoint gRPC clients.
type Client struct {
	grpccli example_servicepb.ExampleServiceClient
	opts    []grpc.CallOption
}

// NewClient instantiates gRPC client for all the Example Service service
// servers.
func NewClient(cc *grpc.ClientConn, opts ...grpc.CallOption) *Client {
	return &Client{
		grpccli: example_servicepb.NewExampleServiceClient(cc),
		opts:    opts,
	}
}

// SayHello calls the "SayHello" function in
// example_servicepb.ExampleServiceClient interface.
func (c *Client) SayHello() goa.Endpoint {
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		inv := goagrpc.NewInvoker(
			BuildSayHelloFunc(c.grpccli, c.opts...),
			nil,
			DecodeSayHelloResponse)
		res, err := inv.Invoke(ctx, v)
		if err != nil {
			return nil, goa.Fault(err.Error())
		}
		return res, nil
	}
}
