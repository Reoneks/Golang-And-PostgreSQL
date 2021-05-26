package product

import "time"

type CommentsDto struct {
	Id        int64     `gorm:"column:id"`
	Text      string    `gorm:"column:text"`
	ProductID int64     `gorm:"column:product_id"`
	CreatedBy int64     `gorm:"column:created_by"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func (CommentsDto) TableName() string {
	return "comments"
}
