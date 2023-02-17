package parser

import (
	"fmt"
	"strings"
	"sync"

	"golang.org/x/net/html"
)

type ParsedData struct {
	Key  string
	Data interface{}
}
type TokenCheck struct {
	Class string
	Tag string
	Key string
	Data interface{}
	IsActive bool
}
func IsToken(t html.Token, tag string, class string) bool {
	if t.Data == tag || tag == "" {
		for _, at := range t.Attr {
			if at.Key == "class" && strings.Contains(string(at.Val), class) {
				return true
			}
		}
	}
	return false
}
func ParseTokenizer(tkn *html.Tokenizer, wg *sync.WaitGroup, tokenChecks []TokenCheck){
	fmt.Println("Parsing with",len(tokenChecks))
	for {
		tt := tkn.Next()
		switch {
		case tt == html.ErrorToken:
			return
		case tt == html.StartTagToken:
			t := tkn.Token()
			for index, tc := range tokenChecks{
				if IsToken(t, tc.Tag, tc.Class) {
					(&tokenChecks[index]).IsActive = true
				}
			} 
			
		case tt == html.TextToken:
			{
				t := tkn.Token()
				for index, tc := range tokenChecks{
					if tc.IsActive {
						if len(t.Data) > 1 {
							switch  tc.Data.(type) {
								case []string:
									value, ok := tc.Data.([]string)
									if ok {
										(&tokenChecks[index]).Data = append(value, strings.TrimSpace(t.Data))
									}
								case []interface{}:
									value, ok := tc.Data.([]interface{})
									if ok {
										(&tokenChecks[index]).Data = append(value, strings.TrimSpace(t.Data))
									}
								case string:
									(&tokenChecks[index]).Data = strings.TrimSpace(t.Data)
							}
						}
					}
				}
			}
		case tt == html.EndTagToken:
			t := tkn.Token()
			for index, tc := range tokenChecks{
				if tc.IsActive && t.Data == tc.Tag {
					(&tokenChecks[index]).IsActive = false
					wg.Done()
				} 
			}
		}
	}
}