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
						Path:   "/api/demo/group1/login",
						Name:   "demo",
						Method: "POST",
						Extra:  docer.KV{},
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
							Array: 0,
						},
						Rsp:   &docer.Model{},
						Group: "Group1",
					},
				},
				Children: map[string]*docer.DocumentGroup{
					"Group2": {
						Name: "Group2",
						Documents: []*docer.Document{
							{
								Path:   "/api/group2/hello",
								Name:   "hello",
								Method: "GET",
								Extra:  docer.KV{},
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
									Array: 0,
								},
								Rsp:   &docer.Model{},
								Group: "",
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
	Postman(f, doc)
	f.Close()
}
