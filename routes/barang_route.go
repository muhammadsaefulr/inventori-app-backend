package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/muhammadsaefulr/inventori-barang/config/db"
	"github.com/muhammadsaefulr/inventori-barang/handler"
	"github.com/muhammadsaefulr/inventori-barang/usecase"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	db, err := config.ConnectDatabase()

	if err != nil {
		panic("Failed To Connect The Database !")
	}

	MainUseCase := usecase.NewBarangUseCase(db)

	v1 := r.Group("/api/v1")
	{
		handler.NewBarangHandler(v1, MainUseCase)
	}
	return r

}
