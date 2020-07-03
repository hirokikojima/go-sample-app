package router

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/hirokikojima/go-sample-app/http/handler"
	"github.com/hirokikojima/go-sample-app/utilities/authorizer"
)

func NewRouter(e *echo.Echo, h handler.AppHandler) {
	e.Use(middleware.CORS())
    e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	/**
	 * 静的ルート
	 */
	e.Static("/public", "public")

	/**
	 * 非認証ルート
	 */
	e.POST("/signup", h.Signup)
	e.POST("/signin", h.Signin)
	e.GET("/users", h.GetUsers)
	e.GET("/user/:id", h.GetUser)
	e.GET("/user/:user_id/services", h.GetServices)
	e.GET("/user/:user_id/service/:id", h.GetService)


	/**
	 * 認証ルート
	 */
	guarded := e.Group("")
	guarded.Use(middleware.JWTWithConfig(authorizer.JwtConfig))

	guarded.GET("/me", h.Me)
	guarded.POST("/user/:user_id/service", h.CreateService)
    guarded.PUT("/user/:user_id/service/:id", h.UpdateService)
    guarded.POST("/user/:user_id/service/:service_id/log", h.CreateLog)
    guarded.PUT("/user/:user_id/service/:service_id/log/:id", h.UpdateLog)
}