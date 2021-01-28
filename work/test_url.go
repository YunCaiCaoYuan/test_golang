package main

import (
	"fmt"
	"net/url"
)

func main() {
	fileUrl := "http://yes-web.oss-cn-shenzhen.aliyuncs.com/upload/beta/icon/5B92/C937/5B92B2D6ED2B1FC937.jpg"
	//fileUrl := "upload/beta/icon/5B92/C937/5B92B2D6ED2B1FC937.jpg"
	u, err := url.Parse(fileUrl)
	if err != nil {
		fmt.Println("err", err)
	}
	fmt.Println("u.Path=", u.Path)

	// 存放去除域名等前缀
	iconUrl := ""
	if u.Path[0] == '/' {
		iconUrl = u.Path[1:]
	} else {
		iconUrl = u.Path
	}
	fmt.Println("iconUrl", iconUrl)

	// fileUrl = "http://yes-web.oss-cn-shenzhen.aliyuncs.com/upload/beta/icon/5B92/C937/5B92B2D6ED2B1FC937.jpg"
	// u.Path = /upload/beta/icon/5B92/C937/5B92B2D6ED2B1FC937.jpg

	//fileUrl := "upload/beta/icon/5B92/C937/5B92B2D6ED2B1FC937.jpg"
	// u.Path = upload/beta/icon/5B92/C937/5B92B2D6ED2B1FC937.jpg

}
