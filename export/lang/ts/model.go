package ts

import (
	"bytes"
	"github.com/cro4k/doc/docer"
	"html/template"
)

var _modelTplText = `interface {{.Name}} {
	{{range .Fields}} {{.Name}}:{{.Type}},
	{{end}}
}
`
var _apiTplText = `	async {{.Name}}(body:{{.Request}},header?:any) {
	return axios.post(this.api.{{.Name}},body,{{headers:header}})
}`

var _modelTpl, _ = template.New("model").Parse(_modelTplText)
var _apiTpl, _ = template.New("api").Parse(_apiTplText)

type Model struct {
	Name   string
	Fields []Field
}

type Field struct {
	Name string
	Type string
}

func ConvertModel(doc *docer.Model) *Model {
	m := &Model{Name: doc.Name}
	for _, v := range doc.Fields {
		var name = v.Name
		var typ = "any"
		if !v.Required {
			name = name + "?"
		}
		switch v.Type {
		case "string":
			typ = "string"
		case "int", "int8", "int16", "int32", "int64":
			fallthrough
		case "uint", "uint8", "uint16", "uint32", "uint64":
			fallthrough
		case "float32", "float64":
			typ = "number"
		case "bool":
			typ = "boolean"
		}
		m.Fields = append(m.Fields, Field{
			Name: name,
			Type: typ,
		})
	}
	return m
}

func (m *Model) Format() string {
	buf := bytes.NewBuffer([]byte{})
	_ = _modelTpl.Execute(buf, m)
	return buf.String()
}

type API struct {
	Name    string
	Request string
}

func (a *API) Format() string {
	buf := bytes.NewBuffer([]byte{})
	_ = _apiTpl.Execute(buf, a)
	return buf.String()
}

func ConvertAPI(doc docer.Document) *API {
	
}
