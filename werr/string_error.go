package werr

import "fmt"

func NewStringError(s string) *StringError {
	return &StringError{s: s}
}

type StringError struct {
	s string
}

func (e *StringError) Error() string {
	return e.s
}

func (e *StringError) String() string {
	return e.Error()
}

func (e *StringError) Wrap() WrappedError {
	return e
}

// IMPL WrappedError

func (e *StringError) Unwrap() error {
	return nil
}

func (e *StringError) WithCause(cause error) *ErrorWithCause {
	return NewErrorWithCause(e, cause)
}

func (e *StringError) Prefix(prefix string) *ErrorWithPrefix {
	return NewErrorWithPrefix(prefix, e)
}

func (e *StringError) Prefixf(prefixFormat string, a ...any) *ErrorWithPrefix {
	return NewErrorWithPrefix(fmt.Sprintf(prefixFormat, a...), e)
}

func (e *StringError) Explain(explanation string) *ErrorWithExplanation {
	return NewErrorWithExplanation(e, explanation)
}

func (e *StringError) Explainf(format string, a ...any) *ErrorWithExplanation {
	return NewErrorWithExplanationf(e, format, a...)
}

func (e *StringError) WithPayload(payload any) *ErrorWithPayload {
	return NewErrorWithPayload(e, payload)
}

func (e *StringError) Format(format string, args ...any) *ErrorFormat {
	return NewErrorFormat(format, e, args...)
}
