package main

import (
	"crypto/md5"
	"fmt"
	"reflect"
)

func main() {
	mn := md5.New()

	mn2 := md5.New()

	if reflect.DeepEqual(&mn, &mn2) {
		fmt.Println("same")
	}
}
