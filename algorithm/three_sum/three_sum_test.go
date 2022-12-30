package three_sum

import (
	"fmt"
	"sort"
	"testing"
)

// 三数之和=0
func threeSum(nums []int) [][]int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	printArr(nums)

	res := make([][]int, 0)
	for i, elem_i := range nums[:len(nums)-2] {
		fmt.Println("i:", i, "elem_i:", elem_i)
		if i >= 1 && elem_i == nums[i-1] { // 去重
			continue
		}
		hashTable := make(map[int]struct{})
		fmt.Printf("elem_j:")
		for _, elem_j := range nums[i+1:] {
			fmt.Printf("%d ", elem_j)
			if _, ok := hashTable[elem_j]; !ok {
				hashTable[-elem_i-elem_j] = struct{}{}
			} else {
				fmt.Printf(" res:{%d, %d, %d} ", elem_i, -elem_i-elem_j, elem_j)
				res = append(res, []int{elem_i, -elem_i - elem_j, elem_j})
			}
		}
		fmt.Println("")
		fmt.Println("hashTable:", hashTable)
		fmt.Println("")
	}

	return res
}

func printArr(nums []int) {
	for _, item := range nums {
		fmt.Printf("%d ", item)
	}
	fmt.Println("")
	fmt.Println("")
}

func Test_thresSum(t *testing.T) {
	nums := []int{-1, 0, 1, 2, -1, -4}
	res := threeSum(nums)
	fmt.Println(res)
}

//-4 -1 -1 0 1 2
//
//i: 0 elem_i: -4
//elem_j:-1 -1 0 1 2
//hashTable: map[2:{} 3:{} 4:{} 5:{}]
//
//i: 1 elem_i: -1
//elem_j:-1 0 1  res:{-1, 0, 1} 2  res:{-1, -1, 2}
//hashTable: map[1:{} 2:{}]
//
//i: 2 elem_i: -1
//i: 3 elem_i: 0
//elem_j:1 2
//hashTable: map[-2:{} -1:{}]
//
//[[-1 0 1] [-1 -1 2]]
