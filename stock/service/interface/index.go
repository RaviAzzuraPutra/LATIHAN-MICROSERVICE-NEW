package service_interface

import (
	"stock/model"
	"stock/request"
)

type StockServiceInterface interface {
	CreateService(request *request.StockRequest) (*model.Stock, error)
	GetService() ([]model.Stock, error)
	GetByIDService(id string) (*model.Stock, error)
	DeleteService(id string) error
	UpdateService(request *request.StockRequest, id string) (*model.Stock, error)
	HapusStockBerdasarkanOrder(productID string, quantity int) error
}
