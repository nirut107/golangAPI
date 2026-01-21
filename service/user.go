package service


import (
	"go-backend/model"
	"go-backend/repository"
)

type UserService struct{
	Repo repository.UserRepository
}

func (s UserService) GetAll() ([]model.User, error) {
	return s.Repo.GetAll()
}

func (s UserService) GetByID(id int) (model.User, error) {
	return s.Repo.GetByID(id)
}

func (s UserService) Create(u model.User) (model.User, error) {
	return s.Repo.Create(u)
}

func (s UserService) Delete(id int) error {
	return s.Repo.Delete(id)
}

func (s UserService) Update(u model.User) (model.User, error) {
	return s.Repo.Update(u)
}