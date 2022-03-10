package docer

import (
	"encoding/json"
	"testing"
)

func TestDocumentGroup(t *testing.T) {
	var documents = Documents{
		{Name: "111", Group: "111"},
		{Name: "111-111", Group: "111/111"},
		{Name: "111-222", Group: "111/222"},
		{Name: "111-111-111", Group: "111/111/111"},
		{Name: "111-111-222", Group: "111/111/222"},
		{Name: "no group"},
	}

	g := documents.Group()
	b, _ := json.Marshal(g)
	t.Log(string(b))
}
