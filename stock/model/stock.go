package model

type Stock struct {
	Id       *string `json:"id" gorm:"column:id;primaryKey;type:uuid;default:gen_random_uuid()"`
	Name     *string `json:"name"`
	Quantity *int    `json:"quantity"`
}
