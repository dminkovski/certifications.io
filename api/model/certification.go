package model

import (
	"encoding/json"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Certification struct {
	_id		  primitive.ObjectID `bson:"_id"`
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
	Description string `json: "description`
}

func (cert Certification) MarshalJSON() ([]byte, error) {
	fmt.Println(cert._id)
	return json.Marshal(map[string]interface{}{
		"_id":		 cert._id,
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
		"description": cert.Description,
	})
}

func (cert *Certification) Create() {
	cert.Created = time.Now().String()
}

func (cert *Certification) AddCourse(course Course) bool {
	if course.Name != "" {
		cert.Courses = append(cert.Courses, course)
		cert.Updated = time.Now().String()
		return true
	}
	return false
}

func (cert *Certification) SetImage(path string) {
	cert.Image = path
}
