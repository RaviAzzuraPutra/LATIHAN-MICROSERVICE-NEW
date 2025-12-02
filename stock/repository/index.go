package repository

import (
	"stock/database"
	"stock/model"

	"gorm.io/gorm"
)

type StockRepositoryDB struct {
	DB *gorm.DB
}

func NewStockRepositoryDB() *StockRepositoryDB {
	return &StockRepositoryDB{
		DB: database.DB,
	}
}

func (repo *StockRepositoryDB) StockRepositoryCreate(stock *model.Stock) error {
	err := repo.DB.Table("stocks").Create(stock).Error

	return err
}

func (repo *StockRepositoryDB) StockRepositoryGet() ([]model.Stock, error) {

	var stock []model.Stock

	err := repo.DB.Table("stocks").Find(&stock).Error

	return stock, err
}

func (repo *StockRepositoryDB) StockRepositoryGetByID(id string) (*model.Stock, error) {
	var stock model.Stock

	err := repo.DB.Table("stocks").Where("id = ?", id).First(&stock).Error

	return &stock, err
}

func (repo *StockRepositoryDB) StockRepositoryDelete(id string) error {
	var stock *model.Stock

	err := repo.DB.Table("stocks").Unscoped().Where("id = ?", id).Delete(&stock).Error

	return err
}

func (repo *StockRepositoryDB) StockRepositoryUpdate(stock *model.Stock, id string) error {
	err := repo.DB.Table("stocks").Where("id = ?", id).Updates(stock).Error

	return err
}
