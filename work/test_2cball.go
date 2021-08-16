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

	ACVal := len(killRedMap1) - (len(rList) - 1)
	killRedMap4 := make(map[int]bool)
	for i := 0; i < len(rList); i++ {
		killRed := rList[i] - ACVal
		if killRed < 0 {
			killRed = 33 + killRed
		}
		if !killRedMap4[killRed] {
			killRedMap4[killRed] = true
		}
	}
	fmt.Println("killRedMap4=", killRedMap4)

	killRedMap5 := make(map[int]bool)
	total := 0
	for i := 0; i < len(rList); i++ {
		total += rList[i]
	}
	divSum := calcDivSum(total)
	for i := 0; i < len(rList); i++ {
		killRed := rList[i] - divSum
		if killRed < 0 {
			killRed = 33 + killRed
		}
		if !killRedMap5[killRed] {
			killRedMap5[killRed] = true
		}
	}
	fmt.Println("killRedMap5=", killRedMap5)

	killRedMap6 := make(map[int]bool)
	tailTotal := 0
	for i := 0; i < len(rList); i++ {
		tailTotal += rList[i] - (rList[i]/10)*10
	}
	tailSum := calcDivSum(tailTotal)
	for i := 0; i < len(rList); i++ {
		killRed := rList[i] - tailSum
		if killRed < 0 {
			killRed = 33 + killRed
		}
		if !killRedMap6[killRed] {
			killRedMap6[killRed] = true
		}
	}
	fmt.Println("killRedMap6=", killRedMap6)
}

func calcDivSum(total int) int {
	divSum := 0
	for {
		if total > 10 {
			shangVal := total / 10
			divSum += total - shangVal*10
			total = shangVal
		} else {
			divSum += total
			break
		}
	}
	return divSum
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
