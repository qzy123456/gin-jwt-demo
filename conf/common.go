package conf

import "errors"

var (
	PageSize          uint    = 10
	Version           string  = "0.3.9"
	Upload            string  = "upload/"
	Dir               string  = "configs/"
	TokenExpired     error  = errors.New("Token is expired")
	TokenNotValidYet error  = errors.New("Token not active yet")
	TokenMalformed   error  = errors.New("That's not even a token")
	TokenInvalid     error  = errors.New("Couldn't handle this token:")
)
