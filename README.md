# search_user

http://74.208.92.108:3005/credit/search_user
- GET:  /search_user
- Archivo configHttp.json donde se configura IP y puerto.
- Archivo configDB.json donde se configura la conección a la base de datos.
		
- Servicio middleware para la seguridad con el manejo de un token (Autoenthication Beaer token):
		- Valida que tenga un token 12345.54321 
		- Autentica al usuario.
		- Verifica a que tiene acceso
