package repository_interface

import "stock/model"

type StockRepositoryInterface interface {
	StockRepositoryCreate(stock *model.Stock) error
	StockRepositoryGet() ([]model.Stock, error)
	StockRepositoryGetByID(id string) (*model.Stock, error)
	StockRepositoryDelete(id string) error
	StockRepositoryUpdate(stock *model.Stock, id string) error
}
