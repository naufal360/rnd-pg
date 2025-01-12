package services

import (
	"payment-gateway/domain"
	"payment-gateway/dto"
	"payment-gateway/provider"
	"payment-gateway/repo"

	"github.com/google/uuid"
)

type TransactionService interface {
	CreateTransaction(payload dto.TransactionRequest) (*provider.MidtransResponse, error)
	UpdateTransactionStatus(orderID, status string) error
	GetTransactions() ([]domain.Transaction, error)
}

type transactionService struct {
	paymentSvc provider.MidtransInterface
	repo       repo.TransactionRepository
}

func NewTransactionService(repo repo.TransactionRepository, paymentSvc provider.MidtransInterface) TransactionService {
	return &transactionService{
		repo:       repo,
		paymentSvc: paymentSvc,
	}
}

func (s *transactionService) CreateTransaction(payload dto.TransactionRequest) (*provider.MidtransResponse, error) {
	orderID := uuid.New().String()
	req := provider.MidtransRequest{
		TransactionDetails: provider.TransactionDetails{
			OrderID:     orderID,
			GrossAmount: int(payload.GrossAmount),
		},
		CustomerDetails: &provider.CustomerDetails{
			FirstName: payload.CustomerFirstName,
			LastName:  payload.CustomerLastName,
			Email:     payload.CustomerEmail,
		},
	}
	resp, transactionTime, err := s.paymentSvc.SendPayment(req)
	if err != nil {
		return nil, err
	}

	transaction := &domain.Transaction{
		OrderID:           orderID,
		CustomerFirstName: req.CustomerDetails.FirstName,
		CustomerLastName:  req.CustomerDetails.LastName,
		CustomerEmail:     req.CustomerDetails.Email,
		GrossAmount:       float64(req.TransactionDetails.GrossAmount),
		PaymentURL:        resp.PaymentURL,
		TransactionTime:   transactionTime,
	}
	err = s.repo.Save(transaction)
	if err != nil {
		return nil, err
	}
	return &resp, err
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
