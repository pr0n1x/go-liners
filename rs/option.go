package rs

import (
	"encoding/json"
	"errors"
)

var ErrUnwrappingOptionNoneValue = errors.New("unwrapping Option::None value")

type Option[T any] interface {
	sealedOption()
	IsSome() bool
	IsNone() bool
	Unwrap() T
	Match() (payloadRef *T, isSome bool)
}

type Some[T any] struct{ Some T }
type None[T any] struct{}

func SetNone[T any]() Option[T] {
	return None[T]{}
}
func SetSome[T any](some T) Option[T] {
	return Some[T]{Some: some}
}

func (_ Some[T]) sealedOption()     {}
func (o Some[T]) Unwrap() T         { return o.Some }
func (_ Some[T]) IsSome() bool      { return true }
func (o Some[T]) IsNone() bool      { return !o.IsSome() }
func (o Some[T]) Match() (*T, bool) { return &o.Some, true }

func (_ None[T]) sealedOption()     {}
func (_ None[T]) Unwrap() T         { panic(ErrUnwrappingOptionNoneValue) }
func (_ None[T]) IsSome() bool      { return false }
func (o None[T]) IsNone() bool      { return !o.IsSome() }
func (_ None[T]) Match() (*T, bool) { return nil, false }
func (_ None[T]) String() string    { return "None" }

func MapOption[FROM any, TO any](from Option[FROM], convert func(f FROM) TO) Option[TO] {
	if from.IsSome() {
		return Some[TO]{Some: convert(from.Unwrap())}
	}
	return None[TO]{}
}

func MapSome[FROM any, TO any](from Some[FROM], convert func(f FROM) TO) Some[TO] {
	return Some[TO]{Some: convert(from.Some)}
}

func OptionPair[T any](res *T, ok bool) Option[T] {
	if ok {
		return Some[T]{Some: *res}
	}
	return None[T]{}
}

func OptionToString[S interface{ String() string }](opt Option[S]) string {
	if s, ok := opt.(Some[S]); ok {
		return s.Some.String()
	}
	return "None"
}

type StringableOption[T interface{ String() string }] struct {
	Option[T]
}

func (o StringableOption[T]) String() string {
	return OptionToString(o.Option)
}

func (o Some[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.Some)
}

func (o None[T]) MarshalJSON() ([]byte, error) {
	return []byte("null"), nil
}
