package main

import (
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



        e.Start(":5050")
}
