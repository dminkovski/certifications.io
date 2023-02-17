package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"sync"

	"github.com/dminkovski/certifications.io/api/model"
	"github.com/dminkovski/certifications.io/api/parser"
	"github.com/dminkovski/certifications.io/api/utils"
	"golang.org/x/net/html"
)


func init(){
	http.HandleFunc("/api/credly", GetFromCredly)
}

/*func main(){
	res, err := http.Get("https://www.credly.com/org/rumos/badge/azure-architecture")
	if err != nil {
		log.Panic(err)
	}
	defer res.Body.Close()
	bodyBytes, err := io.ReadAll(res.Body)
	cert := Parse(string(bodyBytes))
	fmt.Println(cert)
}*/

func GetFromCredly(w http.ResponseWriter, req *http.Request) {
	utils.PrepareResponse(w, nil)
	fmt.Println("Get From Credly", req.Method)
	if req.Method == "POST" {
		var c struct{URL string `json:"url"`}
		err := DecodeJsonBody(w, req, &c)
		if err != nil {
			log.Panic(err)
		}
		fmt.Println(c)
		fmt.Println("Starting Crawler with",c.URL)
		res, err := http.Get(c.URL)
		if err != nil {
			log.Panic(err)
		}
		defer res.Body.Close()
		bodyBytes, err := io.ReadAll(res.Body)
		cert := Parse(string(bodyBytes))	
		out, err := json.Marshal(cert)
		utils.PrepareResponse(w, out)
	}
}

func Parse(text string) model.Certification {
	tkn := html.NewTokenizer(strings.NewReader(text))
	wg := new(sync.WaitGroup)
	tokens := []parser.TokenCheck{
		parser.TokenCheck{
			"cr-badges-badge-skills__skills",
			"ul",
			"skills",
			make([]string, 0),
			false,
		},
		parser.TokenCheck{
			"ac-heading--badge-name-hero",
			"h1",
			"name",
			"",
			false,
		},
		parser.TokenCheck{
			"cr-badges-full-badge__description",
			"p",
			"description",
			"",
			false,
		},
	}
	wg.Add(len(tokens))

	go parser.ParseTokenizer(tkn, wg, tokens)

	wg.Wait()

	fmt.Println(tokens)
	var cert model.Certification
	for _, tkn := range tokens {
		switch tkn.Key {
			case "name":
				val, ok := tkn.Data.(string)
				if ok {
					cert.Name = val
				}
			case "description":
				val, ok := tkn.Data.(string)
				if ok {
					cert.Description = val
				}
			case "skills":
				val, ok := tkn.Data.([]string)
				if ok {
					cert.Skills = val
				}
		}
	}

	return cert
}
