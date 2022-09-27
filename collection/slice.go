package collection

import (
	"kumarishan/goz"
	"kumarishan/goz/builder"
)

type Slice[A any] []A

func (s Slice[A]) Iterator() Iterator[A] {
	return SliceIter(s)
}

type slice[A any] struct {
	idx int
	s   []A
}

func (s *slice[A]) HasNext() bool {
	return len(s.s) != 0 && s.idx < len(s.s)
}

func (s *slice[A]) Next() (A, error) {
	if s.HasNext() {
		v := s.s[s.idx]
		s.idx += 1
		return v, nil
	}
	var zero A
	return zero, goz.ErrNoSuchElement
}

func (s *slice[A]) Iterator() Iterator[A] {
	return SliceIter(s.s)
}

func SliceIter[A any](s []A) Iterator[A] {
	return &slice[A]{
		idx: 0,
		s:   s,
	}
}

type SliceBuilder[A any] struct {
	internal []A
}

func (b *SliceBuilder[A]) AddOne(ele A) builder.Builder[A, Slice[A]] {
	b.internal = append(b.internal, ele)
	return b
}

func (b *SliceBuilder[A]) Clear() builder.Builder[A, Slice[A]] {
	b.internal = []A{}
	return b
}

func (b *SliceBuilder[A]) Result() Slice[A] {
	return b.internal
}
