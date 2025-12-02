package controller

import (
	"stock/request"
	service_interface "stock/service/interface"

	"github.com/gin-gonic/gin"
)

type StockControllerImpl struct {
	stockService service_interface.StockServiceInterface
}

func NewStockController(service service_interface.StockServiceInterface) *StockControllerImpl {
	return &StockControllerImpl{
		stockService: service,
	}
}

func (c *StockControllerImpl) Create(ctx *gin.Context) {
	var stockRequest = new(request.StockRequest)

	errRequest := ctx.ShouldBind(&stockRequest)

	if errRequest != nil {
		ctx.JSON(400, gin.H{
			"Message": "Format Request Tidak Sesuai",
			"Error":   errRequest.Error(),
		})
		return
	}

	errService := c.stockService.CreateService(stockRequest)

	if errService != nil {
		ctx.JSON(500, gin.H{
			"Message": "Terjadi Kesalahan Saat Menambahkan Data",
			"Error":   errService.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"Message": "Berhasil Menambahkan Data",
		"Data":    stockRequest,
	})

}

func (g *StockControllerImpl) Get(ctx *gin.Context) {

	stock, err := g.stockService.GetService()

	if err != nil {
		ctx.JSON(404, gin.H{
			"Message": "Data Kosong",
			"Error":   err,
		})
		return
	}

	ctx.JSON(200, gin.H{
		"Message": "Berhasil Mendapatkan Data Stock",
		"Data":    stock,
	})
}

func (gid *StockControllerImpl) GetByID(ctx *gin.Context) {

	id := ctx.Param("id")

	stock, err := gid.stockService.GetByIDService(id)

	if err != nil {
		ctx.JSON(404, gin.H{
			"Message": "Data Kosong!",
			"Error":   err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"Message": "Berhasil Mendapatkan Data Berdasarkan ID",
		"Data":    stock,
	})
}

func (d *StockControllerImpl) Delete(ctx *gin.Context) {

	id := ctx.Param("id")

	err := d.stockService.DeleteService(id)

	if err != nil {
		ctx.JSON(500, gin.H{
			"Message": "Terjadi Kesalahan Saat Menghapus Data",
			"Error":   err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"Message": "Berhasil Menghapus Data",
	})
}

func (u *StockControllerImpl) Update(ctx *gin.Context) {

	id := ctx.Param("id")

	request := new(request.StockRequest)

	errRequest := ctx.ShouldBind(&request)

	if errRequest != nil {
		ctx.JSON(400, gin.H{
			"Message": "Format Request Tidak Sesuai",
			"Error":   errRequest.Error(),
		})
		return
	}

	err := u.stockService.UpdateService(request, id)

	if err != nil {
		ctx.JSON(500, gin.H{
			"Message": "Terjadi Kesalahan Saat Update Data",
			"Error":   err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"Message": "Berhasil Mengubah Data",
		"Data":    request,
	})

}
