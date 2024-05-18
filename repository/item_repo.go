package repository

import (
	"gorm.io/gorm"
)

type itemRepo struct {
	db *gorm.DB
}

func NewItemRepository(db *gorm.DB) ItemRepository {
	return &itemRepo{db: db}
}

func (r *itemRepo) Create(data ItemOrder) error {
	if res := r.db.Create(&data); res.Error != nil {
		return res.Error
	}
	return nil
}

func (r *itemRepo) Edit(id int, data ItemOrder) error {
	if res := r.db.Model(&ItemOrder{}).Where("item_id = ?", id).Updates(data); res.Error != nil {
		return res.Error
	}
	return nil
}

func (r *itemRepo) Delete(id int) error {
	if res := r.db.Delete(&ItemOrder{}, id); res.Error != nil {
		return res.Error
	}
	return nil
}

func (r *itemRepo) Get(id int) (*ItemOrder, error) {
	var ItemOrder ItemOrder
	if res := r.db.First(&ItemOrder, id); res.Error != nil {
		return nil, res.Error
	}

	return &ItemOrder, nil
}

func (r *itemRepo) Gets() ([]ItemOrder, error) {
	var ItemOrder []ItemOrder

	if res := r.db.Find(&ItemOrder); res.Error != nil {
		return nil, nil
	}

	return ItemOrder, nil

}

func (r *itemRepo) GetsItemFromProduct(data ItemOrder) ([]ItemOrder, error) {

	Items := []ItemOrder{}

	if res := r.db.Where(&ItemOrder{ProductId: data.ProductId}).Find(&Items); res.Error != nil {
		return nil, res.Error
	}

	return Items, nil
}
