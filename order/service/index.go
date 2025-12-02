package service

import (
	"context"
	"errors"
	"fmt"
	kafka_interface "order/apache_kafka/interface"
	"order/model"
	repository_interface "order/repository/interface"
	"order/request"
)

type ServiceOrderImpl struct {
	Repo      repository_interface.RepositoryOrderInterface
	Publisher kafka_interface.OrderPublisherInterface
}

func NewServiceOrder(repo repository_interface.RepositoryOrderInterface, publisher kafka_interface.OrderPublisherInterface) *ServiceOrderImpl {
	return &ServiceOrderImpl{
		Repo:      repo,
		Publisher: publisher,
	}
}

func (service *ServiceOrderImpl) GetService() ([]model.Order, error) {
	order, err := service.Repo.GetOrderRepository()

	if len(order) == 0 {
		return nil, errors.New("DATA MASIH KOSONG")
	}

	if err != nil {
		return nil, errors.New("TERJADI KESALAHAN SAAT MENGAMBIL DATA : " + err.Error())
	}

	return order, err
}

func (service *ServiceOrderImpl) CreateService(request *request.OrderRequest) (*model.Order, error) {
	if request.ProductID == nil || *request.ProductID == "" {
		return nil, errors.New("PRODUCT ID TIDAK BOLEH KOSONG")
	}

	if request.Quantity == nil || *request.Quantity < 0 {
		return nil, errors.New("QUANTITY TIDAK BOLEH KOSONG")
	}

	var order = &model.Order{
		ProductID: request.ProductID,
		Quantity:  request.Quantity,
	}

	err := service.Repo.CreateOrderRepository(order)

	if err != nil {
		return nil, errors.New("TERJADI KESALAHAN SAAT MENAMBAHKAN DATA : " + err.Error())
	}

	ctx := context.Background()

	errPublish := service.Publisher.PublishOrderCreated(ctx, *order.Id, *order.ProductID, *order.Quantity)

	if errPublish != nil {
		fmt.Printf("Warning: Gagal kirim ke Kafka: %v\n", errPublish)
	}

	return order, err
}
