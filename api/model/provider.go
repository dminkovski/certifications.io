package model

import "encoding/json"

type Provider struct {
	Name string `json: "name"`
	Link string `json: "link"`
}

func (provider Provider) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"name": provider.Name,
		"link": provider.Link,
	})
}
