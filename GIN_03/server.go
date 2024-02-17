package main
import (
	"github.com/gin-gonic/gin"
	//"github.com/gin_03/entity"
	"github.com/gin_03/controller"
    "github.com/gin_03/service"
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