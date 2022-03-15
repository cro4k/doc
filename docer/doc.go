package docer

import (
	"github.com/gin-gonic/gin"
	"reflect"
	"runtime"
	"sort"
	"strings"
)

type Doc struct {
	API     string                 // 接口地址，为空时会自动提取
	Name    string                 // 接口名称
	Method  string                 // 接口方法，为空时会自动提取
	Handler gin.HandlerFunc        // API Handler
	Extra   map[string]interface{} // 自定义数据
	Req     interface{}            // 请求参数
	Rsp     interface{}            // 响应参数
	Group   string                 // 接口分组
	Sort    int                    // 排序 默认只按添加顺序排序，如有需求，自行调用sort.Sort()
}

func (d *Doc) SetAPI(api string) {
	if d.API == "" {
		d.API = api
	}
}

func (d *Doc) SetMethod(method string) {
	if d.Method == "" {
		d.Method = method
	}
}

type Document struct {
	API    string                 `json:"api"`
	Name   string                 `json:"name"`
	Method string                 `json:"method"`
	Extra  map[string]interface{} `json:"extra"`
	Req    *Model                 `json:"req"`
	Rsp    *Model                 `json:"rsp"`
	Group  string                 `json:"group"`
	Sort   int                    `json:"sort"`
}

type Manager struct {
	docs []*Doc
}

func newManager() *Manager {
	return &Manager{}
}

func (m *Manager) add(doc ...*Doc) {
	m.docs = append(m.docs, doc...)
}

func (m *Manager) Decode() Documents {
	var documents []*Document
	for _, v := range m.docs {
		if v.Name == "" {
			funcName := runtime.FuncForPC(reflect.ValueOf(v.Handler).Pointer()).Name()
			if tmp := strings.Split(funcName, "."); len(tmp) > 1 {
				v.Name = tmp[len(tmp)-1]
			} else {
				v.Name = funcName
			}
		}
		if v.Extra == nil {
			v.Extra = make(map[string]interface{})
		}
		doc := &Document{
			API:    v.API,
			Name:   v.Name,
			Method: v.Method,
			Extra:  v.Extra,
			Req:    newDecoder(tree{}).decode(v.Req),
			Rsp:    newDecoder(tree{}).decode(v.Rsp),
			Group:  v.Group,
			Sort:   v.Sort,
		}
		documents = append(documents, doc)
	}
	return documents
}

type Documents []*Document

type DocumentGroup struct {
	Name      string                    `json:"name"`
	Documents Documents                 `json:"documents"`
	Children  map[string]*DocumentGroup `json:"children"`
	Root      bool                      `json:"root"`
}

func (a Documents) Len() int           { return len(a) }
func (a Documents) Less(i, j int) bool { return a[i].Sort < a[j].Sort }
func (a Documents) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Documents) Sort() {
	sort.Sort(a)
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
