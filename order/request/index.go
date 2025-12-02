package request

type OrderRequest struct {
	ProductID *string `json:"productid" binding:"required"`
	Quantity  *int    `json:"quantity" binding:"required"`
}
