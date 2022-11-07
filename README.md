# search_user

- GET:  /credit/search_user
- Archivo configHttp.json donde se configura IP y puerto.
- Archivo configDB.json donde se configura la conección a la base de datos.
		
- Servicio middleware para la seguridad con el manejo de un token (Autoenthication Beaer token):
		- Valida que tenga un token XXXXXXXXX.XXXXXXXX 
		- Autentica al usuario.
		- Verifica a que tiene acceso
		
Responde en caso de exito 200 y json: 
{ 
  "code": "search_user_microservice", 
  "message": "Id del usuario", "status_code": 200 }

Responde en caso de error: 
{ 
  "code": "search_user_microservice", 
  "message": "Mensaje a detalle", 
  "status_code": 400 
}
