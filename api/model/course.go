package model

import "encoding/json"

type Course struct {
	Id       int      `json: "id"`
	Link     string   `json: "link"`
	Provider Provider `json: "provider`
	Name     string   `json: "name"`
}

func (course Course) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"id":       course.Id,
		"link":     course.Link,
		"provider": course.Provider,
		"name":     course.Name,
	})
}
