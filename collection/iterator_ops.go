package collection

import "kumarishan/goz/builder"

func Filter[A any, C any](i Iterable[A], b builder.Builder[A, C], op func(A) bool) C {
	var iter = i.Iterator()
	for iter.HasNext() {
		a, _ := iter.Next()
		if op(a) {
			b.AddOne(a)
		}
	}

	return b.Result()
}

func Map[A any, C any, B any](i Iterable[A], b builder.Builder[B, C], op func(A) B) C {
	// now I can't create collection.Map data structure :'(
	var iter = i.Iterator()
	for iter.HasNext() {
		a, _ := iter.Next()
		b.AddOne(op(a))
	}
	return b.Result()
}

func FlatMap[A any, C any, B any](i Iterable[A], b builder.Builder[B, C], op func(A) Iterable[B]) C {
	var iter = i.Iterator()
	for iter.HasNext() {
		a, _ := iter.Next()
		ib := op(a)
		if ib != nil {
			builderAddAll(b, ib)
		}
	}
	return b.Result()
}

func Flatten[A any, C any](i Iterable[Iterable[A]], b builder.Builder[A, C]) C {
	return FlatMap(i, b, func(a Iterable[A]) Iterable[A] { return a })
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

func Take[A any, C any](i Iterable[A], b builder.Builder[A, C], n int) C {
	var iter = i.Iterator()
	var c = 0
	for iter.HasNext() && c < n {
		a, _ := iter.Next()
		b.AddOne(a)
		c++
	}
	return b.Result()
}

func builderAddAll[A any, C any](b builder.Builder[A, C], iter Iterable[A]) builder.Builder[A, C] {
	var it = iter.Iterator()
	for it.HasNext() {
		a, _ := it.Next()
		b.AddOne(a)
	}
	return b
}
