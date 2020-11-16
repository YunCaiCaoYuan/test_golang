package main

import (
        "fmt"
        "regexp"
        )

func main() {
    //str := "111111111"
    //ret := str.match('([//d])\\1{2,}')
    //fmt.Println("ret", ret)

    //MatchString返回的第一个参数是bool类型即匹配结果，第二个参数是error类型
    //sourceStr := `my email is gerrylon@163.com`
    //matched, _ := regexp.MatchString(`[\w-]+@[\w]+(?:\.[\w]+)+`, sourceStr)
    //fmt.Printf("%v", matched) // true

    sourceStr := "231213"
    matched, _ := regexp.MatchString(`(0){5,}|(1){5,}|(1){5,}|(2){8,}`, sourceStr)
    fmt.Printf("%v", matched) // true
}