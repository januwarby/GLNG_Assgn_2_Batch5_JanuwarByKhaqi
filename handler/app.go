package handler

import (
	"h8-assignment-2/docs"
	"h8-assignment-2/infra/database"
	"h8-assignment-2/repository/item_repository/item_pg"
	"h8-assignment-2/repository/order_repository/order_pg"
	"h8-assignment-2/service"

	"github.com/gin-gonic/gin"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
	// swagger embed files
)

func StartApp() {
	database.InitiliazeDatabase()

	db := database.GetDatabaseInstance()

	orderRepo := order_pg.NewOrderPG(db)

	itemRepo := item_pg.NewItemPG(db)

	orderService := service.NewOrderService(orderRepo, itemRepo)

	orderHandler := NewOrderHandler(orderService)

	r := gin.Default()

	docs.SwaggerInfo.Title = "H8 Assignment 2"
	docs.SwaggerInfo.Description = "Ini adalah tugas ke 2 dari kelas Kampus Merdeka"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.Schemes = []string{"http"}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.POST("/orders", orderHandler.CreateOrder)

	r.GET("/orders", orderHandler.GetOrders)

	r.PUT("/orders/:orderId", orderHandler.UpdateOrder)

	r.DELETE("/orders/:orderId", orderHandler.DeleteOrder)

	r.Run(":8080")
}
