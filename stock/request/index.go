package request

type StockRequest struct {
	Name     *string `json:"name" binding:"required"`
	Quantity *int    `json:"quantity" binding:"required"`
}
