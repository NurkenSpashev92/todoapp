package core_errors

import "errors"

var (
	ErrNotFound       = errors.New("not found")
	ErrInvlidArgument = errors.New("invalid argument")
	ErrConflict       = errors.New("conflict")
)
