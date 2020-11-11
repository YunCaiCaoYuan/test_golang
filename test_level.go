package main

import "fmt"

const (
	LevelRange0 = 0
	LevelRange1 = 10
	LevelRange2 = 100
	LevelRange3 = 400
)

// 前面的三个等差阶段的等级对应的基数等级
var baseLevelMap = map[int64]int64{
	LevelRange1: 0,
	LevelRange2: 4,
	LevelRange3: 9,
}
var maxLevelMap = map[int64]int64{
	LevelRange1: 4,
	LevelRange2: 9,
	LevelRange3: 19,
}

// 前面三个等差数列的公差
var diffMap = map[int64]int64{
	LevelRange1: 2,
	LevelRange2: 12,
	LevelRange3: 22,
}

var levelRanges = []int64{
	LevelRange0,
	LevelRange1,
	LevelRange2,
	LevelRange3,
}

func GetLevel(point int64) int64 {
	// 积分大于80240 是100级
	if point >= 80240 {
		return 100
	}
	var level int64
	// 积分小于400 按照每阶段的等差数列计算
	if point < LevelRange3 {
		for i := 1; i < len(levelRanges); i++ {
			temp := levelRanges[i]
			if point < temp {
				level = baseLevelMap[temp] + (point-levelRanges[i-1])/diffMap[temp] + 1
				if level > maxLevelMap[temp] {
					return maxLevelMap[temp]
				} else {
					return level
				}
			}
		}
	}
	// 积分大于400 按照公式算分
	freq := 1.0
	base := 400.0
	for {
		base = (48 * freq) + 2.0 + base
		freq += 0.5
		if point < int64(base) {
			return 20 + level
		}
		level++
	}
	return level
}

func main() {
    //var point int64

    level := GetLevel(10)
    fmt.Println("level=", level)
}


