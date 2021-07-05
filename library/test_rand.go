package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for i := 0; i<10; i++ {
		r := rand.Int()
		fmt.Println(r)
	}

}
