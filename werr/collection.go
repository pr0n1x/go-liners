package werr

import (
	"errors"
	"strings"
)

type Collection struct {
	sep  string
	errs []error
}

func Collect(errs []error, sep ...string) *Collection {
	if len(errs) < 1 {
		return nil
	}
	errList := []error(nil)
	for _, err := range errs {
		if err == nil {
			continue
		}
		errList = append(errList, err)
	}
	if len(errList) < 1 {
		return nil
	}
	res := Collection{errs: errList, sep: "\n"}
	if len(sep) > 0 && sep[0] != "" {
		res.sep = sep[0]
	}
	return &res
}

func (e *Collection) Error() (res string) {
	if e == nil || len(e.errs) == 0 {
		return "empty errors collection"
	}
	msgs := make([]string, len(e.errs))
	for i, err := range e.errs {
		msgs[i] = err.Error()
	}
	return strings.Join(msgs, e.sep)
}

func (e *Collection) Empty() bool {
	return e == nil || len(e.errs) == 0
}

func (e *Collection) Len() int {
	if e == nil {
		return 0
	}
	return len(e.errs)
}

func (e *Collection) Append(err error) *Collection {
	if err == nil {
		return e
	}
	if e.Empty() {
		return &Collection{errs: []error{err}}
	}
	errSlice := append(e.errs, err)
	if cap(errSlice) == cap(e.errs) {
		e.errs = errSlice
		return e
	}
	return &Collection{errs: errSlice}
}

func (e *Collection) Map(convert func(error) error) *Collection {
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
	return &Collection{errs: errs}
}

func (e *Collection) List() []error {
	// Map allocates new slice
	return e.Map(func(err error) error { return err }).errs
}

func (e *Collection) Is(target error) bool {
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
