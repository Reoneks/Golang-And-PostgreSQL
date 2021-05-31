package user

import (
	gm "gorm.io/gorm"
)

type UserRepository interface {
	GetUser(id int64) (*UserDto, error)
	CreateUser(user UserDto) (*UserDto, error)
	UpdateUser(user UserDto) (*UserDto, error)
	DeleteUser(id int64) error
	GetUsers(where string) ([]UserDto, error)
}

type UserRepositoryImpl struct {
	db *gm.DB
}

func (r *UserRepositoryImpl) GetUser(id int64) (*UserDto, error) {
	user := &UserDto{}
	if err := r.db.Where("id = ?", id).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepositoryImpl) CreateUser(user UserDto) (*UserDto, error) {
	user.Status = 1
	if err := r.db.Create(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepositoryImpl) UpdateUser(user UserDto) (*UserDto, error) {
	if err := r.db.Save(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepositoryImpl) DeleteUser(id int64) error {
	if err := r.db.Delete(&UserDto{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *UserRepositoryImpl) GetUsers(where string) (users []UserDto, err error) {
	var findResult *gm.DB = r.db
	if where != "" {
		findResult = findResult.Where(where)
	}
	if err = findResult.Find(&users).Error; err != nil {
		return nil, err
	}
	return
}

func NewUserRepository(db *gm.DB) UserRepository {
	return &UserRepositoryImpl{
		db,
	}
}
