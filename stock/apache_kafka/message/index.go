package message

type MessageFromOrder struct {
	OrderID   string `json:"id"`
	ProductID string `json:"product_id"`
	Quantity  int    `json:"quantity"`
}
