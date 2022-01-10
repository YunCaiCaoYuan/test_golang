package main

import "fmt"

type student struct {
	id	int
	name string
}

type pbStu struct {
	id int
	name string
}

func main() {
	achiList := []*student{{1, "1"}, {2, "2"}}

	achiLangList := make([]*student, 0, len(achiList))
	for _, achi := range achiList {
		achiLang := *achi // 深拷贝
		achiLang.name = "111"

		fmt.Println("achi:", achi, "achiLang:", achiLang)
		achiLangList = append(achiLangList, &achiLang)
	}
	fmt.Println("achiList:")
	for _,v := range achiList {
		fmt.Println(v)

	}
	fmt.Println("achiLangList:")
	for _,v := range achiLangList {
		fmt.Println(v)
	}


	/*
	s1 := "hello"
	s2 := s1
	s2 = "world"
	fmt.Println(s1)
	fmt.Println(s2)
	//hello
	//world

	var stu1 = &student{id: 1, name: "123"}
	var pbStu1 = &pbStu{stu1.id, stu1.name}
	pbStu1.name = "456"

	fmt.Println("stu1:", stu1)
	fmt.Println("pbStu1:", pbStu1)
	*/

	/*
	obj1 := "你好"
	fmt.Println("obj1:", obj1, "len(obj1):", len(obj1), "len([]byte(obj1)):", len([]byte(obj1)), "utf8.RuneCountInString:", utf8.RuneCountInString(obj1))

	obj1 = "nihao"
	fmt.Println("obj1:", obj1, "len(obj1):", len(obj1), "len([]byte(obj1)):", len([]byte(obj1)), "utf8.RuneCountInString:", utf8.RuneCountInString(obj1))

	obj1 = "สบายดีไหม"
	fmt.Println("obj1:", obj1, "len(obj1):", len(obj1), "len([]byte(obj1)):", len([]byte(obj1)), "utf8.RuneCountInString:", utf8.RuneCountInString(obj1))

	obj1 = "Selamat datang!"
	fmt.Println("obj1:", obj1, "len(obj1):", len(obj1), "len([]byte(obj1)):", len([]byte(obj1)), "utf8.RuneCountInString:", utf8.RuneCountInString(obj1))

	obj1 = "Hoan nghênh"
	fmt.Println("obj1:", obj1, "len(obj1):", len(obj1), "len([]byte(obj1)):", len([]byte(obj1)), "utf8.RuneCountInString:", utf8.RuneCountInString(obj1))
	*/

	/*
	obj := "asd.qwe.123"
	objList := strings.Split(obj, ".")
	fmt.Println(objList)
	*/

	/*
		cell := "你好" // BOM
		ret := strings.Trim(cell, " ")
		fmt.Println(ret)
		fmt.Println("len 你好 = ", len(cell))
		fmt.Println("len 123 = ", len("123"))

		fmt.Println(string("") == "1")
	*/

	/*
		rankStr := ""
		num := 1
		rankStr += strconv.Itoa(num)
		fmt.Println(rankStr)
	*/

	/*
		i := small(10)
		println("i=", i)
	*/

	/*
		ret := strings.Contains("新手宝宝注意查收一下", "100")
		fmt.Println(ret)*/
	/*
		str := " 123 "
		ret := strings.TrimSpace(str)
		fmt.Println(ret)
	*/
}

/*
const digits = "0123456789abcdefghijklmnopqrstuvwxyz"
const smallsString = "00010203040506070809" +
	"10111213141516171819" +
	"20212223242526272829" +
	"30313233343536373839" +
	"40414243444546474849" +
	"50515253545556575859" +
	"60616263646566676869" +
	"70717273747576777879" +
	"80818283848586878889" +
	"90919293949596979899"
func small(i int) string {
	if i < 10 {
		return digits[i : i+1]
	}
	return smallsString[i*2 : i*2+2]
}
*/
