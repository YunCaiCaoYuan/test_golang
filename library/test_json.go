package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func checkError(err error) {
	if err != nil{
		log.Fatalln(err)
	}
}

// 状态名称可能是 int 也可能是 string，指定为 json.RawMessage 类型
func main() {
	bin, _ := json.Marshal(1)
	fmt.Println(bin)
	ret := int32(0)
	json.Unmarshal(bin, &ret)
	fmt.Println(ret)

	/*
	records := [][]byte{
		[]byte(`{"status":200, "tag":"one"}`),
		[]byte(`{"status":"ok", "tag":"two"}`),
	}

	for idx, record := range records {
		var result struct {
			StatusCode uint64
			StatusName string
			Status     json.RawMessage `json:"status"`
			Tag        string          `json:"tag"`
		}

		err := json.NewDecoder(bytes.NewReader(record)).Decode(&result)
		checkError(err)

		var name string
		err = json.Unmarshal(result.Status, &name)
		if err == nil {
			result.StatusName = name
		}

		var code uint64
		err = json.Unmarshal(result.Status, &code)
		if err == nil {
			result.StatusCode = code
		}

		fmt.Printf("[%v] result => %+v\n", idx, result)
	}
	 */
}

/*
// struct 中指定字段类型
func main() {
	var data = []byte(`{"status": 200}`)
	var result struct {
		Status uint64 `json:"status"`
	}

	err := json.NewDecoder(bytes.NewReader(data)).Decode(&result)
	checkError(err)
	fmt.Printf("Result: %+v", result)
}
*/
/*
// 你可以使用 string 来存储数值数据，在 decode 时再决定按 int 还是 float 使用
// 将数据转为 decode 为 string
func main() {
	var data = []byte(`{"status": 200}`)
	var result map[string]interface{}
	var decoder = json.NewDecoder(bytes.NewReader(data))
	decoder.UseNumber()
	if err := decoder.Decode(&result); err != nil {
		log.Fatalln(err)
	}
	var status uint64
	err := json.Unmarshal([]byte(result["status"].(json.Number).String()), &status);
	checkError(err)
	fmt.Println("Status value: ", status)
}
 */
/*
// 指定字段类型
func main() {
	var data = []byte(`{"status": 200}`)
	var result map[string]interface{}

	var decoder = json.NewDecoder(bytes.NewReader(data))
	decoder.UseNumber()

	if err := decoder.Decode(&result); err != nil {
		log.Fatalln(err)
	}

	var status, _ = result["status"].(json.Number).Int64()
	fmt.Println("Status value: ", status)
}
*/
/*
// 将 decode 的值转为 int 使用
func main() {
	var data = []byte(`{"status": 200}`)
	var result map[string]interface{}

	if err := json.Unmarshal(data, &result); err != nil {
		log.Fatalln(err)
	}

	var status = uint64(result["status"].(float64))
	fmt.Println("Status value: ", status)
}
*/
/*
func main() {
	var data = []byte(`{"status": 200}`)
	var result map[string]interface{}

	if err := json.Unmarshal(data, &result); err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("%T\n", result["status"])    // float64
	var status = result["status"].(int)    // 类型断言错误
	fmt.Println("Status value: ", status)
}
*/
/*
func main() {
	test := 1234
	ret, _ := json.Marshal(test)
	fmt.Println("ret=", ret) // [49 50 51 52]

	var untest int
	_ = json.Unmarshal(ret, &untest)
	fmt.Println("untest=", untest) // 1234
}
*/

