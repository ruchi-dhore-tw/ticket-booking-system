package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PingHandler() func(context *gin.Context) {
	return func(context *gin.Context) {
		context.Status(http.StatusOK)
	}
}
