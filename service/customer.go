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
	Cost       int    `json:"customer_cost"`
}

type CustomerResonseItem struct {
	CustomerID int `json:"customer_id"`
	ItemID     int `json:"item_id"`
	ItemPrice  int `json:"item_price"`
}

type CustomerWalletReq struct {
	CustomerID   int    `json:"customer_id"`
	CustomerLink string `json:"customer_wallet"`
}

type CustomerWalletRes struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Amount  string `json:"amount"`
	Phone   string `json:"phone"`
	Link    string `json:"gift_link"`
	Time    string `json:"time"`
}

type CustomerItemReq struct {
	CustomerID int `json:"customer_id"`
	ItemID     int `json:"item_id"`
	ItemPrice  int `json:"item_price"`
	ItemUid    int `json:"item_uid"`
	ProductID  int `json:"item_product_id"`
}

type CustomerServiceRepo interface {
	Customer_Create(CustomerRequest) (*CustomerResponse, error)
	Customer_Delete(int) error
	Customer_ChangePassword(CustomerRequest) error
	Customer_AddMoney(CustomerRequest) error
	Customer_Login(CustomerRequest) (*CustomerResponse, error)
	Customer_BuyItem(CustomerItemReq) error
	Customer_GetData(CustomerRequest) (*CustomerResponse, error)
	// CustomerComment(CustomerRequest)
}
