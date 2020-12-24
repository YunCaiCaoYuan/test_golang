package main

import (
	"fmt"
	"regexp"
)

func main() {
	fmt.Println("同号：")
	sourceStr := "99999"
	matched, _ := regexp.MatchString(`(0){5,}|(1){5,}|(1){5,}|(2){8,}|(9){4,}`, sourceStr)
	fmt.Printf("%v\n", matched)

	fmt.Println("连号：")
	/*
	   sourceStr = "01234567"
	   matched, _ = regexp.MatchString(`(?:0(?=1)|1(?=2)|2(?=3)|3(?=4)|4(?=5)|5(?=6)|6(?=7)|7(?=8)|8(?=9)){5,}\d`, sourceStr)
	   fmt.Printf("%v\n", matched)*/

	cityListReg := `(?:0(?=1)|1(?=2)|2(?=3)|3(?=4)|4(?=5)|5(?=6)|6(?=7)|7(?=8)|8(?=9)){5,}\\d`
	contents := "01234567"
	compile, _ := regexp.CompilePOSIX(cityListReg)
	submatch := compile.FindString(contents)
	fmt.Printf("%v\n", submatch)
}
