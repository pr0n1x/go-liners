package werr

import "fmt"

type ErrorWithExplanation struct {
	err         error
	explanation string
}

func NewErrorWithExplanation(err error, explanation string) *ErrorWithExplanation {
	return &ErrorWithExplanation{
		err:         err,
		explanation: explanation,
	}
}

func NewErrorWithExplanationf(err error, format string, a ...any) *ErrorWithExplanation {
	return &ErrorWithExplanation{
		err:         err,
		explanation: fmt.Sprintf(format, a...),
	}
}

func (e *ErrorWithExplanation) Error() string {
	return fmt.Sprintf("%s: %s", e.err.Error(), e.explanation)
}

func (e *ErrorWithExplanation) String() string {
	return e.Error()
}

func (e *ErrorWithExplanation) Wrap() WrappedError {
	return e
}

// IMPL WrappedError

func (e *ErrorWithExplanation) Unwrap() error { return e.err }

func (e *ErrorWithExplanation) WithCause(cause error) *ErrorWithCause {
	return NewErrorWithCause(e, cause)
}

func (e *ErrorWithExplanation) Prefix(prefix string) *ErrorWithPrefix {
	return NewErrorWithPrefix(prefix, e)
}

func (e *ErrorWithExplanation) Prefixf(prefixFormat string, a ...any) *ErrorWithPrefix {
	return NewErrorWithPrefix(fmt.Sprintf(prefixFormat, a...), e)
}

func (e *ErrorWithExplanation) Explain(explanation string) *ErrorWithExplanation {
	return NewErrorWithExplanation(e, explanation)
}

func (e *ErrorWithExplanation) Explainf(format string, a ...any) *ErrorWithExplanation {
	return NewErrorWithExplanationf(e, format, a...)
}

func (e *ErrorWithExplanation) WithPayload(payload any) *ErrorWithPayload {
	return NewErrorWithPayload(e, payload)
}

func (e *ErrorWithExplanation) Format(format string, args ...any) *ErrorFormat {
	return NewErrorFormat(format, e, args...)
}
