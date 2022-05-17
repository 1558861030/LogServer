package controllers

import (
	"LogServer/models"
	"fmt"
	"github.com/gin-gonic/gin"
)

func PutLog(c *gin.Context) {
	date := models.LogType{}
	err := c.BindJSON(&date)

	if err != nil {
		fmt.Println("error:PutLog date to json error")
		fmt.Println(err)
		fmt.Println(&date)
		c.JSON(400, gin.H{"error": "PutLog date to json error"})
		//return
	}
	//date.OtherInfo = append(date.OtherInfo, models.OtherInfo{Key: "keyaaaa", Value: "valueaaaa"})

	c.String(200, "ok")

	go func() {
		resp := models.MySqlMaster.Create(&date)
		if resp.Error != nil {
			fmt.Println(resp.Error)
		}
	}()
}

func Test(c *gin.Context) {
	c.String(200, "ok")
}
