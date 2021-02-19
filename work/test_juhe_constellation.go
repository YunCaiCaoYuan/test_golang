package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	NewType = "today"
	Key     = "d0ed44e55968d40ada422799806d7bd5"
)

//type HttpResult struct {
//	ErrorCode int32  `json:"error_code"`
//	Datetime  string `json:"datetime"`
//}

func main() {
	apiUrl := "http://web.juhe.cn:8080/constellation/getAll"
	param := url.Values{}
	param.Set("consName", "白羊座") // 注意中文编码!!!
	param.Set("type", NewType)
	param.Set("key", Key) // 接口请求Key

	var Url *url.URL
	Url, err := url.Parse(apiUrl)
	if err != nil {
		fmt.Printf("解析url错误:\r\n%v", err)
		return
	}

	Url.RawQuery = param.Encode()
	resp, err := http.Get(Url.String())
	if err != nil {
		fmt.Println("get failed, err:", err)
		return
	}

	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("get resp failed,err:", err)
		return
	}

	//httpResult := new(HttpResult)
	var httpResult map[string]interface{}
	err = json.Unmarshal(b, &httpResult) // 参数无法解析直接丢弃
	if err != nil {
		fmt.Println("json unmarshal failed,err:", err)
		return
	}

	//if httpResult.ErrorCode == 0 {
	//	fmt.Println("datetime:", httpResult.Datetime)
	//}

	errorCode := httpResult["error_code"]
	if errorCode.(float64) == 0 {
		fmt.Println("datetime:", httpResult["datetime"])
		fmt.Println("name:", httpResult["name"])
		fmt.Println("color:", httpResult["color"])
		fmt.Println("health:", httpResult["health"])
		fmt.Println("love:", httpResult["love"])
		fmt.Println("work:", httpResult["work"])
		fmt.Println("money:", httpResult["money"])
		fmt.Println("work:", httpResult["work"])
		fmt.Println("number:", httpResult["number"])
		fmt.Println("summary:", httpResult["summary"])
		fmt.Println("all:", httpResult["all"])
	}
}
