package main

import (
	"fmt"
	"sort"
)

func twoCBall(rList, rList1 sort.IntSlice, blue, blue1 int) {
	sort.Sort(rList)
	//fmt.Println(rList)
	//heartWaterList := make([]int, 0) // 心水
	heartWaterMap := make(map[int]bool)
	killRedMap1 := make(map[int]bool)
	for i := len(rList) - 1; i >= 0; i-- {
		for j := i - 1; j >= 0; j-- {
			heartWater := rList[i] - rList[j]
			//heartWaterList = append(heartWaterList, rList[i]-rList[j])
			//fmt.Println(rList[i], "-", rList[j], "=", rList[i]-rList[j])
			if !heartWaterMap[heartWater] {
				if heartWater < 16 {
					heartWaterMap[heartWater] = true
				}
				killRedMap1[heartWater] = true
			}

		}
	}
	//fmt.Println("heartWaterMap=", heartWaterMap)
	fmt.Println("killRedMap1=", killRedMap1)
	killRedMap2 := make(map[int]bool)
	for i := len(rList) - 1; i >= 0; i-- {
		killRed := rList[i] - rList1[i]
		if killRed < 0 {
			killRed = -killRed
		}
		if !killRedMap2[killRed] {
			killRedMap2[killRed] = true
		}
	}
	fmt.Println("killRedMap2=", killRedMap2)
	killRedMap3 := make(map[int]bool)
	for i := len(rList) - 1; i >= 0; i-- {
		killRed := rList[i] - blue1
		if killRed < 0 {
			killRed = -killRed
		}
		if !killRedMap3[killRed] {
			killRedMap3[killRed] = true
		}
	}
	fmt.Println("killRedMap3=", killRedMap3)
}

func main() {
	rList := [][]int{
		{1, 2, 7, 13, 23, 24},
		{5, 6, 12, 14, 27, 28},
		{4, 6, 16, 24, 26, 33},
	}
	bList := []int{
		12,
		8,
		16,
	}
	twoCBall(rList[0], rList[1], bList[0], bList[1])
}
