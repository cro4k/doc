package docer

import (
	"github.com/gin-gonic/gin"
	"strings"
)

type Doc struct {
	Path    string
	Method  string
	Handler string
	Name    string
	Comment string
	Header  KV
	Extra   KV
	Group   string
	Req     interface{}
	Rsp     interface{}
}

type Manager struct {
	docs []*Doc
}

func newManager() *Manager {
	return &Manager{}
}

func (m *Manager) add(doc ...*Doc) *Manager {
	m.docs = append(m.docs, doc...)
	return m
}

func (m *Manager) Decode() Documents {
	var documents []*Document
	for _, v := range m.docs {
		if v.Name == "" {
			//funcName := runtime.FuncForPC(reflect.ValueOf(v.Handler).Pointer()).Name()
			funcName := v.Handler
			if tmp := strings.Split(funcName, "."); len(tmp) > 1 {
				v.Name = tmp[len(tmp)-1]
			} else {
				v.Name = funcName
			}
		}
		if v.Header == nil {
			v.Header = KV{}
		}
		if v.Extra == nil {
			v.Extra = KV{}
		}
		if v.Group == "" {
			if n := strings.LastIndex(v.Name, "/"); n >= 0 {
				v.Group = v.Name[:n]
				v.Name = v.Name[n+1:]
			}
		}
		doc := &Document{
			Path:    v.Path,
			Name:    v.Name,
			Comment: v.Comment,
			Method:  v.Method,
			Header:  v.Header,
			Extra:   v.Extra,
			Req:     newDecoder(tree{}).decode(v.Req),
			Rsp:     newDecoder(tree{}).decode(v.Rsp),
			Group:   v.Group,
		}
		documents = append(documents, doc)
	}
	return documents
}

func Decode(g *gin.Engine) Documents {
	routers := DecodeGin(g)
	docs := FromAnnotation(routers)
	return newManager().add(docs...).Decode()
}
