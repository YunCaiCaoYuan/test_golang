package main

import "fmt"
import "reflect"

type Bird struct {
	Name string
	LifeExpectance int
}

func (b * Bird) Fly() {
	fmt.Println("I am flying...")
}

func main() {
	sparrow := &Bird{"Sparrow", 3}
	S := reflect.ValueOf(sparrow).Elem()
	typeOfT := S.Type()
	for i := 0; i < S.NumField(); i++ {
		f := S.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i, typeOfT.Field(i).Name, f.Type(), f.Interface())
	}
//	0: Name string = Sparrow
//  1: LifeExpectance int = 3
}

