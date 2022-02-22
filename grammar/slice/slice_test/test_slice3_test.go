package slice_test

import (
	"fmt"
	"testing"
)

func newSlice() []int {
	//return nil

	list := make([]int, 0)
	list = append(list, 1)
	list = append(list, 2)
	//list = append(list, 3)
	return list
}

func Test_GetTop(t *testing.T) {
	list := newSlice()

	if len(list) > 2 {
		list = list[:2]
	}
	fmt.Println("GetTop", list)
}

func array() [1024]int {
	var x [1024]int
	for i := 0; i < len(x); i++ {
		x[i] = i
	}
	return x
}

func slice() []int {
	x := make([]int, 1024)
	for i := 0; i < len(x); i++ {
		x[i] = i
	}
	return x
}

func BenchmarkArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		array()
	}
}

func BenchmarkSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		slice()
	}
}

// go test -bench . -benchmem -gcflags "-N -l"
