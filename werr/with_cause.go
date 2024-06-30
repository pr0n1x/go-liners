package werr

import "fmt"

type ErrorWithCause struct {
	err   error
	cause error
}

func NewErrorWithCause(err error, cause error) *ErrorWithCause {
	if err == nil {
		panic("NewErrorWithCause(err, cause): err is nil")
	}
	if cause == nil {
		panic("NewErrorWithCause(err, cause): cause is nil")
	}
	return &ErrorWithCause{
		err:   err,
		cause: cause,
	}
}

func (e *ErrorWithCause) Error() string {
	return fmt.Sprintf("%s: %s", e.err.Error(), e.cause.Error())
}

func (e *ErrorWithCause) String() string {
	return e.Error()
}

func (e *ErrorWithCause) Wrap() WrappedError {
	return e
}

func (e *ErrorWithCause) Cause() error {
	return e.cause
}

// TODO: test it for infinite recursion before use
//func (e *ErrorWithCause) Is(target error) bool {
//	return HasCause(e.cause, target) || HasCause(e.err, target)
//}

// IMPL WrappedError

func (e *ErrorWithCause) Unwrap() error {
	return e.err
}

func (e *ErrorWithCause) WithCause(cause error) *ErrorWithCause {
	return NewErrorWithCause(e, cause)
}

func (e *ErrorWithCause) Prefix(prefix string) *ErrorWithPrefix {
	return NewErrorWithPrefix(prefix, e)
}

func (e *ErrorWithCause) Prefixf(prefixFormat string, a ...any) *ErrorWithPrefix {
	return NewErrorWithPrefix(fmt.Sprintf(prefixFormat, a...), e)
}

func (e *ErrorWithCause) Explain(explanation string) *ErrorWithExplanation {
	return NewErrorWithExplanation(e, explanation)
}

func (e *ErrorWithCause) Explainf(format string, a ...any) *ErrorWithExplanation {
	return NewErrorWithExplanationf(e, format, a...)
}

func (e *ErrorWithCause) WithPayload(payload any) *ErrorWithPayload {
	return NewErrorWithPayload(e, payload)
}

func (e *ErrorWithCause) Format(format string, args ...any) *ErrorFormat {
	return NewErrorFormat(format, e, args...)
}
