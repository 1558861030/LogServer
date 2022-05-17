package router

import (
	"LogServer/controllers"
	"github.com/gin-gonic/gin"
)

func InitRouter(R *gin.Engine) {
	R.POST("/putlog", controllers.PutLog)
	R.GET("/test", controllers.Test)
}
