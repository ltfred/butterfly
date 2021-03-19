package rest

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	r := gin.New()
	r.GET("/time", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"message": time.Now().Format(time.RFC3339)})
	})
	return r
}
