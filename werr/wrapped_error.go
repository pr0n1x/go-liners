package werr

import "fmt"

type wrappedError struct {
	err error
}

func (e wrappedError) Error() string { return e.err.Error() }

func (e wrappedError) String() string { return e.err.Error() }

func (e wrappedError) Wrap() WrappedError { return e }

// IMPL WrappedError

func (e wrappedError) Unwrap() error { return e.err }

func (e wrappedError) WithCause(cause error) *ErrorWithCause {
	return NewErrorWithCause(e.err, cause)
}

func (e wrappedError) Prefix(prefix string) *ErrorWithPrefix {
	return NewErrorWithPrefix(prefix, e.err)
}

func (e wrappedError) Prefixf(prefixFormat string, a ...any) *ErrorWithPrefix {
	return NewErrorWithPrefix(fmt.Sprintf(prefixFormat, a...), e.err)
}

func (e wrappedError) Explain(explanation string) *ErrorWithExplanation {
	return NewErrorWithExplanation(e.err, explanation)
}

func (e wrappedError) Explainf(format string, a ...any) *ErrorWithExplanation {
	return NewErrorWithExplanationf(e.err, format, a...)
}

func (e wrappedError) WithPayload(payload any) *ErrorWithPayload {
	return NewErrorWithPayload(e.err, payload)
}

func (e wrappedError) Format(format string, args ...any) *ErrorFormat {
	return NewErrorFormat(format, e.err, args...)
}
