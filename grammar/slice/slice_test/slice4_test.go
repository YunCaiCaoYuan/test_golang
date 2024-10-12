package slice

// 切片 不是并发安全的

import (
	"fmt"
	"reflect"
	"sync"
	"testing"
)

func concurrentAppendSliceNotForceIndex() {
	sl := make([]int, 0)
	wg := sync.WaitGroup{}
	for index := 0; index < 100; index++ {
		k := index
		wg.Add(1)
		go func(num int) {
			sl = append(sl, num)
			wg.Done()
		}(k)
	}
	wg.Wait()
	fmt.Println(sl)
	fmt.Printf("final len(sl)=%d cap(sl)=%d\n", len(sl), cap(sl))
}

/*
func main() {
	concurrentAppendSliceNotForceIndex()
	//第一次运行代码后，输出：[2 0 1 5 6 7 8 9 10 4 17 11 12 13 14 15 16 21 18 19 20 23 22 24 25 26 39 27 28 29 30 31 35 55 54 56 57 58 59 60 61 62 64 63 65 66 67 68 69 70 71 72 73 74 75 76 77 78 79 80 81 82 83 84 86 91 92 93 94 96 95 97 98 99]
	//final len(sl)=74 cap(sl)=128
	//第二次运行代码后，输出：省略切片元素输出... final len(sl)=81 cap(sl)=128
	//第二次运行代码后，输出：省略切片元素输出... final len(sl)=77 cap(sl)=128
}
*/

func Test_concurrentAppendSliceNotForceIndex(t *testing.T) {
	concurrentAppendSliceNotForceIndex()
}

func Test_SlicePointerUse(t *testing.T) {
	list := make([]int64, 0)
	fmt.Println("list1:", list)
	slicePointerUse(&list)
	fmt.Println("list2:", list)
}
func slicePointerUse(list *[]int64) {
	// 在函数内部，通过解引用指针来访问和操作切片
	int64s := *list
	int64s = append(int64s, 1)
	fmt.Println("int64s:", int64s)
	fmt.Println("lista:", *list)
}

func Test_SlicePointerUse2(t *testing.T) {
	list := make([]int64, 3)
	fmt.Println("list1:", list)
	slicePointerUse2(&list)
	fmt.Println("list2:", list)
}
func slicePointerUse2(list *[]int64) {
	int64s := *list
	int64s[0] = 1
	//fmt.Println("int64s:", int64s)
	//fmt.Println("lista:", *list)
}

//list1: [0 0 0]
//list2: [1 0 0]

// gorm scan是如何修改切片的？
// 反射
func Test_SlicePointerUse3(t *testing.T) {
	list := make([]int64, 0)
	fmt.Println("list1:", list)
	slicePointerUse3(&list)
	fmt.Println("list2:", list)
}
func slicePointerUse3(list *[]int64) {
	//fmt.Println("TypeOf list:", reflect.TypeOf(list))
	value := reflect.ValueOf(list)
	v := value.Elem()
	fmt.Println("v:", v)
	v.Set(reflect.Append(v, reflect.ValueOf(int64(4))))
}
