package functional_test

import (
	"fmt"
	"kumarishan/goz/collection"
	"kumarishan/goz/functional"
	"testing"
)

type IntAddMonoid struct{}

func (IntAddMonoid) Combine(x int, y int) int {
	return IntAddSemigroup{}.Combine(x, y)
}

func (IntAddMonoid) Empty() int {
	return 0
}

type SliceConcatMonoid[A any] struct{}

func (SliceConcatMonoid[A]) Combine(x []A, y []A) []A {
	return SliceConcatSemigroup[A]{}.Combine(x, y)
}

func (SliceConcatMonoid[A]) Empty() []A {
	return []A{}
}

// of type Monoid[Option[A]]
type OptionMonoid[A any] struct {
	s functional.Semigroup[A]
}

func (o OptionMonoid[A]) Combine(x functional.Option[A], y functional.Option[A]) functional.Option[A] {
	if x.IsEmpty() {
		return y
	} else {
		xv, _ := x.Get()
		if y.IsEmpty() {
			return x
		} else {
			yv, _ := y.Get()
			return functional.SomeOf(o.s.Combine(xv, yv))
		}
	}
}

func (OptionMonoid[A]) Empty() functional.Option[A] {
	return functional.None[A]{}
}

type NonEmptySlice[A any] struct {
	Head A
	Tail []A
}

type NonEmptySliceConcatSemigroup[A any] struct{}

func (NonEmptySliceConcatSemigroup[A]) Combine(x NonEmptySlice[A], y NonEmptySlice[A]) NonEmptySlice[A] {
	var tail []A
	for _, v := range x.Tail {
		tail = append(tail, v)
	}
	tail = append(tail, y.Head)
	for _, v := range y.Tail {
		tail = append(tail, v)
	}

	return NonEmptySlice[A]{
		Head: x.Head,
		Tail: tail,
	}
}

func TestMonoid(t *testing.T) {
	x := 1
	var intAdd functional.Monoid[int] = IntAddMonoid{}
	intAdd.Combine(x, intAdd.Empty())
	intAdd.Combine(intAdd.Empty(), x)

	y := []int{1}
	var sliceConcat functional.Monoid[[]int] = SliceConcatMonoid[int]{}
	sliceConcat.Combine(y, sliceConcat.Empty())
	sliceConcat.Combine(sliceConcat.Empty(), y)

	var nonEmptySliceConcat functional.Monoid[functional.Option[NonEmptySlice[int]]] = OptionMonoid[NonEmptySlice[int]]{NonEmptySliceConcatSemigroup[int]{}}

	p := NonEmptySlice[int]{1, []int{2, 3}}
	q := NonEmptySlice[int]{4, []int{5, 6}}
	op := functional.OptionOf(p)
	oq := functional.OptionOf(q)
	or := functional.None[NonEmptySlice[int]]{}
	os := nonEmptySliceConcat.Combine(op, oq)
	o, err := os.Get()
	fmt.Println(o)
	if err != nil {
		t.Errorf("err=-= nil")
	}

	os = nonEmptySliceConcat.Combine(op, or)
	o, err = os.Get()
	fmt.Println(o)
	if err != nil {
		t.Errorf("err == nil")
	}

	os = nonEmptySliceConcat.Combine(or, op)
	o, err = os.Get()
	fmt.Println(o)
	if err != nil {
		t.Errorf("err == nil")
	}

	os = nonEmptySliceConcat.Combine(or, or)
	o, err = os.Get()
	fmt.Println(o)
	fmt.Println(err)
	if err == nil {
		t.Errorf("err != nil")
	}

	s := [][]int{{1, 2}, {3, 4}, {5, 6, 7}}

	r := functional.MonoidOps[[]int, functional.Monoid[[]int]]{sliceConcat}.CombineAll(collection.Slice[[]int](s))
	fmt.Println(r)

}
