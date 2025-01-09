package services

import (
	"payment-gateway/domain"
	"payment-gateway/repo"
)

type TransactionService interface {
	CreateTransaction(orderID, customerName, customerEmail string, grossAmount float64) (*domain.Transaction, error)
	UpdateTransactionStatus(orderID, status string) error
	GetTransactions() ([]domain.Transaction, error)
}

type transactionService struct {
	repo repo.TransactionRepository
}

func NewTransactionService(repo repo.TransactionRepository) TransactionService {
	return &transactionService{repo: repo}
}

func (s *transactionService) CreateTransaction(orderID, customerName, customerEmail string, grossAmount float64) (*domain.Transaction, error) {
	transaction := &domain.Transaction{
		OrderID:       orderID,
		CustomerName:  customerName,
		CustomerEmail: customerEmail,
		GrossAmount:   grossAmount,
	}
	err := s.repo.Save(transaction)
	return transaction, err
}

func (s *transactionService) UpdateTransactionStatus(orderID, status string) error {
	transaction, err := s.repo.FindByOrderID(orderID)
	if err != nil {
		return err
	}
	transaction.TransactionStatus = status
	return s.repo.Update(transaction)
}

func (s *transactionService) GetTransactions() ([]domain.Transaction, error) {
	return s.repo.GetAll()
}
