package helloworld

import (
	"context"
	"log"

	exampleservice "github.com/harrytucker/hello-world-goa/gen/example_service"
)

// Example Service service example implementation.
// The example methods log the requests and return zero values.
type exampleServicesrvc struct {
	logger *log.Logger
}

// NewExampleService returns the Example Service service implementation.
func NewExampleService(logger *log.Logger) exampleservice.Service {
	return &exampleServicesrvc{logger}
}

// SayHello implements Say Hello.
func (s *exampleServicesrvc) SayHello(ctx context.Context) (res string, err error) {
	s.logger.Print("exampleService.Say Hello")
	return
}
