package main

import (
    "flag"
    "fmt"
    "github.com/labstack/echo"
    "github.com/hirokikojima/go-sample-app/config"
    "github.com/hirokikojima/go-sample-app/http/router"
    "github.com/hirokikojima/go-sample-app/http/handler"
)

func main() {
    isProduction := flag.Bool("production", false, "production is -production option require")

    env := "local"
    if *isProduction {
        env = "production"
    }

    config.NewConfig(env)

    e := echo.New()
    conn := config.NewDBConnection()
    defer func() {
        if err := conn.Close(); err != nil {
            e.Logger.Fatal(fmt.Sprintf("Failed to close: %v", err))
        }
    }()

    h := handler.NewAppHandler(conn)
    router.NewRouter(e, h)

    e.Logger.Fatal(e.Start(":1323"))
}