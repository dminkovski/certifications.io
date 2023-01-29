package model

import (
	"encoding/json"
	"time"
)


type Certification struct {
	Id        int      `json: "id"`
	Link      string   `json: "link"`
	Image     string   `json: "image"`
	Courses   []Course `json: "courses"`
	Name      string   `json: "name"`
	ShortName string   `json: "shortName"`
	Updated   string   `json: "updated"`
	Created   string   `json: "created"`
	Skills    []string `json: "skills"`
	Notes     string   `json: "notes"`
	Price     int      `json: "price`
	MinMonths int      `json: "minMonths"`
	MaxMonths int      `json: "maxMonths"`
}

func (cert Certification) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"id":        cert.Id,
		"link":      cert.Link,
		"image":     cert.Image,
		"courses":   cert.Courses,
		"name":      cert.Name,
		"shortName": cert.ShortName,
		"skills":    cert.Skills,
		"minMonths": cert.MinMonths,
		"maxMonths": cert.MaxMonths,
		"price":     cert.Price,
	})
}

func (cert *Certification) Create() {
	cert.Created = time.Now().String()
}

func (cert *Certification) AddCourse(course Course) bool {
	if course.Id != 0 && course.Name != "" {
		cert.Courses = append(cert.Courses, course)
		cert.Updated = time.Now().String()
		return true
	}
	return false
}

func (cert *Certification) SetImage(path string) {
	cert.Image = path
}
