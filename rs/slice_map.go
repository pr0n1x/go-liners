package rs

type IndexedValue[I Index, T any] struct {
	Index I
	Value T
}

type Item[T any] IndexedValue[int, T]

func Map[S ~[]F, F any, T any](s S, cb func(int, F) T) (res []T) {
	length := len(s)
	if length < 1 {
		return nil
	}
	res = make([]T, 0, length)
	for k, v := range s {
		res = append(res, cb(k, v))
	}
	return res
}

func Enumerate[S ~[]T, T any](s S) []Item[T] {
	return Map(s, func(k int, v T) Item[T] {
		return Item[T]{Index: k, Value: v}
	})
}

func Filter[S ~[]T, T any](s S, predicate func(int, T) bool) (res []T) {
	res = nil
	for k, v := range s {
		if predicate(k, v) {
			res = append(res, v)
		}
	}
	return res
}

func Find[S ~[]T, T any](s S, predicate func(int, T) bool) Option[T] {
	for k, v := range s {
		if predicate(k, v) {
			return Some[T]{Some: v}
		}
	}
	return None[T]{}
}

func FindMap[S ~[]F, F any, T any](s S, predicate func(int, F) Option[T]) Option[T] {
	for k, v := range s {
		found := predicate(k, v)
		if r, ok := found.Match(); ok {
			return Some[T]{Some: *r}
		}
	}
	return None[T]{}
}

func FindR[S ~[]T, T any](s S, predicate func(int, T) (bool, error)) Result[Option[T]] {
	for k, v := range s {
		found, err := predicate(k, v)
		if err != nil {
			return Err[Option[T]]{Err: err}
		}
		if found {
			return Ok[Option[T]]{Ok: Some[T]{Some: v}}
		}
	}
	return Ok[Option[T]]{
		Ok: None[T]{},
	}
}

func Fold[S ~[]F, F any, T any](s S, init T, cb func(accum T, index int, item F) T) T {
	accum := init
	for k, v := range s {
		accum = cb(accum, k, v)
	}
	return accum
}

func FoldR[S ~[]F, F any, T any](s S, init T, cb func(T, int, F) Result[T]) Result[T] {
	accum := init
	for k, v := range s {
		res := cb(accum, k, v)
		if res.IsErr() {
			return res
		}
		accum = res.Unwrap()
	}
	return Ok[T]{Ok: accum}
}
