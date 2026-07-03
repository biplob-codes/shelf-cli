package store

import "errors"

// ErrNotFound is returned when a lookup, update, or delete targets a row
// that doesn't exist.
var ErrNotFound = errors.New("not found")