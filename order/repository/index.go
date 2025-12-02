package repository

import (
	"order/database"
	"order/model"

	"gorm.io/gorm"
)

type OrderRepositoryDB struct {
	DB *gorm.DB
}

func NewReositoryOrderDB() *OrderRepositoryDB {
	return &OrderRepositoryDB{
		DB: database.DB,
	}
}

func (repo *OrderRepositoryDB) GetOrderRepository() ([]model.Order, error) {
	var order []model.Order

	err := repo.DB.Table("orders").Find(&order).Error

	return order, err
}

func (repo *OrderRepositoryDB) CreateOrderRepository(order *model.Order) error {
	err := repo.DB.Table("orders").Create(&order).Error

	return err
}
