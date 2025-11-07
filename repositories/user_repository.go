package repositories

import "fibre-api/models"

// UserRepository defines the interface for user data interactions
type UserRepository interface {
    GetAllUsers() []models.User
    CreateUser(user models.User) models.User
}

type userRepository struct {
    users []models.User
}

func NewUserRepository() UserRepository {
    return &userRepository{
        users: []models.User{},
    }
}

func (r *userRepository) GetAllUsers() []models.User {
    return r.users
}

func (r *userRepository) CreateUser(user models.User) models.User {
    user.ID = len(r.users) + 1
    r.users = append(r.users, user)
    return user
}