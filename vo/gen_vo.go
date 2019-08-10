package vo

type Gen struct {
	//作者
	Author string `json:"author"`
	//dao父类
	DaoParentClass string `json:"daoParentClass"`
	//dao包名
	DaoPkg string `json:"daoPkg"`
	//do父类
	DoParentClass string `json:"doParentClass"`
	//do包名
	DoPkg string `json:"doPkg"`
	//是否生成模块名:on/off
	Module string `json:"module"`
	//service父类
	ServiceParentClass string `json:"serviceParentClass"`
	//service包名
	ServicePkg string `json:"servicePkg"`
	//表名
	Tables []string `json:"tables"`
}
