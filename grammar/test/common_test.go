package test

import (
	"fmt"
	"testing"
)

func Test_plus(t *testing.T) {

	a := 1
	//++a 不支持前置++
	fmt.Println(a)
	a++
}

func Test_divide(t *testing.T) {
	a := 1
	b := 2
	fmt.Printf("%d\n", a/b)
	fmt.Printf("%0.2f\n", float32(a)/float32(b))
}
