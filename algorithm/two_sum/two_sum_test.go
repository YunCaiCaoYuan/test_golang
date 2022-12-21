package two_sum

import (
	"fmt"
	"testing"
	"unsafe"
)

func Test_sizeof(t *testing.T) {
	fmt.Println("sizeof bool", unsafe.Sizeof(true))
	fmt.Println("sizeof struct{}", unsafe.Sizeof(struct{}{}))
}

//sizeof bool 	 ：1
//sizeof struct{}：0

func twoSum(nums []int, target int) []int {
	res := make([]int, 0)
	hashTable := make(map[int]struct{})
	for _, item := range nums {
		hashTable[item] = struct{}{}
	}
	for _, item := range nums {
		y := target - item
		if _, ok := hashTable[y]; ok {
			res = append(res, item, y)
		}
	}
	return res
}

// 官方答案
func twoSumV2(nums []int, target int) []int {
	hashTable := make(map[int]int)
	for i, x := range nums {
		if p, ok := hashTable[target-x]; ok {
			return []int{p, i}
		}
		hashTable[x] = i
	}
	return nil
}

func Test_twoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	res := twoSum(nums, target)
	fmt.Println(res)
}

func Test_twoSumV2(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	res := twoSumV2(nums, target)
	fmt.Println(res)
}
