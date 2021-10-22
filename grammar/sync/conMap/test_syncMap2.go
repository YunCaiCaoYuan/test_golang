package conMap

import (
	"fmt"
	"reflect"
	"sync"
)

type ConcurrentMap struct {
	m         sync.Map
	KeyType   reflect.Type
	ValueType reflect.Type
}

func (cMap *ConcurrentMap) Load(key interface{}) (value interface{}, ok bool) {
	if reflect.TypeOf(key) != cMap.KeyType {
		return
	}
	return cMap.m.Load(key)
}

func (cMap *ConcurrentMap) Store(key, value interface{}) {
	if reflect.TypeOf(key) != cMap.KeyType {
		fmt.Printf("wrong key type: %v", reflect.TypeOf(key))
		return
	}
	if reflect.TypeOf(value) != cMap.ValueType {
		fmt.Printf("wrong value type: %v", reflect.TypeOf(value))
		return
	}
	cMap.m.Store(key, value)
}

func (cMap *ConcurrentMap) Delete(key interface{}) {
	if reflect.TypeOf(key) != cMap.KeyType {
		return
	}
	cMap.m.Delete(key)
}

func main() {
	var cMap = ConcurrentMap{
		KeyType: reflect.TypeOf(int32(123)),
		ValueType: reflect.TypeOf("123"),
	}
	val, ok := cMap.Load(123)
	if ok {
		fmt.Println("val=", val)
	} else {
		cMap.Store(int32(123), "123")
	}

	/*
	cMap.m.Range(func(key, value interface{}) bool {
		fmt.Println("key=", key.(int32), "val=", value.(string))
		return true
	})
	 */
}
