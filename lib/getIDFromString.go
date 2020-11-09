package lib

import (
	"regexp"
)

// GetIdFromString takes a swapiURL and returns a substring after 'api'
func GetIDFromString(swapiURL string) string {
	re := regexp.MustCompile("[0-9]+")
	arr := re.FindAllString(swapiURL, -1)
	if len(arr) > 0 {
		return arr[0]
	}
	return "1"
}
