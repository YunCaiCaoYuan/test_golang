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
