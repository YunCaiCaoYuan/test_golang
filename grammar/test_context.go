package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.TODO()

	ctx2 := context.WithValue(ctx, "key1", "value1")
	if val, ok := ctx2.Value("key1").(string); ok {
		fmt.Println("val:", val)
	}

}
