package controller

import "github.com/thanpawatpiti/gobeer/pkg/services"

type Controller struct {
	service services.Service
}

func NewController(service services.Service) *Controller {
	return &Controller{
		service: service,
	}
}
