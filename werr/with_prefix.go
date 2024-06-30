package werr

import "fmt"

type ErrorWithPrefix struct {
	err    error
	prefix string
}

func NewErrorWithPrefix(prefix string, err error) *ErrorWithPrefix {
	return &ErrorWithPrefix{
		err:    err,
		prefix: prefix,
	}
}

func (e *ErrorWithPrefix) Error() string {
	return fmt.Sprintf("%s: %s", e.prefix, e.err.Error())
}

func (e *ErrorWithPrefix) String() string { return e.Error() }

func (e *ErrorWithPrefix) Wrap() WrappedError { return e }

// IMPL WrappedError

func (e *ErrorWithPrefix) Unwrap() error { return e.err }

func (e *ErrorWithPrefix) WithCause(cause error) *ErrorWithCause {
	return NewErrorWithCause(e, cause)
}

func (e *ErrorWithPrefix) Prefix(prefix string) *ErrorWithPrefix {
	return NewErrorWithPrefix(prefix, e)
}

func (e *ErrorWithPrefix) Prefixf(prefixFormat string, a ...any) *ErrorWithPrefix {
	return NewErrorWithPrefix(fmt.Sprintf(prefixFormat, a...), e)
}

func (e *ErrorWithPrefix) Explain(explanation string) *ErrorWithExplanation {
	return NewErrorWithExplanation(e, explanation)
}

func (e *ErrorWithPrefix) Explainf(format string, a ...any) *ErrorWithExplanation {
	return NewErrorWithExplanationf(e, format, a...)
}

func (e *ErrorWithPrefix) WithPayload(payload any) *ErrorWithPayload {
	return NewErrorWithPayload(e, payload)
}

func (e *ErrorWithPrefix) Format(format string, args ...any) *ErrorFormat {
	return NewErrorFormat(format, e, args...)
}
