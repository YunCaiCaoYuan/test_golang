package max_profit

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
	"time"
)

type MaxProfit struct {
	res int
}

// DFS ? 回溯法 ?
func (this *MaxProfit) maxProfit(prices []int) int {
	leng := len(prices)
	if leng < 2 {
		return 0
	}
	this.res = 0
	this.dfs(prices, 0, leng, 0, this.res)
	return this.res
}
func (this *MaxProfit) dfs(prices []int, index, leng, status, profit int) {
	if index == leng {
		fmt.Println("出口，index:", index, "res", this.res, "profit", profit)
		this.res = int(math.Max(float64(this.res), float64(profit)))
		return
	}

	// 不操作：
	this.dfs(prices, index+1, leng, status, profit)

	// 买入/卖出：
	if status == 0 {
		fmt.Println("买入，index:", index, "res", profit-prices[index])
		this.dfs(prices, index+1, leng, 1, profit-prices[index])
	} else {
		fmt.Println("卖出，index:", index, "res", profit+prices[index])
		this.dfs(prices, index+1, leng, 0, profit+prices[index])
	}
}

func Test_maxProfit(t *testing.T) {
	prices := []int{7, 1, 5, 3, 6, 4}
	maxProfitH := new(MaxProfit)
	res := maxProfitH.maxProfit(prices)
	fmt.Println(res)
}
func Test_maxProfit2(t *testing.T) {
	prices := []int{7, 6}
	maxProfitH := new(MaxProfit)
	res := maxProfitH.maxProfit(prices)
	fmt.Println(res)
}
func Test_maxProfit3(t *testing.T) {
	prices := []int{1, 2, 3, 4, 5}
	maxProfitH := new(MaxProfit)
	res := maxProfitH.maxProfit(prices)
	fmt.Println(res)
}
func Test_maxProfit4(t *testing.T) {
	t1 := time.Now()
	leng := 40
	prices := make([]int, 0, leng)
	for i := 0; i < leng; i++ {
		prices = append(prices, rand.Intn(1000))
	}
	maxProfitH := new(MaxProfit)
	res := maxProfitH.maxProfit(prices)
	fmt.Println(time.Since(t1))
	fmt.Println(res)
}
