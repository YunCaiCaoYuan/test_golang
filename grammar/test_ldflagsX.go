package main

import "fmt"

var Version = "No Version Provided"

func main() {
	fmt.Println("HelloWorld Version is:", Version)
}

//go run -ldflags "-X main.Version=1.5" helloworld.go

//HelloWorld Version is: 1.5
