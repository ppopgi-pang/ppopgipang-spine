package service

import "github.com/ppopgi-pang/ppopgipang-spine/notifications/repository"

type NotificationService struct {
	repo *repository.NotificationRepository
}

func NewNotificationService(repo *repository.NotificationRepository) *NotificationService {
	return &NotificationService{repo: repo}
}
