package checkpoint

import (
	root "checkpoint/gen/root"
	"context"
	"log"
)

// root service example implementation.
// The example methods log the requests and return zero values.
type rootsrvc struct {
	logger *log.Logger
}

var (
	resultHeader root.DefaultResult
)

// NewRoot returns the root service implementation.
func NewRoot(logger *log.Logger) root.Service {
	status := "/swagger-ui/index.html"
	resultHeader.Location = &status
	return &rootsrvc{logger}
}

// Return default redirect
func (s *rootsrvc) Default(ctx context.Context) (res *root.DefaultResult, err error) {
	res = &resultHeader
	s.logger.Print("root.default")
	return
}
