package rs

import (
	"errors"
	"strings"
)

type ErrorCollection struct {
	errs []error
}

func CollectErrors(errs []error) *ErrorCollection {
	if len(errs) < 1 {
		return nil
	}
	res := []error(nil)
	for _, err := range errs {
		if err == nil {
			continue
		}
		res = append(res, err)
	}
	if len(res) < 1 {
		return nil
	}
	return &ErrorCollection{errs: res}
}

func (e *ErrorCollection) Error() (res string) {
	if e == nil || len(e.errs) == 0 {
		return "empty errors collection"
	}
	return strings.Join(Map(e.errs, func(i int, err error) string {
		return err.Error()
	}), "\n")
}

func (e *ErrorCollection) Empty() bool {
	return e == nil || len(e.errs) == 0
}

func (e *ErrorCollection) Len() int {
	if e == nil {
		return 0
	}
	return len(e.errs)
}

func (e *ErrorCollection) Append(err error) *ErrorCollection {
	if err == nil {
		return e
	}
	if e.Empty() {
		return &ErrorCollection{errs: []error{err}}
	}
	errSlice := append(e.errs, err)
	if cap(errSlice) == cap(e.errs) {
		e.errs = errSlice
		return e
	}
	return &ErrorCollection{errs: errSlice}
}

func (e *ErrorCollection) Map(convert func(error) error) *ErrorCollection {
	if e == nil || len(e.errs) < 1 {
		return nil
	}
	errs := []error(nil)
	for _, err := range e.errs {
		if err == nil {
			continue
		}
		errs = append(errs, convert(err))
	}
	if len(errs) < 1 {
		return nil
	}
	return &ErrorCollection{errs: errs}
}

func (e *ErrorCollection) List() []error {
	// Map allocates new slice
	return e.Map(func(err error) error { return err }).errs
}

func (e *ErrorCollection) Is(target error) bool {
	if e == target {
		return true
	}
	if e == nil || len(e.errs) < 1 {
		return target == nil
	}
	for _, err := range e.errs {
		if errors.Is(err, target) {
			return true
		}
	}
	return false
}
