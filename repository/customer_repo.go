package repository

import (
	"fmt"
	"log"

	"gorm.io/gorm"
)

type customerRepo struct {
	db *gorm.DB
}

func NewCustomerDatabase(db *gorm.DB) CustomerRepository {
	return &customerRepo{db: db}
}

func (r *customerRepo) CreateUser(data Customer) error {
	resault := r.db.Create(&data)

	if resault.Error != nil {
		log.Fatalln("Create User Error", resault.Error)
		fmt.Println("errorssssssss")
		return resault.Error
	}
	return nil
}

func (r *customerRepo) DeleteUser(data int) error {
	res := r.db.Delete(&Customer{}, data)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (r *customerRepo) EditUser(id int, data Customer) error {
	res := r.db.Model(&Customer{}).Where("id = ?", id).Updates(data)

	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (r *customerRepo) AddCostUser(id int, money int) error {
	res := r.db.Model(&Customer{}).Where("id = ?", id).Update("const", money)

	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (r *customerRepo) GetUser(data Customer) (*Customer, error) {
	customer := Customer{Email: data.Email}
	res := r.db.First(&customer)

	if res.Error != nil {
		return nil, res.Error
	}

	return &customer, nil
}

func (r *customerRepo) ChangePassword(data Customer) error {
	if err := r.EditUser(int(data.ID), data); err != nil {
		return err
	}
	return nil
}
