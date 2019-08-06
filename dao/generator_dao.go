package dao

import (
	"com.phh/generator/domain"
	"github.com/jinzhu/gorm"
)

//mapper
type GeneratorDao interface {
	//获取表列表
	QueryTableList(name string) []domain.TableName
	//根据表名获取表信息
	GetTableByTableName(tableName string) domain.TableName
	//获取表的列
	GetTableColumnsByTableName(tableName string) []domain.Column
}

func NewGeneratorDao(db *gorm.DB) GeneratorDao {
	return &generatorDaoMapper{db: db}
}

type generatorDaoMapper struct {
	//数据源
	db *gorm.DB
}

func (g *generatorDaoMapper) QueryTableList(name string) []domain.TableName {
	var list []domain.TableName
	db := g.db.Select("TABLE_NAME,TABLE_COMMENT,CREATE_TIME").
		Where("table_schema = (SELECT DATABASE())")
	if name != "" {
		db = db.Where("table_name LIKE ?", "%"+name+"%")
	}
	db.Find(&list)
	return list
}

func (g *generatorDaoMapper) GetTableByTableName(tableName string) domain.TableName {
	var table domain.TableName
	g.db.Select("TABLE_NAME,TABLE_COMMENT,CREATE_TIME").
		Where("table_schema = (SELECT DATABASE()) AND table_name=?", tableName).
		Find(&table)
	return table
}

func (g *generatorDaoMapper) GetTableColumnsByTableName(tableName string) []domain.Column {
	var list []domain.Column
	g.db.Select("COLUMN_NAME,DATA_TYPE,COLUMN_COMMENT,COLUMN_KEY,EXTRA").
		Where("table_name = ? AND table_schema = (SELECT DATABASE())", tableName).
		Find(&list)
	return list
}
