package models

import (
	"time"
	"fmt"
	"encoding/json"
)

const (
	TRANS_KEY = "trans"
	TRANS_INDEX_PREFIX = "TRANS_INDEX_BY_BOOK_ID"
)

type Trans struct {
	Db
	BookId string `json:"book_id"`
	Id string `json:"id"`
	From string `json:"from"`
	To string `json:"to"`
	Setup int64 `json:"setup"`
	Settled int64 `json:"settled"`
	Previous string `json:"previous"`
}

func (this *Trans)SetIndex(){
	_, err := this.Client().Do("zset", TRANS_INDEX_PREFIX+this.BookId, this.Id,this.Setup)
	if err != nil {
		panic(err)
	}
}

func (this *Trans)Genesis(){
	this.Id = this.TransId()
	this.From = "Tony"
	this.Setup = time.Now().Unix()
	this.Save()
}

func (this *Trans)Get(){
	val, _ := this.Client().Do("hget", TRANS_KEY, this.Id)
	if len(val) > 1 {
		*this = this.De(val[1])
	}
}

func (this *Trans)Save(){
	val := this.En()
	this.Client().Do("hset", TRANS_KEY, this.Id, string(val))
	this.SetIndex()
}
func NewTrans() *Trans{
	var tra = new(Trans)
	tra.Id = tra.TransId()
	tra.Setup = time.Now().Unix()
	return tra
}

func (this *Trans)SetteleDown(){
	this.Get()
	this.Settled = time.Now().Unix()
	this.Save()
}

func (this *Trans)TransId()string{
	now := time.Now()
	val, err := this.Client().Do("incr", TRANS_KEY+"_id")
	if err != nil{
		panic(err)
	}
	if len(val)>1 {
		id := fmt.Sprintf("%s%d", val[1], now.Unix()*2)
		return id
	}
	return ""
}

func (this *Trans)Del(){
	this.Get()
	fmt.Println(TRANS_INDEX_PREFIX+this.BookId)
	this.Client().Do("hdel", TRANS_KEY, this.Id)
	val,_ := this.Client().Do("zclear", TRANS_INDEX_PREFIX+this.BookId)
	fmt.Println(val)
}

func (this *Trans)All(data *[]Trans){
	var val []string
	val, err := this.Client().Do("hgetall", TRANS_KEY)
	if err != nil {
		panic(err)
	}
	if len(val)>1 {
		for i:= 1;i < len(val);i+= 2 {
			*data = append(*data, this.De(val[i+1]))
		}
	}
}

func (this *Trans)Query(data *[]Trans){
	var val []string
	val, err := this.Client().Do("zrange", TRANS_INDEX_PREFIX+this.BookId, 0, -1)
	if err != nil {
		panic(err)
	}
	if len(val) > 1 {
		var trans_id_list []string
		for j:= 1;j<len(val);j+=2{
			trans_id_list = append(trans_id_list, val[j])
		}
		n_val, _ := this.Client().Do("multi_hget", TRANS_KEY,trans_id_list)
		if len(n_val) > 1 {
			for i:= 1;i < len(n_val);i+= 2 {
				*data = append(*data, this.De(n_val[i+1]))
			}
		}
	}
}

func (this *Trans)En()[]byte{
	en, _ := json.Marshal(this)
	return en
}

func (this *Trans)De(data string)Trans{
	var tra Trans
	json.Unmarshal([]byte(data), &tra)
	return tra
}
