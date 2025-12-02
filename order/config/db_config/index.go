package db_config

import "os"

type DBConfig struct {
	DB_HOST     string
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string
	DB_PORT     string
	DB_SSL      string
	DB_TIMEZONE string
}

func DB_CONFIG() *DBConfig {
	return &DBConfig{
		DB_HOST:     os.Getenv("ORDER_HOST"),
		DB_USER:     os.Getenv("ORDER_USER"),
		DB_PASSWORD: os.Getenv("ORDER_PASSWORD"),
		DB_NAME:     os.Getenv("ORDER_NAME"),
		DB_PORT:     os.Getenv("ORDER_PORT"),
		DB_SSL:      os.Getenv("ORDER_SSL"),
		DB_TIMEZONE: os.Getenv("ORDER_TIMEZONE"),
	}
}
