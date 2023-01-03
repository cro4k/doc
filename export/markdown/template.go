package markdown

import "text/template"

var markdownTemplate = "## {{.Name}}\n\n" +
	"### 请求地址\n" +
	"> {{.API}}\n\n" +
	"### 请求方式\n" +
	"> {{.Method}}\n\n" +
	"### 请求参数\n" +
	"{{.Req.Name}}\n\n" +
	"Header:\n\n" +
	"| 名称 | 说明 |\n" +
	"|-----|-----|\n" +
	"{{ range $k,$v := .Header }}" +
	"|{{$k}}|{{.}}|\n" +
	"{{end}}\n" +
	"Body:\n\n" +
	"| 字段       | 必填  | 类型     | 备注  | 可选值 |\n" +
	"|----------|-----|--------|-----|-----|\n" +
	"{{ range $i,$field := .Req.Fields }}" +
	"| {{.Name}} | {{.Required}}   | {{.Type}} | {{.Comment}}| {{.Option}}|\n" +
	"{{end}}\n" +
	"### 请求示例\n" +
	"```json\n{{.RequestJSON}}\n```\n" +
	"### 响应数据\n" +
	"{{.Rsp.Name}}\n\n" +
	"| 字段       |  类型     | 备注  | 可选值 |\n" +
	"|----------|--------|-----|-----|\n" +
	"{{ range $i,$field := .Rsp.Fields }}" +
	"| {{.Name}} | {{.Type}} | {{.Comment}}| {{.Option}}|\n" +
	"{{end}}\n" +
	"### 响应示例\n" +
	"```json\n{{.ResponseJSON}}\n```\n\n" +
	"### 模型列表\n" +
	"{{ range $name,$val := .Models }}\n" +
	"{{.Name}}\n\n" +
	"| 字段       | 必填  | 类型     | 备注  | 可选值 |\n" +
	"|----------|-----|--------|-----|-----|\n" +
	"{{ range $i,$field := .Fields }}" +
	"| {{.Name}} | {{.Required}}   | {{.Type}} | {{.Comment}}| {{.Option}}|\n" +
	"{{end}}\n" +
	"{{ end}}\n" +
	"本文档由 [doc](https://github.com/cro4k/doc) 自动生成"

var htmlTemplate = "<!doctype html>\n" +
	"<html lang='zh'>\n" +
	"<head>\n" +
	"    <meta charset='utf-8'/>\n" +
	"    <title>Docer</title>\n" +
	"    <style>\n" +
	"        table {\n" +
	"            border-style: none none solid solid;\n" +
	"            border-color: #aaa;\n" +
	"            border-width: 1px;\n" +
	"            margin-top: 20px;\n" +
	"            border-spacing: 0;\n" +
	"        }\n" +
	"        table th {\n" +
	"            border-style: solid solid none none;\n" +
	"            border-color: #aaa;\n" +
	"            border-width: 1px;\n" +
	"            margin-top: 10px;\n" +
	"            padding: 5px 10px 5px 10px;\n" +
	"        }\n" +
	"        table td {\n" +
	"            border-style: solid solid none none;\n" +
	"            border-color: #aaa;\n" +
	"            border-width: 1px;\n" +
	"            margin-top: 10px;\n" +
	"            padding: 5px 10px 5px 10px;\n" +
	"        }\n" +
	"        blockquote {\n" +
	"            display: block;\n" +
	"            margin-block-start: 1em;\n" +
	"            margin-block-end: 1em;\n" +
	"            margin-inline-start: 10px;\n" +
	"            margin-inline-end: 30px;\n" +
	"        }\n" +
	"        blockquote p {\n" +
	"            border-style: none none none solid;\n" +
	"            padding: 5px 10px 5px 10px;\n" +
	"            border-color: #479393;\n" +
	"            background-color: #47939333;\n" +
	"        }\n" +
	"        #content{\n" +
	"            margin-left: 20px;\n" +
	"        }\n" +
	"    </style>\n" +
	"</head>\n" +
	"<body>\n" +
	"<div id='content'></div>\n" +
	"<script src='https://cdn.jsdelivr.net/npm/marked/marked.min.js'></script>\n" +
	"<script>\n" +
	"    var data = '{{.MD}}'\n" +
	"    document.getElementById('content').innerHTML = marked.parse(data);\n" +
	"</script>\n" +
	"</body>\n" +
	"</html>\n"

var indexTemplate = "<!DOCTYPE html>\n" +
	"<html lang='zh'>\n" +
	"<head>\n" +
	"    <meta charset='UTF-8'>\n" +
	"    <title>Title</title>\n" +
	"</head>\n" +
	"<body>\n" +
	"    <ul>\n" +
	"{{ range $k,$v := .routers }}" +
	"        <li> <a href='{{.link}}'>{{.name}}</a> </li>\n" +
	"{{ end }}" +
	"    </ul>\n" +
	"</body>\n" +
	"</html>\n"

var md, _ = template.New("markdown").Parse(markdownTemplate)
var html, _ = template.New("html").Parse(htmlTemplate)
var index, _ = template.New("index").Parse(indexTemplate)
