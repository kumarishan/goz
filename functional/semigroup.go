package functional

type Semigroup[A any] interface {
	Combine(x A, y A) A
}
