package lib

import (
	"regexp"
)

// GetIdFromString takes a swapiURL and returns a substring after 'api'
func GetIdFromString(swapiURL string) string {
	re := regexp.MustCompile("[0-9]+")
	arr := re.FindAllString(swapiURL, -1)
	return arr[0]
}
