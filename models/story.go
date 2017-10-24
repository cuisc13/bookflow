package models

import (
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
)

type Story struct {

	Db `bson:",omitempty"`
	ID  bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Cid string `bson:"c_id,omitempty" json:"c_id"`
	Cate []string `bson:"cate,omitempty" json:"cate"`
	Title string `bson:"title,omitempty" json:"title"`
	Author string `bson:"author,omitempty" json:"author"`
	Body string `bson:"body,omitempty" json:"body"`
}

func (s *Story)CName()string{
	return "articles"
}

func (s *Story)Database()*mgo.Database{
	return _MgoSession.DB("tonghuadb")
}

func (s *Story)Collection(cname string) *mgo.Collection{
	return s.Database().C(cname)
}

func (s *Story)Find(cname string, query, selector interface{})*mgo.Query{
	return s.Database().C(cname).Find(query).Select(selector)
}

func (s *Story) GetPageData(skip, limit int, query,selector interface{}, data *[]Story)(n int, err error){
	q :=s.Find(s.CName(), query, selector)
	n, err = q.Count()
	err = q.Skip(skip).Limit(limit).All(data)
	return
}

// GetSingleData 获取单条数据
func (s *Story) GetSingleData() (Story, error) {
	var story Story
	err := s.Collection(s.CName()).FindId(s.ID).One(&story)
	return story, err
}
