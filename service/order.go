package service

type OrderReq struct {
	OrderID int `json:"order_id"`
	Done    int `json:"order_done"`
}

type OrderRes struct {
	Message string
}

type OrderService interface {
	OrderFail(OrderReq) (*OrderRes, error)
	OrderSuccess(OrderReq) (*OrderRes, error)
}
