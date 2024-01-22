package errors

import "github.com/pkg/errors"

// Auth
var (
	AuthTokenInvalid      = errors.New("auth token is invalid")
	AuthTokenExpired      = errors.New("auth token is expired")
	AuthTokenNotValidYet  = errors.New("auth token not active yet")
	AuthTokenMalformed    = errors.New("auth token is malformed")
	AuthTokenGenerateFail = errors.New("failed to generate auth token")
)
