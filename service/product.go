package service

type ProductReqAndRes struct {
	ProductID          int    `json:"product_id"`
	ProductDescription string `json:"product_des"`
	ProductName        string `json:"product_name"`
	ProductImage       string `json:"product_image"`
}

type ProductService interface {
	GetProducts() ([]ProductReqAndRes, error)
	GetProduct(int) (*ProductReqAndRes, error)
	EditProduct(int, ProductReqAndRes) error
	DeleteProduct(int) error
	NewProduct(ProductReqAndRes) (*ProductReqAndRes, error)
}

func ProductResInit(id int, des string, name string, image string) ProductReqAndRes {
	return ProductReqAndRes{
		ProductID:          id,
		ProductName:        name,
		ProductDescription: des,
		ProductImage:       image,
	}

}
