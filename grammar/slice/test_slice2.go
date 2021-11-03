package main

import (
	"fmt"
	"reflect"
)

func GetSlice(src, dest interface{}, length int) error {
	srcRv := reflect.ValueOf(src)
	if srcRv.Kind() != reflect.Slice {
		return fmt.Errorf("src 参数必须是切片")
	}
	srcElemType := srcRv.Type()

	destPrv := reflect.ValueOf(dest)
	if destPrv.Kind() != reflect.Ptr || destPrv.IsNil() {
		return fmt.Errorf("dest 必须是一个指向切片的指针")
	}
	destRv := destPrv.Elem()
	if destRv.Kind() != reflect.Slice {
		return fmt.Errorf("dest 必须是一个指向切片的指针")
	}
	if length < 0 {
		return fmt.Errorf("l 必须大于等于0")
	}

	destElemType := destRv.Type()

	if srcElemType != destElemType {
		return fmt.Errorf("src 和 dest 切片类型不相同")
	}

	l := srcRv.Len()
	if l > length {
		l = length
	}
	destRv.Set(srcRv.Slice(0, l))

	return nil
}


func main() {
	var src = []int{1,2,3,4}
	dest := make([]int, 0)
	err := GetSlice(src, &dest, 2)
	if err != nil {
		fmt.Println("err=", err)
	}
	src[0]=10

	fmt.Printf("%+v\n", dest)
}
