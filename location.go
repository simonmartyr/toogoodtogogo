package toogoodtogo

type Location struct {
	Address  Address `json:"address"`
	Location Origin  `json:"location"`
}

type Address struct {
	Country     Country `json:"country"`
	AddressLine string  `json:"address_line"`
	City        string  `json:"city"`
	PostalCode  string  `json:"postal_code"`
}

type Country struct {
	IsoCode string `json:"iso_code"`
	Name    string `json:"name"`
}
