package handler

import (
	"net/http"
	"strconv"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/hirokikojima/go-sample-app/models"
	"github.com/hirokikojima/go-sample-app/utilities/database"
	"github.com/hirokikojima/go-sample-app/utilities/storage"
)

type ServiceHandler interface {
	GetServices(c echo.Context) error
	GetService(c echo.Context) error
	CreateService(c echo.Context) error
	UpdateService(c echo.Context) error
}

type serviceHandler struct {
	conn *gorm.DB
}

func NewServiceHandler(conn *gorm.DB) ServiceHandler {
	return &serviceHandler{conn}
}

func (h *serviceHandler) GetServices(c echo.Context) error {
	db := database.ConnectDatabase()
	defer db.Close()

	cc := NewCustomContext(c)

	service := models.Service{}

	if len(cc.Param("id")) > 0 {
		userID, err := strconv.ParseUint(cc.Param("id"), 10, 64)
		if err != nil {
			return err
		}
		service.UserID = uint(userID)
	}

	services := service.GetServices(db)

	return cc.JSON(http.StatusOK, services)
}

func (h *serviceHandler) GetService(c echo.Context) error {
	cc := NewCustomContext(c)
	return cc.JSON(http.StatusOK, echo.Map{})
}

func (h *serviceHandler) CreateService(c echo.Context) error {
	db := database.ConnectDatabase()
	defer db.Close()

	cc := NewCustomContext(c)

	user := cc.CurrentUser(db)

	file, err := cc.FormFile("thumbnail")
	if err != nil {
		return err
	}

	if err := storage.LocalPut(file); err != nil {
		return err
	}

	service := new(models.Service)
	if err := c.Bind(service); err != nil {
		return err
	}

	service.UserID = user.ID
	service.Thumbnail = storage.Prefix + file.Filename
	service.CreateService(db)

	return cc.JSON(http.StatusOK, service)
}

func (h *serviceHandler) UpdateService(c echo.Context) error {
	db := database.ConnectDatabase()
	defer db.Close()

	cc := NewCustomContext(c)

	user := cc.CurrentUser(db)

	if len(cc.Param("id")) == 0 {
		return &echo.HTTPError{
            Code:    http.StatusUnauthorized,
            Message: "invalid email or password",
        }
	}
		
	serviceID, err := strconv.ParseUint(cc.Param("id"), 10, 64)
	if err != nil {
		return err
	}
	
	s := models.Service{
		Model: gorm.Model{ID: uint(serviceID)},
		UserID: user.ID,
	}

	service := s.GetService(db)
	if err := c.Bind(service); err != nil {
		return err
	}

	service.UpdateService(db)

	return cc.JSON(http.StatusOK, service)
}