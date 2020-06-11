package responses

import (
	"github.com/hirokikojima/go-sample-app/models"
)

type userResponse struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
}

func NewUserResponse(user models.User) *userResponse {
	return &userResponse {
		ID: user.ID,
		Name: user.Name,
		Email: user.Email,
	}
}