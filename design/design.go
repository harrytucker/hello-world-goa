package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("Hello world!", func() {
	Title("Hello world!")
	Description("Example API that demonstrates REST & gRPC.")
	Server("Hello world!", func() {
		Services("Example Service")

		Host("localhost", func() {
			URI("http://localhost:8000")
			URI("grpc://localhost:8080")
		})
	})
})

// Service defines a group of remotely accessible methods that are hosted together.
// The service DSL makes it possible to define the methods, their input and output
// as well as the errors they may return independently of the underlying transport (HTTP or gRPC).
var _ = Service("Example Service", func() {
	Description("The calc service performs operations on numbers.")

	// Method defines a single service method.
	Method("Say Hello", func() {
		Result(String)

		HTTP(func() {
			// route and method declaration
			GET("/hello/world")

			// define response code
			Response(StatusOK) // 200
		})

		GRPC(func() {
		})
	})

	Files("/openapi.json", "./gen/http/openapi.json")
})
