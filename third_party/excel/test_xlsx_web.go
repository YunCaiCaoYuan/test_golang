package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tealeg/xlsx"
)

type DownloadRoleInfoBo struct {
	Name        string    `json:"name"`
	Level       int       `json:"level"`
	Description string    `json:"description"`
	CreateTime  time.Time `json:"createTime"`
}

// API处理器函数

func DownRolesHandler(c *gin.Context) {
	var roleData []DownloadRoleInfoBo

	// 略过向 roleData 添加数据过程

	var res []interface{}
	for _, role := range roleData {
		res = append(res, &DownloadRoleInfoBo{
			Name:        role.Name,
			Level:       role.Level,
			Description: role.Description,
			CreateTime:  role.CreateTime,
		})
	}
	content := ToExcel([]string{`角色名称`, `角色级别`, `描述`, `创建日期`}, res)
	ResponseXls(c, content, "角色数据")
}

// 生成io.ReadSeeker  参数 titleList 为Excel表头，dataList 为数据

func ToExcel(titleList []string, dataList []interface{}) (content io.ReadSeeker) {
	// 生成一个新的文件
	file := xlsx.NewFile()
	// 添加sheet页
	sheet, _ := file.AddSheet("Sheet1")
	// 插入表头
	titleRow := sheet.AddRow()
	for _, v := range titleList {
		cell := titleRow.AddCell()
		cell.Value = v
	}
	// 插入内容
	for _, v := range dataList {
		row := sheet.AddRow()
		row.WriteStruct(v, -1)
	}

	var buffer bytes.Buffer
	_ = file.Write(&buffer)
	content = bytes.NewReader(buffer.Bytes())
	return
}

// 向前端返回Excel文件
// 参数 content 为上面生成的io.ReadSeeker， fileTag 为返回前端的文件名

func ResponseXls(c *gin.Context, content io.ReadSeeker, fileTag string) {
	fileName := fmt.Sprintf("%s%s%s.xlsx", NowTime(), `-`, fileTag)
	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, fileName))
	c.Writer.Header().Add("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	http.ServeContent(c.Writer, c.Request, fileName, time.Now(), content)
}

func NowTime() string {
	return "123"
}

func main() {
	r := gin.Default()
	r.GET("/download", DownRolesHandler)
	r.Run()
}
