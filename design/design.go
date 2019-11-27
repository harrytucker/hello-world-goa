package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("Hello world!", func() {
	Title("Hello world!")
	Description("Example API that demonstrates REST & gRPC.")
	Server("Hello world!", func() {
		Services("example", "openapi", "ipaddr")

		Host("localhost", func() {
			URI("http://localhost:8000")
		})
	})
})

// Service defines a group of remotely accessible methods that are hosted together.
// The service DSL makes it possible to define the methods, their input and output
// as well as the errors they may return independently of the underlying transport (HTTP or gRPC).
var _ = Service("example", func() {
	Description("The example service returns a hello world message")

	// Method defines a single service method.
	Method("hello", func() {
		Result(String)

		HTTP(func() {
			// route and method declaration
			GET("/helloworld")

			// define response code
			Response(StatusOK) // 200
		})
	})
})

// We can define multiple services in a single design for some logical separation
// of responsibilities.
//
// this has an impact on code generation, you'll need to run goa example again
// to get a stub file for the new service (after you've generated the backend)
var _ = Service("openapi", func() {
	Description("Download the OpenAPI specification for the hello world API.")

	Files("/openapi.json", "./gen/http/openapi.json")
})

var _ = Service("ipaddr", func() {
	Method("ip", func() {
		Description("Returns the public IP address of the requester")

		Result(String)

		HTTP(func() {
			GET("/ip")

			Response(StatusOK)
		})
	})
})
