package collection_test

import (
	"fmt"
	"kumarishan/goz/collection"
	"testing"
)

func TestSlice(t *testing.T) {
	var s collection.Slice[int] = []int{1, 2}
	s = collection.Filter[int, collection.Slice[int]](s, &collection.SliceBuilder[int]{}, func(i int) bool { return i%2 == 0 })
	fmt.Printf("%v\n", s)

	mapper := func(i int) string { return fmt.Sprintf("%d", i) }
	var sStr = collection.Map[int, collection.Slice[string], string](s, &collection.SliceBuilder[string]{}, mapper)
	fmt.Printf("%v\n", sStr)

}
