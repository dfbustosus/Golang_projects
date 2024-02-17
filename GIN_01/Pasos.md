1. Dependencies
```bash
go mod init gitlab.com/pragmaticreviews/gloang-gin-poc
```
2. Instalar atributos
```bash
go get github.com/gin-gonic/gin
```
3. Inicializar el server
```go
package main
import "github.com/gin-gonic/gin"

func main(){
	server:= gin.Default()	
	
	server.GET("/test",func(ctx *gin.Context){
		ctx.JSON(200, gin.H{
			"mensaje" : "Ok!",
		})
	})
	
	server.Run(":8080")

}
```
4. Inicializar
```bash
go run .\server.go
```
5. Ir a postman y testear `http://localhost:8080/test` y enviar en el body `{"message":"Hello"}` y deberas obtener un `{"mensaje": "Ok!"}`
6. Definimos carpeta entity y creamos el archivo `video.go`
```go
package entity
type Video struct{
	Titulo string `json:"titulo"`
	Descripcion string `json:"descripcion"`
	URL string `json:"url"`
}
```
7. Creamos el folder service y creamos el archivo `video-service.go` y ponemos dentro
```go
package service
type VideoService interface {
	// Funciones de interes
	Save(entity.Video) entity.Video // Return el nuevo video
	FindAll() []entity.Video // Devuelve un set de videos
}
type videoService struct {
	videos []entity.Video // Propiedad de slice of videos
}

func New()  VideoService{
	return &videoService{} // Devolver el pointer al struct
}

// Definir el struct dentro de las funciones
func (service *videoService) Save(video entity.Video) entity.Video {
	service.videos = append(service.videos, video) // Agregar el video a la estructura
	return video
}
func (service *videoService) FindAll() []entity.Video {
	return service.videos
}
```
8. Creamos un nuevo folder `controller` y dentro `video-controller.go` con:
```go
package controller

type VideoController interface{
	FindAll() []entity.Video // Funcion devolver slice of video
	Save(ctx *gin.Context) entity.Video
}

type controller struct {
	service service.VideoService
}

func New(service service.VideoService) VideoController {
	return controller {
		service: service,
	}
}

func (c *controller) FindAll() []entity.Video {
	return c.service.New().FindAll()
}
func (c *controller) Save(ctx *gin.Context) {
	var video entity.Video
	ctx.BindJSON(&video)
	c.service.Save(video)
	return video
}
```
9. Volver al server.go y cambiar a:
```go
package main
package main
import (
	"github.com/gin-gonic/gin"
	"gitlab.com/pragmaticreviews/golang-gin-poc/controller"
	"gitlab.com/pragmaticreviews/golang-gin-poc/service"
)

var(
	videoService service.VideoService = service.New()
	videoController controller.VideoController = controller.New(videoService)
)
func main(){
	server:= gin.Default()	
	
	server.GET("/test",func(ctx *gin.Context){
		ctx.JSON(200, gin.H{
			"mensaje" : "Ok!",
		})
	})

	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.FindAll())
	})

	server.POST("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.Save(ctx))
	})
	
	server.Run(":8080")

}
```
10. Instalar dependencias externas si no ha sinicializado el proyect
```bash
go get gitlab.com/pragmaticreviews/golang-gin-poc/controller
go get gitlab.com/pragmaticreviews/golang-gin-poc/service
```
11. Ahora podemos correr de nuevo
```bash
go run .\server.go
```
12. Ir a postman y probar el metodo GET `http://localhost:8080/videos` te debe devolver un array empty
13. Ir a postman y probar el metodo POST `http://localhost:8080/videos` con el body 

```json
{
    "title": "David Titulo",
    "description": "Descripcion 1",
    "url": "https//myweb.com"
}
```
Al ejecutar te deberia devolver el registro insertado
```json
{
    "title": "David Titulo",
    "description": "Descripcion 1",
    "url": "https//myweb.com"
}
```
14. Si vuelver a ir al GET `http://localhost:8080/videos` te deberan salir todos los registros insertados

FIN!