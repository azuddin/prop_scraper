package main

type iPropertyData struct {
	Data iPropertyDatum `json:"data"`
}

type iPropertyDatum struct {
	ACSListing iPropertyACSListing `json:"ascListings"`
}

type iPropertyACSListing struct {
	Items []iPropertyItem `json:"items"`
	TotalCount int `json:"totalCount"`
	NextPageToken int `json:"nextPageToken"`
}

type iPropertyItem struct {
	Id string `json:"id"`
	Address iPropertyAddress `json:"address"`
	Prices []iPropertyPrice `json:"prices"`
}

type iPropertyAddress struct {
	FormattedAddress string `json:"formattedAddress"`
	Latitude float64 `json:"lat"`
	Longitude float64 `json:"lng"`
}

type iPropertyPrice struct {
	Type string `json:"type"`
	Currency string `json:"currency"`
	Min int `json:"min"`
	Max int `json:"max"`
	MonthlyPayment int `json:"monthlyPayment"`
}