package service

import (
	"com.phh/generator/utils/strutil"
	"fmt"
	"os"
	"testing"
	"text/template"
)

type Column struct {
	Type string
	Name string
}

//测试模版
func Test_tpl_entity(t *testing.T) {

	dataMap := make(map[string]interface{})
	dataMap["package"] = "com.phh"
	dataMap["author"] = "phh"
	dataMap["comments"] = "测试"
	dataMap["datetime"] = "2019-08-06 15:06:00"
	dataMap["className"] = "User"
	var columns []Column
	columns = append(columns, Column{"String", "userName"})
	columns = append(columns, Column{"Integer", "age"})
	columns = append(columns, Column{"Date", "birthday"})
	fmt.Println(columns)
	dataMap["columns"] = columns

	tpl, err := template.ParseFiles("../resource/tpl/entity.java.html")
	if err != nil {
		panic(err)
	}
	tpl.Execute(os.Stdout, dataMap)
}

func Test_tpl_mapper_java(t *testing.T) {
	dataMap := make(map[string]interface{})
	dataMap["package"] = "com.phh"
	dataMap["author"] = "phh"
	dataMap["comments"] = "测试"
	dataMap["datetime"] = "2019-08-06 15:06:00"
	dataMap["className"] = "User"
	baseMapper := "com.phh.base.mapper.BaseMapper"
	dataMap["baseMapper"] = baseMapper
	fmt.Println(strutil.IndexRune(baseMapper, "."))
	fmt.Println(strutil.LastIndexRune(baseMapper, "."))
	dataMap["baseMapperName"] = strutil.SubStr1(baseMapper, strutil.LastIndexRune(baseMapper, ".")+1)
	tpl, err := template.ParseFiles("../resource/tpl/mapper.java.html")
	if err != nil {
		panic(err)
	}
	tpl.Execute(os.Stdout, dataMap)
}
