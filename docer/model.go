package docer

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
