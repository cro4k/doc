package docer

import (
	"strings"
)

type Document struct {
	Path   string
	Name   string
	Method string
	Req    *Model
	Rsp    *Model
	Group  string
	Header KV
	Extra  KV
}

type Documents []*Document

type DocumentGroup struct {
	Name      string                    `json:"name"`
	Documents Documents                 `json:"documents"`
	Children  map[string]*DocumentGroup `json:"children"`
	Root      bool                      `json:"root"`
}

func (a Documents) Group(seps ...string) *DocumentGroup {
	var sep = "/"
	if len(seps) > 0 {
		sep = seps[0]
	}
	var groups = &DocumentGroup{Children: make(map[string]*DocumentGroup), Documents: make([]*Document, 0), Root: true}
	for _, doc := range a {
		var ptr = groups
		names := strings.Split(doc.Group, sep)
		for n, name := range names {
			if ptr.Children[name] == nil {
				ptr.Children[name] = &DocumentGroup{Name: name, Children: make(map[string]*DocumentGroup), Documents: make([]*Document, 0)}
			}
			ptr = ptr.Children[name]
			if n == len(names)-1 {
				ptr.Documents = append(ptr.Documents, doc)
			}
		}
	}
	return groups
}
