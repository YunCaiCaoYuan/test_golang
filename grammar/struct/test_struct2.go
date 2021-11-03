package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (self *Person) init(name string, age int) {
	self.Name = name
	self.Age = age
}

func (self *Person) String() string {
	return self.Name

}

func main() {
	var person1 = new(Person)
	person1.init("wd", 22)
	//(*person1).init("wd",22)
	fmt.Println(person1) //&{wd 22}
}
