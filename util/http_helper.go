package util

import (
	"encoding/json"
	"net/http"
)

// Returns a map of the json response
func GetApiData(url string) map[string]interface{} {
	resp, err := http.Get(url)
	if err != nil {
		PrintError(err)
		return nil
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		PrintError(err)
		return nil
	}

	return result
}
