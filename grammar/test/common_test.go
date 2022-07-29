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

func Test_int_format(t *testing.T) {
	a := 1
	fmt.Printf("%8d\n", a)
}

// 空格美化
func Test_blank(t *testing.T) {
	fmt.Println("你好\u30001")
	fmt.Println("你\u3000\u30001")
	fmt.Println("\u3000\u3000\u30001")
}

func Test_float(t *testing.T) {
	var a int32
	a = 1
	var b int32
	b = 2
	var c float32
	c = float32(a / b) // 0 错误
	fmt.Println(c)

	c = float32(a) / float32(b) // 0.5 正确
	fmt.Println(c)
}
