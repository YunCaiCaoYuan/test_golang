package test

import (
	"fmt"
	"testing"
)

type Key struct {
	appType uint32
	cmdId   uint16
}

// map key可以使用结构体? 可以
// why? 可以比较都可作为key type

//包括：
/*
boolean
numeric
string
pointer
channel
interface
structs、arrays
*/

func Test_struct_as_key(t *testing.T) {
	Map := make(map[Key]int, 0)
	key1 := Key{appType: 1, cmdId: 1}
	Map[key1] = 1
	key2 := Key{appType: 2, cmdId: 2}
	Map[key2] = 2
	fmt.Println(Map)

	tmpKey := Key{appType: 1, cmdId: 1}
	if v, ok := Map[tmpKey]; ok {
		fmt.Println("found :", v)
	}
}

// 结构体比较
func Test_struct_comparison(t *testing.T) {
	key1 := Key{1, 1}
	key2 := Key{1, 1}
	key3 := Key{3, 3}
	fmt.Println("key1 == key2", key1 == key2) // true
	fmt.Println("key1 == key3", key1 == key3) // false
}
