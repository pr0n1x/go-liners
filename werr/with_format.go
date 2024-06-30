package werr

import (
	"fmt"
)

type ErrorFormat struct {
	err    error
	format string
	args   []any
}

func NewErrorFormat(format string, err error, args ...any) *ErrorFormat {
	return &ErrorFormat{
		err:    err,
		format: format,
		args:   args,
	}
}

func (e *ErrorFormat) Error() string {
	args := make([]any, len(e.args)+1)
	args[0] = e.err
	for i, arg := range e.args {
		args[i+1] = arg
	}
	return fmt.Errorf(e.format, args...).Error()
}

func (e *ErrorFormat) String() string { return e.Error() }

func (e *ErrorFormat) Wrap() WrappedError { return e }

// IMPL WrappedError

func (e *ErrorFormat) Unwrap() error { return e.err }

func (e *ErrorFormat) WithCause(cause error) *ErrorWithCause {
	return NewErrorWithCause(e, cause)
}

func (e *ErrorFormat) Prefix(prefix string) *ErrorWithPrefix {
	return NewErrorWithPrefix(prefix, e)
}

func (e *ErrorFormat) Prefixf(prefixFormat string, a ...any) *ErrorWithPrefix {
	return NewErrorWithPrefix(fmt.Sprintf(prefixFormat, a...), e)
}

func (e *ErrorFormat) Explain(explanation string) *ErrorWithExplanation {
	return NewErrorWithExplanation(e, explanation)
}

func (e *ErrorFormat) Explainf(format string, a ...any) *ErrorWithExplanation {
	return NewErrorWithExplanationf(e, format, a...)
}

func (e *ErrorFormat) WithPayload(payload any) *ErrorWithPayload {
	return NewErrorWithPayload(e, payload)
}

func (e *ErrorFormat) Format(format string, args ...any) *ErrorFormat {
	return NewErrorFormat(format, e, args...)
}
