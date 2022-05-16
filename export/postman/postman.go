package postman

const (
	SchemaV2_1_0 = "https://schema.getpostman.com/json/collection/v2.1.0/collection.json"
)

type Info struct {
	PostmanID string `json:"_postman_id"`
	Name      string `json:"name"`
	Schema    string `json:"schema"`
}

type Item struct {
	Name        string         `json:"name"`
	Request     ItemRequest    `json:"request,omitempty"`
	Response    []ItemResponse `json:"response,omitempty"`
	Item        []*Item        `json:"item,omitempty"`
	Description string         `json:"description"`
}

type ItemRequest struct {
	Method      string       `json:"method"`
	Header      []ItemHeader `json:"header"`
	Body        ItemBody     `json:"body"`
	Url         ItemUrl      `json:"url"`
	Description string       `json:"description"`
}

type ItemResponse struct {
}

type ItemHeader struct {
	Key   string `json:"key"`
	Value string `json:"value"`
	Type  string `json:"type"`
}

type ItemBody struct {
	Mode    string `json:"mode"`
	Raw     string `json:"raw"`
	Options struct {
		Raw struct {
			Language string `json:"language"`
		} `json:"raw"`
	} `json:"options"`
}

type ItemUrl struct {
	Raw      string   `json:"raw"`
	Protocol string   `json:"protocol"`
	Host     []string `json:"host"`
	Path     []string `json:"path"`
}

type Document struct {
	Info Info    `json:"info"`
	Item []*Item `json:"item"`
}
