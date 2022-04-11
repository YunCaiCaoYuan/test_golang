package test

import (
	"fmt"
	"testing"
	"unicode/utf8"
)

func Test_RuneCountInString(t *testing.T) {
	var s string
	s = "123"
	fmt.Printf("%s len:%d\n", s, utf8.RuneCountInString(s))
	s = "你好"
	fmt.Printf("%s len:%d\n", s, utf8.RuneCountInString(s))
	s = "123你好"
	fmt.Printf("%s len:%d\n", s, utf8.RuneCountInString(s))
}

//123 len:3
//你好 len:2
//123你好 len:5
