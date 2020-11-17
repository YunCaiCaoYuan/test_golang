package main

import (
	"fmt"
	//"regexp"
	//"strconv"
)

func IsGoodNum(num int64) bool {
	//tNum := num
	var arr = make([]int64, 0)
	for {
		if num >= 10 {
			i := num % 10
			num = num / 10
			arr = append(arr, i)
		} else {
			arr = append(arr, num)
			break
		}
	}
	var nums = make([]int64, 0)
	for i := len(arr) - 1; i >= 0; i-- {
		nums = append(nums, arr[i])
	}
	// 同号>=5位，如xxx66666
	if (nums[0] == nums[1] && nums[0] == nums[2] && nums[0] == nums[3] && nums[0] == nums[4]) ||
		(nums[1] == nums[2] && nums[1] == nums[3] && nums[1] == nums[4] && nums[1] == nums[5]) ||
		(nums[2] == nums[3] && nums[2] == nums[4] && nums[2] == nums[5] && nums[2] == nums[6]) ||
		(nums[3] == nums[4] && nums[3] == nums[5] && nums[3] == nums[6] && nums[3] == nums[7]) {
		return true
	}
	// 连号>=5位，如xxx12345
	if (nums[0] == nums[1]-1 && nums[0] == nums[2]-2 && nums[0] == nums[3]-3 && nums[0] == nums[4]-4) ||
		(nums[1] == nums[2]-1 && nums[1] == nums[3]-2 && nums[1] == nums[4]-3 && nums[1] == nums[5]-4) ||
		(nums[2] == nums[3]-1 && nums[2] == nums[4]-2 && nums[2] == nums[5]-3 && nums[2] == nums[6]-4) ||
		(nums[3] == nums[4]-1 && nums[3] == nums[5]-2 && nums[3] == nums[6]-3 && nums[3] == nums[7]-4) {
		return true
	}
	// 4连号+4同号，4同号+4连号，如12346666,22225432
	if (nums[0] == nums[1]-1 && nums[0] == nums[2]-2 && nums[0] == nums[3]-3 && nums[4] == nums[5] && nums[5] == nums[6] && nums[6] == nums[7]) ||
		(nums[0] == nums[1] && nums[0] == nums[2] && nums[0] == nums[3] && nums[4] == nums[5]-1 && nums[4] == nums[6]-2 && nums[4] == nums[7]-3) {
		return true
	}
	// 4连号+4连号，如12348765,12342345,54329876,87654567
	if (nums[0] == nums[1]-1 && nums[0] == nums[2]-2 && nums[0] == nums[3]-3 && nums[4] == nums[7]+3 && nums[5] == nums[7]+2 && nums[6] == nums[7]+1) ||
		(nums[0] == nums[1]-1 && nums[0] == nums[2]-2 && nums[0] == nums[3]-3 && nums[4] == nums[5]-1 && nums[4] == nums[6]-2 && nums[4] == nums[7]-3) ||
		(nums[0] == nums[1]+1 && nums[0] == nums[2]+2 && nums[0] == nums[3]+3 && nums[4] == nums[7]+3 && nums[5] == nums[7]+2 && nums[6] == nums[7]+1) ||
		(nums[0] == nums[1]+1 && nums[0] == nums[2]+2 && nums[0] == nums[3]+3 && nums[4] == nums[5]-1 && nums[4] == nums[6]-2 && nums[4] == nums[7]-3) {
		return true
	}
	// 4同号+4同号，如55558888
	if nums[0] == nums[1] && nums[0] == nums[2] && nums[0] == nums[3] && nums[4] == nums[5] && nums[4] == nums[6] && nums[4] == nums[7] {
		return true
	}
	// 3连号+4同号，如123x6666,321x6666,2222x543,2222x345
	if (nums[0] == nums[1]-1 && nums[0] == nums[2]-2 && nums[4] == nums[5] && nums[4] == nums[6] && nums[4] == nums[7]) ||
		(nums[0] == nums[1]+1 && nums[0] == nums[2]+2 && nums[4] == nums[5] && nums[4] == nums[6] && nums[4] == nums[7]) ||
		(nums[0] == nums[1] && nums[0] == nums[2] && nums[0] == nums[3] && nums[5] == nums[6]+1 && nums[5] == nums[7]+2) ||
		(nums[0] == nums[1] && nums[0] == nums[2] && nums[0] == nums[3] && nums[5] == nums[6]-1 && nums[5] == nums[7]-2) {
		return true
	}
	//4连号+3同号，如1234666x，8765333x，3456x111,6543x111
	if (nums[0] == nums[1]-1 && nums[0] == nums[2]-2 && nums[0] == nums[3]-3 && nums[4] == nums[5] && nums[4] == nums[6]) ||
		(nums[0] == nums[1]+1 && nums[0] == nums[2]+2 && nums[0] == nums[3]+3 && nums[4] == nums[5] && nums[4] == nums[6]) ||
		(nums[0] == nums[1]-1 && nums[0] == nums[2]-2 && nums[0] == nums[3]-3 && nums[5] == nums[6] && nums[5] == nums[7]) ||
		(nums[0] == nums[1]+1 && nums[0] == nums[2]+2 && nums[0] == nums[3]+3 && nums[5] == nums[6] && nums[5] == nums[7]) {
		return true
	}
	// 双同连号，如11223344,55667788
	if nums[0] == nums[1] && nums[2] == nums[3] && nums[4] == nums[5] && nums[6] == nums[7] && nums[0] == nums[2]-1 && nums[0] == nums[4]-2 && nums[0] == nums[6]-3 {
		return true
	}
	// 4连数列组，如12121212,89898989，46464646，10101010
	if nums[0] == nums[2] && nums[0] == nums[4] && nums[0] == nums[6] && nums[1] == nums[3] && nums[1] == nums[5] && nums[1] == nums[7] {
		return true
	}
	// 2连数列组，如15471547,96879687，10001000
	if nums[0] == nums[4] && nums[1] == nums[5] && nums[2] == nums[6] && nums[3] == nums[7] {
		return true
	}
	// 含特殊数字5201314，如5201314x，x5201314
	// 含特殊数字5211314，如5211314x，x5211314
	if nums[0] == 5 && nums[1] == 2 && nums[2] == 0 && nums[3] == 1 && nums[4] == 3 && nums[5] == 1 && nums[6] == 4 ||
		nums[1] == 5 && nums[2] == 2 && nums[3] == 0 && nums[4] == 1 && nums[5] == 3 && nums[6] == 1 && nums[7] == 4 ||
		nums[0] == 5 && nums[1] == 2 && nums[2] == 1 && nums[3] == 1 && nums[4] == 3 && nums[5] == 1 && nums[6] == 4 ||
		nums[1] == 5 && nums[2] == 2 && nums[3] == 1 && nums[4] == 1 && nums[5] == 3 && nums[6] == 1 && nums[7] == 4 {
		return true
	}
	//特殊数字520重复，如 520520xx, 520x520x, x520x520, x520520x, xx520520, 520xx520
	if nums[0] == 5 && nums[1] == 2 && nums[2] == 0 && nums[3] == 5 && nums[4] == 2 && nums[5] == 0 ||
		nums[0] == 5 && nums[1] == 2 && nums[2] == 0 && nums[4] == 5 && nums[5] == 2 && nums[6] == 0 ||
		nums[1] == 5 && nums[2] == 2 && nums[3] == 0 && nums[5] == 5 && nums[6] == 2 && nums[7] == 0 ||
		nums[1] == 5 && nums[2] == 2 && nums[3] == 0 && nums[4] == 5 && nums[5] == 2 && nums[6] == 0 ||
		nums[2] == 5 && nums[3] == 2 && nums[4] == 0 && nums[5] == 5 && nums[6] == 2 && nums[7] == 0 ||
		nums[0] == 5 && nums[1] == 2 && nums[2] == 0 && nums[5] == 5 && nums[6] == 2 && nums[7] == 0 {
		return true
	}
	//numStr := strconv.Itoa(int(tNum))
	//re := regexp.MustCompile(`(520[\d]?){2,}`)
	//if re.MatchString(numStr) {
	//	return true
	//}
	//特殊数字521重复，如521x521x，x521x521，xx521521
	if nums[0] == 5 && nums[1] == 2 && nums[2] == 1 && nums[3] == 5 && nums[4] == 2 && nums[5] == 1 ||
		nums[0] == 5 && nums[1] == 2 && nums[2] == 1 && nums[4] == 5 && nums[5] == 2 && nums[6] == 1 ||
		nums[1] == 5 && nums[2] == 2 && nums[3] == 1 && nums[5] == 5 && nums[6] == 2 && nums[7] == 1 ||
		nums[1] == 5 && nums[2] == 2 && nums[3] == 1 && nums[4] == 5 && nums[5] == 2 && nums[6] == 1 ||
		nums[2] == 5 && nums[3] == 2 && nums[4] == 1 && nums[5] == 5 && nums[6] == 2 && nums[7] == 1 ||
		nums[0] == 5 && nums[1] == 2 && nums[2] == 1 && nums[5] == 5 && nums[6] == 2 && nums[7] == 1 {
		return true
	}
	//re = regexp.MustCompile(`(521[\d]?){2,}`)
	//if re.MatchString(numStr) {
	//	return true
	//}
	return false
}

func main() {
	var numberList = [...]int64{
		// 连号>=5
		12366666, 97777799, 99999999, 88888811,
		// 顺号>=5
		12312345, 91234599, 12345678,
		// 4顺号+4连号
		12346666, 22225432, 22225432,
		// 4顺号+4顺号
		12348765, 12342345, 54329876, 87654567,
		// 4连号+4连号
		55558888, 99996666,
		// 3顺号+4连号
		12396666, 32196666, 22229543, 22229345,
		// 4顺号+3连号
		12346669, 87653339, 34569111, 65439111,
		//双连顺
		11223344, 55667788,
		//4连数列
		12121212, 89898989, 46464646, 10101010,
		//2连数列
		15471547, 96879687, 10001000,
		52013149, 95201314,
		52113149, 95211314,
		52052099, 52045205, 85209520, 15205201, 34520520, 52034520,
		52152199, 52145215, 85219521, 15215211, 34521521, 52134521}

	for _, num := range numberList {
		if IsGoodNum(num) {
			fmt.Printf("%d is good number\n", num)
		} else {
			fmt.Printf("%d not good number\n", num)
		}
	}
}
