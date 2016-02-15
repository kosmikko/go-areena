package areena

import "errors"

var (
	// ErrNotFound - resource was not found
	ErrNotFound       = errors.New("Resource not found")
	ErrMediaIDMissing = errors.New("Failed to find available media id")
)
