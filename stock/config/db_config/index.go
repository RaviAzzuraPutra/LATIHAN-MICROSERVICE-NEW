package db_config

import "os"

type DBConfig struct {
	HOST     string
	USER     string
	PASSWORD string
	NAME     string
	PORT     string
	SSL      string
	TIMEZONE string
}

func DB_Config() *DBConfig {
	return &DBConfig{
		HOST:     os.Getenv("STOCK_HOST"),
		USER:     os.Getenv("STOCK_USER"),
		PASSWORD: os.Getenv("STOCK_PASSWORD"),
		NAME:     os.Getenv("STOCK_NAME"),
		PORT:     os.Getenv("STOCK_PORT"),
		SSL:      os.Getenv("STOCK_SSL"),
		TIMEZONE: os.Getenv("STOCK_TIMEZONE"),
	}
}
