package main

import (
    "github.com/joho/godotenv"
    "github.com/labstack/echo"
    "github.com/hirokikojima/go-sample-app/handlers"
)

func main() {
    err := godotenv.Load()
    if err != nil {
        panic(err)
    }

    e := echo.New()

    e.POST("/users", handlers.CreateUser)
    e.GET("/users/:id", handlers.GetUser)
    e.PUT("/users/:id", handlers.UpdateUser)
    e.DELETE("/users/:id", handlers.DeleteUser)

    e.Logger.Fatal(e.Start(":1323"))
}