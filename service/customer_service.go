package service

import (
	"github.com/Termpao/helps"
	"github.com/Termpao/repository"
)

type customerService struct {
	repo repository.CustomerRepository
}

func NewCustomerService(repo repository.CustomerRepository) CustomerServiceRepo {
	return &customerService{repo: repo}
}

func (s *customerService) Customer_Create(data CustomerRequest) (*CustomerResponse, error) {
	hashpassword, _ := helps.HashPassword(data.Password)
	customerReq := repository.Customer{
		Email:    data.Email,
		Password: hashpassword,
		Username: data.Username,
	}
	err := s.repo.CreateUser(customerReq)

	if err != nil {
		return nil, nil
	}

	return &CustomerResponse{
		CustomerID: int(customerReq.ID),
		Email:      customerReq.Email,
	}, nil
}

func (s *customerService) Customer_Delete(id int) error {

	err := s.repo.DeleteUser(id)

	if err != nil {
		return err
	}

	return nil
}

func (s *customerService) Customer_ChangePassword(data CustomerRequest) error {
	hashpassword, _ := helps.HashPassword(data.Password)
	err := s.repo.EditUser(data.ID, repository.Customer{Password: hashpassword})
	if err != nil {
		return err
	}

	return nil
}

func (s *customerService) Customer_AddMoney(data CustomerRequest) error {
	if err := s.repo.AddCostUser(data.ID, data.Cost); err != nil {
		return err
	}

	return nil
}

func (s *customerService) Customer_Login(data CustomerRequest) (*CustomerResponse, error) {
	customer, err := s.repo.GetUser(repository.Customer{Email: data.Email})

	if err != nil {
		return nil, err
	}

	if check := helps.CheckHashPassword(data.Password, customer.Password); !check {
		return nil, nil
	}

	return &CustomerResponse{
		CustomerID: int(customer.ID),
		Email:      customer.Email,
		Username:   customer.Username,
	}, nil

}
