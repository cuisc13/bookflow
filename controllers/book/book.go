package book

import (
    "net/http"
    "github.com/labstack/echo"
    "cn/bkread/booktrans/models"
    "fmt"
)

type BookController struct {

}

func (this *BookController) Info(c echo.Context) error{
    bid := c.QueryParam("bid")
    var book models.Book
    book.Id = bid
    book, err := book.Hget()

    if err != nil {
        return err
    }
    return c.JSON(http.StatusOK, book)
}

// api/v1/book/
func (this *BookController) RegBook(c echo.Context) error {
    type req struct {
        models.Book
        UserId string `json:"user_id"`
    }
    var r = new(req)
    if err := c.Bind(r); err != nil {
        return err
    }

    r.Book.SetId(r.UserId)

    if err := r.Hset(); err != nil {
        return err
    }
    tra := new(models.Trans)
    tra.BookId = r.Id
    tra.To = r.UserId
    tra.Genesis()
    return c.JSON(http.StatusOK, r.Book)
}

// api/v1/book/rm
func (this *BookController) Rm(c echo.Context)error{
    type req struct {
        models.Book
        UserId string `json:"user_id"`
    }
    r := new(req)
    if err := c.Bind(r); err != nil {
        return err
    }
    if err:= r.Del();err != nil {
        return err
    }
    return c.JSON(http.StatusOK, r.Book)
}

func (this *BookController)List(c echo.Context)error {
    var book = new(models.Book)
    var book_list = []models.Book{}
    if err := book.All(&book_list); err != nil {
        //return c.JSON(http.StatusOK, []interface{}{})
    }
    fmt.Println(new(models.Trans).TransId())
    return c.JSON(http.StatusOK, book_list)
}
