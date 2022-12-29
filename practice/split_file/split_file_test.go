package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
	"testing"
)

func Test_createBigFile(t *testing.T) {
	file, err := os.Create("big_file.txt")
	if err != nil {
		fmt.Println(err)
	}

	intMap := make(map[int64]struct{})
	for i := 1; i <= 10*1000*1000; i++ {
		intMap[rand.Int63n(math.MaxInt32)] = struct{}{}
	}
	fmt.Println(len(intMap))

	for num, _ := range intMap {
		_, err := file.Write([]byte(strconv.Itoa(int(num)) + "\n"))
		if err != nil {
			fmt.Println(err)
		}
	}
}
