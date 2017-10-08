package index

import (
    "net/http"
    "github.com/labstack/echo"
)

type Index struct {

}

func (this *Index)Hello(c echo.Context) error {
    return c.String(http.StatusOK, "Hello, 世界")
}
