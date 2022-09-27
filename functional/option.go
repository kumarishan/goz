package functional

import (
	"kumarishan/goz"
	"reflect"
)

type Option[A any] interface {
	IsDefined() bool
	NonEmpty() bool
	IsEmpty() bool
	Get() (A, error)
	GetOrElse(a A) A
	OrElse(func() A) A
}

type Some[A any] struct {
	v A
}

type None[A any] struct{}

//////// Some Methods ////////

func (Some[A]) IsDefined() bool {
	return true
}
func (Some[A]) NonEmpty() bool {
	return true
}
func (Some[A]) IsEmpty() bool {
	return false
}
func (s Some[A]) Get() (A, error) {
	return s.v, nil
}
func (s Some[A]) GetOrElse(a A) A {
	return s.v
}
func (s Some[A]) OrElse(alternative func() A) A {
	return s.v
}

//////// None Methods ////////

func (None[A]) IsDefined() bool {
	return false
}
func (None[A]) NonEmpty() bool {
	return false
}
func (None[A]) IsEmpty() bool {
	return true
}
func (None[A]) Get() (A, error) {
	var zero A
	return zero, goz.ErrNoSuchElement
}
func (None[A]) GetOrElse(a A) A {
	return a
}
func (None[A]) OrElse(alternative func() A) A {
	return alternative()
}

////////// Factory Methods //////////

func OptionOf[A any](v A) Option[A] {
	if reflect.ValueOf(&v).Elem().IsZero() {
		return None[A]{}
	}

	return Some[A]{v}
}

func SomeOf[A any](v A) Option[A] {
	return Some[A]{v}
}
