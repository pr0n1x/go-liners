package werr

import "errors"

func HasCause(err error, target error) bool {
	current := err
	for {
		if errors.Is(current, target) {
			return true
		}
		var cause error = nil
		if x, ok := current.(interface{ Cause() error }); ok {
			cause = x.Cause()
		}
		if cause == nil {
			if x, ok := current.(interface{ Unwrap() error }); ok {
				current = x.Unwrap()
				continue
			} else {
				return false
			}
		}
		current = cause
	}
}

func AsCause(err error, target any) bool {
	current := err
	for {
		if //goland:noinspection GoErrorsAs
		errors.As(current, target) {
			return true
		}
		var cause error = nil
		if x, ok := current.(interface{ Cause() error }); ok {
			cause = x.Cause()
		}
		if cause == nil {
			if x, ok := current.(interface{ Unwrap() error }); ok {
				current = x.Unwrap()
				continue
			} else {
				return false
			}
		}
		current = cause
	}
}
