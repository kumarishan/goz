package goz

type S[A any] []A

func (s S[A]) Iterator() Iterator[A] {
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
	return zero, ErrNoSuchElement
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
