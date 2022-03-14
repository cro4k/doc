package markdown

import "text/template"

var markdownTemplate = "## {{.Name}}\n\n" +
	"### 说明 \n" +
	"> 请求/响应中的子模型请参阅末尾模型列表\n\n" +
	"### 请求地址\n" +
	"> {{.API}}\n\n" +
	"### 请求方式\n" +
	"> {{.Method}}\n\n" +
	"### 请求参数\n" +
	"{{.Req.Name}}\n\n" +
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

var htmlTemplate = "<!doctype html>" +
	"<html lang='zh'>" +
	"<head>" +
	"    <meta charset='utf-8'/>" +
	"    <title>Marked in the browser</title>" +
	"    <style>" +
	"        table {" +
	"            border-style: none none solid solid;" +
	"            border-color: #aaa;" +
	"            border-width: 1px;" +
	"            margin-top: 20px;" +
	"            border-spacing: 0;" +
	"        }" +
	"        table th {" +
	"            border-style: solid solid none none;" +
	"            border-color: #aaa;" +
	"            border-width: 1px;" +
	"            margin-top: 10px;" +
	"            padding: 5px 10px 5px 10px;" +
	"        }" +
	"        table td {" +
	"            border-style: solid solid none none;" +
	"            border-color: #aaa;" +
	"            border-width: 1px;" +
	"            margin-top: 10px;" +
	"            padding: 5px 10px 5px 10px;" +
	"        }" +
	"        blockquote {" +
	"            display: block;" +
	"            margin-block-start: 1em;" +
	"            margin-block-end: 1em;" +
	"            margin-inline-start: 10px;" +
	"            margin-inline-end: 30px;" +
	"        }" +
	"        blockquote p {" +
	"            border-style: none none none solid;" +
	"            padding: 5px 10px 5px 10px;" +
	"            border-color: #479393;" +
	"            background-color: #47939333;" +
	"        }" +
	"        #content{" +
	"            margin-left: 20px;" +
	"        }" +
	"    </style>" +
	"</head>" +
	"<body>" +
	"<div id='content'></div>" +
	"<script src='https://cdn.jsdelivr.net/npm/marked/marked.min.js'></script>" +
	"<script>" +
	"    document.getElementById('content').innerHTML = marked.parse('{{.MD}}');" +
	"</script>" +
	"</body>" +
	"</html>"

var md, _ = template.New("markdown").Parse(markdownTemplate)
var html, _ = template.New("html").Parse(htmlTemplate)
