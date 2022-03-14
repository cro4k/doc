package markdown

import (
	"bytes"
	"fmt"
	"github.com/cro4k/doc/docer"
	"io"
	"os"
	"strings"
)

type Data struct {
	docer.Document
	Models       map[string]*docer.Model
	RequestJSON  string
	ResponseJSON string
}

func expand(models map[string]*docer.Model, m *docer.Model) {
	if m == nil {
		return
	}
	models[m.Name] = m
	for _, v := range m.Fields {
		if v.Sub != nil {
			expand(models, v.Sub)
		}
	}
}

// Markdown
// to markdown
func Markdown(w io.Writer, doc *docer.Document) error {
	data := Data{Document: *doc}
	data.Models = make(map[string]*docer.Model)

	for _, field := range data.Req.Fields {
		expand(data.Models, field.Sub)
	}

	for _, field := range data.Rsp.Fields {
		expand(data.Models, field.Sub)
	}

	data.RequestJSON = data.Req.ExampleJSON()
	data.ResponseJSON = data.Rsp.ExampleJSON()
	return md.Execute(w, data)
}

type HTMLData struct {
	MD   string //markdown
	Path string //页面路径
}

// HTML
// 导出html（基于第三方库 https://github.com/markedjs/marked 来解析 markdown）
func HTML(w io.Writer, doc *docer.Document, path string) error {
	b := bytes.NewBuffer([]byte{})
	err := Markdown(b, doc)
	if err != nil {
		return err
	}
	return html.Execute(w, HTMLData{
		MD:   string(b.Bytes()),
		Path: path,
	})
}

func Export(root string, docs *docer.DocumentGroup, toHtml ...bool) error {
	return export(root+"/"+docs.Name, docs, toHtml...)
}

func export(path string, group *docer.DocumentGroup, toHtml ...bool) error {
	path = strings.TrimRight(path, "/")

	os.MkdirAll(path, 0777)
	for _, v := range group.Documents {
		if err := exportDocument(path, v, toHtml...); err != nil {
			return err
		}
	}
	for _, child := range group.Children {
		if err := export(path+"/"+child.Name, child); err != nil {
			return err
		}
	}

	return nil
}

func exportDocument(path string, doc *docer.Document, toHtml ...bool) error {
	filename := fmt.Sprintf("%s/%s.md", path, doc.Name)
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	if len(toHtml) > 0 && toHtml[0] {
		return HTML(f, doc, path)
	}
	return Markdown(f, doc)
}
