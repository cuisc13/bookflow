package main

import (
    "github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
    "cn/bkread/booktrans/controllers"
    "cn/bkread/booktrans/controllers/index"
	"cn/bkread/booktrans/view"
)

func main(){
    e := echo.New()

	e.Static("/static", "static")

	e.Renderer = view.NewTemplate()

	ind := new(index.Index)
	e.GET("/", ind.Hello)
	e.GET("/vue", ind.Vue)

    g := e.Group("/api").Group("/v1")
    controllers.Register(g)

	e.Use(middleware.Logger())
	e.Logger.Fatal(e.Start(":9091"))

}
