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