package service

import (
	"assignment-7/entity"
	"assignment-7/repository"
	"log"
)

type ServiceInterface interface {
	Create(*entity.Orders) (*entity.Orders, error)
	GetById(int) (*entity.Orders, error)
	GetAll() (*[]entity.Orders, error)
	Update(*entity.Orders) (*entity.Orders, error)
	Delete(int)
}

type service struct {
}

var Serv ServiceInterface = &service{}

func (s *service) Create(userReq *entity.Orders) (*entity.Orders, error) {

	res, err := repository.Repo.Create(userReq)

	if err != nil {
		log.Fatal(err.Error())
	}

	return res, err
}

func (s *service) GetById(ordersId int) (*entity.Orders, error) {
	res, err := repository.Repo.GetById(ordersId)

	if err != nil {
		return nil, err
	}

	return res, err
}

func (s *service) GetAll() (*[]entity.Orders, error) {
	res, err := repository.Repo.GetAll()

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *service) Update(ordersReq *entity.Orders) (*entity.Orders, error) {
	res, err := repository.Repo.Update(ordersReq)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *service) Delete(ordersId int) {
	repository.Repo.Delete(ordersId)
}
