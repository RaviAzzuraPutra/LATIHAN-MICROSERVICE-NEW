package database

import (
	"fmt"
	"log"
	"stock/config/db_config"
	"stock/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {

	var errorConnection error

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s client_encoding=UTF8",
		db_config.DB_Config().HOST, db_config.DB_Config().USER, db_config.DB_Config().PASSWORD, db_config.DB_Config().NAME, db_config.DB_Config().PORT, db_config.DB_Config().SSL, db_config.DB_Config().TIMEZONE)

	DB, errorConnection = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if errorConnection != nil {
		log.Printf("TERJADI KESALAHAN SAAT KONEK KE DATABASE" + errorConnection.Error())
	}

	errMigrate := DB.AutoMigrate(&model.Stock{})

	if errMigrate != nil {
		log.Printf("TERJADI KESALAHAN SAAT MIGRASI DATABASE" + errMigrate.Error())
	}

	if DB == nil {
		log.Panic("TERJADI KESALAHAN DI DATABASE!!")
	}

	log.Printf("BERHASIL TERHUBUNG KE DATABASE ✅")
}
