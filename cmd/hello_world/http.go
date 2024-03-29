package main

import (
	"context"
	"log"
	"net/http"
	"net/url"
	"os"
	"sync"
	"time"

	example "github.com/harrytucker/hello-world-goa/gen/example"
	examplesvr "github.com/harrytucker/hello-world-goa/gen/http/example/server"
	ipaddrsvr "github.com/harrytucker/hello-world-goa/gen/http/ipaddr/server"
	openapisvr "github.com/harrytucker/hello-world-goa/gen/http/openapi/server"
	ipaddr "github.com/harrytucker/hello-world-goa/gen/ipaddr"
	goahttp "goa.design/goa/v3/http"
	httpmdlwr "goa.design/goa/v3/http/middleware"
	"goa.design/goa/v3/middleware"
)

// handleHTTPServer starts configures and starts a HTTP server on the given
// URL. It shuts down the server if any error is received in the error channel.
func handleHTTPServer(ctx context.Context, u *url.URL, exampleEndpoints *example.Endpoints, ipaddrEndpoints *ipaddr.Endpoints, wg *sync.WaitGroup, errc chan error, logger *log.Logger, debug bool) {

	// Setup goa log adapter.
	var (
		adapter middleware.Logger
	)
	{
		adapter = middleware.NewLogger(logger)
	}

	// Provide the transport specific request decoder and response encoder.
	// The goa http package has built-in support for JSON, XML and gob.
	// Other encodings can be used by providing the corresponding functions,
	// see goa.design/implement/encoding.
	var (
		dec = goahttp.RequestDecoder
		enc = goahttp.ResponseEncoder
	)

	// Build the service HTTP request multiplexer and configure it to serve
	// HTTP requests to the service endpoints.
	var mux goahttp.Muxer
	{
		mux = goahttp.NewMuxer()
	}

	// Wrap the endpoints with the transport specific layers. The generated
	// server packages contains code generated from the design which maps
	// the service input and output data structures to HTTP requests and
	// responses.
	var (
		exampleServer *examplesvr.Server
		openapiServer *openapisvr.Server
		ipaddrServer  *ipaddrsvr.Server
	)
	{
		eh := errorHandler(logger)
		exampleServer = examplesvr.New(exampleEndpoints, mux, dec, enc, eh)
		openapiServer = openapisvr.New(nil, mux, dec, enc, eh)
		ipaddrServer = ipaddrsvr.New(ipaddrEndpoints, mux, dec, enc, eh)
	}
	// Configure the mux.
	examplesvr.Mount(mux, exampleServer)
	openapisvr.Mount(mux)
	ipaddrsvr.Mount(mux, ipaddrServer)

	// Wrap the multiplexer with additional middlewares. Middlewares mounted
	// here apply to all the service endpoints.
	var handler http.Handler = mux
	{
		if debug {
			handler = httpmdlwr.Debug(mux, os.Stdout)(handler)
		}
		handler = httpmdlwr.Log(adapter)(handler)
		handler = httpmdlwr.RequestID()(handler)
	}

	// Start HTTP server using default configuration, change the code to
	// configure the server as required by your service.
	srv := &http.Server{Addr: u.Host, Handler: handler}
	for _, m := range exampleServer.Mounts {
		logger.Printf("HTTP %q mounted on %s %s", m.Method, m.Verb, m.Pattern)
	}
	for _, m := range openapiServer.Mounts {
		logger.Printf("HTTP %q mounted on %s %s", m.Method, m.Verb, m.Pattern)
	}
	for _, m := range ipaddrServer.Mounts {
		logger.Printf("HTTP %q mounted on %s %s", m.Method, m.Verb, m.Pattern)
	}

	(*wg).Add(1)
	go func() {
		defer (*wg).Done()

		// Start HTTP server in a separate goroutine.
		go func() {
			logger.Printf("HTTP server listening on %q", u.Host)
			errc <- srv.ListenAndServe()
		}()

		<-ctx.Done()
		logger.Printf("shutting down HTTP server at %q", u.Host)

		// Shutdown gracefully with a 30s timeout.
		ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
		defer cancel()

		srv.Shutdown(ctx)
	}()
}

// errorHandler returns a function that writes and logs the given error.
// The function also writes and logs the error unique ID so that it's possible
// to correlate.
func errorHandler(logger *log.Logger) func(context.Context, http.ResponseWriter, error) {
	return func(ctx context.Context, w http.ResponseWriter, err error) {
		id := ctx.Value(middleware.RequestIDKey).(string)
		w.Write([]byte("[" + id + "] encoding: " + err.Error()))
		logger.Printf("[%s] ERROR: %s", id, err.Error())
	}
}
