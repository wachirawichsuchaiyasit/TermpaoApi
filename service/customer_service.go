package service

import (
	"errors"

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

func (s *customerService) Customer_GetUser(id int) (*CustomerResponse, error) {

	customer, err := s.repo.GetUser(repository.Customer{
		ID: uint(id),
	})

	if err != nil {
		return nil, err
	}

	return &CustomerResponse{
		CustomerID: int(customer.ID),
		Email:      customer.Email,
		Username:   customer.Username,
		Cost:       int(customer.Cost),
	}, nil
}

func (s *customerService) Customer_AddMoney(data CustomerRequest) error {

	customer, err := s.Customer_GetUser(data.ID)

	if err != nil {
		return err
	}

	data.Cost += customer.Cost
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
	// fmt.Println(customer)
	if check := helps.CheckHashPassword(data.Password, customer.Password); !check {
		return nil, errors.New("password not match")
	}

	return &CustomerResponse{
		CustomerID: int(customer.ID),
		Email:      customer.Email,
		Username:   customer.Username,
	}, nil

}

func (s *customerService) Customer_BuyItem(data CustomerItemReq) error {

	customer, item, err := s.repo.GetDataItemAndUser(repository.Customer{
		ID: uint(data.CustomerID),
	}, data.ItemID)

	if err != nil {
		return err
	}

	if customer.Cost < int32(item.ItemPrice) {
		return errors.New("not enough money")
	}

	err = s.repo.AddOrder(repository.Order{
		OrderPrice:      data.ItemPrice,
		OrderUid:        data.ItemUid,
		OrderCustomerID: data.CustomerID,
		OrderProductID:  data.ProductID,
	})

	if err != nil {
		return err
	}

	return nil
}

func (s *customerService) Customer_GetData(data CustomerRequest) (*CustomerResponse, error) {

	customer, err := s.repo.GetUser(repository.Customer{
		Email: data.Email,
	})

	if err != nil {
		return nil, nil
	}

	return &CustomerResponse{
		CustomerID: int(customer.ID),
		Email:      customer.Email,
		Username:   customer.Username,
		Cost:       int(customer.Cost),
	}, nil
}
