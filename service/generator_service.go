package service

import (
	"com.phh/generator/dao"
	"com.phh/generator/domain"
)

type GeneratorService interface {
	//获取表列表
	QueryTableList(name string) []domain.TableName
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
