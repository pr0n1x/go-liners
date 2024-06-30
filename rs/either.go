package rs

type Either[L any, R any] interface {
	sealedEither()
	IsLeft() bool
	IsRight() bool
	UnwrapLeft() L
	UnwrapRight() R
	Match() (*L, *R)
}

type Left[L any, R any] struct{ Left L }
type Right[L any, R any] struct{ Right R }

func (l Left[L, R]) sealedEither()   {}
func (l Left[L, R]) IsLeft() bool    { return true }
func (l Left[L, R]) IsRight() bool   { return false }
func (l Left[L, R]) UnwrapLeft() L   { return l.Left }
func (l Left[L, R]) UnwrapRight() R  { panic("right value of Either is empty") }
func (l Left[L, R]) Match() (*L, *R) { return &l.Left, nil }

func (r Right[L, R]) sealedEither()   {}
func (r Right[L, R]) IsLeft() bool    { return false }
func (r Right[L, R]) IsRight() bool   { return true }
func (r Right[L, R]) UnwrapLeft() L   { panic("left value of Either is empty") }
func (r Right[L, R]) UnwrapRight() R  { return r.Right }
func (r Right[L, R]) Match() (*L, *R) { return nil, &r.Right }
