package controller

import "github.com/ppopgi-pang/ppopgipang-spine/notifications/service"

type NotificationController struct {
	service *service.NotificationService
}

func NewNotificationController(notificationsService *service.NotificationService) *NotificationController {
	return &NotificationController{service: notificationsService}
}
