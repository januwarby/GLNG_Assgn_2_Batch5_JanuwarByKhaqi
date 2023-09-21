package handler

import (
	"h8-assignment-2/infra/database"
	"h8-assignment-2/repository/order_repository/order_pg"
	"h8-assignment-2/service"

	"github.com/gin-gonic/gin"
)

func StartApp() {
	database.InitiliazeDatabase()

	db := database.GetDatabaseInstance()

	orderRepo := order_pg.NewOrderPG(db)

	orderService := service.NewOrderService(orderRepo)

	orderHandler := NewOrderHandler(orderService)

	r := gin.Default()

	r.POST("/orders", orderHandler.CreateOrder)

	r.GET("/orders", orderHandler.GetOrders)

	r.Run(":8080")

}
