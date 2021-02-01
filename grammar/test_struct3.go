package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name  string
	Age   int64
	wight int64
	high  int64
	score int64
}

func main() {
	var stu1 = new(Student)
	fmt.Printf("%p\n", &stu1.Name)
	fmt.Printf("%p\n", &stu1.Age)
	fmt.Printf("%p\n", &stu1.wight)
	fmt.Printf("%p\n", &stu1.high)
	fmt.Printf("%p\n", &stu1.score)
	typ := reflect.TypeOf(Student{})
	fmt.Printf("Struct is %d bytes long\n", typ.Size())
	// We can run through the fields in the structure in order
	n := typ.NumField()
	for i := 0; i < n; i++ {
		field := typ.Field(i)
		fmt.Printf("%s at offset %v, size=%d, align=%d\n",
			field.Name, field.Offset, field.Type.Size(),
			field.Type.Align())
	}
}

//0xc000090000
//0xc000090010
//0xc000090018
//0xc000090020
//0xc000090028
//Struct is 48 bytes long
//Name at offset 0, size=16, align=8
//Age at offset 16, size=8, align=8
//wight at offset 24, size=8, align=8
//high at offset 32, size=8, align=8
//score at offset 40, size=8, align=8
