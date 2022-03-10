### 基于 [gin](https://github.com/gin-gonic/gin) 的自动化文档工具（开发中）
Automatic generate API Documentation in Golang
#### Usage:
```
package main

import (
    "encoding/json"
    "fmt"
    "github.com/cro4k/doc/docer"
    "github.com/gin-gonic/gin"
)

type Request struct {
    Username string `json:"username" doc:"required;this is username"`
    Type     int    `json:"type"     doc:"this is type;option:1,2,3"` 
}

type Response struct {
    Code    int     `json:"code"`
    Message string  `json:"message"`
    Username string `json:"username"`
    Password string `json:"-" doc:"-"`
}

func main() {
    eng := docer.Wrap(gin.Default())

    eng.GET("/api/version", &docer.Doc{
        API: "",          // 接口地址，为空时会自动提取
        Method: "",       // 接口方法，为空时会自动提取
        Name: "api demo", // 接口名称
        Req: &Request{},  // 请求参数 
        Rsp: &Response{}, // 响应参数
        Extra: map[string]string{}, // 自定义数据
        Handler: func(ctx *gin.Context) { //api handler （如果name为空，会自动反射读取handler的方法名，不建议使用匿名函数）
        },
    })

    group := eng.Group("/api")
    group.POST("/hello", &docer.Doc{
        Handler: func(ctx *gin.Context) { 
        },
    })

    docs := eng.Docs().Decode()
    b, _ := json.MarshalIndent(docs, "", "  ")
    fmt.Println(string(b))
}

```

#### Features:
- 自动提取api地址
- 自动解析请求响应参数
- 只需要封装Handler方法，尽量降低代码侵入