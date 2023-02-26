package delivery

import "errors"

var (
	InvalidDataError = errors.New("invalid user data")
	ServerError      = errors.New("something went wrong")
)
