// Code generated by goa v3.0.7, DO NOT EDIT.
//
// ipaddr HTTP server
//
// Command:
// $ goa gen github.com/harrytucker/hello-world-goa/design

package server

import (
	"context"
	"net/http"

	ipaddr "github.com/harrytucker/hello-world-goa/gen/ipaddr"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Server lists the ipaddr service endpoint HTTP handlers.
type Server struct {
	Mounts []*MountPoint
	IP     http.Handler
}

// ErrorNamer is an interface implemented by generated error structs that
// exposes the name of the error as defined in the design.
type ErrorNamer interface {
	ErrorName() string
}

// MountPoint holds information about the mounted endpoints.
type MountPoint struct {
	// Method is the name of the service method served by the mounted HTTP handler.
	Method string
	// Verb is the HTTP method used to match requests to the mounted handler.
	Verb string
	// Pattern is the HTTP request path pattern used to match requests to the
	// mounted handler.
	Pattern string
}

// New instantiates HTTP handlers for all the ipaddr service endpoints.
func New(
	e *ipaddr.Endpoints,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
	eh func(context.Context, http.ResponseWriter, error),
) *Server {
	return &Server{
		Mounts: []*MountPoint{
			{"IP", "GET", "/ip"},
		},
		IP: NewIPHandler(e.IP, mux, dec, enc, eh),
	}
}

// Service returns the name of the service served.
func (s *Server) Service() string { return "ipaddr" }

// Use wraps the server handlers with the given middleware.
func (s *Server) Use(m func(http.Handler) http.Handler) {
	s.IP = m(s.IP)
}

// Mount configures the mux to serve the ipaddr endpoints.
func Mount(mux goahttp.Muxer, h *Server) {
	MountIPHandler(mux, h.IP)
}

// MountIPHandler configures the mux to serve the "ipaddr" service "ip"
// endpoint.
func MountIPHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/ip", f)
}

// NewIPHandler creates a HTTP handler which loads the HTTP request and calls
// the "ipaddr" service "ip" endpoint.
func NewIPHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
	eh func(context.Context, http.ResponseWriter, error),
) http.Handler {
	var (
		encodeResponse = EncodeIPResponse(enc)
		encodeError    = goahttp.ErrorEncoder(enc)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "ip")
		ctx = context.WithValue(ctx, goa.ServiceKey, "ipaddr")
		var err error

		res, err := endpoint(ctx, nil)

		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			eh(ctx, w, err)
		}
	})
}
