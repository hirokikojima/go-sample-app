package handlers

import (
	"net/http"
	"strconv"
	"github.com/labstack/echo"
	"github.com/hirokikojima/go-sample-app/models"
	"github.com/hirokikojima/go-sample-app/utilities/database"
)

func GetLogs(c echo.Context) error {
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

func CreateLog(c echo.Context) error {
	db := database.ConnectDatabase()
	defer db.Close()

	cc := NewCustomContext(c)

	user := cc.CurrentUser(db)

	log := new(models.Log)
	if err := c.Bind(log); err != nil {
		return err
	}

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
	
	// log.CreateLog(db)

	return cc.JSON(http.StatusOK, log)
}