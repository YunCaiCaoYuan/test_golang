package main

import (
	"fmt"
	"github.com/speps/go-hashids"
)

//salt 盐值
const salt = "salt"

//Encode 混淆
func Encode(data int) string {
	hd := hashids.NewData()
	hd.Salt = salt
	h := hashids.NewWithData(hd)
	e, _ := h.Encode([]int{data})
	return e
}

//Decode 还原混淆
func Decode(data string) int {
	hd := hashids.NewData()
	hd.Salt = salt
	h := hashids.NewWithData(hd)
	e, _ := h.DecodeWithError(data)
	return e[0]
}

func main() {
	id := 769460962
	str := Encode(id)
	fmt.Println("Encode(id):", str)
	fmt.Println("Decode(id):", Decode(str))
}
