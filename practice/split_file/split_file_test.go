package main

import (
	"fmt"
	"os"
	"strconv"
	"testing"
)

func Test_createBigFile(t *testing.T) {
	file, err := os.Create("big_file.txt")
	if err != nil {
		fmt.Println(err)
	}
	for i := 1; i <= 10*1000*1000; i++ {
		_, err := file.Write([]byte(strconv.Itoa(i) + "\n"))
		if err != nil {
			fmt.Println(err)
		}
	}
}
