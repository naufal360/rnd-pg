package services

import (
	"payment-gateway/domain"
	"payment-gateway/dto"
	"payment-gateway/repo"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	CreateUser(dto dto.CreateUserDTO) (*domain.User, error)
	GetUsers() ([]domain.User, error)
}

type userService struct {
	repo repo.UserRepository
}

func NewUserService(repo repo.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) CreateUser(data dto.CreateUserDTO) (*domain.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &domain.User{
		Name:     data.Name,
		Email:    data.Email,
		Password: string(hashedPassword),
	}
	err = s.repo.Save(user)
	return user, err
}

func (s *userService) GetUsers() ([]domain.User, error) {
	return s.repo.GetAll()
}
