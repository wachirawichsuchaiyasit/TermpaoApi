package service

type ItemRes struct {
	ItemID          uint   `json:"item_id"`
	ItemName        string `json:"item_name"`
	ItemPrice       int    `json:"item_price"`
	ItemDescription string `json:"item_description"`
	ItemProdctId    int    `json:"product_id"`
}

type ItemService interface {
	GetItems() ([]ItemRes, error)
	GetItem(int) (*ItemRes, error)
	CreateItem(ItemRes) error
	EditItem(int, ItemRes) error
	DeleteItem(int) error
}
