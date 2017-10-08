package main

import (
    "github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
    "cn/bkread/booktrans/controllers"
)

func main(){
    e := echo.New()
    g := e.Group("/api").Group("/v1")
    controllers.Register(g)
	e.Use(middleware.Logger())
	e.Logger.Fatal(e.Start(":9091"))

}
