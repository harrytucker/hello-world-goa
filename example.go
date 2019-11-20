package helloworld

import (
	"context"
	"log"

	example "github.com/harrytucker/hello-world-goa/gen/example"
)

// example service example implementation.
// The example methods log the requests and return zero values.
type examplesrvc struct {
	logger *log.Logger
}

// NewExample returns the example service implementation.
func NewExample(logger *log.Logger) example.Service {
	return &examplesrvc{logger}
}

// Hello implements Hello.
func (s *examplesrvc) Hello(ctx context.Context) (res string, err error) {
	s.logger.Print("example.Say Hello")
	return "Hello world!", nil
}
