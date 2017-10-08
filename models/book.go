package models

import (
	"encoding/json"
	"errors"
	"crypto/sha256"
	"fmt"
	"time"
)

const (
	BOOK_KEY = "books"
)

type Book struct {
	Db	`json:",omitempty"`
	Id string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
}

func (bk *Book)Hset() error {
	_, err := bk.Client().Do("hset", BOOK_KEY, bk.Id, string(bk.En()))
	return err
}

func (bk *Book)Hget()(Book, error){
	var val []string
	var book Book
	val, err := bk.Client().Do("hget", BOOK_KEY, bk.Id)
	if err != nil {
		return book,err
	}
	if len(val) > 1 {
		content := val[1]
		err = json.Unmarshal([]byte(content), &book)
	}else{
		err = errors.New("没找到")
	}
	return book, err
}

func (bk *Book)All(data *[]Book)(error){
	var val []string
	val, err := bk.Client().Do("hgetall", BOOK_KEY)
	if err != nil {
		return err
	}
	if len(val) > 1 {
		for i:= 1;i< len(val); i += 2{
			var book Book
			json.Unmarshal([]byte(val[i+1]),&book)
			*data = append(*data, book)
		}
	}
	return errors.New("空")
}

func (bk *Book)Del() error {
	_, err := bk.Client().Do("hdel", BOOK_KEY, bk.Id)
	return err
}

func (bk *Book)SetId(uid string)string{
	now := time.Now()
	stamp := now.Unix()
	var id string
	cry := sha256.Sum256([]byte(fmt.Sprintf("%s%s%x", bk.Isbn, uid, stamp)))
	id = fmt.Sprintf("%x", cry)
	bk.Id = id
	return id
}
func (bk *Book)En()[]byte{
	en, _ := json.Marshal(bk)
	return en
}
