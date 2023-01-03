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
	for _, v := range m.GetFields() {
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

	for _, field := range data.Req.GetFields() {
		expand(data.Models, field.Sub)
	}

	for _, field := range data.Rsp.GetFields() {
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

	data := string(b.Bytes())
	data = strings.ReplaceAll(data, "\n", "\\n")
	data = strings.ReplaceAll(data, "\t", "\\t")
	return html.Execute(w, HTMLData{
		MD:   data,
		Path: path,
	})
}

func Export(root string, docs *docer.DocumentGroup, toHtml ...bool) error {
	//os.RemoveAll(root)
	if err := export(root+"/"+docs.Name, docs, toHtml...); err != nil {
		return err
	}
	return nil
}

func export(path string, group *docer.DocumentGroup, toHtml ...bool) error {
	path = strings.TrimRight(path, "/")
	if path == "" {
		path = "."
	}

	var routers []map[string]interface{}
	os.MkdirAll(path, 0777)
	for _, doc := range group.Documents {
		routers = append(routers, map[string]interface{}{
			"link": fmt.Sprintf("./%s.html", doc.Name),
			"name": doc.Name + ".html",
		})
		if err := exportDocument(path, doc, toHtml...); err != nil {
			return err
		}
	}

	for _, child := range group.Children {
		var name = child.Name
		if name == "" {
			name = "default"
		}
		os.MkdirAll(path+"/"+name, 0777)
		if err := export(path+"/"+name, child, toHtml...); err != nil {
			return err
		}
		routers = append(routers, map[string]interface{}{
			"link": fmt.Sprintf("./%s/index.html", name),
			"name": name,
		})
	}

	if len(toHtml) > 0 && toHtml[0] {
		name := fmt.Sprintf("%s/index.html", path)
		os.Remove(name)
		indexFile, err := os.OpenFile(name, os.O_CREATE|os.O_RDWR, 0644)
		if err != nil {
			return err
		}
		defer indexFile.Close()
		return index.Execute(indexFile, map[string]interface{}{"routers": routers})
	}
	return nil
}

func exportDocument(path string, doc *docer.Document, toHtml ...bool) error {
	var filename string
	if len(toHtml) > 0 && toHtml[0] {
		filename = fmt.Sprintf("%s/%s.html", path, doc.Name)
	} else {
		filename = fmt.Sprintf("%s/%s.md", path, doc.Name)
	}
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
