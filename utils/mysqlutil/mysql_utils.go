package mysqlutil

import "strings"

var mysqlJavaTypeMap map[string]string
var mysqlJdbcTypeMap map[string]string

func init() {
	//mysql类型-java类型映射
	mysqlJavaTypeMap = map[string]string{}
	//长整型
	mysqlJavaTypeMap["bigint"] = "Long"
	//整型
	mysqlJavaTypeMap["int"] = "Integer"
	mysqlJavaTypeMap["tinyint"] = "Integer"
	mysqlJavaTypeMap["smallint"] = "Integer"
	mysqlJavaTypeMap["mediumint"] = "Integer"
	mysqlJavaTypeMap["integer"] = "Integer"
	//小数
	mysqlJavaTypeMap["float"] = "Float"
	mysqlJavaTypeMap["double"] = "Double"
	mysqlJavaTypeMap["decimal"] = "Double"
	//bool
	mysqlJavaTypeMap["bit"] = "Boolean"
	//字符串
	mysqlJavaTypeMap["char"] = "String"
	mysqlJavaTypeMap["varchar"] = "String"
	mysqlJavaTypeMap["tinytext"] = "String"
	mysqlJavaTypeMap["text"] = "String"
	mysqlJavaTypeMap["mediumtext"] = "String"
	mysqlJavaTypeMap["longtext"] = "String"
	//日期
	mysqlJavaTypeMap["date"] = "Date"
	mysqlJavaTypeMap["datetime"] = "Date"
	mysqlJavaTypeMap["timestamp"] = "Date"
	//其它统一为字符串
	//-------------------------------------------------------
	//mysql类型与jdbcType对应
	mysqlJdbcTypeMap = map[string]string{}
	mysqlJdbcTypeMap["bigint"] = "BIGINT"
	mysqlJdbcTypeMap["decimal"] = "DECIMAL"
	mysqlJdbcTypeMap["double"] = "DOUBLE"
	mysqlJdbcTypeMap["float"] = "FLOAT"
	mysqlJdbcTypeMap["int"] = "INTEGER"
	mysqlJdbcTypeMap["integer"] = "INTEGER"
	mysqlJdbcTypeMap["boolean"] = "BIT"
	mysqlJdbcTypeMap["bit"] = "BIT"
	mysqlJdbcTypeMap["tinyint"] = "TINYINT"
	mysqlJdbcTypeMap["smallint"] = "SMALLINT"
	mysqlJdbcTypeMap["numeric"] = "NUMERIC"
	mysqlJdbcTypeMap["blob"] = "BLOB"
	mysqlJdbcTypeMap["char"] = "CHAR"
	mysqlJdbcTypeMap["clob"] = "CLOB"
	mysqlJdbcTypeMap["real"] = "REAL"
	mysqlJdbcTypeMap["date"] = "DATE"
	mysqlJdbcTypeMap["time"] = "TIME"
	mysqlJdbcTypeMap["year"] = "VARCHAR"
	mysqlJdbcTypeMap["timestamp"] = "TIMESTAMP"
	mysqlJdbcTypeMap["datetime"] = "TIMESTAMP"
	mysqlJdbcTypeMap["varchar"] = "VARCHAR"
	mysqlJdbcTypeMap["tinytext"] = "VARCHAR"
	mysqlJdbcTypeMap["text"] = "VARCHAR"
	mysqlJdbcTypeMap["mediumtext"] = "LONGVARCHAR"
	mysqlJdbcTypeMap["longtext"] = "LONGVARCHAR"
}

func GetJavaType(mysqlType string) string {
	javaType := mysqlJavaTypeMap[mysqlType]
	if javaType == "" {
		return "String"
	}
	return javaType
}

func GetJdbcType(mysqlType string) string {
	jdbcType := mysqlJdbcTypeMap[mysqlType]
	if jdbcType == "" {
		return strings.ToUpper(mysqlType)
	}
	return jdbcType
}
