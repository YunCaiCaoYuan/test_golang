package interview

import (
	"fmt"
	"testing"
)

// 1
// 2
// 3
func Test_iter_map(t *testing.T) {
	mapInt := make(map[int]int)
	mapInt[1] = 1
	mapInt[2] = 2
	mapInt[3] = 3
	for k, v := range mapInt {
		go fmt.Println(k, v)
		// 使用的是其数值
	}

	select {}
}

//3 3
//3 3
//3 3
func Test_iter_map_2(t *testing.T) {
	mapInt := make(map[int]int)
	mapInt[1] = 1
	mapInt[2] = 2
	mapInt[3] = 3
	for k, v := range mapInt {
		go func() {
			fmt.Println(k, v)
			// 闭包中，k,v是其引用
		}()
	}

	select {}
}
