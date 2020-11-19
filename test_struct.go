package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	var p1 Person
	p1.name = "zhangsan"
	p1.age = 18
	fmt.Printf("This is %s, %d years old\n", p1.name, p1.age)

	p2 := new(Person)
	p2.name = "lisi"
	p2.age = 20
	(*p2).age = 23 //这种写法也是合法的
	fmt.Printf("This is %s, %d years old\n", p2.name, p2.age)

	p3 := Person{"wangwu", 25}
	fmt.Printf("This is %s, %d years old\n", p3.name, p3.age)
}
