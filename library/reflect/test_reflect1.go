package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string
	Age  int
}

func main() {
	var stu Student

	typeOfStu := reflect.TypeOf(stu)

	fmt.Println(typeOfStu.Name(), typeOfStu.Kind())
}
