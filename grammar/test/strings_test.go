package test

import (
	"fmt"
	"strings"
	"testing"
)

func Test_trim(t *testing.T) {
	bindStr := "+8618780170596"
	bindReal := strings.Trim(bindStr, "+")
	fmt.Println(bindReal)
}

func Test_trim2(t *testing.T) {
	bindStr := "+18780170596"
	bindReal := strings.Trim(bindStr, "+")
	fmt.Println(bindReal)
}

func Test_trim3(t *testing.T) {
	bindStr := "18780170596"
	bindReal := strings.Trim(bindStr, "+")
	fmt.Println(bindReal)
}

func Test_string(t *testing.T) {
	intS := 123         // rune
	str := string(intS) // 按ascii值去转换了
	fmt.Println(str)
}

func Test_empty(t *testing.T) {
	s := ""
	ret := strings.Split(s, ",")
	fmt.Println("ret:", ret, "len(ret)", len(ret)) // ret: [] len(ret) 1
	for _, item := range ret {
		fmt.Println("item:", item)
	}
}

func Test_1(t *testing.T) {
	s := "1"
	ret := strings.Split(s, ",")
	fmt.Println("ret:", ret, "len(ret)", len(ret))
	for _, item := range ret {
		fmt.Println("item:", item)
	}
}

func Test_12(t *testing.T) {
	s := "1,2"
	ret := strings.Split(s, ",")
	fmt.Println("ret:", ret, "len(ret)", len(ret))
	for _, item := range ret {
		fmt.Println("item:", item)
	}
}
