package main

import (
	"github.com/cro4k/doc/docer"
	"github.com/cro4k/doc/export/markdown"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	eng := docer.Wrap(gin.Default())

	eng.GET("/api/version", Version)

	group := eng.Group("/api")
	group.POST("/hello", Hello)

	//b, _ := json.MarshalIndent(eng.Docs().Decode().Group(), "", "  ")
	//fmt.Println(string(b))

	//http.ListenAndServe(":8080", eng)

	err := markdown.Export("./boom/view", eng.Docs().Decode().Group(), true)
	log.Println(err)
}

var (
	Hello = &docer.Doc{
		Req: &Request{},
		Rsp: &Response{},
		Extra: map[string]interface{}{
			"备注": "",
		},
		Handler: hello,
		Group:   "",
	}

	Version = &docer.Doc{
		Name: "version",
		Rsp:  &Response{},
		Handler: func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, map[string]interface{}{
				"version": "v0.0.1",
			})
		},
		Group: "demo",
	}
)

func hello(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "hello",
	})
}

type Request struct {
	ID    string     `json:"id"  doc:"required;id"`
	Type  [][]int    `json:"type" doc:"request type;option:0,1"`
	Child *Request   `json:"child"`
	Data  *Data      `json:"data2"`
	Info  [][][]Info `json:"info"`
}

type Response struct {
	Code    int    `json:"code"    doc:""`
	Message string `json:"message" doc:""`
}

type Data struct {
	Name string `json:"name"`
	Info *Info  `json:"info"`
}

type Info struct {
	Hello string `json:"hello"`
}
