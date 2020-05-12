package authorizer

import (
	"time"
	"golang.org/x/crypto/bcrypt"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/hirokikojima/go-sample-app/models"
)

type jwtCustomClaims struct {
    UID  uint    `json:"uid"`
    Name string `json:"name"`
    jwt.StandardClaims
}

var JwtConfig = middleware.JWTConfig{
    Claims:     &jwtCustomClaims{},
    SigningKey: []byte("secret"),
}

func EncryptPassword(password string) (*string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	encrypted := string(bytes)
	return &encrypted, nil
}

func ComparePassword(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GenerateSignedToken(user *models.User) (string, error) {
	claims := &jwtCustomClaims{
		user.ID,
		user.Name,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("secret"))
}

func GetClaims(c echo.Context) *jwtCustomClaims {
	user := c.Get("user").(*jwt.Token)
	return user.Claims.(*jwtCustomClaims)
}