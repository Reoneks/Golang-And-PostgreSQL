package config

import (
	"github.com/caarlos0/env"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type DBConfig struct {
	DbUrl string `env:"DB_URL"`
}

func (c *ConfigImpl) DBClient() *gorm.DB {
	if c.dbClient != nil {
		return c.dbClient
	}

	c.Lock()
	defer c.Unlock()

	dbConfig := &DBConfig{}
	if err := env.Parse(dbConfig); err != nil {
		panic(err)
	}

	client, err := gorm.Open("postgres", dbConfig.DbUrl)
	if err != nil {
		panic(err)
	}

	c.dbClient = client

	return client
}
