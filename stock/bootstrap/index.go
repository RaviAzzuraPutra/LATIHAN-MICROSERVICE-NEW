package bootstrap

import (
	"context"
	"stock/apache_kafka/consumer"
	"stock/config"
	"stock/config/app_config"
	"stock/config/kafka_config"
	"stock/controller"
	"stock/database"
	"stock/repository"
	"stock/router"
	"stock/service"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func InitAPP() {
	errEnv := godotenv.Load()

	if errEnv != nil {
		panic("ERROR SAAT LOAD ENV")
	}

	config.Init_Config()

	database.Connect()

	stockRepository := repository.NewStockRepositoryDB()
	stockService := service.NewStockService(stockRepository)
	stockController := controller.NewStockController(stockService)

	consumer := consumer.NewStockConsumer(kafka_config.Kafka_Broker, kafka_config.Kafka_Topic, stockService)

	go func() {
		ctx := context.Background()
		consumer.StartDataConsumer(ctx)
	}()

	app := gin.Default()

	app.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"Message": "Aplikasi Berjalan Dengan Baik 👍",
		})
	})

	router.StockRouter(app, stockController)

	app.Run(app_config.APP)
}
