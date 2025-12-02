package message

import "time"

type OrderMessage struct {
	OrderID   string `json:"id"`
	ProductID string `json:"product_id"`
	Quantity  int    `json:"quantity"`
	TimeStamp time.Time
}
