package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	test := 1234
	ret, _ := json.Marshal(test)
	fmt.Println("ret=", ret) // [49 50 51 52]

	var untest int
	_ = json.Unmarshal(ret, &untest)
	fmt.Println("untest=", untest) // 1234
}
