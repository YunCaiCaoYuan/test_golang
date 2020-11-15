package main

import  "fmt"

func test1(args ...string) { //可以接受任意个string参数
    for _, v:= range args{
        fmt.Println(v)
    }
}

func main(){
/*
var strss= []string{
        "qwr",
        "234",
        "yui",
        "cvbc",
    }
    test1(strss...) //切片被打散传入
*/
     var strss= []string{
            "qwr",
            "234",
            "yui",

        }
        var strss2= []string{
            "qqq",
            "aaa",
            "zzz",
            "zzz",
        }
    strss=append(strss,strss2...) //strss2的元素被打散一个个append进strss
    fmt.Println(strss)
}

