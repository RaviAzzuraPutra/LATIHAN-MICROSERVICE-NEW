package consumer

import (
	"context"
	"encoding/json"
	"log"
	"stock/apache_kafka/message"
	"stock/config/kafka_config"
	service_interface "stock/service/interface"
	"time"

	"github.com/segmentio/kafka-go"
)

type StockConsumer struct {
	Reader  *kafka.Reader
	Service service_interface.StockServiceInterface
}

func NewStockConsumer(brokerURL string, topic string, service service_interface.StockServiceInterface) *StockConsumer {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:        []string{brokerURL},
		Topic:          topic,
		GroupID:        kafka_config.GroupID,
		MinBytes:       10e3,
		MaxBytes:       10e6,
		QueueCapacity:  1,
		CommitInterval: time.Second,
	})

	return &StockConsumer{
		Reader:  r,
		Service: service,
	}
}

func (c *StockConsumer) StartDataConsumer(ctx context.Context) {
	for {
		m, err := c.Reader.ReadMessage(ctx)

		if err != nil {
			log.Printf("⚠️ Gagal membaca pesan (Kafka belum siap / koneksi gagal): %v", err)

			time.Sleep(10 * time.Second)
			continue
		}

		var msg message.MessageFromOrder
		errJson := json.Unmarshal(m.Value, &msg)

		if errJson != nil {
			log.Printf("GAGAL MEMPARSING JSON", errJson)
			continue
		}

		errService := c.Service.HapusStockBerdasarkanOrder(msg.ProductID, msg.Quantity)

		if errService != nil {
			log.Printf("❌ Gagal update stok: %v\n", errService)
		} else {
			log.Printf("BERHASIL MELAKUKAN UPDATE STOCK DAN KAFKA BERHASIL")
		}
	}
}
