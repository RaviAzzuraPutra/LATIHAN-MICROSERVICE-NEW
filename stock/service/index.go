package service

import (
	"errors"
	"stock/model"
	repository_interface "stock/repository/interface"
	"stock/request"
)

type StockServiceImpl struct {
	Repo repository_interface.StockRepositoryInterface
}

func NewStockService(repo repository_interface.StockRepositoryInterface) *StockServiceImpl {
	return &StockServiceImpl{
		Repo: repo,
	}
}

func (s *StockServiceImpl) CreateService(request *request.StockRequest) error {

	if request.Name == nil {
		return errors.New("nama barang tidak boleh kosong")
	}

	if request.Quantity == nil {
		return errors.New("quantity tidak boleh kosong")
	}

	stock := &model.Stock{
		Name:     request.Name,
		Quantity: request.Quantity,
	}

	return s.Repo.StockRepositoryCreate(stock)
}

func (s *StockServiceImpl) GetService() ([]model.Stock, error) {

	stock, err := s.Repo.StockRepositoryGet()

	if err != nil {
		return nil, err
	}

	if len(stock) == 0 {
		return nil, errors.New("data stok masih kosong")
	}

	return stock, err
}

func (s *StockServiceImpl) GetByIDService(id string) (*model.Stock, error) {
	stock, err := s.Repo.StockRepositoryGetByID(id)

	if err != nil {
		return nil, err
	}

	if stock == nil {
		return nil, errors.New("data stok tidak ditemukan")
	}

	return stock, err
}

func (s *StockServiceImpl) DeleteService(id string) error {

	stock, err := s.Repo.StockRepositoryGetByID(id)

	if err != nil {
		return err
	}

	if stock == nil {
		return nil
	}

	errDelete := s.Repo.StockRepositoryDelete(id)

	if errDelete != nil {
		return errors.New("TERJADI KESALAHAN SAAT DELETE")
	}

	return nil
}

func (s *StockServiceImpl) UpdateService(request *request.StockRequest, id string) error {
	findStock, err := s.Repo.StockRepositoryGetByID(id)

	if err != nil {
		return err
	}

	if findStock == nil {
		return errors.New("data tidak ditemukan")
	}

	if request.Name == nil {
		return errors.New("nama barang tidak boleh kosong")
	}

	if request.Quantity == nil {
		return errors.New("quantity tidak boleh kosong")
	}

	stock := &model.Stock{
		Name:     request.Name,
		Quantity: request.Quantity,
	}

	errUpdate := s.Repo.StockRepositoryUpdate(stock, id)

	if errUpdate != nil {
		return errUpdate
	}

	return errUpdate
}

func (s *StockServiceImpl) HapusStockBerdasarkanOrder(productID string, quantity int) error {
	stock, err := s.Repo.StockRepositoryGetByID(productID)

	if err != nil {
		return err
	}

	if stock == nil {
		return errors.New("DATA STOCK TIDAK ADA")
	}

	currentQuantity := *stock.Quantity

	newQuantity := currentQuantity - quantity

	stock.Quantity = &newQuantity

	errUpdate := s.Repo.StockRepositoryUpdate(stock, productID)

	if errUpdate != nil {
		return errUpdate
	}

	return nil
}
