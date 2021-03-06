package story

import (
	"github.com/labstack/echo"
	"net/http"
	"fmt"
	"io"
	"strings"
	"time"
	"io/ioutil"
	"cn/bkread/booktrans/models"
	"strconv"
	"gopkg.in/mgo.v2/bson"
)

type Controller_stroy struct {

}

type reBody struct {
	tok, ctp, cuid, spd, per, tex, lan string
}
func (rb *reBody)String()(s string){
	s = fmt.Sprintf("tok=%s&ctp=%s&cuid=%s&spd=%s&per=%s&tex=%s&lan=%s",
		rb.tok,rb.ctp,rb.cuid,rb.spd,rb.per, rb.tex, rb.lan)
	return
}
func (rb *reBody) Read() io.Reader {
	return strings.NewReader(rb.String())
}

func getAudio(r io.Reader)(res []byte){
	url := "http://tsn.baidu.com/text2audio"
	req, _ := http.NewRequest("POST", url, r)
	client := &http.Client{
		Timeout: 20 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	res, _ = ioutil.ReadAll(resp.Body)
	return
}

func (this *Controller_stroy)Tell(c echo.Context)error {
	const JIE int = 1*200
	id := c.QueryParam("id")
	story := new(models.Story)
	story.ID = bson.ObjectIdHex(id)
	*story, _ = story.GetSingleData()
	text := story.Body
	//text = "我是谁"

	body := reBody{
		tok:"24.56f654b217197fc345e95f004bd98067.2592000.1517095892.282335-10241789",
		ctp:"1",
		cuid:"iid",
		spd:"3",
		per:"2",
		tex:text,
		lan:"zh",
	}
	c.Response().Header().Set(echo.HeaderContentType, "audio/mp3")
	c.Response().WriteHeader(http.StatusOK)

	textRune := []rune(text)
	if len(textRune) > JIE {
		step := len(textRune)/JIE
		for i:= 0; i < step; i++ {
			var sub []byte
			if i == step-1 { // 判断是不是最后一步
				body.tex = string(textRune[i*JIE:])
			}else{
				body.tex = string(textRune[i*JIE:(i+1)*JIE-1])
			}
			blob := strings.NewReader(body.String())
			sub = getAudio(blob)
			c.Response().Write(sub)
			c.Response().Flush()
		}

	}else {
		blob := strings.NewReader(body.String())
		c.Response().Write(getAudio(blob))
		c.Response().Flush()
	}

	return nil
}


func (this *Controller_stroy)List(c echo.Context)error{
	story := new(models.Story)
	storyList := []models.Story{}
	pn, _ := strconv.Atoi(c.QueryParam("pn"))
	p, _ := strconv.Atoi(c.QueryParam("p"))
	skip := (p-1)*pn
	selector := map[string]int{"cate":1,"title":1, "body":1}
	story.GetPageData(skip, pn, nil, selector, &storyList)

	//time.Sleep(10*time.Second)

	return c.JSON(http.StatusOK, storyList)
}

func (this *Controller_stroy)Search(c echo.Context)error{
	story := new(models.Story)
	storyList := []models.Story{}
	pn, _ := strconv.Atoi(c.QueryParam("pn"))
	p, _ := strconv.Atoi(c.QueryParam("p"))
	kw := c.QueryParam("title")
	skip := (p-1)*pn
	selector := map[string]int{"cate":1,"title":1, "body":1}
	query := map[string]interface{}{"title":map[string]string{"$regex":kw}}
	total, _ := story.GetPageData(skip, pn, query, selector, &storyList)

	return c.JSON(http.StatusOK, map[string]interface{}{"total":total, "story_list":storyList})
}
