package werr

import "fmt"

func StringError(s string) WrappedError {
	return &stringError{s: s}
}

type stringError struct {
	s string
}

func (e *stringError) Error() string {
	return e.s
}

func (e *stringError) String() string {
	return e.Error()
}

func (e *stringError) Wrap() WrappedError {
	return e
}

// IMPL WrappedError

func (e *stringError) Unwrap() error {
	return nil
}

func (e *stringError) WithCause(cause error) *ErrorWithCause {
	return NewErrorWithCause(e, cause)
}

func (e *stringError) Prefix(prefix string) *ErrorWithPrefix {
	return NewErrorWithPrefix(prefix, e)
}

func (e *stringError) Prefixf(prefixFormat string, a ...any) *ErrorWithPrefix {
	return NewErrorWithPrefix(fmt.Sprintf(prefixFormat, a...), e)
}

func (e *stringError) Explain(explanation string) *ErrorWithExplanation {
	return NewErrorWithExplanation(e, explanation)
}

func (e *stringError) Explainf(format string, a ...any) *ErrorWithExplanation {
	return NewErrorWithExplanationf(e, format, a...)
}

func (e *stringError) WithPayload(payload any) *ErrorWithPayload {
	return NewErrorWithPayload(e, payload)
}

func (e *stringError) Format(format string, args ...any) *ErrorFormat {
	return NewErrorFormat(format, e, args...)
}
