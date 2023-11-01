package services

import (
	gormmodel "assingment/graph/gormModel"
	"assingment/graph/model"
)

func (s *DbConnStruct) CreateCompany(newComp model.NewCompany) (*model.Company, error) {
	var job []*gormmodel.NewJob
	for _,sJob := range newComp.Jobs{
		nJob:= gormmodel.NewJob{
			Title: sJob.Title,
		}
		job = append(job,&nJob)
	}
	comp := gormmodel.NewCompany{
		Name: newComp.Name,
		City: newComp.City,
		Jobs: job,
	}
	err := s.db.Create(&comp).Error
	if err != nil {
		return &model.Company{}, err
	}
	return &model.Company{
		CompanyID: &comp.CompanyID,
		Name: comp.Name,
		City: comp.City,
	}, nil
}
