package sqrt

import (
	"fmt"
	"testing"
)

// 二分法
func sqrt(x int) int {
	if x == 0 || x == 1 {
		return x
	}

	l := 1
	r := x
	res := 0
	for l <= r {
		m := l + (r-l)/2
		if m == x/m {
			fmt.Println("1——", "m:", m, "l:", l, "r:", r)
			return m
		} else if m > x/m {
			fmt.Println("2——", "m:", m, "l:", l, "r:", r)
			r = m - 1
		} else {
			fmt.Println("3——", "m:", m, "l:", l, "r:", r)
			l = m + 1
			res = m
		}

	}
	return res
}

func Test_sqrt(t *testing.T) {
	res := sqrt(4)
	fmt.Println(res)
}

func Test_sqrt2(t *testing.T) {
	res := sqrt(8)
	fmt.Println(res)
}
