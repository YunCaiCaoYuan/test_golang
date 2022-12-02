package test

import (
	"fmt"
	"testing"
)

func fn(in []int) {
	in = append(in, 5)
}

func Test_slice(t *testing.T) {
	slice := make([]int, 0, 5)

	slice = append(slice, 1)
	fmt.Println(slice, len(slice), cap(slice))

	fn(slice)
	fmt.Println(slice, len(slice), cap(slice))

	s1 := slice[0:5]
	fmt.Println(s1, len(s1), cap(s1))
}

// 我的
// 1 5
// 2 5 (分析，切片相当于是底层数据结构的一个封装，slice自己的指向位置没有发生变化）
// 0 0 (答错，非智力因素-粗心）

//[1] 1 5
//[1] 1 5
//[1 5 0 0 0] 5 5
