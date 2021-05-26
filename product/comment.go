package product

import "time"

type Comments struct {
	Id        int64     `json:"id"`
	Text      string    `json:"text"`
	ProductID int64     `json:"product_id"`
	CreatedBy int64     `json:"created_by"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
