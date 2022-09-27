package goz

////// Iterator //////

type Iterator[A any] interface {
	HasNext() bool
	Next() (A, error)
	Iterator() Iterator[A]
}

type IteratorOps[A any, B any] struct {
	Iterator[A]
}

func (i IteratorOps[A, B]) Filter(op func(A) bool) Iterator[A] {
	return nil
}

//// Iterable /////

type Iterable[A any] interface {
	Iterator() Iterator[A]
}

type IterableOps[A any, B any, C Iterable[A], CB Iterable[B], F IterableFactory[B, CB]] interface {
	Iterable[A]

	Filter(func(A) bool) C
	Map(func(A) B) CB
	FlatMap(func(A) Iterable[B]) CB
}

type IterableFactory[A any, C Iterable[A]] interface {
	From(source Iterable[A]) C
	Empty() C
}

func FoldLeft[A any, B any](i Iterable[A], z B, op func(B, A) B) B {
	var iter = i.Iterator()
	var s = z
	for iter.HasNext() {
		a, _ := iter.Next()
		s = op(s, a)
	}
	return s
}
