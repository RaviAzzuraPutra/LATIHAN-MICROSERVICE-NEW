package router

import (
	"stock/controller"

	"github.com/gin-gonic/gin"
)

func StockRouter(app *gin.Engine, stockController *controller.StockControllerImpl) {
	route := app

	stock := route.Group("/stock")

	stock.POST("/add-stock", stockController.Create)
	stock.GET("/", stockController.Get)
	stock.GET("/:id", stockController.GetByID)
	stock.DELETE("/delete/:id", stockController.Delete)
	stock.PUT("/update/:id", stockController.Update)
}
