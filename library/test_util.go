package main

import (
	"fmt"
	"math/rand"
	//"crypto/rand"
)

func main() {
	b := make([]byte, 20)
	fmt.Println(b)

	_, err := rand.Read(b)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(b)
	/*
		for i := 0; i < 10; i++ {
			r := rand.New(rand.NewSource(time.Now().UnixNano()))
			fmt.Printf("%d ", r.Int31())
		}
		fmt.Println("")
		for i := 0; i < 10; i++ {
			//fmt.Printf("%d ", rand.Int31n(1000))

			ret := rand.Int31n(1000)
			//ret := time.Second*time.Duration(rand.Int31n(1000))
			fmt.Println(ret)
		}
	*/

	/*
		ret := rand.Int31n(1000)
		ret := time.Second*time.Duration(rand.Int31n(1000))
		fmt.Println(ret)
		fmt.Println(int64(ret))
	*/
	/*
		ret, _ := strconv.ParseInt("null", 10, 64)
		fmt.Println(ret)
	*/

	/*
		testMap := make(map[int32]bool)
		testMap[1] = true

		if testMap[2] {
			fmt.Println("found")
		}
		fmt.Println("not found")
		fmt.Println("testMap[2]", testMap[2])
	*/
}
