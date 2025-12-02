package bootstrap

import (
	"log"
	"order/apache_kafka/producer"
	"order/config"
	"order/config/app_config"
	"order/config/kafka_config"
	"order/controller"
	"order/database"
	"order/repository"
	"order/router"
	"order/service"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func InitApp() {
	errEnv := godotenv.Load()

	if errEnv != nil {
		panic("TERJADI KESALAHAN SAAT LOAD ENV")
	}

	config.Init_Config()

	database.Connect()

	orderPublisher := producer.NewOrderProducer(kafka_config.Kafka_Broker, kafka_config.Kafka_Topic)

	defer func() {
		if err := orderPublisher.Writer.Close(); err != nil {
			log.Fatalf("Gagal menutup Kafka writer: %v", err)
		}
	}()

	NewRepository := repository.NewReositoryOrderDB()
	NewService := service.NewServiceOrder(NewRepository, orderPublisher)
	NewController := controller.NewControllerOrder(NewService)

	app := gin.Default()

	app.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"Message": "Aplikasi Berjalan Dengan Baik",
		})
	})

	router.Router(app, NewController)

	app.Run(app_config.PORT)
}
