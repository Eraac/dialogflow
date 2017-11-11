package dialogflow

import "errors"

// List errors
var (
	ErrNotFound = errors.New("key not found")
	ErrCastFail = errors.New("cast fail")
)
