package errors_test

import (
	"testing"

	"github.com/maragudk/errors"
)

func TestWrap(t *testing.T) {
	t.Run("returns wrapped error", func(t *testing.T) {
		err := errors.Wrap(errors.New("oops"), "wrap this")
		if err == nil {
			t.FailNow()
		}
		if err.Error() != "wrap this: oops" {
			t.FailNow()
		}
		unwrappedErr := errors.Unwrap(err)
		if unwrappedErr == nil {
			t.FailNow()
		}
		if unwrappedErr.Error() != "oops" {
			t.FailNow()
		}
	})

	t.Run("can use format parameters", func(t *testing.T) {
		err := errors.Wrap(errors.New("whoops"), "error in %v", 123)
		if err == nil {
			t.FailNow()
		}
		if err.Error() != "error in 123: whoops" {
			t.FailNow()
		}
	})
}
