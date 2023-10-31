package services

import (
	"assingment/graph/model"
	"context"
)

type Service interface {
	CreateUser(ctx context.Context, nu model.NewUser) (*model.User, error)
	ViewUser(ctx context.Context)([]*model.User,error)
}

type Store struct {
	Service
}

func NewStore(s Service) Store {
	return Store{Service: s}
}
