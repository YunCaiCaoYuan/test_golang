package max_sliding_window

// 滑动窗口的最大值

// 标准解法：双端队列

func maxSlidingWindow(nums []int, k int) []int {
	window := make([]int, 0)
	res := make([]int, 0)
	for index, val := range nums {
		if len(res) < k {
			res = append(res)
		}
	}
}
