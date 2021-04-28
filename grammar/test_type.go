package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var v interface{}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 10; i++{
		v = i
		if (r.Intn(100) % 2) == 0 {
			v = "hello"
		}

		if _, ok := v.(int); ok {
			fmt.Printf("%d\n", v)
		}
	}
}

/*
type data struct {

	num     int
	fp      float32
	complex complex64
	str     string
	char    rune
	yes     bool
	events  <-chan string
	handler interface{}
	ref     *byte
	raw     [10]byte
}
*/
type data struct {
	num    int
	checks [10]func() bool        // 无法比较
	doIt   func() bool        // 无法比较
	m      map[string]string    // 无法比较
	bytes  []byte            // 无法比较
}
/*
func main() {
	v1 := data{}
	v2 := data{}
	fmt.Println("v1 == v2: ", v1 == v2)    // true
}
*/
/*
// 比较相等运算符无法比较的元素
func main() {
	v1 := data{}
	v2 := data{}
	fmt.Println("v1 == v2: ", reflect.DeepEqual(v1, v2))    // true

	m1 := map[string]string{"one": "a", "two": "b"}
	m2 := map[string]string{"two": "b", "one": "a"}
	fmt.Println("v1 == v2: ", reflect.DeepEqual(m1, m2))    // true

	s1 := []int{1, 2, 3}
	s2 := []int{1, 2}
	// 注意两个 slice 相等，值和顺序必须一致
	fmt.Println("v1 == v2: ", reflect.DeepEqual(s1, s2))    // true
}
 */
/*
func main() {
	var b1 []byte = nil
	b2 := []byte{}
	fmt.Println("b1 == b2: ", reflect.DeepEqual(b1, b2))    // false
}
 */
/*
func main() {
	var str = "one"
	var in interface{} = "one"
	fmt.Println("str == in: ", reflect.DeepEqual(str, in))    // true

	v1 := []string{"one", "two"}
	v2 := []string{"two", "one"}
	fmt.Println("v1 == v2: ", reflect.DeepEqual(v1, v2))    // false

	data := map[string]interface{}{
		"code":  200,
		"value": []string{"one", "two"},
	}
	encoded, _ := json.Marshal(data)
	var decoded map[string]interface{}
	json.Unmarshal(encoded, &decoded)
	fmt.Println("data == decoded: ", reflect.DeepEqual(data, decoded))    // false
}
 */
/*
func main() {
	var b1 []byte = nil
	b2 := []byte{}

	// b1 与 b2 长度相等、有相同的字节序
	// nil 与 slice 在字节上是相同的
	fmt.Println("b1 == b2: ", bytes.Equal(b1, b2))    // true
}
*/
/*
// 类型以字段形式直接嵌入
type myLocker struct {
	sync.Mutex
}

func main() {
	var locker myLocker
	locker.Lock()
	locker.Unlock()
}
 */
/*
type myLocker sync.Locker

func main() {
	var locker myLocker
	locker.Lock()
	locker.Unlock()
}
 */

// 定义 Mutex 的自定义类型
/*
type myMutex sync.Mutex

func main() {
	var mtx myMutex
	mtx.Lock()
	mtx.UnLock()
}
 */
