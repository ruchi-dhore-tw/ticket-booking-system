package app

import (
	"github.com/gin-gonic/gin"
	"github.com/ruchi-dhore-tw/ticket-booking-system/handler"
)

func addRoutes() *gin.Engine {
	router := gin.Default()

	router.GET("/", handler.PingHandler())
	router.POST("/book", handler.BookTicketHandler())

	return router
}
