package max_sliding_window

import (
	"fmt"
	"testing"
)

// 滑动窗口的最大值

// 标准解法：双端队列

func maxSlidingWindow(nums []int, k int) []int {
	window := make([]int, 0)
	res := make([]int, 0)
	for i, x := range nums {
		if i >= k && window[0] <= i-k { // 窗口启动
			window = window[1:]
		}

		// 如果进来的数字大于队尾元素，则移除队尾
		// 循环执行
		for len(window) > 0 && nums[window[len(window)-1]] <= x {
			window = window[:len(window)-1]
		}

		window = append(window, i)

		if i >= k-1 {
			res = append(res, nums[window[0]])
		}

	}
	return res
}

func Test_maxSlidingWindow(t *testing.T) {
	nums := []int{1, 3, -1, -3, 5, 3, 6}
	k := 3
	res := maxSlidingWindow(nums, k)
	fmt.Println(res)
}

// [3 3 5 5 6]
