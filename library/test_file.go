package main

import (
	"bufio" //缓存IO
	"fmt"
	"io"
	"io/ioutil" //io 工具包
	"os"
)

func check(e error) {
	if e != nil {
		fmt.Println("e=", e)
		panic(e)
	}
}

/**
 * 判断文件是否存在  存在返回 true 不存在返回false
 */
func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

func main() {
	var wireteString = "测试"
	var f *os.File

	/***************************** 第一种方式: 使用 io.WriteString 写入文件 ***********************************************/

	var filename = "./output1.txt"
	var err1 error
	if checkFileIsExist(filename) { //如果文件存在
		f, err1 = os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0666) //打开文件
		fmt.Println("文件存在")
	} else {
		f, err1 = os.Create(filename) //创建文件
		fmt.Println("文件不存在")
	}
	check(err1)
	n, err1 := io.WriteString(f, wireteString) //写入文件(字符串)
	check(err1)
	fmt.Printf(" 写入%d个字节", n)

	/*****************************  第二种方式: 使用 ioutil.WriteFile 写入文件 ***********************************************/
	var d1 = []byte(wireteString)
	err2 := ioutil.WriteFile("./output2.txt", d1, 0666) //写入文件(字节数组)
	check(err2)

	/*****************************  第三种方式:  使用 File(Write,WriteString) 写入文件 ***********************************************/
	f, err3 := os.Create("./output3.txt") //创建文件
	check(err3)
	defer f.Close()
	n2, err3 := f.Write(d1) //写入文件(字节数组)
	check(err3)
	fmt.Printf(" 写入%d个字节", n2)
	n3, err3 := f.WriteString("writesn") //写入文件(字节数组)
	fmt.Printf(" 写入%d个字节", n3)
	f.Sync()

	/***************************** 第四种方式:  使用 bufio.NewWriter 写入文件 ***********************************************/
	w := bufio.NewWriter(f) //创建新的 Writer 对象
	n4, err3 := w.WriteString("bufferedn")
	fmt.Printf(" 写入%d个字节\n", n4)
	w.Flush()
	f.Close()
}
