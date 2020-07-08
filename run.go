package main

import (
	"github.com/aosfather/bingo_mvc/context"
	"github.com/aosfather/bingo_mvc/fasthttp"
)

func main() {
	dispatch := &fasthttp.FastHTTPDispatcher{}
	boot := context.Boot{}
	boot.Init(dispatch, load)
	boot.Start()
}

func load() []interface{} {
	return []interface{}{&DictionaryManager{}, &SearchEngine{}, &QueryController{}}
}
