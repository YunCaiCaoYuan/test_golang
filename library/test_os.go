package main

import (
	"fmt"
	"os"
)

func main() {
	osFile := os.Getenv("TARS_SO")
	if osFile == "" {
		fmt.Println("TARS_SO not set")
	}
	fmt.Println(osFile)
}
