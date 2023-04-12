package service

import (
	"context"
	"encoding/json"
	"github.com/jutionck/grpc-sinegmapay-server/repository"
)

type SinegmaPayService struct {
	repo repository.SinegmaPayRepository
	UnimplementedSinegmaPaymentServer
}

// ambil function yang sudah ke generate -> sinegmapay_grpc.pb.go
func (s *SinegmaPayService) CheckBalance(ctx context.Context, in *CheckBalanceMessage) (*ResultMessage, error) {
	customer, err := s.repo.GetBy(in.CustomerId)
	if err != nil {
		return nil, err
	}
	custMarshal, err := json.Marshal(customer)
	if err != nil {
		return nil, err
	}
	resultMessage := &ResultMessage{
		Result: string(custMarshal),
		Error:  nil,
	}
	return resultMessage, nil
}

func (s *SinegmaPayService) DoPayment(ctx context.Context, in *PaymentMessage) (*ResultMessage, error) {
	customer, err := s.repo.GetBy(in.CustomerId)
	if err != nil {
		return nil, err
	}
	errResult := &ResultMessage{
		Result: "Failed",
		Error: &Error{
			Code:    "400",
			Message: "Insufficient Balance",
		},
	}
	if customer.Balance < in.Amount {
		return errResult, nil
	}

	successResult := &ResultMessage{
		Result: "SUCCESS",
		Error:  nil,
	}
	return successResult, nil
}

func NewSinegmaPayService(repo repository.SinegmaPayRepository) *SinegmaPayService {
	service := new(SinegmaPayService)
	service.repo = repo
	return service
}
