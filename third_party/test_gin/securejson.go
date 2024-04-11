package main

// json劫持又称json hijacking，攻击过程有点类似于csrf，只不过csrf只管发送http请求，但是
// json-hijsck的目的是获取敏感数据。

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 你也可以使用自己的 SecureJSON 前缀
	// r.SecureJsonPrefix(")]}',\n")

	r.GET("/someJSON", func(c *gin.Context) {
		names := []string{"lena", "austin", "foo"}

		// 将输出：while(1);["lena","austin","foo"]
		c.SecureJSON(http.StatusOK, names)
	})

	// 监听并在 0.0.0.0:8080 上启动服务
	r.Run(":8080")
}

// 测试：
// ➜  bin curl http://localhost:8080/someJSON
// while(1);["lena","austin","foo"]%
