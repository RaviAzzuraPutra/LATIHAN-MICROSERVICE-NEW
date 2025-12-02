package controller

import (
	"order/request"
	service_interfcae "order/service/interface"

	"github.com/gin-gonic/gin"
)

type OrderControllerImpl struct {
	Service service_interfcae.OrderServiceInterface
}

func NewControllerOrder(service service_interfcae.OrderServiceInterface) *OrderControllerImpl {
	return &OrderControllerImpl{
		Service: service,
	}
}

func (c *OrderControllerImpl) GetController(ctx *gin.Context) {
	order, err := c.Service.GetService()

	if err != nil {
		ctx.JSON(500, gin.H{
			"Message": "Terjadi Kesalahan Saat Mengambil Data",
			"Error":   err.Error(),
		})
		return
	}

	if len(order) == 0 {
		ctx.JSON(404, gin.H{
			"Message": "DATA MASIH KOSONG",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"Message": "Berhasil Mendapatkan Data Order",
		"Data":    order,
	})
}

func (c *OrderControllerImpl) CreateController(ctx *gin.Context) {
	var OrderRequest = new(request.OrderRequest)

	errRequest := ctx.ShouldBind(&OrderRequest)

	if errRequest != nil {
		ctx.JSON(400, gin.H{
			"Message": "Format Request Tidak Sesuai",
			"Error":   errRequest.Error(),
		})
		return
	}

	if OrderRequest.ProductID == nil {
		ctx.JSON(400, gin.H{
			"Message": "Product ID Harus Terisi",
		})
		return
	}

	if OrderRequest.Quantity == nil {
		ctx.JSON(400, gin.H{
			"Message": "Quantity Harus Terisi",
		})
		return
	}

	err := c.Service.CreateService(OrderRequest)

	if err != nil {
		ctx.JSON(500, gin.H{
			"Message": "Terjadi Kesalahan Saat Menambahkan Data",
			"Error":   err.Error(),
		})
		return
	}

	ctx.JSON(201, gin.H{
		"Message": "Berhasil Menambahkan Data",
		"Data":    OrderRequest,
	})
}
