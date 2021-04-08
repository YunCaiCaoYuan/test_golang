package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// var chinese = "我是中国人， I am Chinese"
	// 输出，注意在golang中一个汉字占3个byte
	var chinese = "我"
	fmt.Println("chinese length", len(chinese))
	fmt.Println("chinese word length", len([]rune(chinese)))
	fmt.Println("chinese word length", utf8.RuneCountInString(chinese))
}
