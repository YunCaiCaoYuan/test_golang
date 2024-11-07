package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

func main() {
	// 替换 clientId、accessToken、macKey 参数
	// clientId 参数在 TapDC 后台查看
	clientId := "ifv7wj6he6amcdwl6m"
	// TapTap 登录成功后 TapSDK 返回的 kid
	kid := "1/jvsVxFC6-PIUiXvZVVtv1hogX5q9Z1y_rp-AtVjE3iyHikXXfd_2h-i0wLmc9UjLJwhH6fQ8cvGrklONdvy2J5YfoqzV0ewGPMSLkQIkRv_xaLaYPariWbrkP1MtG2b4CzR1KHvuSCJHewCmTFZmsyNGojTJr5t75f5Nc8j-jjCYeDtFO0-XFI_J7kzktswzzsmISt7cx49QVess-VbaQcU31pEDb_OA03I28H5ehIvqQ0CQdf1LieLyONcH97l1IEU39AirioF_KGJccVG64QsgWmzxLPwmfTurw4cwBPo04yuXnas4YI5haE2UxtckNCpagP19drtGW57-HaAdww"
	// TapTap 登录成功后 TapSDK 返回的 mac_key
	macKey := "fTCuDUDDmNny7a36EWbhUDLaqpoDMQu2hCi9qAJ5"

	// 随机数，正式上线请替换
	nonce := "8IBTHwOdqNKAWeKl7plt66=="
	// 时间戳转换成字符串
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	// 请求 url 相关
	reqHost := "open.tapapis.cn"
	reqURI := "/account/profile/v1?client_id=" + clientId
	reqURL := "https://" + reqHost + reqURI

	macStr := timestamp + "\n" + nonce + "\n" + "GET" + "\n" + reqURI + "\n" + reqHost + "\n" + "443" + "\n\n"
	mac := hmacSha1(macStr, macKey)
	authorization := "MAC id=" + "\"" + kid + "\"" + "," + "ts=" + "\"" + timestamp + "\"" + "," + "nonce=" + "\"" + nonce + "\"" + "," + "mac=" + "\"" + mac + "\""

	client := http.Client{}
	req, err := http.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// 添加请求头
	req.Header.Add("Authorization", authorization)
	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(respBody))
}

/*
HMAC-SHA1 签名
*/
func hmacSha1(valStr, keyStr string) string {
	key := []byte(keyStr)
	mac := hmac.New(sha1.New, key)
	mac.Write([]byte(valStr))

	// 进行 Base64 编码
	return base64.StdEncoding.EncodeToString(mac.Sum(nil))
}
