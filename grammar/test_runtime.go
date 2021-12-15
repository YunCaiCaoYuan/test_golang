package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println(runtime.GOMAXPROCS(-1))    // 4
	fmt.Println(runtime.NumCPU())    		  // 4
	runtime.GOMAXPROCS(20)
	fmt.Println(runtime.GOMAXPROCS(-1))    // 20
	runtime.GOMAXPROCS(300)
	fmt.Println(runtime.GOMAXPROCS(-1))    // Go 1.9.2 // 300

	//exec.Command()
	//syscall.StartProcess()
}


