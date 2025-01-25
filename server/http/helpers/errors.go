package helpers

import (
	"errors"
)

var (
	ErrRequiredBodyNotFound = errors.New("missing required body data")
)
