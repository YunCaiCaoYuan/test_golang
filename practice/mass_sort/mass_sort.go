package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

//100M空间，10G数据，排序

// 分割数据10G数据到100个文件里面
// 每个文件排好序，写入文件，分别取每个文件第一个元素，然后排序，取最小值，够100个，然后批量写入一个大文件末尾
// 依次执行上面的逻辑，直到结束

func massSort() {
	infile := "./big_file"
	size := 1024 * 1024
	file, err := os.Open(infile)
	if err != nil {
		fmt.Println(err, "failed to open:", infile)
		return
	}
	defer file.Close()

	finfo, err := file.Stat()
	if err != nil {
		fmt.Println("get file info failed", file)
	}
	buf := make([]byte, size)
	num := (int(finfo.Size()) + size - 1) / size
	for i := 0; i < num; i++ {
		newFileName := finfo.Name() + strconv.Itoa(i)
		newfile, err := os.Create(newFileName)
		if err != nil {
			fmt.Println("failed to create file", newfile)
		} else {
			fmt.Println("create file:", newFileName)
		}
		n, err := file.Read(buf)
		if err != nil && err != io.EOF {
			fmt.Println(err, "failed to read from ", file)
			break
		}
		wBuf := buf[:n]
		newfile.Write(wBuf)
	}

}

func main() {
	massSort()
}
