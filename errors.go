// Package errors provides Wrap in addition to the functions in the stdlib errors package.
// It's otherwise a drop-in replacement for the stdlib package.
package errors

import (
	"errors"
	"fmt"
)

// As calls errors.As.
func As(err error, target interface{}) bool {
	return errors.As(err, target)
}

// Is calls errors.Is.
func Is(err, target error) bool {
	return errors.Is(err, target)
}

// New calls errors.New.
func New(text string) error {
	return errors.New(text)
}

// Unwrap calls errors.Unwrap.
func Unwrap(err error) error {
	return errors.Unwrap(err)
}

// Wrap an error with a message, or return nil if err is nil.
func Wrap(err error, message string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	var fmtArgs []interface{}
	fmtArgs = append(fmtArgs, args...)
	fmtArgs = append(fmtArgs, err)
	return fmt.Errorf(message+": %w", fmtArgs...)
}
