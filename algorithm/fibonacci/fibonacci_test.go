package fibonacci

import (
	"fmt"
	"testing"
)

// 时间复杂度：O( n)
//            2
func fib(n int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

// 55
func Test_fib(t *testing.T) {
	ret := fib(10)
	fmt.Println(ret)
}

// 保存中间结果
func fibV2(n int, memory map[int]int) int {
	if n == 0 || n == 1 {
		return n
	}
	val, ok := memory[n]
	if !ok {
		memory[n] = fibV2(n-1, memory) + fibV2(n-2, memory)
		return memory[n]
	}
	return val
}

//55
func Test_fibV2(t *testing.T) {
	memory := make(map[int]int, 0)
	ret := fibV2(10, memory)
	fmt.Println(ret)
}

// 递推（动态规划）
// 时间复杂度：O(n)
func fibV3(n int) int {
	F := make([]int, n+1)
	F[0] = 0
	F[1] = 1
	for i := 2; i <= n; i++ {
		F[i] = F[i-1] + F[i-2]
	}
	return F[n]
}
func Test_fibV3(t *testing.T) {
	ret := fibV3(10)
	fmt.Println(ret)
}
