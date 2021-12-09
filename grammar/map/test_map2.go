package main

import "fmt"

func main() {
	inputOpt := make(map[string]string)
	if len(inputOpt["region"]) != 0 {
		fmt.Println("region not empty!")
	} else {
		fmt.Println("region empty!")
	}

	switch inputOpt["region"] {
	case "Tmp":
		fmt.Println("Tmp")
	default:
		fmt.Println("default")
	}
}

/*
// fatal error: concurrent map writes

func main() {
	fmt.Println("start...")
	testMap := make(map[int]int)
	wg := sync.WaitGroup{}
	for i := 0; i < 10000; i ++ {
		testMap[i] = i
	}


	for i := 0; i < 10000; i ++ {
		wg.Add(1)
		go func(key int) {
			delete(testMap, key)
		}(i)
		wg.Done()
	}

	wg.Wait()
	fmt.Println("finish...")
}
*/
