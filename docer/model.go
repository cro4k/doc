package docer

import (
	"fmt"
	"strings"
)

type Field struct {
	Name     string `json:"name"`
	Type     string `json:"type"`
	Required bool   `json:"required"`
	Comment  string `json:"comment"`
	Option   string `json:"option"`
	Sub      *Model `json:"sub"`
	sub      string
}

func (f *Field) SetName(name string) {
	if f.Name == "" {
		f.Name = name
	}
}

type Model struct {
	Name   string   `json:"name"`
	Fields []*Field `json:"fields"`
	Array  bool     `json:"array"`
}

type Example struct {
	Type uint8 `doc:"required;this is field comment;option:1,2,3"`
}

func (m *Model) GetFields() []*Field {
	if m == nil {
		return nil
	}
	return m.Fields
}

func (m *Model) exampleJSON(prefix string) string {
	var fields []string
	for _, field := range m.GetFields() {
		var text = fmt.Sprintf("\t\"%s\": ", field.Name)
		switch strings.TrimLeft(field.Type, "[]") {
		case "string":
			text = text + "\"\""
		case "int", "int8", "int16", "int32", "int64", "uint", "uint8", "uint16", "uint32", "uint64", "float64", "float32":
			text = text + "0"
		case "bool":
			text = text + "false"
		case "Time":
			text = text + "\"2006-01-02 15:04:05\""
		default:
			if field.Sub != nil {
				text = text + field.Sub.exampleJSON(prefix+"\t")
			} else {
				text = text + "{}"
			}
		}

		var wrap = text
		if n := strings.Count(field.Type, "[]"); n > 0 {
			wrap = strings.Repeat("[", n) + wrap + strings.Repeat("]", n)
		}
		fields = append(fields, prefix+text)
	}
	if m.Array {
		return "[\n" + strings.Join(fields, ",\n") + "\n]"
	} else {
		return "{\n" + strings.Join(fields, ",\n") + "\n" + prefix + "}"
	}
}

func (m *Model) ExampleJSON() string {
	return m.exampleJSON("")
}
