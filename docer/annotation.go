package docer

import (
	"github.com/cro4k/annotation/core"
	"strings"
)

var elements map[string]*core.Element

func Init(ele map[string]*core.Element) {
	elements = ele
}

func FromAnnotation(routers []*Router) []*Doc {
	var docs []*Doc
	for _, router := range routers {
		ele := elements[router.FuncName]
		var doc = &Doc{
			Path:    router.Path,
			Method:  router.Method,
			Handler: router.FuncName,
		}
		if ele == nil {
			docs = append(docs, doc)
			continue
		}
		for _, ann := range ele.Annotations {
			n := strings.Index(ann.Raw, " ")
			if n <= 0 {
				continue
			}
			var key = ann.Raw[:n]
			var content = strings.TrimSpace(ann.Raw[n+1:])
			switch key {
			case "req":
				if len(ann.Relation) > 0 {
					doc.Req = ann.Relation[0]
				}
			case "rsp":
				if len(ann.Relation) > 0 {
					doc.Rsp = ann.Relation[0]
				}
			case "name":
				doc.Name = content
			case "comment":
				doc.Comment = content
			case "group":
				doc.Group = content
			case "header":
				if doc.Header == nil {
					doc.Header = KV{}
				}
				k, v := splitKV(content)
				doc.Header.Add(k, v)
			case "extra":
				if doc.Extra == nil {
					doc.Extra = KV{}
				}
				k, v := splitKV(content)
				doc.Extra.Add(k, v)
			}

			if strings.HasPrefix(ann.Raw, "req") {
				if len(ann.Relation) > 0 {
					doc.Req = ann.Relation[0]
				}
			} else if strings.HasPrefix(ann.Raw, "rsp") {
				if len(ann.Relation) > 0 {
					doc.Rsp = ann.Relation[0]
				}
			}
		}
		docs = append(docs, doc)
	}
	return docs
}

func splitKV(content string) (string, string) {
	n := strings.Index(content, ":")
	if n <= 0 {
		return content, ""
	}
	key := strings.TrimSpace(content[:n])
	val := strings.TrimSpace(content[n+1:])
	return key, val
}
