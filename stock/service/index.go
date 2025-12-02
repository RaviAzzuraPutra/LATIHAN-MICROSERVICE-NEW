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

func (s *StockServiceImpl) CreateService(request *request.StockRequest) (*model.Stock, error) {

	if request.Name == nil || *request.Name == "" {
		return nil, errors.New("nama barang tidak boleh kosong")
	}

	if request.Quantity == nil || *request.Quantity < 0 {
		return nil, errors.New("quantity tidak boleh kosong")
	}

	stock := &model.Stock{
		Name:     request.Name,
		Quantity: request.Quantity,
	}

	errCreate := s.Repo.StockRepositoryCreate(stock)

	if errCreate != nil {
		return nil, errors.New("TERJADI KESALAHAN SAAT MENAMBAHKAN DATA" + errCreate.Error())
	}

	return stock, errCreate
}

func (s *StockServiceImpl) GetService() ([]model.Stock, error) {

	stock, err := s.Repo.StockRepositoryGet()

	if err != nil {
		return nil, errors.New("TERJADI KESALAHAN SAAT MENGAMBIL DATA" + err.Error())
	}

	if len(stock) == 0 {
		return nil, errors.New("data stok masih kosong")
	}

	return stock, err
}

func (s *StockServiceImpl) GetByIDService(id string) (*model.Stock, error) {
	stock, err := s.Repo.StockRepositoryGetByID(id)

	if err != nil {
		return nil, errors.New("TERJADI KESALAHAN SAAT MENCARI DATA: " + err.Error())
	}

	if stock == nil {
		return nil, errors.New("data stok tidak ditemukan")
	}

	return stock, err
}

func (s *StockServiceImpl) DeleteService(id string) error {

	stock, err := s.Repo.StockRepositoryGetByID(id)

	if err != nil {
		return errors.New("TERJADI KESALAHAN SAAT MENCARI DATA: " + err.Error())
	}

	if stock == nil {
		return errors.New("DATA STOCK KOSONG")
	}

	errDelete := s.Repo.StockRepositoryDelete(id)

	if errDelete != nil {
		return errors.New("TERJADI KESALAHAN SAAT DELETE" + errDelete.Error())
	}

	return nil
}

func (s *StockServiceImpl) UpdateService(request *request.StockRequest, id string) (*model.Stock, error) {
	findStock, err := s.Repo.StockRepositoryGetByID(id)

	if err != nil {
		return nil, errors.New("Terjadi Kesalahan Saat Mengambil Data: " + err.Error())
	}

	if findStock == nil {
		return nil, errors.New("data tidak ditemukan")
	}

	if request.Name == nil || *request.Name == "" {
		return nil, errors.New("nama barang tidak boleh kosong")
	}

	if request.Quantity == nil || *request.Quantity < 0 {
		return nil, errors.New("quantity tidak boleh kosong")
	}

	stock := &model.Stock{
		Name:     request.Name,
		Quantity: request.Quantity,
	}

	errUpdate := s.Repo.StockRepositoryUpdate(stock, id)

	if errUpdate != nil {
		return nil, errUpdate
	}

	return stock, errUpdate
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
