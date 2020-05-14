package presenters

import (
	"github.com/hirokikojima/go-sample-app/models"
)

type userPresenter struct {
	user models.User
}

type userResponse struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
}

func NewUserPresenter(user models.User) *userPresenter {
	return &userPresenter{ user: user }
}

func (i userPresenter) Convert() *userResponse {
	return &userResponse {
		ID: i.user.ID,
		Name: i.user.Name,
		Email: i.user.Email,
	}
}