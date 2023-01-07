package max_profit

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

// 贪心算法: 在每一步总是做出在当前看来最好的选择
func maxProfitV3(prices []int) int {
	leng := len(prices)
	if leng < 2 {
		return 0
	}

	res := 0
	for i := 1; i < leng; i++ {
		diff := prices[i] - prices[i-1]
		if diff > 0 {
			res += diff
		}
	}
	return res
}

// 7
func Test_maxProfitV3(t *testing.T) {
	prices := []int{7, 1, 5, 3, 6, 4}
	res := maxProfitV3(prices)
	fmt.Println(res)
}
func Test_maxProfitV32(t *testing.T) {
	prices := []int{7, 6}
	res := maxProfitV3(prices)
	fmt.Println(res)
}

//4
func Test_maxProfitV33(t *testing.T) {
	prices := []int{1, 2, 3, 4, 5}
	res := maxProfitV3(prices)
	fmt.Println(res)
}
func Test_maxProfitV34(t *testing.T) {
	t1 := time.Now()
	leng := 40
	prices := make([]int, 0, leng)
	for i := 0; i < leng; i++ {
		prices = append(prices, rand.Intn(1000))
	}
	res := maxProfitV3(prices)
	fmt.Println(time.Since(t1))
	fmt.Println(res)
}
