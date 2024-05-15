package repository

type Product struct {
	ProductID          uint `gorm:"primarykey";autoIncrement:true`
	ProductName        string
	ProductDescription string
	ProductImage       string
}

type ProductRepository interface {
	CreateProduct(Product) error
	DeleteProduct(int) error
	EditProduct(int, Product) error
	Get(int) (*Product, error)
	GetAll() ([]Product, error)
}
