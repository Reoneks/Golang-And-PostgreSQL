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

type Comments struct {
	Id        int64     `json:"id"`
	Text      string    `json:"text"`
	ProductID int64     `json:"product_id"`
	CreatedBy int64     `json:"created_by"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FromCommentsDto(commentsDto CommentsDto) Comments {
	return Comments(commentsDto)
}

func FromCommentsDtos(commentsDtos []CommentsDto) (comments []Comments) {
	for _, dto := range commentsDtos {
		comments = append(comments, Comments(dto))
	}
	return
}

func ToCommentsDto(comments Comments) CommentsDto {
	return CommentsDto(comments)
}

func ToCommentsDtos(comments []Comments) (commentsDto []CommentsDto) {
	for _, dto := range comments {
		commentsDto = append(commentsDto, CommentsDto(dto))
	}
	return
}

func (CommentsDto) TableName() string {
	return "comments"
}
