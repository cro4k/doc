package goclient

import "net/http"

var c = new(http.Client)

func SetClient(client *http.Client) {
	c = client
}
