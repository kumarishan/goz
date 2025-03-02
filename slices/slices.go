package slices

func Filter[A any](s []A, op func(A) bool) []A {
	var res []A
	for _, a := range s {
		if op(a) {
			res = append(res, a)
		}
	}
	return res
}

func Map[A any, B any](s []A, op func(A) B) []B {
	var res []B
	for _, a := range s {
		res = append(res, op(a))
	}
	return res
}

func FlatMap[A any, B any](s []A, op func(A) []B) []B {
	var res []B
	for _, a := range s {
		res = append(res, op(a)...)
	}
	return res
}

func Flatten[A any](s [][]A) []A {
	return FlatMap(s, func(a []A) []A { return a })
}

func FoldLeft[A any, B any](s []A, z B, op func(B, A) B) B {
	var res B = z
	for _, a := range s {
		res = op(res, a)
	}
	return res
}

func Take[A any, C any](s []A, n int) []A {
	var res []A
	c := 0
	for _, a := range s {
		res = append(res, a)
		c++
		if c == n {
			break
		}
	}
	return res
}
