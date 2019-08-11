package vo

type Gen struct {
	//作者
	Author string `json:"author"`
	//dao父类
	DaoParentClass string `json:"daoParentClass"`
	//dao包名
	DaoPkg string `json:"daoPkg"`
	//dao名称后辍
	DaoSuffix string `json:"daoSuffix"`
	//do父类
	DoParentClass string `json:"doParentClass"`
	//do包名
	DoPkg string `json:"doPkg"`
	//do名称后辍
	DoSuffix string `json:"doSuffix"`
	//是否生成模块名:on/off
	Module string `json:"module"`
	//service父类
	ServiceParentClass string `json:"serviceParentClass"`
	//service包名
	ServicePkg string `json:"servicePkg"`
	//service名称后辍
	ServiceSuffix string `json:"serviceSuffix"`
	//mapper.xml名称辍
	MapperXmlSuffix string `json:"mapperXmlSuffix"`
	//表名
	Tables []string `json:"tables"`
}
