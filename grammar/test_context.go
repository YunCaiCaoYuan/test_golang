package main

import (
	"context"
	"fmt"
)

type ctxOutgoingKey struct{}

func FromOutgoingContext(ctx context.Context) (md map[string]string, ok bool) {
	md, ok = ctx.Value(ctxOutgoingKey{}).(map[string]string)
	return
}

// NewOutgoingContext creates a new context with outgoing md attached.
func NewOutgoingContext(ctx context.Context, ctxmap map[string]string) context.Context {
	return context.WithValue(ctx, ctxOutgoingKey{}, ctxmap)
}

func main() {
	//var ctxMap1 map[string]string  // panic: assignment to entry in nil map
	ctxMap1 := make(map[string]string)
	ctxMap1["123"] = "123"
	ctx1 := NewOutgoingContext(context.TODO(), ctxMap1)

	fromCtxMap1, _ := FromOutgoingContext(ctx1)
	fmt.Println(fromCtxMap1)

	/*
	ctx := context.TODO()
	ctx2 := context.WithValue(ctx, "key1", "value1")
	if val, ok := ctx2.Value("key1").(string); ok {
		fmt.Println("val:", val)
	}
	*/
}
