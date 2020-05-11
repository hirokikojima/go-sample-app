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

    e.POST("/signup", handlers.Signup)
    e.POST("/login", handlers.Login)
    
    guarded := e.Group("")
    guarded.Use(middleware.JWTWithConfig(authorizer.JwtConfig))

    e.POST("/users", handlers.CreateUser)
    e.GET("/users/:id", handlers.GetUser)
    e.PUT("/users/:id", handlers.UpdateUser)
    e.DELETE("/users/:id", handlers.DeleteUser)

    e.Logger.Fatal(e.Start(":1323"))
}