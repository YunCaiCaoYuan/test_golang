package my_pow

import (
	"fmt"
	"testing"
)

func myPow(x float64, n int) float64 {
	if n == 0 {
		fmt.Println("x0:", x, "n:", n)
		return 1
	}
	if n < 0 {
		return 1 / myPow(x, -n)
	}
	if n%2 == 1 {
		fmt.Println("x1:", x, "n:", n)
		return x * myPow(x, n-1)
	}
	fmt.Println("x2:", x, "n:", n)
	return myPow(x*x, n/2)
}

func Test_myPow1(t *testing.T) {
	fmt.Println(myPow(2, 10))
}
func Test_myPow2(t *testing.T) {
	fmt.Println(myPow(2.1, 3))
	// 9.261
	//	9.261000000000001
}
func Test_myPow3(t *testing.T) {
	fmt.Println(myPow(2, -2))
}

// 非递归
func myPowV2(x float64, n int) float64 {
	if n < 0 {
		x = 1 / x
		n = -n
	}
	var pow float64 = 1
	for n > 0 {
		if n&1 > 0 {
			fmt.Println("内部，x:", x, "n:", n, "pow:", pow)
			pow *= x
		}
		fmt.Println("x:", x, "n:", n)
		x *= x
		n >>= 1
	}
	return pow
}

func Test_myPowV21(t *testing.T) {
	fmt.Println(myPowV2(2, 10))
}
func Test_myPowV22(t *testing.T) {
	fmt.Println(myPowV2(2.1, 3))
}
func Test_myPowV23(t *testing.T) {
	fmt.Println(myPowV2(2, -2))
}
