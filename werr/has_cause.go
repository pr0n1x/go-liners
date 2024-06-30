package werr

import "errors"

func HasCause(err error, target error) bool {
	cause := err
	for {
		if errors.Is(cause, target) {
			return true
		}
		if x, ok := cause.(interface{ Cause() error }); ok {
			cause = x.Cause()
		} else {
			return false
		}
	}
}
