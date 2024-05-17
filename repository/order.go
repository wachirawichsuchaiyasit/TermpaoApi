package repository

import "time"

type Order struct {
	OrderID         int `gorm:"primarykey";autoincrement:true`
	OrderPrice      int
	OrderUid        int
	OrderCustomerID int
	OrderProductID  int
	OrderTime       time.Time
	OrderDone       int `gorm:"default:0"`
}

type OrderRepository interface {
	Done(Order) error
}
