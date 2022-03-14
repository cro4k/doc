package docer

import "testing"

type ExampleRequest struct {
	Name string      `json:"name"`
	Data ExampleData `json:"data"`
}

type ExampleData struct {
	Value int `json:"value"`
}

func TestModelExampleJSON(t *testing.T) {
	m := newDecoder(tree{}).decode(&ExampleRequest{})
	t.Log("\n" + m.exampleJSON(""))
}
