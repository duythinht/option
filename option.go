package option

type Option[T any] interface {
	Unwrap() T
	IsSome() bool
	IsNone() bool
	UnwrapOr(d T) T
}

// Some of value
type SomeOf[T any] struct {
	v T
}

func (s SomeOf[T]) Unwrap() T {
	return s.v
}

func (s SomeOf[T]) IsSome() bool {
	return true
}

func (s SomeOf[T]) IsNone() bool {
	return false
}

func (s SomeOf[T]) UnwrapOr(d T) T {
	return s.v
}

type NoneOf[T any] struct{}

func (s NoneOf[T]) Unwrap() T {
	panic("None is not unwrapable!")
}

func (s NoneOf[T]) IsSome() bool {
	return false
}

func (s NoneOf[T]) IsNone() bool {
	return true
}

func (s NoneOf[T]) UnwrapOr(d T) T {
	return d
}

func Some[T any](v T) Option[T] {
	return SomeOf[T]{v}
}

func None[T any]() Option[T] {
	return NoneOf[T]{}
}
