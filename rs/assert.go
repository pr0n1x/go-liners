package rs

import "fmt"

func Must[T any](t T, err error) T {
	if err != nil {
		panic(err)
	}
	return t
}

func Trust(err error, msg ...string) {
	if err != nil {
		if len(msg) > 0 {
			panic(fmt.Errorf("%s: %w", msg[0], err))
		}
	}
}

func NotNil[T any](v *T, msg ...string) *T {
	if v == nil {
		if len(msg) > 0 {
			panic(msg[0])
		}
		panic("value is nil")
	}
	return v
}

func NotEmptySlice[S ~[]T, T any](s S, msg ...string) S {
	if len(s) < 1 {
		if len(msg) > 0 {
			panic(msg[0])
		}
		panic("slice is empty")
	}
	return s
}

func NotEmptyMap[M ~map[K]V, K comparable, V any](m M, msg ...string) M {
	if len(m) < 1 {
		if len(msg) > 0 {
			panic(msg[0])
		}
		panic("map is empty")
	}
	return m
}

func NotZero[N Number](n N, msg ...string) N {
	if n == 0 {
		if len(msg) > 0 {
			panic(msg[0])
		}
		panic("value is zero")
	}
	return n
}

func LogIfErr(err error, msg ...string) bool {
	if err != nil {
		if len(msg) > 0 {
			fmt.Printf("%s: %s\n", msg[0], err.Error())
		} else {
			fmt.Println(err.Error())
		}
		return true
	}
	return false
}
