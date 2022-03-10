package docer

import (
	"reflect"
	"strings"
)

// 记录结构体链路，避免指针无限递归
// eg:
// type Model struct{
//   Child *Model
// }
type tree []string

func (t *tree) join(name string) {
	*t = append(*t, name)
}

func (t tree) contain(name string) bool {
	for _, v := range t {
		if name == v {
			return true
		}
	}
	return false
}

type decoder struct {
	t tree
}

func newDecoder(t tree) *decoder {
	return &decoder{
		t: t,
	}
}

//解析结构体
func (d *decoder) decode(model interface{}) *Model {
	if model == nil {
		return &Model{}
	}
	t := reflect.TypeOf(model)
	v := reflect.ValueOf(model)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		v = v.Elem()
	}
	m := new(Model)
	m.Name = t.Name()
	d.t.join(t.Name())

	rt, n := realType(t, 0)
	m.Array = n > 0

	switch rt.Kind() {
	//case reflect.Map:
	case reflect.Struct:
		for i := 0; i < t.NumField(); i++ {
			if t.Field(i).Tag.Get("doc") == "-" {
				continue
			}
			m.Fields = append(m.Fields, d.decodeField(t.Field(i), v.Field(i)))
		}
	default:
	}
	return m
}

// 解析字段
func (d *decoder) decodeField(t reflect.StructField, v reflect.Value) *Field {
	var f = new(Field)
	f.Name = t.Tag.Get("json")
	f.SetName(t.Name)
	decodeDocTag(f, t.Tag.Get("doc"))
	rt, n := realType(v.Type(), 0)
	switch rt.Kind() {
	case reflect.Struct:
		f.Type = rt.Name()
		if !d.t.contain(rt.Name()) {
			f.Sub = newDecoder(d.t).decode(reflect.New(rt).Interface())
		}
	default:
		f.Type = rt.Name()
	}
	if n > 0 {
		f.Type = strings.Repeat("[]", n) + f.Type
	}
	return f
}

// 解析doc标签
func decodeDocTag(f *Field, doc string) {
	fields := strings.Split(doc, ";")
	for _, v := range fields {
		v = strings.TrimSpace(v)
		tmp := strings.Split(v, ":")
		switch len(tmp) {
		case 0:
			continue
		case 1:
			if v == "required" {
				f.Required = true
			} else {
				f.Comment = v
			}
		case 2:
			switch tmp[0] {
			case "option":
				if len(tmp) > 1 {
					f.Option = tmp[1]
				}
			}
		default:
			continue
		}
	}
}

// 获取指针/切片/数组内的真实类型
func realType(v reflect.Type, i int) (reflect.Type, int) {
	switch v.Kind() {
	case reflect.Slice, reflect.Array:
		i++
		fallthrough
	case reflect.Ptr:
		return realType(v.Elem(), i)
	default:
		return v, i
	}
}
