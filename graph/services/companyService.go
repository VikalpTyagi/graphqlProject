package services

import (
	gormmodel "assingment/graph/gormModel"
	"assingment/graph/model"
)

func (s *DbConnStruct) CreateCompany(newComp model.NewCompany) (*model.Company, error) {
	var job []*gormmodel.NewJob
	comp := gormmodel.NewCompany{
		Name: newComp.Name,
		City: newComp.City,
		Jobs: append(job,newComp.Jobs),
	}
	err := s.db.Create(&comp).Error
	if err != nil {
		return &model.Company{}, err
	}
	return comp, nil
}
