package controllers

import (
    "github.com/labstack/echo"
    "cn/bkread/booktrans/controllers/index"
    "cn/bkread/booktrans/controllers/book"
    "cn/bkread/booktrans/controllers/trans"
)

func Register(g *echo.Group){
    ind := new(index.Index)
    bk := new(book.BookController)

    index_group := g.Group("/index")
    {
        index_group.GET("/", ind.Hello)
    }

    book_group := g.Group("/book")
    {
        book_group.GET("/info", bk.Info)
        book_group.POST("/reg", bk.RegBook)
        book_group.POST("/rm", bk.Rm)
        book_group.GET("/list", bk.List)
    }

    tra := new(trans.TransController)

    trans_group := g.Group("/trans")
    {
        trans_group.POST("/setup", tra.SetUp)
        trans_group.GET("/list", tra.List)
        trans_group.POST("/rm", tra.Rm)
        trans_group.GET("/list_bookid", tra.List_bookid)
    }
}
