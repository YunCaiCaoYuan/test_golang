package slice

// 切片 不是并发安全的

import (
	"fmt"
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
