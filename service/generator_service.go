package service

import (
	"com.phh/generator/dao"
	"com.phh/generator/domain"
	"com.phh/generator/utils/dateutil"
	"com.phh/generator/utils/fileutil"
	"com.phh/generator/utils/mysqlutil"
	"com.phh/generator/utils/strutil"
	"com.phh/generator/vo"
	"fmt"
	"github.com/kataras/iris/core/errors"
	"github.com/rs/zerolog/log"
	"os"
	"strings"
	"text/template"
	"time"
)

type GeneratorService interface {
	//获取表列表
	QueryTableList(name string) []domain.TableName

	//生成代码模版文件
	Generate(gen vo.Gen) (filePath string, err error)
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
func (g *generatorService) Generate(gen vo.Gen) (filePath string, err error) {
	setDefaulSuffix(&gen)
	//模版函数
	funcMap := getFuncMap()
	tmplDir := "./resource/tpl/"
	files, err := fileutil.GetFileNameList(tmplDir)
	if err != nil {
		fmt.Println(err)
		log.Error().Msg(TEMPLATE_NOT_FOUND.Error())
		return "", TEMPLATE_NOT_FOUND
	}
	var filePaths []string
	for _, v := range files {
		filePaths = append(filePaths, tmplDir+v)
	}
	tmpl, err := template.New("mybatisTmpl").Funcs(funcMap).ParseFiles(filePaths...)
	if err != nil {
		log.Error().Err(err).Msg(TEMPLATE_LOAD_ERROR.Error())
		return "", TEMPLATE_LOAD_ERROR
	}
	date := dateutil.FormatTime(time.Now(), "yyyy-MM-dd")
	//生成文件根目录
	rootDir := "./mybatis_tmpl/"
	_ = os.RemoveAll(rootDir)
	for _, tableName := range gen.Tables {
		//模版参数
		dataMap := map[string]interface{}{}
		dataMap["gen"] = gen
		dataMap["date"] = date
		dataMap["primaryKeyName"] = "id" //主键实体映射统一为：id
		hasServiceInterface := gen.HasServiceInterface == "on"
		dataMap["hasServiceInterface"] = hasServiceInterface
		dataMap["hasDate"] = false
		dataMap["hasBigDecimal"] = false
		columns := g.genDao.GetTableColumnsByTableName(tableName)
		dataMap["columnNumber"] = len(columns)
		//转换列名与类型
		for i, col := range columns {
			//mysql类型转换为java类型
			columns[i].JavaType = mysqlutil.GetJavaType(col.DataType)
			//mysql字段名称转换为符合java名称
			javaName := strutil.UnderLineToCamelcase(col.Name)
			javaName = strutil.FirstLetterToLower(javaName)
			columns[i].JavaName = javaName
			if columns[i].JavaType == "Date" {
				dataMap["hasDate"] = true
			}
			if columns[i].JavaType == "BigDecimal" {
				dataMap["hasBigDecimal"] = true
			}
			//主键
			if col.Key == "PRI" || i == 0 {
				//没有主键就是第一列
				dataMap["primaryKeyJdbcType"] = col.DataType
				dataMap["primaryKeyJavaType"] = columns[i].JavaType
				dataMap["primaryKeyColumn"] = col.Name
			}
		}
		table := g.genDao.GetTableByTableName(tableName)
		dataMap["columns"] = columns
		dataMap["table"] = table
		//表名对应java名称：下划线转换为驼峰，首字母大写
		javaName := tableName
		//是否生成模块名
		hasModule := gen.HasModule == "on"
		//模块名称
		mod := ""
		if hasModule {
			//解析第一个下划线前的词
			lineIdx := strutil.IndexRune(tableName, "_")
			if lineIdx <= 0 {
				hasModule = false
			} else if lineIdx < (len(tableName) - 1) {
				//if 确保下划线不是最后一个字符
				mod = strutil.SubStr2(tableName, 0, lineIdx)
				//截取第一个下划线之后部分为java名称
				javaName = strutil.SubStr1(tableName, lineIdx+1)
			} else {
				hasModule = false
			}
		}
		//是否生成模块名
		dataMap["hasModule"] = hasModule
		//模块名称
		dataMap["mod"] = mod
		//把表名下划线转换为驼峰
		javaName = strutil.UnderLineToCamelcase(javaName)
		javaName = strutil.FirstLetterToUpper(javaName)
		dataMap["javaName"] = javaName
		dataMap["serialVersionUID"] = time.Now().UnixNano()
		//生成文件
		modPath := "/" + mod
		if !strings.HasSuffix(modPath, "/") {
			modPath = modPath + "/"
		}
		doDir := rootDir + strings.ReplaceAll(gen.DoPkg, ".", "/") + modPath
		daoDir := rootDir + strings.ReplaceAll(gen.DaoPkg, ".", "/") + modPath
		serviceDir := rootDir + strings.ReplaceAll(gen.ServicePkg, ".", "/") + modPath
		serviceImplDir := rootDir + strings.ReplaceAll(gen.ServicePkg, ".", "/") + modPath + "impl/"
		mapperXmlDir := rootDir + "/mapperXml/" + modPath
		_ = os.MkdirAll(doDir, os.ModeDir|os.ModePerm)
		_ = os.MkdirAll(daoDir, os.ModeDir|os.ModePerm)
		_ = os.MkdirAll(serviceDir, os.ModeDir|os.ModePerm)
		_ = os.MkdirAll(mapperXmlDir, os.ModeDir|os.ModePerm)
		if hasServiceInterface {
			_ = os.MkdirAll(serviceImplDir, os.ModeDir|os.ModePerm)
		}
		filePath := ""
		for _, v := range files {
			if strings.Contains(v, "do.java") {
				filePath = doDir + javaName + gen.DoSuffix + ".java"
			} else if strings.Contains(v, "dao.java") {
				filePath = daoDir + javaName + gen.DaoSuffix + ".java"
			} else if strings.Contains(v, "service.java") {
				if hasServiceInterface {
					//如果有接口，实现在就在impl子目录下
					filePath = serviceImplDir + javaName + gen.ServiceSuffix + "Impl.java"
				} else {
					filePath = serviceDir + javaName + gen.ServiceSuffix + ".java"
				}
			} else if strings.Contains(v, "service.interface.java") && hasServiceInterface {
				//service接口
				filePath = serviceDir + javaName + gen.ServiceSuffix + ".java"
			} else if strings.Contains(v, "mapper.xml") {
				filePath = mapperXmlDir + javaName + gen.MapperXmlSuffix + ".xml"
			} else {
				//未启用的模版
				continue
			}
			file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0777)
			if err != nil {
				fmt.Println(err)
				return "", FILE_OPEN_ERROR
			}
			err = tmpl.ExecuteTemplate(file, v, dataMap)
			if err != nil {
				fmt.Println(err)
				return "", TEMPLATE_RENDER_ERROR
			}
			file.Close()
		}
	}
	//打包zip
	fileName := "./mybatis-tmpl-out.zip"
	zipFile, err := os.Create(fileName)
	fmt.Println(err)
	defer zipFile.Close()
	err = fileutil.Zip(rootDir, zipFile)
	if err != nil {
		fmt.Println(err)
		return "", errors.New(fmt.Sprintf("生成文件成功，但打包zip失败，可在目录(%s)找到文件", rootDir))
	}
	return fileName, nil
}

func setDefaulSuffix(gen *vo.Gen) {
	//dao,do,service名称后辍
	if gen.DoSuffix == "" {
		gen.DoSuffix = "DO"
	}
	if gen.DaoSuffix == "" {
		gen.DaoSuffix = "DAO"
	}
	if gen.ServiceSuffix == "" {
		gen.ServiceSuffix = "Service"
	}
	if gen.MapperXmlSuffix == "" {
		gen.MapperXmlSuffix = "Mapper"
	}
}

func getFuncMap() template.FuncMap {
	return template.FuncMap{
		"FirstToLower": func(str string) string {
			return strutil.FirstLetterToLower(str)
		},
		"FirstToUpper": func(str string) string {
			return strutil.FirstLetterToUpper(str)
		},
		"ClassName": func(fullClass string) string {
			dotIdx := strutil.LastIndexRune(fullClass, ".")
			return strutil.SubStr1(fullClass, dotIdx+1)
		},
		"ToUpper": func(str string) string {
			return strings.ToUpper(str)
		},
		"ToLower": func(str string) string {
			return strings.ToLower(str)
		},
		"Add": func(a int, b int) int {
			return a + b
		},
		"Minus": func(a int, b int) int {
			return a - b
		},
	}
}
