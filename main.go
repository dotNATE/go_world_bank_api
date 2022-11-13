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

type CountryInfoMeta struct {
	Id       string
	Iso2Code string
	Value    string
}

type CountryInfo struct {
	Id          string
	Iso2Code    string
	Name        string
	Region      CountryInfoMeta
	AdminRegion CountryInfoMeta
	IncomeLevel CountryInfoMeta
	LendingType CountryInfoMeta
	CapitalCity string
	Longitude   string
	Latitude    string
}

type APIResponse struct {
	PageInfo    map[string]int
	CountryInfo []CountryInfo
}

func (resp *APIResponse) UnmarshalJSON(buf []byte) error {
	tmp := []interface{}{&resp.PageInfo, &resp.CountryInfo}

	err := json.Unmarshal(buf, &tmp)
	if err != nil {
		return nil
	}

	return nil
}

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
	data := tmp.CountryInfo[0]

	fmt.Printf("\nName: %s\n", data.Name)
	fmt.Printf("Region: %s\n", data.Region.Value)
	fmt.Printf("Capital City: %s\n", data.CapitalCity)
	fmt.Printf("Longitude: %s\n", data.Longitude)
	fmt.Printf("Latitude: %s\n", data.Latitude)
}
