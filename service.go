package main

import (
	"encoding/json"
	"github.com/aosfather/bingo_mvc"
	rs "github.com/aosfather/bingo_utils/redis"
	"log"
)

type QueryParameters struct {
	Name       string            `json:"dict"` //字典名称
	Parameters map[string]string `json：“parameters”`
	Orders     []Order           //排序设置
}

//排序设置
type Order struct {
	Field string //字段
	dec   bool   // 降序

}

//提交索引的数据
type RawData struct {
	Name  string            `json:"dict"`  //字典名称
	Datas []json.RawMessage `json:"datas"` //字典数据
}

//查询接口
type QueryController struct {
	SE    *SearchEngine `mapper:"name(query);url(/query);method(POST);style(JSON)" Inject:""`
	Index string        `mapper:"name(add);url(/add);method(POST);style(JSON)"`
}

func (this *QueryController) GetHandles() bingo_mvc.HandleMap {
	result := bingo_mvc.NewHandleMap()
	result.Add("query", this.query, &QueryParameters{})
	result.Add("add", this.add, &RawData{})
	return result
}

func (this *QueryController) query(a interface{}) interface{} {
	parameters := a.(*QueryParameters)
	var fields []rs.Field
	for k, v := range parameters.Parameters {
		fields = append(fields, rs.Field{Key: k, Value: v})
	}

	return this.SE.Search(parameters.Name, fields...)
}

func (this *QueryController) add(a interface{}) interface{} {
	raw := a.(*RawData)
	for _, v := range raw.Datas {
		this.SE.Add(raw.Name, v)
		log.Println(string(v))
	}

	return ""
}
