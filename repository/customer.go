package repository

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	Username string
	Password string
	Email    string
	Verify   int32 `gorm:"default:0"`
	Const    int32 `gorm:"default:0"`
	Admin    int32 `gorm:"default:0"`
}

type CustomerRepository interface {
	CreateUser(Customer) error
	EditUser(int, Customer) error
	DeleteUser(int) error
	AddCostUser(int, int) error
	GetUser(Customer) (*Customer, error)
	ChangePassword(Customer) error
}
