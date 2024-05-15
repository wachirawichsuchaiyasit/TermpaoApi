package service

type CustomerRequest struct {
	ID       int    `json:"customer_id"`
	Email    string `json:"customer_email"`
	Password string `json:"customer_password"`
	Username string `json:"customer_username"`
	Cost     int    `json:"customer_cost"`
}

type CustomerResponse struct {
	CustomerID int    `json:"customer_id"`
	Email      string `json:"customer_email"`
	Username   string `json:"customer_username"`
}

type CustomerServiceRepo interface {
	Customer_Create(CustomerRequest) (*CustomerResponse, error)
	Customer_Delete(int) error
	Customer_ChangePassword(CustomerRequest) error
	Customer_AddMoney(CustomerRequest) error
	Customer_Login(CustomerRequest) (*CustomerResponse, error)
	// CustomerComment(CustomerRequest)
}
