package product

import (
	"time"

	gm "gorm.io/gorm"
)

type CommentsRepository interface {
	CreateComment(user CommentsDto) (*CommentsDto, error)
	UpdateComment(user CommentsDto) (*CommentsDto, error)
	DeleteComment(id int64) error
}

type CommentsRepositoryImpl struct {
	db *gm.DB
}

func (r *CommentsRepositoryImpl) CreateComment(comment CommentsDto) (*CommentsDto, error) {
	comment.CreatedAt = time.Now()
	comment.UpdatedAt = time.Now()
	if err := r.db.Create(&comment).Error; err != nil {
		return nil, err
	}
	return &comment, nil
}

func (r *CommentsRepositoryImpl) UpdateComment(comment CommentsDto) (*CommentsDto, error) {
	var comments *CommentsDto
	if err := r.db.Where("id = ?", comment.Id).First(&comments).Error; err != nil {
		return nil, err
	}
	comment.CreatedAt = comments.CreatedAt
	comment.UpdatedAt = time.Now()
	if err := r.db.Save(&comment).Error; err != nil {
		return nil, err
	}
	return &comment, nil
}

func (r *CommentsRepositoryImpl) DeleteComment(id int64) error {
	if err := r.db.Delete(&CommentsDto{}, id).Error; err != nil {
		return err
	}
	return nil
}

func NewCommentsRepository(db *gm.DB) CommentsRepository {
	return &CommentsRepositoryImpl{
		db,
	}
}
