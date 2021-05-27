package config

import (
	"net/url"
	"sync"

	"github.com/go-chi/jwtauth"
	"gorm.io/gorm"
)

type Config interface {
	DBClient() *gorm.DB
	JWT() *jwtauth.JWTAuth
	ServerAddress() *url.URL
}

type ConfigImpl struct {
	sync.Mutex

	//internal objects
	dbClient *gorm.DB
	jwt      *jwtauth.JWTAuth
	url      *url.URL
}

func NewConfig() Config {
	return &ConfigImpl{
		Mutex: sync.Mutex{},
	}
}
