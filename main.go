package main

import (
        "github.com/alvaroenriqueds/ws_potencie/controllers"
        "github.com/labstack/echo"
        "github.com/labstack/echo/middleware"
)

func main()  {
        e := echo.New()

        e.Use(middleware.Logger())
        e.Use(middleware.Recover())
        e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
                AllowOrigins:   []string{"*"},
                AllowMethods:   []string{echo.GET, echo.POST},
        }))

        e.POST("/track", controllers.Tracking)

        e.Start(":5050")
}
