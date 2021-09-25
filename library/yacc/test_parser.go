package main

import (
	"fmt"
	"go/parser"
	"go/token"
)

func main() {
	fset := token.NewFileSet() // 相对于fset的position
	src := `package foo

import (
	"fmt"
	"time"
)

func bar() {
	fmt.Println(time.Now())
}`

	// 解析src但在处理导入后停止。
	f, err := parser.ParseFile(fset, "", src, parser.ImportsOnly)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 从文件CAST打印导入。
	for _, s := range f.Imports {
		fmt.Println(s.Path.Value)
	}

}
