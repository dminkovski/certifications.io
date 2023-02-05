package database

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/dminkovski/certifications.io/api/model"
)

type ImgLoader struct {
	index int
	url   string
}


// Load Certifications into Database from JSON
func LoadCertifications() int {
	certsPath := "./database/data/certifications.json"
	certifications := make([]model.Certification, 0)
	file, err := os.ReadFile(certsPath)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal([]byte(file), &certifications)
	if err != nil {
		log.Fatal(err)
	}

	docs := make([]interface{}, 0)
	ch := make(chan ImgLoader)
	var wg sync.WaitGroup
	for index, c := range certifications {
		wg.Add(1)
		go func() {
			imagePath := GetImage(c.ShortName, c.Image)
			ob := ImgLoader{
				index: index,
				url:   imagePath,
			}
			ch <- ob
			wg.Done()
		}()

	}
	go func() {
		wg.Wait()
		defer close(ch)
	}()
	for result := range ch {
		certifications[result.index].SetImage(result.url)
		docs = append(docs, interface{}(
			certifications[result.index],
		))
	}
	return LoadDatabaseWithJSON(docs)
}

func SaveCertification(cert model.Certification) error {
	certsPath := "./database/data/created.json"
	certifications := make([]model.Certification, 1)
	file, err := os.ReadFile(certsPath)
	if err != nil {
		log.Fatal(err)
		return err
	}
	err = json.Unmarshal([]byte(file), &certifications)
	if err != nil {
		log.Fatal(err)
		return err
	}
	certifications = append(certifications, cert)
	results, err := json.Marshal(certifications)
	err = os.WriteFile(certsPath, results, 0666)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

// Download the referenced Image from the json of internet
func GetImage(shortName string, url string) string {
	outputPath := fmt.Sprintf("./assets/%s.png", shortName)
	imagePath := fmt.Sprintf("/assets/img/%s.png", shortName)
	_, err := os.Open(outputPath)
	if errors.Is(err, os.ErrNotExist) {
		resp, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
			return ""
		}
		defer resp.Body.Close()
		file, err := os.Create(outputPath)
		if err != nil {
			log.Fatal(err)
			return ""
		}
		_, err = io.Copy(file, resp.Body)
		if err != nil {
			log.Fatal(err)
			return ""
		}
	} else {
		fmt.Println("Image is already in assets: ",shortName)
	}
	return imagePath
}

func LoadDatabaseWithJSON(data []interface{}) int {
	fmt.Println("Connecting to Database and saving Certifications")
	client := Connect()
	col := client.Database("db").Collection("Certifications")
	result, err := col.InsertMany(context.TODO(), data)
	if err != nil {
		log.Panic(err)
	}
	for _, id := range result.InsertedIDs {
		fmt.Printf("Inserted document with _id: %v\n", id)
	}
	err = Disconnect(client)
	if err != nil {
		log.Panic(err)
	}
	return len(result.InsertedIDs)
}
