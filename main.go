package main

import (
	"test_search_user/controller"
	"test_search_user/logs"
	"test_search_user/utils"

	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.DebugMode)
	router := gin.Default()
	router.Use(cors.Default())
	Router(router)
	sIpPort := utils.GetIpPuerto()
	if err := router.Run(sIpPort); err != nil {
		logs.EscribirLineaLog("Error al levantar el sever ..." + err.Error())
		panic(fmt.Errorf("Fatal Error Description: %s ", err))
	}
}

func Router(router *gin.Engine) {
	v1 := router.Group("/credit")
	{
		v1.GET("/search_user", controller.SearchUser)
	}
}
