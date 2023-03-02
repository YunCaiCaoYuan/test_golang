package interview

import (
	"fmt"
	"testing"
)

//panic: runtime error: index out of range [0] with length 0 [recovered]
//panic: runtime error: index out of range [0] with length 0
func Test_over_slice(t *testing.T) {
	list := make([]int, 0, 10)
	list[0] = 1
}

//panic: runtime error: index out of range [5] with length 3 [recovered]
//panic: runtime error: index out of range [5] with length 3
func Test_over_array(t *testing.T) {
	array := []int{1, 2, 3}
	fmt.Println(array[5])
}

//panic: interface conversion: interface {} is string, not int [recovered]
//panic: interface conversion: interface {} is string, not int
func Test_assert(t *testing.T) {
	var name interface{} = "frank"
	a := name.(int)
	fmt.Println(a)
}
