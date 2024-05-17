package repository

type Customer struct {
	ID       uint `gorm:"primarykey";"autoincrement:true"`
	Username string
	Password string
	Email    string
	Verify   int32 `gorm:"default:0"`
	Cost     int32 `gorm:"default:0"`
	Admin    int32 `gorm:"default:0"`
}

type CustomerRepository interface {
	CreateUser(Customer) error
	EditUser(int, Customer) error
	DeleteUser(int) error
	AddCostUser(int, int) error
	GetUser(Customer) (*Customer, error)
	ChangePassword(Customer) error
	GetDataItemAndUser(Customer, int) (*Customer, *ItemOrder, error)
	AddOrder(Order) error
}
