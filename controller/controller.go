package controller

import (
	"fmt"
	"strings"
	"test_search_user/logs"
	"test_search_user/searchuser"

	"github.com/gin-gonic/gin"
)

//=======================================================
func SearchUser(c *gin.Context) {
	//logs.EscribirLineaLog("CreditAssigment ...")
	sToken := ""

	if len(c.Request.Header["Authorization"]) > 0 {
		sToken = c.Request.Header["Authorization"][0]
		if sToken == "Bearer" {
			c.JSON(400, gin.H{"code": "search_user_microservice", "message": "Se requiere de un token.", "status_code": 400})
			return
		}
		//sToken = strings.Split(c.Request.Header["Authorization"][0], " ")[1]
		sToken = strings.Replace(sToken, "Bearer", "", -1)
		sToken = strings.Trim(sToken, " ")
		aToken := strings.Split(sToken, ".")
		//Buscar usuario
		if sToken != "" {
			sTmp, blnBuscar := searchuser.ChkIdUser(aToken[0])
			if blnBuscar == true {
				if len(aToken) > 1 {
					sChk, blnChk := searchuser.ChkApp(aToken[1])
					fmt.Println(sChk)
					if blnChk == true {
						c.JSON(200, gin.H{"code": "search_user_microservice", "message": sTmp, "status_code": 200})
					} else {
						c.JSON(400, gin.H{"code": "search_user_microservice", "message": "No tiene acceso a este servicio", "status_code": 400})
					}
				} else {
					c.JSON(400, gin.H{"code": "search_user_microservice", "message": "Se requiere de un token.", "status_code": 400})
				}
			} else {
				logs.EscribirLineaLog("_GetIdUser_Err token no encontrado.")
				c.JSON(400, gin.H{"code": "api_error", "message": "Token no encontrado.", "status_code": 400})
			}
		} else {
			c.JSON(400, gin.H{"code": "api_error", "message": "Se requiere de un token.", "status_code": 400})

		}
	} else {
		c.JSON(400, gin.H{"code": "api_error", "message": "Se requiere de un token.", "status_code": 400})
		return
	}

} //FIN de CreditAssigment
