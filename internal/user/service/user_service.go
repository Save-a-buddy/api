package service

import (
	"log"
	"save-a-buddy-api/internal/user/repository"
	"save-a-buddy-api/model"
)

type IUserService interface {
	FindUsers() (model.Users, error)
}

type UserService struct {
	userRepository repository.IUserRepository
}

func New(repository repository.IUserRepository) IUserService {
	return UserService{userRepository: repository}
}

func (us UserService) FindUsers() (model.Users, error) {
	users, err := us.userRepository.FindUsersDb()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return users, nil
}
