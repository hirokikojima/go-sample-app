package handler

import (
	"net/http"
	"strconv"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/hirokikojima/go-sample-app/models"
	"github.com/hirokikojima/go-sample-app/utilities/database"
)

type LogHandler interface {
	GetLogs(c echo.Context) error
	CreateLog(c echo.Context) error
	UpdateLog(c echo.Context) error
}

type logHandler struct {
	conn *gorm.DB
}

func NewLogHandler(conn *gorm.DB) LogHandler {
	return &logHandler{conn}
}

func (h *logHandler) GetLogs(c echo.Context) error {
	db := database.ConnectDatabase()
	defer db.Close()

	cc := NewCustomContext(c)

	log := models.Log{}

	if len(cc.Param("id")) > 0 {
		serviceID, err := strconv.ParseUint(cc.Param("id"), 10, 64)
		if err != nil {
			return err
		}
		log.ServiceID = uint(serviceID)
	}

	logs := log.GetLogs(db)

	return cc.JSON(http.StatusOK, logs)
}

func (h *logHandler) CreateLog(c echo.Context) error {
	db := database.ConnectDatabase()
	defer db.Close()

	cc := NewCustomContext(c)

	// user := cc.CurrentUser(db)

	// s := models.Service{
	// 	Model : gorm.Model{ID: log.ServiceID},
	// 	UserID: user.ID,
	// }

	// s := service.Find(db)
	// if s == nil {
	// 	return &echo.HTTPError{
    //         Code:    http.StatusUnauthorized,
    //         Message: "invalid email or password",
    //     }
	// }
	
	log := new(models.Log)
	if err := c.Bind(log); err != nil {
		return err
	}

	// log.CreateLog(db)

	return cc.JSON(http.StatusOK, log)
}

func (h *logHandler) UpdateLog(c echo.Context) error {
	db := database.ConnectDatabase()
	defer db.Close()

	cc := NewCustomContext(c)

	return cc.JSON(http.StatusOK, echo.Map{})
}