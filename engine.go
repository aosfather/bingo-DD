package main

import (
	"encoding/json"
	"fmt"
	rs "github.com/aosfather/bingo_utils/redis"
	"log"
)

type SearchEngine struct {
	Host    string      `Value:"dd.redis.addr"`
	DBIndex int         `Value:"dd.redis.db"`
	Pwd     string      `Value:"dd.redis.pwd"`
	Meta    MetaService `Inject:""`
	se      *rs.SearchEngine
}

func (this *SearchEngine) Init() {
	option := rs.SearchOption{10, 1000, this.Host, this.DBIndex, this.Pwd}
	this.se = &rs.SearchEngine{}
	this.se.Init(option)

}

func (this *SearchEngine) Search(name string, field ...rs.Field) *rs.PageSearchResult {
	dict := this.Meta.GetDictionary(name)
	if dict.Code == "" {
		return nil
	}
	this.se.CreateIndex(name)
	rs := this.se.Search(name, field...)
	return rs

}

//增加索引
func (this *SearchEngine) Add(name string, raw []byte) {
	if name == "" || len(raw) == 0 {
		return
	}
	//构建源对象
	s := rs.SourceObject{}
	s.TargetObject = rs.TargetObject{Id: "", Data: raw}
	dictionary := this.Meta.GetDictionary(name)
	//构建字段取值作为索引字段
	fields := make(map[string][]string)
	datas := make(map[string]string)
	json.Unmarshal(raw, &datas)
	log.Println(datas)
	log.Println(dictionary)
	for _, field := range dictionary.IndexFields {
		log.Println(field)
		fields[field] = []string{fmt.Sprintf("%v", datas[field])}
	}
	s.Fields = fields
	this.se.LoadSource(name, &s)
	log.Println("add to index %s", name)
}
