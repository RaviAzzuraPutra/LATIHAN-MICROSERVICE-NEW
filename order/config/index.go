package config

import (
	"order/config/app_config"
	"order/config/db_config"
	"order/config/kafka_config"
)

func Init_Config() {
	app_config.APP_CONFIG()
	db_config.DB_CONFIG()
	kafka_config.KAFKA_CONFIG()
}
