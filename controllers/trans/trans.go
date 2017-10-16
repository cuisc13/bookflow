package trans

import (
	"github.com/labstack/echo"
	"net/http"
	"cn/bkread/booktrans/models"
	"time"
)

type TransController struct {

}

func (this *TransController)SetUp(c echo.Context)error{
	tra := models.NewTrans()
	c.Bind(tra)
	tra.Save()
	return c.JSON(http.StatusOK, tra)
}

func (this *TransController)Settle(c echo.Context)error{
	tra := new(models.Trans)
	tra.Id = c.QueryParam("id")
	tra.SetteleDown()
	return c.JSON(http.StatusOK, tra)
}

func (this *TransController)List(c echo.Context)error{
	type rTrans struct {
		models.Trans
		SetupFmt string `json:"setup_fmt"`
	}
	tra := new(models.Trans)
	var rtra_list []rTrans = []rTrans{}
	var tra_list = []models.Trans{}
	tra.All(&tra_list)
	for _,t := range tra_list{
		rt := new(rTrans)
		rt.SetupFmt = time.Unix(t.Setup, 0).Format("2006-01-02 03:04:05")
		rt.Trans = t
		rtra_list = append(rtra_list, *rt)
	}
	return c.JSON(http.StatusOK, rtra_list)
}

func (this *TransController)List_bookid(c echo.Context)error{

	tra := new(models.Trans)
	tra.BookId = c.QueryParam("book_id")
	var tra_list []models.Trans = []models.Trans{}
	tra.Query(&tra_list)
	return c.JSON(http.StatusOK,tra_list)
}

func (this *TransController)Rm(c echo.Context)error{
	tra := new(models.Trans)
	c.Bind(tra)
	tra.Del()

	return c.JSON(http.StatusOK, tra)
}