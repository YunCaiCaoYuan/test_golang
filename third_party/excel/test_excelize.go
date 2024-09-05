package main

import (
	"bytes"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
)

type fruit struct {
	ID    uint
	Name  string
	Price float64
}

func getFruits() []*fruit {
	return []*fruit{
		&fruit{
			ID:    1,
			Name:  "苹果",
			Price: 8,
		},
		&fruit{
			ID:    2,
			Name:  "雪梨",
			Price: 5,
		},
		&fruit{
			ID:    3,
			Name:  "香蕉",
			Price: 3,
		},
	}
}

func main() {
	r := gin.Default()
	r.GET("/download", download)
	r.Run()
}

func download(c *gin.Context) {
	f := excelize.NewFile() // 设置单元格的值
	// 这里设置表头
	f.SetCellValue("Sheet1", "A1", "序号")
	f.SetCellValue("Sheet1", "B1", "名称")
	f.SetCellValue("Sheet1", "C1", "价格")

	line := 1

	fruits := getFruits()

	// 设置列宽
	f.SetColWidth("Sheet1", "A", "C", 20)

	// 设置表头行高
	f.SetRowHeight("Sheet1", 1, 30)

	// 循环写入数据
	for _, v := range fruits {
		line++
		f.SetCellValue("Sheet1", fmt.Sprintf("A%d", line), v.ID)
		f.SetCellValue("Sheet1", fmt.Sprintf("B%d", line), v.Name)
		f.SetCellValue("Sheet1", fmt.Sprintf("C%d", line), v.Price)
	}

	// 保存文件
	//if err := f.SaveAs("fruits.xlsx"); err != nil {
	//	fmt.Println(err)
	//}

	var buffer bytes.Buffer
	_ = f.Write(&buffer)
	content := bytes.NewReader(buffer.Bytes())

	fileName := fmt.Sprintf("%s.xlsx", "角色数据")
	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, fileName))
	c.Writer.Header().Add("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	http.ServeContent(c.Writer, c.Request, fileName, time.Now(), content)
}
