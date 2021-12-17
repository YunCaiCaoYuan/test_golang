package main

import "fmt"

type ChannelRel struct {
	ID int64 `json:"id,omitempty" gorm:"column:id"` // 自增 id
}

type testIntSlice struct {
	intSlice []int
}

func main() {
	//var intSlice []int
	//intSlice := make([]int, 0)
	var tis testIntSlice
	tis.intSlice = append(tis.intSlice, 1)
	fmt.Println(tis.intSlice)

	/*
	slice := []int{1,2,3}
	fmt.Println(slice[:])
	fmt.Println(slice)
	 */

	/*
	slice := []int{10, 20, 30, 40}
	for index, value := range slice {
		fmt.Printf("value = %d , value-addr = %x , slice-addr = %x\n", value, &value, &slice[index])
	}
	 */

	/*
	slice := make([]byte, 3)
	n := copy(slice, "abcdef")
	fmt.Println(n,slice)
	 */

	/*
	array := []int{10, 20, 30, 40}
	slice := make([]int, 6)
	n := copy(slice, array)
	fmt.Println(n,slice)
	 */

	/*
	list := []int{3,1,4,15,9,2,6,5,31}
	//slice := list[2:5:7] // [4 15 9]
	slice := list[2:5] // [4 15 9]
	//slice = append(slice, 123)
	fmt.Println(slice)
	fmt.Println(list)
	fmt.Println(cap(slice))
	 */

	/*
	list := make([]int32, 1)
	list = append(list, 100)
	fmt.Println("list=", list)
	fmt.Println("len(list)=", len(list), " cap(list)=", cap(list))
	//list= [0 100]
	//len(list)= 2  cap(list)= 2
	 */

	/*
	list := make([]int32, 7)
	fmt.Println("len(list)=", len(list), " cap(list)=", cap(list))
	list2 := make([]int32, 0, 7)
	fmt.Println("len(list2)=", len(list2), " cap(list2)=", cap(list2))
	 */
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
/*
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
 */
/*
// 错误使用 slice 的拼接示例
func main() {
	path := []byte("AAAA/BBBBBBBBB")
	sepIndex := bytes.IndexByte(path, '/')
	println(sepIndex)							// 4

	dir1 := path[:sepIndex]
	dir2 := path[sepIndex+1:]
	println("dir1: ", string(dir1))        // AAAA
	println("dir2: ", string(dir2))        // BBBBBBBBB

	dir1 = append(dir1, "suffix"...)
	println("current path: ", string(path))    // AAAAsuffixBBBB

	path = bytes.Join([][]byte{dir1, dir2}, []byte{'/'})
	println("dir1: ", string(dir1))        // AAAAsuffix
	println("dir2: ", string(dir2))        // uffixBBBB

	println("new path: ", string(path))    // AAAAsuffix/uffixBBBB    // 错误结果
}
 */
/*
// 使用 full slice expression
func main() {

	path := []byte("AAAA/BBBBBBBBB")
	sepIndex := bytes.IndexByte(path, '/') // 4
	dir1 := path[:sepIndex:sepIndex]        // 此时 cap(dir1) 指定为4， 而不是先前的 16
	dir2 := path[sepIndex+1:]
	dir1 = append(dir1, "suffix"...)

	path = bytes.Join([][]byte{dir1, dir2}, []byte{'/'})
	println("dir1: ", string(dir1))        // AAAAsuffix
	println("dir2: ", string(dir2))        // BBBBBBBBB
	println("new path: ", string(path))    // AAAAsuffix/BBBBBBBBB
}
*/

// 超过容量将重新分配数组来拷贝值、重新存储
/*
func main() {
	s1 := []int{1, 2, 3}
	fmt.Println(len(s1), cap(s1), s1)    // 3 3 [1 2 3 ]

	s2 := s1[1:]
	fmt.Println(len(s2), cap(s2), s2)    // 2 2 [2 3]

	for i := range s2 {
		s2[i] += 20
	}
	// 此时的 s1 与 s2 是指向同一个底层数组的
	fmt.Println(s1)        // [1 22 23]
	fmt.Println(s2)        // [22 23]

	s2 = append(s2, 4)    // 向容量为 2 的 s2 中再追加元素，此时将分配新数组来存

	for i := range s2 {
		s2[i] += 10
	}
	fmt.Println(s1)        // [1 22 23]    // 此时的 s1 不再更新，为旧数据
	fmt.Println(s2)        // [32 33 14]
}
 */
