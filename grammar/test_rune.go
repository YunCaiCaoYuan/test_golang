package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// var chinese = "我是中国人， I am Chinese"
	// 输出，注意在golang中一个汉字占3个byte
	/*
	var chinese = "我"
	fmt.Println("chinese length", len(chinese))
	fmt.Println("chinese word length", len([]rune(chinese)))
	fmt.Println("chinese word length", utf8.RuneCountInString(chinese))*/

	char := "é"
	fmt.Println(len(char))    // 2
	fmt.Println(utf8.RuneCountInString(char))    // 1
	fmt.Println("cafe\u0301")    // café    // 法文的 cafe，实际上是两个 rune 的组合
}
