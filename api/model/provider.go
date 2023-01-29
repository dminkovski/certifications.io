package model

import "encoding/json"

type Provider struct {
	Id   int    `json: "id"`
	Name string `json: "name"`
	Link string `json: "link"`
}

func (provider Provider) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"id":   provider.Id,
		"name": provider.Name,
		"link": provider.Link,
	})
}
