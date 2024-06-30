package werr

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
)

type ErrorWithPayload struct {
	err     error
	payload any
}

func NewErrorWithPayload(err error, payload any) *ErrorWithPayload {
	return &ErrorWithPayload{
		err:     err,
		payload: payload,
	}
}

func (e *ErrorWithPayload) Error() string {
	return fmt.Sprintf("%s: %s", e.err.Error(), spew.Sdump(e.payload))
}

func (e *ErrorWithPayload) Wrap() WrappedError { return e }

func (e *ErrorWithPayload) Payload() any { return e.payload }

// IMPL WrappedError

func (e *ErrorWithPayload) Unwrap() error { return e.err }

func (e *ErrorWithPayload) WithCause(cause error) *ErrorWithCause {
	return NewErrorWithCause(e, cause)
}

func (e *ErrorWithPayload) Prefix(prefix string) *ErrorWithPrefix {
	return NewErrorWithPrefix(prefix, e)
}

func (e *ErrorWithPayload) Prefixf(prefixFormat string, a ...any) *ErrorWithPrefix {
	return NewErrorWithPrefix(fmt.Sprintf(prefixFormat, a...), e)
}

func (e *ErrorWithPayload) Explain(explanation string) *ErrorWithExplanation {
	return NewErrorWithExplanation(e, explanation)
}

func (e *ErrorWithPayload) Explainf(format string, a ...any) *ErrorWithExplanation {
	return NewErrorWithExplanationf(e, format, a...)
}

func (e *ErrorWithPayload) WithPayload(payload any) *ErrorWithPayload {
	return NewErrorWithPayload(e, payload)
}

func (e *ErrorWithPayload) Format(format string, args ...any) *ErrorFormat {
	return NewErrorFormat(format, e, args...)
}
