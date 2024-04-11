package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// json全称json with padding，即“带回调的JSON”，它是json的一种“使用模式”。通过jsonp可以绕过浏览器的同源策略,
// 进行跨域请求。

func main() {
	r := gin.Default()

	r.GET("/JSONP", func(c *gin.Context) {
		data := map[string]interface{}{
			"foo": "bar",
		}

		// /JSONP?callback=x
		// 将输出：x({\"foo\":\"bar\"})
		c.JSONP(http.StatusOK, data)
	})

	// 监听并在 0.0.0.0:8080 上启动服务
	r.Run(":8080")
}
