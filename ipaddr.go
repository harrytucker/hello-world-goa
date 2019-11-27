package helloworld

import (
	"context"
	"log"

	ipaddr "github.com/harrytucker/hello-world-goa/gen/ipaddr"
)

// ipaddr service example implementation.
// The example methods log the requests and return zero values.
type ipaddrsrvc struct {
	logger *log.Logger
}

// NewIpaddr returns the ipaddr service implementation.
func NewIpaddr(logger *log.Logger) ipaddr.Service {
	return &ipaddrsrvc{logger}
}

// Returns the public IP address of the requester
func (s *ipaddrsrvc) IP(ctx context.Context) (res string, err error) {
	s.logger.Print("ipaddr.ip")
	s.logger.Print(ctx)
	return "your ip address here: as a service", nil
}
