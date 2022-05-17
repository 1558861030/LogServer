package main

import (
	"LogServer/models"
	"LogServer/router"
	"github.com/gin-gonic/gin"
)

func main() {
	//连接数据库
	models.GormSql()

	//新建gin路由
	Gin := gin.New()
	gin.DisableConsoleColor()
	//注册路由
	router.InitRouter(Gin)
	Gin.Run(":9100")
}
