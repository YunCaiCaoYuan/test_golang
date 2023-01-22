package sqrt

import (
	"fmt"
	"math"
	"testing"
)

// 二分法
// 添加精度参数
func sqrtV2(x float64, precision float64) float64 {
	if x == 0 || x == 1 {
		return x
	}

	l := 1.0
	r := x
	res := 0.0
	for l <= r {
		m := l + (r-l)/2

		if math.Abs(m-x/m) <= precision {
			fmt.Println("1——", "m:", m, "l:", l, "r:", r)
			return m
		} else if m > x/m {
			fmt.Println("2——", "m:", m, "l:", l, "r:", r)
			r = m
		} else {
			fmt.Println("3——", "m:", m, "l:", l, "r:", r)
			l = m
		}

		//if m == x/m {
		//	return m
		//} else if m > x/m {
		//	r = m - 1
		//} else {
		//	l = m + 1
		//	res = m
		//}
		//if math.Abs(m-x/m) <= precision {
		//	return res
		//}
	}
	return res
}

func Test_sqrtV2(t *testing.T) {
	res := sqrtV2(4, 1e-5)
	fmt.Println(res)
}

func Test_sqrtV2_2(t *testing.T) {
	res := sqrtV2(8, 1e-3)
	fmt.Println(res)
}
