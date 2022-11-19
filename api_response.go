package main

import (
	"encoding/json"
	"fmt"
)

type APIResponse struct {
	PageInfo    PageInfo
	CountryInfo []CountryInfo
}

type PageInfo struct {
	Page    float64
	Pages   float64
	PerPage string
	Total   float64
	Message []InfoMeta
}

type CountryInfo struct {
	Id          string
	Iso2Code    string
	Name        string
	Region      InfoMeta
	AdminRegion InfoMeta
	IncomeLevel InfoMeta
	LendingType InfoMeta
	CapitalCity string
	Longitude   string
	Latitude    string
}

type InfoMeta struct {
	Id       string
	Iso2Code string
	Value    string
}

func (resp *APIResponse) UnmarshalJSON(buf []byte) error {
	tmp := []interface{}{&resp.PageInfo, &resp.CountryInfo}

	err := json.Unmarshal(buf, &tmp)
	if err != nil {
		return nil
	}

	return nil
}

func (c *CountryInfo) PrintData() {
	fmt.Printf("\nName: %s\n", c.Name)
	fmt.Printf("Region: %s\n", c.Region.Value)
	fmt.Printf("Capital City: %s\n", c.CapitalCity)
	fmt.Printf("Longitude: %s\n", c.Longitude)
	fmt.Printf("Latitude: %s\n", c.Latitude)
}

func HandleUnmarshall(b []byte) (*APIResponse, error) {
	var tmp APIResponse

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling json: %s", err.Error())
	}
	if len(tmp.PageInfo.Message) > 0 {
		return nil, fmt.Errorf("invalid country code. please check your country code and try again")
	}

	return &tmp, nil
}
