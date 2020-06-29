package main

import (
    "github.com/labstack/echo"
    "github.com/labstack/echo/middleware"
    "github.com/hirokikojima/go-sample-app/handlers"
    "github.com/hirokikojima/go-sample-app/utilities/authorizer"
)

func main() {
    e := echo.New()

    e.Use(middleware.CORS())
    e.Use(middleware.Logger())
	e.Use(middleware.Recover())

    e.Static("/public", "public")
    e.POST("/signup", handlers.Signup)
    e.POST("/login", handlers.Login)
    e.GET("/user", handlers.GetUsers)
    e.GET("/user/:id", handlers.GetUser)
    e.GET("/service", handlers.GetServices)
    e.GET("/service/:id", handlers.GetServices)

    guarded := e.Group("")
    guarded.Use(middleware.JWTWithConfig(authorizer.JwtConfig))

    guarded.GET("/me", handlers.Me)
    guarded.POST("/service", handlers.CreateService)
    guarded.PUT("/service/:id", handlers.UpdateService)

    e.Logger.Fatal(e.Start(":1323"))
}