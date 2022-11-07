package searchuser

import (
	"strconv"
	"test_search_user/database"
	"test_search_user/logs"
)

type Usuario struct {
	id    int
	token string
	st    string
}

//===================================================
func ChkIdUser(sToken string) (sIdUser string, blnOK bool) {
	db, err := database.CnnDB()
	if err != nil {
		//fmt.Printf("Error obteniendo base de datos: %v", err)
		logs.EscribirLineaLog("ChkIdUser_Error cnn DB " + err.Error())
		return "", false
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		logs.EscribirLineaLog("ChkIdUser_Err: " + err.Error())
		return "", false
	} else {
		defer db.Close()
		sIdUsr := ""
		blnOK := false
		if sToken != "" {
			//Registrar liga
			sSQL := "SELECT id, token, st "
			sSQL += "FROM usr "
			sSQL += "WHERE (token='" + sToken + "')"
			//sSQL += "WHERE (token=?)"
			//sTmp := "'" + sToken + "'"
			//resDat, err := db.Query(sSQL, sTmp)
			resDat, err := db.Query(sSQL)
			if err != nil {
				logs.EscribirLineaLog("SQL: " + sSQL)
				logs.EscribirLineaLog("No fue posible encontrar el token: " + err.Error())
				return "", false
			} else {
				//Recorrer los registros
				for resDat.Next() {
					var datUsr = Usuario{}
					resDat.Scan(
						&datUsr.id, &datUsr.token, &datUsr.st,
					)
					sIdUsr = strconv.Itoa(datUsr.id)
					blnOK = true
				}
				//return sIdUsr, true
				if blnOK == true {
					return sIdUsr, true
				} else {
					return "", false
				}
			}
		} else {
			return "Se requiere de un token", false
		}
	}
} //FIN de ChkIdUser

func ChkApp(sToken string) (sIdUser string, blnOK bool) {
	db, err := database.CnnDB()
	if err != nil {
		//fmt.Printf("Error obteniendo base de datos: %v", err)
		logs.EscribirLineaLog("ChkApp_Error cnn DB " + err.Error())
		return "", false
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		logs.EscribirLineaLog("ChkApp_Err: " + err.Error())
		return "", false
	} else {
		defer db.Close()
		sIdUsr := ""
		blnOK := false
		if sToken != "" {
			//Registrar liga
			sSQL := "SELECT id, token, st "
			sSQL += "FROM app "
			sSQL += "WHERE (token='" + sToken + "')"
			//sSQL += "WHERE (token=?)"
			//sTmp := "'" + sToken + "'"
			//resDat, err := db.Query(sSQL, sTmp)
			resDat, err := db.Query(sSQL)
			if err != nil {
				logs.EscribirLineaLog("SQL: " + sSQL)
				logs.EscribirLineaLog("No fue posible encontrar el token: " + err.Error())
				return "", false
			} else {
				//Recorrer los registros
				for resDat.Next() {
					var datUsr = Usuario{}
					resDat.Scan(
						&datUsr.id, &datUsr.token, &datUsr.st,
					)
					sIdUsr = strconv.Itoa(datUsr.id)
					blnOK = true
				}
				if blnOK == true {
					return sIdUsr, true
				} else {
					return "", false
				}
				//return sIdUsr, true
			}
		} else {
			return "Se requiere de un token", false
		}
	}
} //FIN de ChkApp
