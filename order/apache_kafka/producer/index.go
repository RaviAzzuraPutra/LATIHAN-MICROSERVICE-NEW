package producer

import (
	"context"
	"encoding/json"
	"fmt"
	"order/apache_kafka/message"
	"time"

	"github.com/segmentio/kafka-go"
)

type OrderProducerImpl struct {
	Writer *kafka.Writer
}

func NewOrderProducer(brokerURL string, topic string) *OrderProducerImpl {
	write := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{brokerURL},
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	})

	return &OrderProducerImpl{
		Writer: write,
	}
}

func (p *OrderProducerImpl) PublishOrderCreated(ctx context.Context, orderID string, productID string, quantity int) error {
	event := &message.OrderMessage{
		OrderID:   orderID,
		ProductID: productID,
		Quantity:  quantity,
		TimeStamp: time.Now(),
	}

	messageBody, err := json.Marshal(event)

	if err != nil {
		return fmt.Errorf("Terjadi Kesalahan Saat Serialisasi Pesan", err)
	}

	msg := kafka.Message{
		Key:   []byte(orderID),
		Value: messageBody,
	}

	err = p.Writer.WriteMessages(ctx, msg)

	if err != nil {
		return fmt.Errorf("gagal mengirim pesan ke Kafka: %w", err)
	}

	return nil

}
