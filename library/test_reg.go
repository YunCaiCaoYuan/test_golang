package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	//fmt.Println("同号:")
	re := regexp.MustCompile(`([\\d])\\1{2,}`)
	//fmt.Println(re.MatchString("888888"))
	//fmt.Println(re.MatchString("231568"))

	fmt.Println("特殊数字重复:")
	re = regexp.MustCompile(`(520){2,}`)
	//str := string(95205201)
	str := strconv.Itoa(95205201) //数字变成字符串
	fmt.Println("str=", str)
	fmt.Println(re.MatchString(str))
	fmt.Println(re.MatchString(strconv.Itoa(95201234)))

	//fmt.Println("含特殊数字:")
	//re = regexp.MustCompile(`(5201314)+`)
	//fmt.Println(re.MatchString("5201314x"))
	//fmt.Println(re.MatchString("x5201314"))
	//fmt.Println(re.MatchString("52091314"))

}
