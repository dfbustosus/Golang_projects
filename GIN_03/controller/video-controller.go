package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gin_03/entity"
	"github.com/gin_03/service"
)

type VideoController interface{
	FindAll() []entity.Video // Funcion devolver slice of video
	Save(ctx *gin.Context) entity.Video
}

type controller struct {
	service service.VideoService
}

func New(service service.VideoService) VideoController {
	return &controller{
		service: service,
	}
}

func (c *controller) FindAll() []entity.Video {
	return c.service.FindAll()
}
func (c *controller) Save(ctx *gin.Context) entity.Video{
	var video entity.Video
	ctx.BindJSON(&video)
	c.service.Save(video)
	return video
}