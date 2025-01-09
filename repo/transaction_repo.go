package repo

import (
	"payment-gateway/domain"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	Save(transaction *domain.Transaction) error
	Update(transaction *domain.Transaction) error
	FindByOrderID(orderID string) (*domain.Transaction, error)
	GetAll() ([]domain.Transaction, error)
}

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepository{db: db}
}

func (r *transactionRepository) Save(transaction *domain.Transaction) error {
	return r.db.Create(transaction).Error
}

func (r *transactionRepository) Update(transaction *domain.Transaction) error {
	return r.db.Save(transaction).Error
}

func (r *transactionRepository) FindByOrderID(orderID string) (*domain.Transaction, error) {
	var transaction domain.Transaction
	err := r.db.Where("order_id = ?", orderID).First(&transaction).Error
	return &transaction, err
}

func (r *transactionRepository) GetAll() ([]domain.Transaction, error) {
	var transactions []domain.Transaction
	err := r.db.Find(&transactions).Error
	return transactions, err
}
