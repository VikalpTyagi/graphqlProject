package services

import (
	gormmodel "assingment/graph/gormModel"
	"assingment/graph/model"
	"context"
)

func (s *DbConnStruct) CreateUser(ctx context.Context, nu model.NewUser) (*model.User, error) {

	u := gormmodel.NewUser{
		Name:  nu.Email,
		Email: nu.Email,
	}
	err := s.db.Create(&u).Error
	if err != nil {
		return &model.User{}, err
	}
	return &model.User{
		Name:  u.Email,
		Email: u.Email,
	}, nil
}

func (s *DbConnStruct) ViewUser(ctx context.Context)([]*model.User,error){
	var listUser []*gormmodel.NewUser
	tx := s.db.WithContext(ctx)
	err := tx.Find(&listUser).Error
	if err != nil {
		return nil, err
	}
	var mUser []*model.User
	for _,user := range listUser{
		mUser{user}
	}

	return mUser, nil
}