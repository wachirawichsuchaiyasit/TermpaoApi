package service

import (
	"fmt"

	"github.com/Termpao/repository"
)

type productService struct {
	repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) ProductService {
	return &productService{repo: repo}
}

func (s *productService) NewProduct(data ProductReqAndRes) (*ProductReqAndRes, error) {
	product := repository.Product{
		ProductName:        data.ProductName,
		ProductImage:       data.ProductImage,
		ProductDescription: data.ProductDescription,
	}
	if err := s.repo.CreateProduct(product); err != nil {
		return nil, nil
	}
	fmt.Println("This is Product", product)
	return &ProductReqAndRes{
		ProductID:          int(product.ProductID),
		ProductName:        product.ProductName,
		ProductDescription: data.ProductDescription,
		ProductImage:       data.ProductImage,
	}, nil
}

func (s *productService) DeleteProduct(id int) error {
	if err := s.repo.DeleteProduct(id); err != nil {
		return err
	}
	return nil
}

func (s *productService) EditProduct(id int, data ProductReqAndRes) error {
	product := repository.Product{
		ProductName:        data.ProductName,
		ProductImage:       data.ProductImage,
		ProductDescription: data.ProductDescription,
	}

	if err := s.repo.EditProduct(id, product); err != nil {
		return nil
	}
	return nil
}

func (s *productService) GetProduct(id int) (*ProductReqAndRes, error) {
	product, err := s.repo.Get(id)

	if err != nil {
		return nil, err
	}

	return &ProductReqAndRes{ProductID: int(product.ProductID), ProductDescription: product.ProductDescription, ProductName: product.ProductName, ProductImage: product.ProductName}, nil
}

func (s *productService) GetProducts() ([]ProductReqAndRes, error) {
	products, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	var newproducts []ProductReqAndRes

	for _, product := range products {
		newproducts = append(newproducts, ProductReqAndRes{
			ProductID:          int(product.ProductID),
			ProductName:        product.ProductName,
			ProductDescription: product.ProductDescription,
			ProductImage:       product.ProductImage,
		})
	}
	return newproducts, nil
}
