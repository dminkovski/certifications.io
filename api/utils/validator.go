package utils

import "regexp"

func IsValidURL(url string) bool {
	re := regexp.MustCompile(`https:\/\/\S*\.[a-z]{2,3}`)
	return re.Match([]byte(url))
}
