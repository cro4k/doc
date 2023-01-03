package postman

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/cro4k/doc/docer"
	"github.com/cro4k/doc/export/markdown"
	"github.com/google/uuid"
	"io"
	"os"
	"strings"
)

func Postman(w io.Writer, doc *docer.DocumentGroup) error {
	item := &Item{}
	buildGroup(item, doc)
	data := Document{
		Info: Info{
			PostmanID: uuid.New().String(),
			Name:      doc.Name,
			Schema:    SchemaV2_1_0,
		},
		Item: item.Item,
	}
	return json.NewEncoder(w).Encode(data)
}

func Export(filename string, doc *docer.DocumentGroup) error {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	return Postman(file, doc)
}

func buildGroup(parent *Item, doc *docer.DocumentGroup) {
	item := &Item{}
	item.Name = doc.Name
	for _, v := range doc.Documents {
		item.Item = append(item.Item, buildItem(v))
	}
	for _, v := range doc.Children {
		buildGroup(item, v)
	}
	parent.Item = append(parent.Item, item)
}

func buildItem(doc *docer.Document) *Item {
	item := &Item{}
	item.Name = doc.Name
	item.Request = ItemRequest{
		Method: strings.ToUpper(doc.Method),
		Body:   ItemBody{},
		Url: ItemUrl{
			Raw:      "{{protocol}}://" + doc.Path,
			Protocol: "{{protocol}}",
			Host:     []string{"{{host}}"},
			Path:     strings.Split(doc.Path, "/"),
		},
	}
	item.Request.Body.Mode = "raw"
	item.Request.Body.Raw = doc.Req.ExampleJSON()
	item.Request.Body.Options.Raw.Language = "json"

	buf := bytes.NewBuffer([]byte{})
	markdown.Markdown(buf, doc)
	item.Request.Description = buf.String()

	var headers []ItemHeader

	//if doc.Header == nil {
	//	if h := doc.Extra["header"]; h != nil {
	//		doc.Header = h.(map[string]interface{})
	//	}
	//}
	if len(doc.Header) > 0 {
		for k, v := range doc.Header {
			header := ItemHeader{
				Key:   k,
				Value: fmt.Sprintf("%s", v),
				Type:  "text",
			}
			headers = append(headers, header)
		}
	}
	item.Request.Header = headers
	return item
}
