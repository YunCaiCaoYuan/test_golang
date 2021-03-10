package main

// !!! 未通过
import (
	"os/exec"
	"testing"
	"errors"
	"strings"
	"fmt"
)

//执行python脚本
func CmdPythonSaveImageDpi(filePath, newFilePath string) (err error) {
	args := []string{"main.py", filePath, newFilePath}
	out, err := exec.Command("python", args...).Output()
	if err != nil {
		return
	}
	result := string(out)
	if strings.Index(result, "success") != 0 {
		err = errors.New(fmt.Sprintf("main.py error：%s", result))
	}
	return
}

//test测试
func TestCmdPython(t *testing.T) {
	//test.txt的内容为图片的base64字符串
	filePath := "test.txt"
	newFileName := "test.jpg"
	err :=CmdPythonSaveImageDpi(filePath,newFileName)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("转换成功")
}

