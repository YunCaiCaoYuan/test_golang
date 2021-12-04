package main

import "fmt"

func echo(nums []int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			fmt.Println("echo...", n)
			out <- n
		}
		close(out)
		fmt.Println("close echo...")
	}()
	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			fmt.Println("sq...", n*n)
			out <- n * n
		}
		close(out)
		fmt.Println("close sq...")
	}()
	return out
}

func odd(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			if n%2 != 0 {
				fmt.Println("odd...", n)
				out <- n
			}
		}
		close(out)
		fmt.Println("close odd...")
	}()
	return out
}

func sum(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		var sum = 0
		for n := range in {
			sum += n
		}
		fmt.Println("sum...", sum)
		out <- sum
		close(out)
		fmt.Println("close sum...")
	}()
	return out
}

type EchoFunc func([]int) <-chan int
type PipeFunc func(<-chan int) <-chan int

func pipeline(nums []int, echo EchoFunc, pipeFns ...PipeFunc) <-chan int {
	ch := echo(nums)
	for i := range pipeFns {
		ch = pipeFns[i](ch)
	}
	return ch
}

func main() {

	var nums = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for n := range pipeline(nums, echo, odd, sq, sum) {
		fmt.Println(n)
	}

	/*
		var nums = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		for n := range sum(sq(odd(echo(nums)))) {
			fmt.Println(n)
		}
	*/
}
