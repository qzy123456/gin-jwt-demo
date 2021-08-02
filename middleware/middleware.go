package middleware

import (
	"jwtDemo/servcie"
)

type Middleware struct {
	Service *servcie.Service
}

// New middleware service
func New(s *servcie.Service) *Middleware {
	return &Middleware{
		Service: s,
	}
}

