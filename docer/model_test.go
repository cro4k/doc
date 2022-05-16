package docer

import (
	"testing"
	"time"
)

type ExampleRequest struct {
	Name      string      `json:"name"`
	Data      ExampleData `json:"data"`
	CreatedAt time.Time   `json:"created_at,omitempty"`
}

type ExampleData struct {
	Value int `json:"value"`
	Data  Sub `json:"data"`
}

type Sub struct {
	Value string `json:"value"`
}

func TestModelExampleJSON(t *testing.T) {
	m := newDecoder(tree{}).decode(&ExampleRequest{})
	for _, f := range m.Fields {
		t.Log(f.Name, f.Type)
	}
	t.Log("\n" + m.exampleJSON(""))
}
