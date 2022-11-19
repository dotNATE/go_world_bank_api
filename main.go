package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func GetCountryCodeFromUser(s *bufio.Scanner) string {
	fmt.Println("\nPlease enter a valid ISO country code...")
	s.Scan()
	return s.Text()
}

func GenerateAPIUrl(countryCode string) string {
	return "http://api.worldbank.org/v2/country/" + countryCode + "?format=json"
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	countryCode := GetCountryCodeFromUser(scanner)
	url := GenerateAPIUrl(countryCode)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Error getting country info from World Bank API: %s", err.Error())
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading body of JSON response: %s", err.Error())
	}

	data, err := HandleUnmarshall(body)
	if err != nil {
		log.Fatal(err)
	}

	data.CountryInfo[0].PrintData()
}
