package service

import (
	"com.phh/generator/dao"
	"com.phh/generator/domain"
	"com.phh/generator/utils/dateutil"
	"com.phh/generator/utils/mysqlutil"
	"com.phh/generator/utils/strutil"
	"com.phh/generator/vo"
	"fmt"
	"github.com/rs/zerolog/log"
	"os"
	"text/template"
	"time"
)

type GeneratorService interface {
	//获取表列表
	QueryTableList(name string) []domain.TableName

	//生成代码模版文件
	Generate(gen vo.Gen) error
}

func NewGeneratorService(dao dao.GeneratorDao) GeneratorService {
	return &generatorService{
		genDao: dao,
	}
}

type generatorService struct {
	genDao dao.GeneratorDao
}

//获取表列表
func (g *generatorService) QueryTableList(name string) []domain.TableName {
	return g.genDao.QueryTableList(name)
}

//生成代码模版文件
func (g *generatorService) Generate(gen vo.Gen) error {
	fmt.Println(gen)

	//加载模版
	funcMap := template.FuncMap{
		"FirstToLower": func(str string) string {
			return strutil.FirstLetterToLower(str)
		},
		"FirstToUpper": func(str string) string {
			return strutil.FirstLetterToUpper(str)
		},
	}
	files := []string{
		"./resource/tpl/entity.java.html",
	}
	tmpl, err := template.New("").Funcs(funcMap).ParseFiles(files...)
	if err != nil {
		fmt.Println(err)
		log.Error().Msg(TEMPLATE_NOT_FOUND.Error())
		return TEMPLATE_NOT_FOUND
	}
	date := dateutil.FormatTime(time.Now(), "yyyy-MM-dd")
	dataMap := map[string]interface{}{}
	dataMap["gen"] = gen
	dataMap["date"] = date
	for _, tableName := range gen.Tables {
		columns := g.genDao.GetTableColumnsByTableName(tableName)
		//转换列名与类型
		for i, col := range columns {
			//mysql类型转换为java类型
			columns[i].JavaType = mysqlutil.GetJavaType(col.DataType)
			//mysql字段名称转换为符合java名称
			javaName := strutil.UnderLineToCamelcase(col.Name)
			javaName = strutil.FirstLetterToLower(javaName)
			columns[i].JavaName = javaName
		}
		table := g.genDao.GetTableByTableName(tableName)
		dataMap["columns"] = columns
		dataMap["table"] = table
		//表名对应java名称：下划线转换为驼峰，首字母大写
		//TODO 是否生成模块名
		javaName := strutil.UnderLineToCamelcase(tableName)
		javaName = strutil.FirstLetterToUpper(javaName)
		dataMap["javaName"] = javaName
		dataMap["serialVersionUID"] = time.Now().UnixNano()
		err := tmpl.ExecuteTemplate(os.Stdout, "entity.java.html", dataMap)
		fmt.Println(err)
		//生成实体

		//生成dao

		//生成mapper.xml

	}

	return nil
}
