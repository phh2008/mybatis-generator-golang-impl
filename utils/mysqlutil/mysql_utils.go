package mysqlutil

var mysqlJavaTypeMap map[string]string

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
	mysqlJavaTypeMap["decimal"] = "BigDecimal"
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
}

func GetJavaType(mysqlType string) string {
	javaType := mysqlJavaTypeMap[mysqlType]
	if javaType == "" {
		return "String"
	}
	return javaType
}
