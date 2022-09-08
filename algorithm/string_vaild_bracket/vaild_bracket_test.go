package string_vaild_bracket

import (
	"fmt"
	"testing"
)

// 有效的括号
// [] {} ()
// 栈

var Stack []byte

func New() {
	Stack = make([]byte, 0)
}

func PushStack(c byte) {
	Stack = append(Stack, c)
}
func IsStackEmpty() bool {
	return len(Stack) == 0
}
func PopStack() {
	if len(Stack) == 1 {
		Stack = make([]byte, 0)
		return
	}
	Stack = Stack[:len(Stack)-1]
}
func TopStack() byte {
	return Stack[len(Stack)-1]
}

func ValidBracket(str []byte) bool {
	for _, c := range str {
		if c == '{' || c == '[' || c == '(' {
			PushStack(c)
		}
		if c == '}' || c == ']' || c == ')' {
			if !IsStackEmpty() {
				t := TopStack()
				if c == '}' && t == '{' || c == ']' && t == '[' || c == ')' && t == 'c' {
					PopStack()
					continue
				}
				return false // 右出现了，左边不匹配
			}
			return false //右出现了，空了
		}
	}
	return len(Stack) == 0 // 判断结束了,未消耗完
}

func Test_ValidBracket(t *testing.T) {
	New()
	str := []byte{'}', '{'}
	fmt.Println(str, " isValid: ", ValidBracket(str), Stack)

	New()
	str = []byte{'{', '}', '{'}
	fmt.Println(str, " isValid: ", ValidBracket(str), Stack)

	New()
	fmt.Println("Stack start", Stack)
	str = []byte{'{', '}'}
	fmt.Println(str, " isValid: ", ValidBracket(str), Stack)

}
