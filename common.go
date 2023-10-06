package toogoodtogo

type MonetaryValue struct {
	Code       string `json:"code"`
	MinorUnits int    `json:"minor_units"`
	Decimals   int    `json:"decimals"`
}

type ImageValue struct {
	PictureId              string `json:"picture_id"`
	CurrentUrl             string `json:"current_url"`
	IsAutomaticallyCreated bool   `json:"is_automatically_created"`
}

type Interval struct {
	Start string `json:"start"`
	End   string `json:"end"`
}
