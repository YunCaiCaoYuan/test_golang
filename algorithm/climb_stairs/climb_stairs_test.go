package climb_stairs

// 解法1:递归
// 时间复杂度:O(2n)
// 空间复杂度:O(n)
func climb_stairs(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return climb_stairs(n-1) + climb_stairs(n-2)
}

// 解法1+:记忆化递归
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func climb_stairs_plus(n int) int {
	memo := make([]int, n+1)
	return _climb_stairs_plus(n, memo)
}

func _climb_stairs_plus(n int, memo []int) int {
	if memo[n] > 0 {
		return memo[n]
	}
	if n == 1 {
		memo[n] = 1
	} else if n == 2 {
		memo[n] = 2
	} else {
		memo[n] = _climb_stairs_plus(n-1, memo) + _climb_stairs_plus(n-2, memo)
	}
	return memo[n]
}

// 解法2：动态规划
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func climb_stairs_v2(n int) int {
	if n == 1 {
		return 1
	}
	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2
	for i := 3; i < n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

// 解法2+：动态规划+优化存储
func climb_stairs_v2_plus(n int) int {
	if n == 1 {
		return 1
	}
	first := 1
	second := 2
	for i := 3; i <= n; i++ {
		thirdTmp := first + second
		first = second
		second = thirdTmp
	}
	return second
}
