package controllers

import (
    "github.com/labstack/echo"
    // "cn/bkread/booktrans/controllers/book"
    // "cn/bkread/booktrans/controllers/trans"
    "cn/bkread/booktrans/controllers/story"
)

func Register(g *echo.Group){
    //ind := new(index.Index)
	//
    //index_group := g.Group("/index")
    //{
    //    index_group.GET("/", ind.Hello)
    //}

    // bk := new(book.BookController)
    // book_group := g.Group("/book")
    // {
    //     book_group.GET("/info", bk.Info)
    //     book_group.POST("/reg", bk.RegBook)
    //     book_group.POST("/rm", bk.Rm)
    //     book_group.GET("/list", bk.List)
    //     book_group.GET("/suggest", bk.Suggest)
    // }

    // tra := new(trans.TransController)

    // trans_group := g.Group("/trans")
    // {
    //     trans_group.POST("/setup", tra.SetUp)
    //     trans_group.GET("/settle", tra.Settle)
    //     trans_group.GET("/list", tra.List)
    //     trans_group.POST("/rm", tra.Rm)
    //     trans_group.GET("/list_bookid", tra.List_bookid)
    // }

    story := new(story.Controller_stroy)

    story_group := g.Group("/story")
    {
        story_group.GET("/tell", story.Tell)
        story_group.GET("/list", story.List)
        story_group.GET("/search", story.Search)
    }
}
