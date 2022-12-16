package model

import (
	"time"
)

type Certification struct {
	Id int 				`json: "id"`
	Link string 		`json: "link"`
	Image string 		`json: "image"`
	Courses []string		`json: "courses"`
	Name string			`json: "name"`
	ShortName string	`json: "shortName"`
	Updated string	`json: "updated"`
	Created string	`json: "created"`
	Skills []string	`json: "skills"`
	Notes string	`json: "notes"`
	Price int 		`json: "price`
	MinMonths int 	`json: "minMonths"`
	MaxMonths int 	`json: "maxMonths"`
}


func (cert *Certification) Create(){
	cert.Created = time.Now().String()
}

func (cert *Certification) AddCourse(course string) bool{
	if course != "" {
		cert.Courses = append(cert.Courses, course)
		cert.Updated = time.Now().String()
		return true
	}
	return false
}

func (cert *Certification) SetImage(path string) {
	cert.Image = path
}
