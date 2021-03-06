package service

import (
	"github.com/TestardR/Go-Gin-CRUD/entity"
)

type VideoService interface {
	Save(entity.Video) entity.Video
	FindAll() []entity.Video
	FindOne(id string) entity.Video
}

type videoService struct {
	videos []entity.Video
	video  entity.Video
}

func New() VideoService {
	return &videoService{}
}

func (service *videoService) Save(video entity.Video) entity.Video {
	service.videos = append(service.videos, video)
	return video
}

func (service *videoService) FindAll() []entity.Video {
	return service.videos
}

func (service *videoService) FindOne(id string) entity.Video {
	video := service.video
	for i := 0; i < len(service.videos); i++ {
		if id == service.videos[i].ID {
			video = service.videos[i]
		}
	}

	return video
}
