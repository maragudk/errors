// Package errors provides [Wrap] in addition to the functions in the [errors] package.
// It also provides [Newf] which is a wrapper around [fmt.Errorf].
package errors

import (
	"errors"
	"fmt"
)

// As calls [errors.As].
func As(err error, target interface{}) bool {
	return errors.As(err, target)
}

// Is calls [errors.Is].
func Is(err, target error) bool {
	return errors.Is(err, target)
}

// Join calls [errors.Join].
func Join(errs ...error) error {
	return errors.Join(errs...)
}

// New calls [errors.New].
func New(text string) error {
	return errors.New(text)
}

// Newf calls [fmt.Errorf].
func Newf(format string, a ...interface{}) error {
	return fmt.Errorf(format, a...)
}

// Unwrap calls [errors.Unwrap].
func Unwrap(err error) error {
	return errors.Unwrap(err)
}

// Wrap an error with a message, or return nil if err is nil.
func Wrap(err error, format string, a ...interface{}) error {
	if err == nil {
		return nil
	}
	var fmtArgs []interface{}
	fmtArgs = append(fmtArgs, a...)
	fmtArgs = append(fmtArgs, err)
	return fmt.Errorf(format+": %w", fmtArgs...)
}
