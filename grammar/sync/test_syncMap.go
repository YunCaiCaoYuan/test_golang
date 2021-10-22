package main

import (
	"fmt"
	"sync"
)

type IntStrMap struct {
	m sync.Map
}

func (iMap *IntStrMap) Delete(key int) {
	iMap.m.Delete(key)
}

func (iMap *IntStrMap) Load(key int) (value string, ok bool) {
	v, ok := iMap.m.Load(key)
	if v != nil {
		value = v.(string)
	}
	return
}

func (iMap *IntStrMap) LoadOrStore(key int, value string) (actual string, loaded bool) {
	a, loaded := iMap.m.LoadOrStore(key, value)
	actual = a.(string)
	return
}

func (iMap *IntStrMap) Range(f func(key int, value string) bool) {
	f1 := func(key, value interface{}) bool {
		return f(key.(int), value.(string))
	}
	iMap.m.Range(f1)
}

func (iMap *IntStrMap) Store(key int, value string) {
	iMap.m.Store(key, value)
}

func main() {
	var isMap IntStrMap
	val, ok := isMap.Load(123)
	if ok {
		fmt.Println("val=", val)
	} else {
		isMap.Store(123, "123")
	}

	isMap.Range(func(key int, value string) bool {
		fmt.Println("key=", key, "value=", value)
		return true
	})
}
