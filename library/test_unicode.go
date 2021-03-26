package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := "\u4eba\u5de5\u968f\u4fbf\u5403\u5403"
	fmt.Println(strconv.Quote(a))
}
