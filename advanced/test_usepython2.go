package main

// 成功！！！

import (
	"fmt"
	"github.com/sbinet/go-python"
	"os"
)



func init()  {
	err := python.Initialize()
	if err != nil {
		panic(err.Error())
	}
}

func main()  {
	// 例子1:
	/*
	rc := python.Py_Main(os.Args)
	os.Exit(rc)
	*/

	// 例子2:打印路径
	/*
	m := python.PyImport_ImportModule("sys")
	if m == nil {
		fmt.Println("import error")
		return
	}
	path := m.GetAttrString("path")
	if path == nil {
		fmt.Println("get path error")
		return
	}
	size := python.PyList_GET_SIZE(path)
	fmt.Println("size=",size)
	for i := 0; i < size; i++ {
		item := python.PyList_GET_ITEM(path, i)
		s := python.PyString_AsString(item)
		fmt.Println(s)
	}
	 */

	// 例子3:加入当前路径
	/*
	m := python.PyImport_ImportModule("sys")
	if m == nil {
		fmt.Println("import error")
		return
	}
	path := m.GetAttrString("path")
	if path == nil {
		fmt.Println("get path error")
		return
	}
	// 加入当前目录，空串表示当前目录
	currentDir := python.PyString_FromString("")
	fmt.Println("currentDir=",currentDir)
	err := python.PyList_Insert(path, 0, currentDir)
	fmt.Println("err=",err)
	size := python.PyList_GET_SIZE(path)
	fmt.Println("size=",size)
	for i := 0; i < size; i++ {
		item := python.PyList_GET_ITEM(path, i)
		s := python.PyString_AsString(item)
		fmt.Println(s)
	}
	 */

	// 例子4:
	m := python.PyImport_ImportModule("sys")
	if m == nil {
		fmt.Println("import error1")
		return
	}
	path := m.GetAttrString("path")
	if path == nil {
		fmt.Println("get path error")
		return
	}
	// 加入当前目录，空串表示当前目录
	currentDir := python.PyString_FromString("")
	python.PyList_Insert(path, 0, currentDir)

	//dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	// 1、改路径：代码写相对路径/绝对路径
	// 2、编辑器，加环境变量

	//wd, _ := os.Getwd()
	//fmt.Println("exec dir=", wd)

	dir := os.Getenv("DIR")
	fmt.Println(dir)
	if dir != ""{
		os.Chdir(dir)
	} else {
		os.Chdir("/Users/sunbin/study/test_golang/advanced")
	}

	m = python.PyImport_ImportModule("reqbaidu")
	if m == nil {
		fmt.Println("import error2")
		return
	}
	touchBaidu := m.GetAttrString("touch_baidu")
	if touchBaidu == nil {
		fmt.Println("get touch_baidu error")
		return
	}
	res := touchBaidu.CallFunction()
	if res == nil {
		fmt.Println("call touch_baidu error")
		return
	}
	statusCode := res.GetAttrString("status_code")
	content := res.GetAttrString("content")
	fmt.Println(python.PyInt_AS_LONG(statusCode))
	fmt.Println(python.PyString_AS_STRING(content))

}
