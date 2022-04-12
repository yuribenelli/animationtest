package configuration

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	DB *gorm.DB
}

var currentconfig Config

func InitConfig() {
	dsn := "host=localhost user=animationtest password=Mindsrl dbname=animationtest port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	currentconfig = Config{DB: db}
}

func Current() Config {
	return currentconfig
}
