package functional

import "kumarishan/goz/collection"

type Monoid[A any] interface {
	Semigroup[A]

	Empty() A
}

type MonoidOps[A any, M Monoid[A]] struct {
	M M
}

func (o MonoidOps[A, M]) Combine(x A, y A) A {
	return o.M.Combine(x, y)
}

func (o MonoidOps[A, M]) Empty() A {
	return o.M.Empty()
}

func (o MonoidOps[A, M]) CombineAll(iter collection.Iterable[A]) A {
	return collection.FoldLeft(iter, o.Empty(), o.Combine)
}
