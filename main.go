package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("\nPlease enter a valid ISO country code...")
	scanner.Scan()
	countryCode := scanner.Text()
	url := "http://api.worldbank.org/v2/country/" + countryCode + "?format=json"

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Error getting country info from World Bank API: %s", err.Error())
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading body of JSON response: %s", err.Error())
	}

	var tmp APIResponse
	err = json.Unmarshal([]byte(body), &tmp)
	if err != nil {
		log.Fatalf("Error unmarshalling JSON: %s", err.Error())
	}
	if len(tmp.PageInfo.Message) > 0 {
		log.Fatalf("Invalid country code. Please check your country code and try again.")
	}

	data := tmp.CountryInfo[0]
	fmt.Printf("\nName: %s\n", data.Name)
	fmt.Printf("Region: %s\n", data.Region.Value)
	fmt.Printf("Capital City: %s\n", data.CapitalCity)
	fmt.Printf("Longitude: %s\n", data.Longitude)
	fmt.Printf("Latitude: %s\n", data.Latitude)
}
