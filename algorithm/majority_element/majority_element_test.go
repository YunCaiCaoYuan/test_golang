package majority_element

import (
	"fmt"
	"testing"
)

// 投票算法
func majorityElement(nums []int) int {
	countL := make([]int, 0, len(nums))
	candidateL := make([]int, 0, len(nums))

	count := 0
	candidate := 0
	for _, num := range nums {
		if count == 0 {
			candidate = num
		}

		if num == candidate {
			count += 1
		} else {
			count += -1
		}

		candidateL = append(candidateL, candidate)
		countL = append(countL, count)
	}
	printArr("nums:", nums)
	printArr("candidateL:", candidateL)
	printArr("countL:", countL)
	return candidate
}

// 时间复杂度: O(N)
// 空间复杂度: O(1)

func printArr(tag string, list []int) {
	fmt.Printf("%11s ", tag)
	for _, item := range list {
		fmt.Printf("%d ", item)
	}
	fmt.Println("")
}

func Test_majorityElement1(t *testing.T) {
	nums := []int{3, 2, 3}
	fmt.Println(majorityElement(nums))
}

func Test_majorityElement2(t *testing.T) {
	nums := []int{2, 2, 1, 1, 1, 2, 2}
	fmt.Println(majorityElement(nums))
}
