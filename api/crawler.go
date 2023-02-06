package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"

	"golang.org/x/net/html"
)

type ParsedData struct {
	key  string
	data interface{}
}

func main() {
	fmt.Println("Starting Crawler")
	res, err := http.Get("https://www.credly.com/org/microsoft-certification/badge/microsoft-certified-azure-fundamentals")
	if err != nil {
		log.Panic(err)
	}
	defer res.Body.Close()
	file, err := os.Create("credly.html")
	if err != nil {
		log.Panic(err)
	}
	_, err = io.Copy(file, res.Body)
	if err != nil {
		log.Panic(err)
	}
	
	body, err := ioutil.ReadFile("credly.html")
	if err != nil {
		log.Panic(err)
	}
	Parse(string(body))

}

func ParseSkills(tkn *html.Tokenizer, ch chan ParsedData, wg *sync.WaitGroup) {
	fmt.Println("Parse Skills")
	skills := make([]string, 0)
	var isSkillList bool
	for {
		tt := tkn.Next()
		switch {
		case tt == html.ErrorToken:
			return
		case tt == html.StartTagToken:
			t := tkn.Token()
			if t.Data == "ul" {
				for _, at := range t.Attr {
					if at.Key == "class" && strings.Contains(string(at.Val), "cr-badges-badge-skills__skills") {
						isSkillList = true
					}
				}
			}
		case tt == html.TextToken:
			{
				t := tkn.Token()
				if isSkillList {
					skills = append(skills, strings.TrimSpace(t.Data))
				}
			}
		case tt == html.EndTagToken:
			t := tkn.Token()
			if isSkillList && t.Data == "ul" {
				isSkillList = false
				ch <- ParsedData{
					key:  "skills",
					data: skills,
				}
				defer wg.Done()
			}
		}
	}
}

func ParseName(tkn *html.Tokenizer, ch chan ParsedData, wg *sync.WaitGroup) string {
	fmt.Println("Parse Name")
	var name string
	var isName bool
	for {
		tt := tkn.Next()
		switch {
			case tt == html.ErrorToken:
				return ""
			case tt == html.StartTagToken:
				t := tkn.Token()
				if t.Data == "h1" {
					for _, at := range t.Attr {
						if at.Key == "class" && strings.Contains(string(at.Val), "ac-heading--badge-name-hero") {
							isName = true
						}
					}
				}
			case tt == html.TextToken:
				{
					t := tkn.Token()
					if isName {
						name = strings.TrimSpace(t.Data)
					}
				}
			case tt == html.EndTagToken:
				t := tkn.Token()
				if isName && t.Data == "h1" {
					isName = false
					ch <- ParsedData{
						key:  "name",
						data: name,
					}
					defer wg.Done()
				}
			}
	}
}

func Parse(text string) {
	tkn1 := html.NewTokenizer(strings.NewReader(text))
	tkn2 := html.NewTokenizer(strings.NewReader(text))
	const parserCount = 2
	ch := make(chan ParsedData, parserCount)
	wg := new(sync.WaitGroup)

	wg.Add(parserCount)
	go ParseSkills(tkn1, ch, wg)
	go ParseName(tkn2, ch, wg)
	wg.Wait()
	for i := 0; i < parserCount;i++{
		fmt.Println(<-ch)
	}
	
}
