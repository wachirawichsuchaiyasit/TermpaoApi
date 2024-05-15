package repository

import (
	"gorm.io/gorm"
)

type productRepo struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepo{db: db}
}

func (r *productRepo) CreateProduct(data Product) error {
	if res := r.db.Create(&data); res.Error != nil {
		return res.Error
	}
	return nil
}

func (r *productRepo) EditProduct(id int, data Product) error {
	if res := r.db.Model(&Product{}).Where("product_id = ?", id).Updates(data); res.Error != nil {
		return res.Error
	}
	return nil
}

func (r *productRepo) DeleteProduct(id int) error {
	if res := r.db.Delete(&Product{}, id); res.Error != nil {
		return res.Error
	}
	return nil
}

func (r *productRepo) Get(id int) (*Product, error) {
	var product Product
	if res := r.db.First(&product, id); res.Error != nil {
		return nil, res.Error
	}

	return &product, nil
}

func (r *productRepo) GetAll() ([]Product, error) {
	var product []Product

	if res := r.db.Find(&product); res.Error != nil {
		return nil, nil
	}

	return product, nil

}
