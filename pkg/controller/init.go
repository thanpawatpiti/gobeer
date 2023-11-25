package controller

import (
	"github.com/thanpawatpiti/gobeer/pkg/domain"
)

type Controller struct {
	service domain.BeerSerives
}

func NewController(service domain.BeerSerives) *Controller {
	return &Controller{
		service: service,
	}
}
