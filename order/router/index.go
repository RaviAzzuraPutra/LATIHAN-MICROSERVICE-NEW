package router

import (
	"order/controller"

	"github.com/gin-gonic/gin"
)

func Router(app *gin.Engine, orderController *controller.OrderControllerImpl) {
	route := app

	order := route.Group("/order")

	order.GET("/", orderController.GetController)
	order.POST("/add-order", orderController.CreateController)
}
