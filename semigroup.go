package goz

type Semigroup[A any] interface {
	Combine(x A, y A) A
}
