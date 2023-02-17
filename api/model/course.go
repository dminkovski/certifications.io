package model

import "encoding/json"

type Course struct {
	Link     string   `json: "link"`
	Provider Provider `json: "provider`
	Name     string   `json: "name"`
}

func (course Course) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"link":     course.Link,
		"provider": course.Provider,
		"name":     course.Name,
	})
}

type CourseDTO struct {
	Course Course `json: "course"`
	CertificationId string `json: "certification_id"`
}

func (courseDTO CourseDTO) MarshalJSON() ([]byte, error){
	return json.Marshal(map[string]interface{}{
		"course": courseDTO.Course,
		"certification_id": courseDTO.CertificationId,
	})
}