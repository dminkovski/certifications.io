package model

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func LoadCertifications() []Certification {
	certsPath := "./model/data/certifications.json"
	certifications := make([]Certification, 1)
	file, err := os.ReadFile(certsPath)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal([]byte(file), &certifications)
	if err != nil {
		log.Fatal(err)
	}
	for index, c := range certifications {
		certifications[index].SetImage(GetImage(c.ShortName, c.Image))
	}
	return certifications
}

func GetImage(shortName string, url string) string {
	outputPath := fmt.Sprintf("./assets/%s.png", shortName)
	imagePath := fmt.Sprintf("/assets/img/%s.png", shortName)
	_, err := os.Open(outputPath)
	if errors.Is(err, os.ErrNotExist) {
		resp, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		file, err := os.Create(outputPath)
		if err != nil {
			log.Fatal(err)
		}
		_, err = io.Copy(file, resp.Body)
		if err != nil {
			log.Fatal(err)
		}
	}
	return imagePath
}
