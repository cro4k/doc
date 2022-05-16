package postman

import (
	"github.com/cro4k/doc/docer"
	"os"
	"testing"
)

func TestPostman(t *testing.T) {
	doc := &docer.DocumentGroup{
		Name:      "Demo",
		Documents: nil,
		Children: map[string]*docer.DocumentGroup{
			"Group1": {
				Name: "Group1",
				Documents: []*docer.Document{
					{
						API:    "/api/demo/group1/login",
						Name:   "demo",
						Method: "POST",
						Extra:  map[string]interface{}{},
						Req: &docer.Model{
							Name: "Login",
							Fields: []*docer.Field{
								{
									Name:     "username",
									Type:     "string",
									Required: true,
									Comment:  "用户名",
									Option:   "",
								},
							},
							Array: false,
						},
						Rsp:   &docer.Model{},
						Group: "Group1",
						Sort:  0,
					},
				},
				Children: map[string]*docer.DocumentGroup{
					"Group2": {
						Name: "Group2",
						Documents: []*docer.Document{
							{
								API:    "/api/group2/hello",
								Name:   "hello",
								Method: "GET",
								Extra:  map[string]interface{}{},
								Req: &docer.Model{
									Name: "Login",
									Fields: []*docer.Field{
										{
											Name:     "id",
											Type:     "string",
											Required: true,
											Comment:  "id",
											Option:   "",
										},
									},
									Array: false,
								},
								Rsp:   &docer.Model{},
								Group: "",
								Sort:  0,
							},
						},
						Children: nil,
						Root:     false,
					},
				},
				Root: false,
			},
		},
		Root: true,
	}
	f, _ := os.OpenFile("a.json", os.O_CREATE|os.O_RDWR, 0644)
	ExportPostman(f, doc)
	f.Close()
}
