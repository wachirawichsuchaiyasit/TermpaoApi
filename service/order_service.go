package service

import "github.com/Termpao/repository"

type orderService struct {
	repo repository.OrderRepository
}

func NewOrderService(repo repository.OrderRepository) OrderService {
	return &orderService{repo: repo}
}

func (s *orderService) OrderSuccess(data OrderReq) (*OrderRes, error) {

	err := s.repo.Done(repository.Order{
		OrderID:   data.OrderID,
		OrderDone: data.Done,
	})

	if err != nil {
		return nil, err
	}

	return &OrderRes{
		Message: "OrderSuccess",
	}, nil
}

func (s *orderService) OrderFail(data OrderReq) (*OrderRes, error) {

	err := s.repo.Done(repository.Order{
		OrderID:   data.OrderID,
		OrderDone: data.Done,
	})

	if err != nil {
		return nil, err
	}

	return &OrderRes{
		Message: "Order Fail",
	}, nil
}
