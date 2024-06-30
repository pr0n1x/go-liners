package rs

import "fmt"

type Result[T any] interface {
	sealedResult()
	IsOk() bool
	IsErr() bool
	Check() error
	Unwrap() T
	Expect(msg string) T
	Expectf(format string, a ...any) T
	Match() (*T, error)
}

type Ok[T any] struct{ Ok T }
type Err[T any] struct{ Err error }

func (r Ok[T]) sealedResult()                {}
func (_ Ok[T]) IsOk() bool                   { return true }
func (_ Ok[T]) IsErr() bool                  { return false }
func (r Ok[T]) Unwrap() T                    { return r.Ok }
func (r Ok[T]) Expect(_ string) T            { return r.Ok }
func (r Ok[T]) Expectf(_ string, _ ...any) T { return r.Ok }
func (r Ok[T]) Check() error                 { return nil }
func (r Ok[T]) Match() (*T, error)           { return &r.Ok, nil }

func (_ Err[T]) sealedResult() {}
func (_ Err[T]) IsOk() bool    { return false }
func (_ Err[T]) IsErr() bool   { return true }
func (r Err[T]) Unwrap() T     { panic(r.Err) }
func (r Err[T]) Expect(msg string) T {
	panic(fmt.Sprintf("%s: %s", msg, r.Err.Error()))
}
func (r Err[T]) Expectf(format string, a ...any) T {
	panic(fmt.Sprintf("%s: %s", fmt.Sprintf(format, a...), r.Err.Error()))
}
func (r Err[T]) Check() error {
	if r.Err == nil {
		panic("Result::Err is nil")
	}
	return r.Err
}
func (r Err[T]) Match() (*T, error) { return nil, r.Err }
func (r Err[T]) String() string     { return fmt.Sprintf("Error: %s", r.Err.Error()) }

func MapResult[FROM any, TO any](from Result[FROM], convert func(f FROM) TO) Result[TO] {
	if err := from.Check(); err != nil {
		return Err[TO]{Err: err}
	}
	return Ok[TO]{Ok: convert(from.Unwrap())}
}

func MapOk[FROM any, TO any](from Ok[FROM], convert func(f FROM) TO) Ok[TO] {
	return Ok[TO]{Ok: convert(from.Ok)}
}

func ResultPair[T any](res T, err error) Result[T] {
	if err != nil {
		return Err[T]{Err: err}
	}
	return Ok[T]{Ok: res}
}

func UnwrapResults[S ~[]Result[T], T any](s S) (res []T) {
	res = nil
	for _, v := range s {
		res = append(res, v.Unwrap())
	}
	return res
}

func SplitResults[S ~[]Result[T], T any](s S) (results []T, errs []error) {
	results = nil
	errs = nil
	for _, r := range s {
		switch v := r.(type) {
		case Ok[T]:
			results = append(results, v.Ok)
		case Err[T]:
			errs = append(errs, v.Err)
		}
	}
	return results, errs
}

func ResultToString[S interface{ String() string }](res Result[S]) string {
	switch r := res.(type) {
	case Ok[S]:
		return r.Ok.String()
	case Err[S]:
		return fmt.Sprintf("Error: %s", r.Err.Error())
	}
	panic("unreachable")
}
