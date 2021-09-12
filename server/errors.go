package server

import (
	"errors"
)

var (
	ErrInvalidJSON        = errors.New("invalid json")
	ErrConvertStringToInt = errors.New("error converting string to an int")
	ErrFileOpen           = errors.New("error when opening a file")
	ErrMissingArgs        = errors.New("missing two args from command line ")
)
