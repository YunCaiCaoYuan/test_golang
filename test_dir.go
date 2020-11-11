package main

import "fmt"
import "os"

func main() {
    str,_ := os.Getwd()
    fmt.Println("str=", str)
    //fmt.Println("str.type=", str.type) // 错误
    fmt.Printf("%T\n", str)
    fmt.Println("new dir=", str + "/sunbin/test.txt")

}
