package repository

import "gorm.io/gorm"

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{db: db}
}

func (r *orderRepository) Done(data Order) error {

	if res := r.db.Save(&Order{OrderID: data.OrderID, OrderDone: 1}); res.Error != nil {
		return res.Error
	}

	return nil
}
