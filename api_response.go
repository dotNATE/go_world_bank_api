package main

import "encoding/json"

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
