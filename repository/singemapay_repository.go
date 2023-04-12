package repository

import (
	"fmt"
	"github.com/jutionck/grpc-sinegmapay-server/model"
)

type SinegmaPayRepository interface {
	GetBy(id int32) (model.Customer, error)
}

type sinegmaPayRepository struct {
	db []model.Customer
}

func (s *sinegmaPayRepository) GetBy(id int32) (model.Customer, error) {
	for _, v := range s.db {
		if v.Id == id {
			return v, nil
		}
		return model.Customer{}, fmt.Errorf("customer with ID: %d not found", id)
	}
	return model.Customer{}, nil
}

func NewSinegmaPayRepository() SinegmaPayRepository {
	repo := new(sinegmaPayRepository)
	repo.db = []model.Customer{
		{Id: 1, Balance: 100000},
		{Id: 2, Balance: 50000},
		{Id: 3, Balance: 2500},
	}
	return repo
}
