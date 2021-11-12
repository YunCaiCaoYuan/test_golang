package test

import (
	. "bou.ke/monkey"
	"fmt"
	"testing"
)

func IsEqual(a, b string) bool {
	fmt.Println("not monkey")
	fmt.Println("not monkey")
	fmt.Println("not monkey")
	fmt.Println("not monkey")
	fmt.Println("not monkey")
	fmt.Println("not monkey")
	fmt.Println("not monkey")
	fmt.Println("not monkey")
	fmt.Println("not monkey")
	fmt.Println("not monkey")
	fmt.Println("not monkey")
	fmt.Println("not monkey")
	return a == b
}

func TestIsEqual(t *testing.T) {
			guard := Patch(IsEqual, func(_, _ string) bool {
				fmt.Println("monkey...")
				return true
			})
			defer guard.Unpatch()
			ok := IsEqual("hello", "world")
			if ok {
				fmt.Println("ok")
			} else {
				fmt.Println("not ok")
			}
}

// 编译器内联掉了，把函数改复杂一些，就不会被内联了
//[root@sunbinaliyun tmp]# go test -v monkey_test.go
//=== RUN   TestIsEqual
//not monkey
//not ok
//--- PASS: TestIsEqual (0.00s)
//PASS
//ok  	command-line-arguments	0.002s

//[root@sunbinaliyun tmp]# go test -gcflags "all=-N -l" -v monkey_test.go
//=== RUN   TestIsEqual
//monkey...
//ok
//--- PASS: TestIsEqual (0.00s)
//PASS
//ok  	command-line-arguments	0.002s
