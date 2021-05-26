package config

import (
	"github.com/jinzhu/gorm"
	"sync"
)

type Config interface {
	DBClient() *gorm.DB
	JWT() *jwtauth.JWTAuth
}

type ConfigImpl struct {
	sync.Mutex

	//internal objects
	dbClient *gorm.DB
	jwt      *jwtauth.JWTAuth
}

func NewConfig() Config {
	return &ConfigImpl{
		Mutex: sync.Mutex{},
	}
}
