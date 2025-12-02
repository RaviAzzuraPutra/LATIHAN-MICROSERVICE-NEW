package config

import (
	"stock/config/app_config"
	"stock/config/db_config"
	"stock/config/kafka_config"
)

func Init_Config() {
	app_config.App_Config()
	db_config.DB_Config()
	kafka_config.KAFKA_CONFIG()
}
