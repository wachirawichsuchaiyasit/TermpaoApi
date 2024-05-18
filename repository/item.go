package repository

type ItemOrder struct {
	ItemID          uint `gorm:"primarykey";autoIncrement:true`
	ItemName        string
	ItemPrice       int
	ItemDescription string
	ProductId       int
}

type ItemRepository interface {
	Create(ItemOrder) error
	Edit(int, ItemOrder) error
	Delete(int) error
	Get(int) (*ItemOrder, error)
	Gets() ([]ItemOrder, error)
	GetsItemFromProduct(ItemOrder) ([]ItemOrder, error)
}
