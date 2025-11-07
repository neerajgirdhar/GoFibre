package services

import (
    "fibre-api/models"
    "fibre-api/repositories"
)

type UserService interface {
    GetAllUsers() []models.User
    CreateUser(user models.User) models.User
}

type userService struct {
    repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
    return &userService{repo: repo}
}

func (s *userService) GetAllUsers() []models.User {
    return s.repo.GetAllUsers()
}

func (s *userService) CreateUser(user models.User) models.User {
    return s.repo.CreateUser(user)
}