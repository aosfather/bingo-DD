package main

import (
	"fmt"
	"github.com/aosfather/bingo_utils/files"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

/**
  字典类型模型定义
*/
type DataDictionary struct {
	Code        string   //唯一编码
	Label       string   //说明
	Paginated   bool     //是否分页
	IndexFields []string `yaml:"fields"` //用于索引，搜索的字段
}

type MetaService interface {
	GetDictionary(key string) DataDictionary
}

//字典模型管理
type DictionaryManager struct {
	Path        string `Value:"dd.path"`
	dictionarys map[string]DataDictionary
}

func (this *DictionaryManager) Init() {
	this.dictionarys = make(map[string]DataDictionary)
	this.load()
}

func (this *DictionaryManager) GetDictionary(key string) DataDictionary {
	return this.dictionarys[key]
}

func (this *DictionaryManager) load() {
	log.Println(this.Path)
	this.loadDir(this.Path)
}

//从目录扫描加载
func (this *DictionaryManager) loadDir(p string) {
	if files.IsFileExist(p) {
		log.Println("start loading")
		fs, err := ioutil.ReadDir(p)
		if err != nil {
			log.Println(err.Error())
			return
		}
		for _, f := range fs {
			fname := fmt.Sprintf("%s/%s", p, f.Name())
			if f.IsDir() {
				this.loadDir(fname)
			} else {
				this.AddFromFile(fname)
			}
		}
	}
}

//从文件中加载
func (this *DictionaryManager) AddFromFile(f string) {
	if files.IsFileExist(f) {
		yamlFile, err := ioutil.ReadFile(f)
		if err != nil {
			log.Println(err.Error())
		} else {
			c := DataDictionary{}
			err = yaml.Unmarshal(yamlFile, &c)
			if err != nil {
				log.Println(err.Error())
			} else {
				this.dictionarys[c.Code] = c
			}
		}

	}

}
