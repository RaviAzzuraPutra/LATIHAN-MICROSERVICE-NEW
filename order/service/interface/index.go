package service_interfcae

import (
	"order/model"
	"order/request"
)

type OrderServiceInterface interface {
	GetService() ([]model.Order, error)
	CreateService(request *request.OrderRequest) (*model.Order, error)
}
