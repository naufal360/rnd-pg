package repo

import (
	"payment-gateway/domain"

	"gorm.io/gorm"
)

type UserRepository interface {
	Save(user *domain.User) error
	GetAll() ([]domain.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Save(user *domain.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) GetAll() ([]domain.User, error) {
	var users []domain.User
	err := r.db.Find(&users).Error
	return users, err
}
