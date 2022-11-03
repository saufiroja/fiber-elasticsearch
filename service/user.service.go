package service

import (
	"elasticsearch/fiber-elasticsearch/entity"
	"elasticsearch/fiber-elasticsearch/interfaces"
)

type UserService struct {
	UserRepository interfaces.UserRepository
}

func NewUserService(userRepository interfaces.UserRepository) interfaces.UserService {
	return &UserService{
		UserRepository: userRepository,
	}
}

func (u *UserService) CreateUser(user *entity.User) (entity.User, error) {
	return u.UserRepository.CreateUser(user)
}

func (u *UserService) FindAllUser() (map[string]any, error) {
	return u.UserRepository.FindAllUser()
}

func (u *UserService) DeleteUser(id string) error {
	return u.UserRepository.DeleteUser(id)
}

func (u *UserService) FindUserById(id string) (map[string]any, error) {
	return u.UserRepository.FindUserById(id)
}

func (u *UserService) UpdateUser(id string, user *entity.User) error {
	return u.UserRepository.UpdateUser(id, user)
}

func (u *UserService) SearchUser(query string) (map[string]any, error) {
	return u.UserRepository.SearchUser(query)
}
