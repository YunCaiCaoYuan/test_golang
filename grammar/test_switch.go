package main

import (
	"fmt"
)

func main() {
	var typ int
	fmt.Scan(&typ)

	switch typ {
	case 1:
		fmt.Println("1...")
	case 2:
		fmt.Println("2...")
	}
	// 匹配不到的，并不报错
}
