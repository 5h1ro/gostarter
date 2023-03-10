package service

import (
	"gostarter/server/model"

	"gorm.io/gorm"
)

type PostService struct {
	DB *gorm.DB
}

func (service PostService) CreatePost(title, content string, userID uint) model.Post {
	return model.Post{
		Title:   title,
		Content: content,
		UserID:  userID,
	}
}
