package fileutil

import (
	"fmt"
	"os"
	"testing"
)

func TestGetFileNameList(t *testing.T) {
	list, _ := GetFileNameList("../../resource/tpl")
	fmt.Println(list)
}

//创建目录
func Test_createDir(t *testing.T) {

	err := os.MkdirAll("./com/phh/dao", os.ModeDir)
	fmt.Println(err)
}

func Test_createFile(t *testing.T) {
	f, err := os.OpenFile("./com/phh/dao/xx.text", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0777)
	if err != nil {
		fmt.Println(err)
		return
	}
	f.WriteString("新内容")

}
