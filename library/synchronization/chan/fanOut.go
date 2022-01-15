package main

import (
	"fmt"
	"reflect"
)

func fanOut(ch <-chan interface{}, out []chan interface{}, async bool) {
	go func() {
		defer func() { //退出时关闭所有的输出chan
			for i := 0; i < len(out); i++ {
				fmt.Println("fanOut close...")
				close(out[i])
			}
		}()

		for v := range ch { // 从输入chan中读取数据
			v := v
			for i := 0; i < len(out); i++ {
				i := i
				if async { //异步
					go func() {
						out[i] <- v // 放入到输出chan中,异步方式
					}()
				} else {
					out[i] <- v // 放入到输出chan中，同步方式
				}
			}
		}
	}()
}


func main() {
	//start := time.Now()

	in := make(chan interface{})
	var outs []chan interface{}
	for i := 1; i<= 10; i++ {
		outs = append(outs, make(chan interface{}))
	}

	var cases []reflect.SelectCase
	for _, c := range outs {
		cases = append(cases, reflect.SelectCase{
			Dir: reflect.SelectRecv,
			Chan: reflect.ValueOf(c),
		})
	}

	// 被观察者
	go func() {
		for {
			in <- 1
		}
	}()

	// 观察者
	go fanOut(in, outs, true)

	for len(cases) > 0 {
		i, v, ok := reflect.Select(cases)
		if !ok {
			fmt.Println("channel close!")
			cases = append(cases[:i], cases[i+1:]...)
			continue
		}
		fmt.Println("i:", i, "v:", v)
	}

	//fmt.Printf("done after %v", time.Since(start))
}
