package main

import "reflect"

type test struct {
	Count int
	Price int
}

func (u *test) TableName() string {
	return "channel_skip_id"
}

func main() {
	const test = 4
	println(reflect.TypeOf(test).Name())

	// 测试不用make给切片分配内存，testList=nil
	//var testList []*test
	//fmt.Println("testList-len=", len(testList)) // len=0
	//if testList == nil {
	//	fmt.Println("1111")
	//}

	// 测试var定义结构体，是否可以调用方法，可以
	//var testVar test
	//testVar.Count = 10
	//fmt.Println("TableName=", testVar.TableName())
	//fmt.Println("testVar.Count=", testVar.Count)

	// 测试切片长度，len=0
	//testList := make([]*test, 0)
	//fmt.Println("testList.len=", len(testList))
	//fmt.Println("testList.len=", len(nil))

	// test switch
	//num := 1
	//switch num {
	//case 1	:
	//	//fmt.Println("1")
	//case 2	:
	//	fmt.Println("2")
	//default:
	//	fmt.Println("default")
	//}

	// test 格式化
	//sql := fmt.Sprintf("head '%s' num '%d' ", "a string ", 10000)
	//fmt.Println("sql : ", sql)
	//
	//sql2 := fmt.Sprintf("head %s num %d ", "a string ", 10000)
	//fmt.Println("sql2 : ", sql2)

	//Emoji
	//r := regexp.MustCompile("^[ a-zA-Z0-9\u4e00-\u9fa5\u1F300-\u1F5FF\u1F600-\u1F64F\u1F680-\u1F6FF\u2600-\u2B55]{1,4}$")
	//name := "ἰ0"
	//if r.MatchString(name) {
	//	fmt.Println("ok!!!")
	//}
}
