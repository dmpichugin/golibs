package errors

import (
	"github.com/cockroachdb/errors"
)

func As(err error, target interface{}) bool {
	return errors.As(err, target)
}

func Wrap(err error, message string) error {
	return errors.Wrap(err, message)
}

func Wrapf(err error, format string, args ...interface{}) error {
	return errors.Wrapf(err, format, args...)
}

func Cause(err error) error {
	return errors.Cause(err)
}

func Errorf(format string, args ...interface{}) error {
	return errors.Errorf(format, args...)
}

func Is(err, target error) bool {
	return errors.Is(err, target)
}

func New(message string) error {
	return errors.New(message)
}

func Unwrap(err error) error {
	return errors.Unwrap(err)
}

func WithMessage(err error, message string) error {
	return errors.WithMessage(err, message)
}

func WithMessagef(err error, format string, args ...interface{}) error {
	return errors.WithMessagef(err, format, args...)
}

func WithStack(err error) error {
	return errors.WithStack(err)
}
