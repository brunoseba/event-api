# EJERCICIO BACKEND Servicio APi Rest con Go

_Acá va un párrafo que describa lo que es el proyecto_

Se realiza una aplicación web con una API  que le permita a una compañía interactuar con la interfaz de usuario. La aplicación permitirá a un administrador gestionar eventos (crear, eliminar o editar eventos), y los eventos tendrán un título, una breve descripción, una descripción detallada, una fecha y hora, un organizador, un lugar y un estado (borrador o publicado).

## Comenzando 🚀

_Estas instrucciones te permitirán obtener una copia del proyecto en funcionamiento en tu máquina local para propósitos de desarrollo y pruebas._

* Clonar el repositorio
* Ingresar a la captepa cmd
* Ejecuta: (para devantar el servicio)
```
go run main.go
```

* Ejecutar: (para crear el servicio)
```
go build main.go
```


### Pre-requisitos 📋

_Que necesitas para correr el sistema_

Tener instalado Go con la version 1.19 en adelante, para verificar version ejecute en terminal:

```
go version
```

Tener MongoDb instalado, verifique corriendo el comando:
```
mongo --version
```


### Docker 🔧
_Para correr el sistema con docker debe tener instalado docker en el sisttema_

* Verificar corriendo el comando
```
docker -v 
```

Luego de clonar el repositorio e ingrasar a la carpeta principal del proyecto,

* Correr el comando
```
docker compose up
```
(esto ejecutara docker creando los contenedores necesarios)

## Ejecutando el servicio ⚙️

### Create a new user (server response: 201)
```shell script
curl -X POST \
  http://localhost:8080/register \
  -d '{
	"username": "usuario"
}'
```
* Para crear un usuario Admin debe enviar el campo 'isAdmin' con valor true

### Create a new event (server response: 201)
*Este ejemplo crea un evento (con campos minimos) de estado publicado
```shell script
curl -X POST \
  http://localhost:8080/event \
  -d '{
	"title": "Evento"
  "date": "2023-03-30T18:30:20-03:00"
  "description_shot": "Creacion de evento"
  "state": "publicado"
}'
```

### Get User (Obtiene user por ID)
```shell script
curl -X GET \
  http://localhost:8080/user/ID-user
```

### Get Event (Obtiene evento por ID)
```shell script
curl -X GET \
  http://localhost:8080/user/ID-user
```


## Construido con 🛠️

_Herramientas que se utilizo para crear el proyecto_
* [GinGonic](https://github.com/gin-gonic/gin)
* GoDoc
* Client Rest:[Postman](https://www.postman.com/)


## Autor ✒️

* **Bruno Sebastian Riotorto**
