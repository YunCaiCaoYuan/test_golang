package max_profit

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
	"time"
)

// 动态规划
func maxProfitV2(prices []int) int {
	leng := len(prices)
	if leng < 2 {
		return 0
	}

	dp := make([][]int, leng)
	for i := 0; i < leng; i++ {
		dp[i] = make([]int, 2)
	}

	// dp[i][j]表示下标为i的这一天，持股状态为j时，手上拥有的最大现金数
	dp[0][0] = 0
	dp[0][1] = -prices[0]

	// 0 持有现金
	// 1 持有股票
	for i := 1; i < leng; i++ {
		// 状态转移方程
		dp[i][0] = int(math.Max(float64(dp[i-1][0]), float64(dp[i-1][1]+prices[i]))) // 不操作 | 卖出
		dp[i][1] = int(math.Max(float64(dp[i-1][1]), float64(dp[i-1][0]-prices[i]))) // 不操作 | 买入
	}
	printArr2Row(dp)
	return dp[leng-1][0]
}

func printArr2(list [][]int) {
	fmt.Println("len(list):", len(list))
	for i := 0; i < len(list); i++ {
		for j := 0; j < len(list[i]); j++ {
			fmt.Printf("%d ", list[i][j])
		}
		fmt.Println("")
	}
}

func printArr2Row(list [][]int) {
	for i := 0; i < 2; i++ {
		for j := 0; j < len(list); j++ {
			fmt.Printf("%3d ", list[j][i])
		}
		fmt.Println("")
	}
}

// 7
func Test_maxProfitV2(t *testing.T) {
	prices := []int{7, 1, 5, 3, 6, 4}
	res := maxProfitV2(prices)
	fmt.Println(res)
}
func Test_maxProfitV22(t *testing.T) {
	prices := []int{7, 6}
	res := maxProfitV2(prices)
	fmt.Println(res)
}

//4
func Test_maxProfitV23(t *testing.T) {
	prices := []int{1, 2, 3, 4, 5}
	res := maxProfitV2(prices)
	fmt.Println(res)
}
func Test_maxProfitV24(t *testing.T) {
	t1 := time.Now()
	leng := 40
	prices := make([]int, 0, leng)
	for i := 0; i < leng; i++ {
		prices = append(prices, rand.Intn(1000))
	}
	res := maxProfitV2(prices)
	fmt.Println(time.Since(t1))
	fmt.Println(res)
}
