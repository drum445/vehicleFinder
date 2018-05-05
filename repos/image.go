package repos

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type ImageResponse struct {
	URL string
}

func GetImage(vehicleID int) string {
	// format our URL to include the vehicle's ID
	URL := "https://izrite.com:5555/image/%v"
	url := fmt.Sprintf(URL, vehicleID)

	// make the api call for image
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return ""
	}

	defer resp.Body.Close()

	// decode it and return the URL
	var response ImageResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	return response.URL
}
