package lib

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

// GetJSONwg gets a single endpoint with a json response and defers to a wait group
func GetJSONwg(url string, target interface{}, wg *sync.WaitGroup) error {
	defer wg.Done()
	fmt.Println("Url", url)
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&target)
	return err
}

// GetJSON response without a wait group
func GetJSON(url string, target interface{}) error {
	fmt.Println("Url", url)
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&target)
	return err
}
