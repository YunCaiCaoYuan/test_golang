package main

import (
	"context"
	"fmt"
)

func main()  {
	ctx := context.Background()
	fmt.Printf("ctx=%T\n", ctx)
	fmt.Println("ctx", ctx)
}
