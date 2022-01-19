package router

import (
	"assignment-7/controller"
	"assignment-7/db"

	"github.com/gin-gonic/gin"
)

var PORT = ":8080"

func init() {
	db.NewDB()
}

func StartRouter() {
	route := gin.Default()

	ordersRoute := route.Group("/")
	{
		ordersRoute.POST("/orders", controller.OrdersCreate)
		ordersRoute.GET("/orders/:ordersId", controller.GetById)
		ordersRoute.GET("/orders", controller.GetAll)
		ordersRoute.PUT("/orders/:ordersId/:itemId", controller.Update)
		ordersRoute.DELETE("/orders/:ordersId", controller.Delete)
	}

	route.Run(PORT)
}
