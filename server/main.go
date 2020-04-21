package main

import (
    "github.com/labstack/echo"
    "github.com/labstack/echo/middleware"
    "github.com/hirokikojima/go-sample-app/handlers"
)

func main() {
    e := echo.New()

    e.Use(middleware.CORS())

    e.POST("/signup", handlers.Signup)
    e.POST("/login", handlers.Login)
    e.POST("/users", handlers.CreateUser)
    e.GET("/users/:id", handlers.GetUser)
    e.PUT("/users/:id", handlers.UpdateUser)
    e.DELETE("/users/:id", handlers.DeleteUser)

    e.Logger.Fatal(e.Start(":1323"))
}