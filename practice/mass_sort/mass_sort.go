package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

const Path = "/Users/sunbin/Github/test_golang/practice/mass_sort/"

//100M空间，10G数据，排序

// 分割数据10G数据到100个文件里面
// 每个文件排好序，写入文件，分别取每个文件第一个元素，然后排序，取最小值，够100个，然后批量写入一个大文件末尾
// 依次执行上面的逻辑，直到结束

// 固定长度切割有问题，会把一个数据切成两部分
func massSort() {
	infile := "./big_file.txt"
	size := 32
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

		// 排序
		numStrL := strings.Split(string(buf[:n]), "\n")
		numL := make([]int, 0, len(numStrL))
		for _, item := range numStrL {
			num, _ := strconv.Atoi(item)
			if num != 0 {
				numL = append(numL, num)
			}
		}
		fmt.Println(numL)
		sort.Slice(numL, func(i, j int) bool {
			return numL[i] < numL[j]
		})

		newfile.Write(wBuf)
	}

}

func ReadAndSplitFile(filePath string) error {
	f, err := os.Open(filePath)
	defer f.Close()
	if err != nil {
		fmt.Println(err)
		return err
	}
	buf := bufio.NewReader(f)

	groupNum := 10
	groupList := make([]int, 0, groupNum)
	index := 0
	for {
		line, _, err := buf.ReadLine()
		if err == nil {
			lineStr := strings.TrimSpace(string(line))
			num, _ := strconv.Atoi(lineStr)
			//fmt.Println("num:", num, "err:", err)
			groupList = append(groupList, num)
		}
		if len(groupList) == groupNum || (err == io.EOF && len(groupList) > 0) {
			newFileName := "text" + strconv.Itoa(index)
			newfile, err := os.Create(Path + newFileName)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("create file", newFileName)
			}
			sort.Slice(groupList, func(i, j int) bool {
				return groupList[i] < groupList[j]
			})
			wList := make([]string, 0)
			for _, item := range groupList {
				wList = append(wList, strconv.Itoa(item))
			}
			wListStr := strings.Join(wList, "\n")
			//fmt.Println(wListStr)
			newfile.WriteString(wListStr)
			index++
			groupList = groupList[:0]
		}

		if err == io.EOF {
			break
		}
	}

	return nil
}

func SortAndMerge(fileNum int) {
	fileHandle := make([]*bufio.Reader, 0, fileNum)
	for i := 0; i < fileNum; i++ {
		fileName := "text" + strconv.Itoa(i)
		f, err := os.Open(Path + fileName)
		if err != nil {
			fmt.Println("err", err)
			return
		}
		buf := bufio.NewReader(f)
		fileHandle = append(fileHandle, buf)
	}

	fs, err := os.Create(Path + "/text_sorted")
	if err != nil {
		fmt.Println("err", err)
		return
	}
	//wBuf := bufio.NewWriter(fs)

	minList := make([]int, 0, fileNum)
	for _, fileH := range fileHandle {
		line, _, err := fileH.ReadLine()
		if err == nil {
			numStr := strings.TrimSpace(string(line))
			num, _ := strconv.Atoi(numStr)
			minList = append(minList, num)
		}
	}

	var count int
	for {
		if len(fileHandle) <= 0 {
			break
		}
		minNum, idx := min(minList)
		minNumStr := strconv.Itoa(minNum)
		fmt.Println(minNumStr)
		//wBuf.WriteString(minNumStr) todo ???
		fs.WriteString(minNumStr + "\n")
		line, _, err := fileHandle[idx].ReadLine()
		if err == nil {
			numStr := strings.TrimSpace(string(line))
			num, _ := strconv.Atoi(numStr)
			minList[idx] = num
		} else {
			minList = deleteNum(minList, idx)
			fileHandle = deleteFile(fileHandle, idx)
		}

		count++
		if count > 100 {
			break
		}
	}

	// todo 关闭文件
}

func min(list []int) (int, int) {
	minN := math.MaxInt64
	index := 0
	for idx, num := range list {
		if num < minN {
			minN = num
			index = idx
		}
	}
	return minN, index
}

func deleteNum(minList []int, idx int) []int {
	//fmt.Println("idx:", idx, "minList:", minList)
	return append(minList[0:idx], minList[idx+1:]...)
}

func deleteFile(handList []*bufio.Reader, idx int) []*bufio.Reader {
	//fmt.Println("idx:", idx)
	return append(handList[0:idx], handList[idx+1:]...)
}

func main() {
	ReadAndSplitFile(Path + "big_file.txt")
	SortAndMerge(10)
}
