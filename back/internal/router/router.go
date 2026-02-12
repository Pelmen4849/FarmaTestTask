package router

import (
	"medical_farm/back/internal/handler"
	"medical_farm/back/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(
	drugHandler *handler.DrugHandler,
	orderHandler *handler.OrderHandler,
) *gin.Engine {
	r := gin.Default()

	// Глобальные middleware
	r.Use(middleware.CORSMiddleware())
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// Публичные маршруты
	api := r.Group("/api")
	{
		// Товары
		api.GET("/drugs", drugHandler.GetAvailableDrugs)
		api.GET("/drugs/:id", drugHandler.GetDrugByID)

		// Заказы
		api.POST("/orders", orderHandler.CreateOrder)
	}

	return r
}
