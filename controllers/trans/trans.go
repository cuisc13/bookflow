package trans

import (
	"github.com/labstack/echo"
	"net/http"
	"cn/bkread/booktrans/models"
)

type TransController struct {

}

func (this *TransController)SetUp(c echo.Context)error{
	tra := models.NewTrans()
	c.Bind(tra)
	return c.JSON(http.StatusOK, tra)
}

func (this *TransController)List(c echo.Context)error{
	tra := new(models.Trans)
	var tra_list []models.Trans = []models.Trans{}
	tra.All(&tra_list)
	return c.JSON(http.StatusOK, tra_list)
}

func (this *TransController)List_bookid(c echo.Context)error{

	tra := new(models.Trans)
	tra.BookId = c.QueryParam("book_id")
	var tra_list []models.Trans = []models.Trans{}
	tra.Query(&tra_list)
	return c.JSON(http.StatusOK, tra_list)
}

func (this *TransController)Rm(c echo.Context)error{
	tra := new(models.Trans)
	c.Bind(tra)
	tra.Del()

	return c.JSON(http.StatusOK, tra)
}