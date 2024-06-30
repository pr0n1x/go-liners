package werr

import (
	"fmt"
)

type ErrorWrapper interface {
	WithCause(cause error) *ErrorWithCause
	Prefix(prefix string) *ErrorWithPrefix
	Prefixf(prefixFormat string, a ...any) *ErrorWithPrefix
	Explain(explanation string) *ErrorWithExplanation
	Explainf(format string, a ...any) *ErrorWithExplanation
	WithPayload(payload any) *ErrorWithPayload
	Format(format string, args ...any) *ErrorFormat
}
type WrappedError interface {
	error
	Unwrap() error
	ErrorWrapper
}

type ToWrappedError interface {
	Wrap() WrappedError
}

func Wrap(err error) WrappedError {
	if x, ok := err.(ToWrappedError); ok {
		return x.Wrap()
	}
	return wrappedError{err: err}
}

func New(s string) WrappedError {
	return StringError(s)
}

func Errorf(format string, args ...any) WrappedError {
	return StringError(fmt.Sprintf(format, args...))
}
