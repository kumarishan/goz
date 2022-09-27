package functional_test

import (
	"kumarishan/goz/functional"
	"testing"
)

type IntAddSemigroup struct{}

func (IntAddSemigroup) Combine(x int, y int) int {
	return x + y
}

type SliceConcatSemigroup[A any] struct{}

func (SliceConcatSemigroup[A]) Combine(x []A, y []A) []A {
	var r []A
	for _, v := range x {
		r = append(r, v)
	}

	for _, v := range y {
		r = append(r, v)
	}
	return r
}

func TestSemigroup(t *testing.T) {
	x := 1
	y := 2
	z := 3

	var intAdd functional.Semigroup[int] = IntAddSemigroup{}

	intAdd.Combine(x, y)
	intAdd.Combine(x, intAdd.Combine(y, z))
	intAdd.Combine(intAdd.Combine(x, y), z)

	p := []int{1}
	q := []int{2}
	r := []int{3}

	var sliceConcat functional.Semigroup[[]int] = SliceConcatSemigroup[int]{}
	sliceConcat.Combine(p, q)
	sliceConcat.Combine(p, sliceConcat.Combine(q, r))
	sliceConcat.Combine(sliceConcat.Combine(p, q), r)
}
