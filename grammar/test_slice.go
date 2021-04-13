package main

import "fmt"

type ChannelRel struct {
	ID int64 `json:"id,omitempty" gorm:"column:id"` // 自增 id
}
/*
func main() {
	//rels := make([]*ChannelRel, 0)
	//fmt.Println(len(rels))
	//fmt.Println(rels)

	slice := []int{1,2}
	fmt.Println(slice[:4]) // panic: runtime error: slice bounds out of range
}
*/
/*
func get() []byte {
	raw := make([]byte, 10000)
	fmt.Println(len(raw), cap(raw), &raw[0])    // 10000 10000 0xc420080000
	return raw[:3]    // 重新分配容量为 10000 的 slice
}
func main() {
	data := get()
	fmt.Println(len(data), cap(data), &data[0])    // 3 10000 0xc420080000
}
 */
func get() (res []byte) {
	raw := make([]byte, 10000)
	fmt.Println(len(raw), cap(raw), &raw[0])    // 10000 10000 0xc420080000
	res = make([]byte, 3)
	copy(res, raw[:3])
	return
}
func main() {
	data := get()
	fmt.Println(len(data), cap(data), &data[0])    // 3 3 0xc4200160b8
}
