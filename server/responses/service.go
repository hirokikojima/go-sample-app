package responses

import (
	"github.com/hirokikojima/go-sample-app/models"
)

type serviceResponse struct {
	ID       uint   `json:"id"`
	UserID   uint   `json:"user_id"`
	Title    string `json:"title"`
	Body     string `json:"body"`
}

func NewServiceResponse(service models.Service) *serviceResponse {
	return &serviceResponse {
		ID: service.ID,
		UserID: service.UserID,
		Title: service.Title,
		Body: service.Body,
	}
}