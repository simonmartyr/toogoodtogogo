package toogoodtogo

type Store struct {
	StoreId        string     `json:"store_id"`
	StoreName      string     `json:"store_name"`
	Branch         string     `json:"branch"`
	Description    string     `json:"description"`
	Website        string     `json:"website"`
	StoreLocation  Location   `json:"store_location"`
	LogoPicture    ImageValue `json:"logo_picture"`
	StoreTimeZone  string     `json:"store_time_zone"`
	Hidden         bool       `json:"hidden"`
	FavoriteCount  int        `json:"favorite_count"`
	WeCare         bool       `json:"we_care"`
	Distance       float64    `json:"distance"`
	CoverPicture   ImageValue `json:"cover_picture"`
	IsManufacturer bool       `json:"is_manufacturer"`
}
