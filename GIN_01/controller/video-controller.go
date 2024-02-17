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