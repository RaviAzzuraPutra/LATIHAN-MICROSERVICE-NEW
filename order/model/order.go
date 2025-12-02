package model

type Order struct {
	Id        *string `json:"id" gorm:"column:id;primaryKey;type:uuid;default:gen_random_uuid()"`
	ProductID *string `json:"productid"`
	Quantity  *int    `json:"quantity"`
}
