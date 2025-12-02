package database

import (
	"fmt"
	"log"
	"order/config/db_config"
	"order/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	var errorConnect error

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s client_encoding=UTF8",
		db_config.DB_CONFIG().DB_HOST,
		db_config.DB_CONFIG().DB_USER,
		db_config.DB_CONFIG().DB_PASSWORD,
		db_config.DB_CONFIG().DB_NAME,
		db_config.DB_CONFIG().DB_PORT,
		db_config.DB_CONFIG().DB_SSL,
		db_config.DB_CONFIG().DB_TIMEZONE)

	DB, errorConnect = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if errorConnect != nil {
		log.Fatalf("TERJADI KESALAHAN SAAT KONEK KE DATABASE" + errorConnect.Error())
		return
	}

	if DB == nil {
		log.Fatalf("TERJADI KESALAHAN DI DATABASE")
		return
	}

	errMigrate := DB.AutoMigrate(&model.Order{})

	if errMigrate != nil {
		log.Fatalf("TERJADI KESALAHAN SAAT MIGRASI" + errMigrate.Error())
		return
	}

	log.Printf("BERHASIL TERHUBUNG KE DATABASE ✅")

}
