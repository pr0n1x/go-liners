package rs

import (
	"reflect"
	"sync"
)

type copier[T any] interface {
	Copy() T
}

type RWLock[T copier[T]] struct {
	lock    sync.RWMutex
	isNil   bool
	nilable bool
	value   T
}

func NewRWLock[T copier[T]](v T) RWLock[T] {
	var (
		value   T
		nilable bool
		isNil   bool
	)
	val := reflect.ValueOf(v)
	kind := val.Kind()
	switch kind {
	case
		reflect.Chan,
		reflect.Func,
		reflect.Map,
		reflect.Pointer,
		reflect.UnsafePointer,
		reflect.Interface,
		reflect.Slice:
		nilable = true
	default:
		nilable = false
	}
	isNil = nilable && val.IsNil()
	if isNil {
		value = v
	} else {
		value = v.Copy()
	}
	return RWLock[T]{
		lock:    sync.RWMutex{},
		nilable: nilable,
		isNil:   isNil,
		value:   value,
	}
}

func (s *RWLock[T]) Set(v T) {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.nilable {
		if s.isNil = reflect.ValueOf(v).IsNil(); s.isNil {
			s.value = v
			return
		}
	}
	s.value = v.Copy()
}

func (s *RWLock[T]) Get() T {
	s.lock.RLock()
	defer s.lock.RUnlock()
	if s.isNil {
		return s.value
	}
	return s.value.Copy()
}
