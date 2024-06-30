package rs

import "fmt"

type Ref[T any] struct {
	ptr *T
}

func NewRef[T any](ptr *T, err ...string) Ref[T] {
	if ptr == nil {
		if len(err) > 0 {
			panic(fmt.Errorf("pointer is nil: %s", err[0]))
		}
		panic("pointer is nil")
	}
	return Ref[T]{ptr: ptr}
}

func (r Ref[T]) Ptr() *T {
	return r.ptr
}

func (r Ref[T]) Deref() T {
	return *r.ptr
}

func Ptr[T any](v T) *T {
	return &v
}
