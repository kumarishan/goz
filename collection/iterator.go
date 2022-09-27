package collection

type Iterable[A any] interface {
	Iterator() Iterator[A]
}

type Iterator[A any] interface {
	HasNext() bool
	Next() (A, error)
	Iterator() Iterator[A]
}
