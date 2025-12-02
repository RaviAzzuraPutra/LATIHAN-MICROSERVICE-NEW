package kafka_interface

import "context"

type OrderPublisherInterface interface {
	PublishOrderCreated(ctx context.Context, orderID string, productID string, quantity int) error
}
