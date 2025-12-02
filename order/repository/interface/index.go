package repository_interface

import "order/model"

type RepositoryOrderInterface interface {
	GetOrderRepository() ([]model.Order, error)
	CreateOrderRepository(order *model.Order) error
}
