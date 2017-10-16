package index

import (
    "net/http"
    "github.com/labstack/echo"
)

type Index struct {

}

func (this *Index)Hello(c echo.Context) error {
    return c.Render(http.StatusOK,  "hello.html", "World")
    //return c.Render(http.StatusOK, "index.html", "world")
}
func (this *Index)Vue(c echo.Context)error {
    return c.Render(http.StatusOK, "index.html", "")
}