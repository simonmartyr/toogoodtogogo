package toogoodtogo

type SalesTax struct {
	TaxDescription string  `json:"tax_description"`
	TaxPercentage  float64 `json:"tax_percentage"`
}

type Badge struct {
	BadgeType   string `json:"badge_type"`
	RatingGroup string `json:"rating_group"`
	Percentage  int    `json:"percentage"`
	UserCount   int    `json:"user_count"`
	MonthCount  int    `json:"month_count"`
}

type AverageRating struct {
	AverageOverallRating float64 `json:"average_overall_rating"`
	RatingCount          int     `json:"rating_count"`
	MonthCount           int     `json:"month_count"`
}

type Item struct {
	Item           ItemDetails `json:"item"`
	Store          Store       `json:"store"`
	DisplayName    string      `json:"display_name"`
	PickupInterval Interval    `json:"pickup_interval"`
	PickupLocation Location    `json:"pickup_location"`
	PurchaseEnd    string      `json:"purchase_end"`
	ItemsAvailable int         `json:"items_available"`
	Distance       float64     `json:"distance"`
	Favorite       bool        `json:"favorite"`
	InSalesWindow  bool        `json:"in_sales_window"`
	NewItem        bool        `json:"new_item"`
	ItemType       string      `json:"item_type"`
	MatchesFilters bool        `json:"matches_filters"`
}

type ItemsQueryFilter struct {
	Latitude       float64
	Longitude      float64
	Radius         int
	PageSize       int
	Page           int
	Discover       bool
	FavoritesOnly  bool
	ItemCategories []string
	PickupEarliest string
	PickupLatest   string
	SearchPhrase   string
	WithStockOnly  bool
	HiddenOnly     bool
	WeCareOnly     bool
}

type ItemsQueryRequest struct {
	//UserId         string   `json:"user_id"`
	Origin         Origin   `json:"origin,omitempty"`
	Radius         int      `json:"radius,omitempty"`
	PageSize       int      `json:"page_size,omitempty"`
	Page           int      `json:"page,omitempty"`
	Discover       bool     `json:"discover,omitempty"`
	FavoritesOnly  bool     `json:"favorites_only,omitempty"`
	ItemCategories []string `json:"item_categories,omitempty"`
	PickupEarliest string   `json:"pickup_earliest,omitempty"`
	PickupLatest   string   `json:"pickup_latest,omitempty"`
	SearchPhrase   string   `json:"search_phrase,omitempty"`
	WithStockOnly  bool     `json:"with_stock_only,omitempty"`
	HiddenOnly     bool     `json:"hidden_only,omitempty"`
	WeCareOnly     bool     `json:"we_care_only,omitempty"`
}

type ItemsQueryResult struct {
	Items []Item `json:"items"`
}

type ItemDetails struct {
	ItemId                   string        `json:"item_id"`
	SalesTaxes               []SalesTax    `json:"sales_taxes"`
	TaxAmount                MonetaryValue `json:"tax_amount"`
	PriceExcludingTaxes      MonetaryValue `json:"price_excluding_taxes"`
	PriceIncludingTaxes      MonetaryValue `json:"price_including_taxes"`
	ValueExcludingTaxes      MonetaryValue `json:"value_excluding_taxes"`
	ValueIncludingTaxes      MonetaryValue `json:"value_including_taxes"`
	TaxationPolicy           string        `json:"taxation_policy"`
	ShowSalesTaxes           bool          `json:"show_sales_taxes"`
	ItemPrice                MonetaryValue `json:"item_price"`
	ItemValue                MonetaryValue `json:"item_value"`
	CoverPicture             ImageValue    `json:"cover_picture"`
	LogoPicture              ImageValue    `json:"logo_picture"`
	Name                     string        `json:"name"`
	Description              string        `json:"description"`
	FoodHandlingInstructions string        `json:"food_handling_instructions"`
	CanUserSupplyPackaging   bool          `json:"can_user_supply_packaging"`
	CollectionInfo           string        `json:"collection_info"`
	//	Diet
	ItemCategory          string        `json:"item_category"`
	Buffet                bool          `json:"buffet"`
	Badges                []Badge       `json:"badges"`
	PositiveRatingReasons []string      `json:"positive_rating_reasons"`
	AverageOverallRating  AverageRating `json:"average_overall_rating"`
	FavoriteCount         int           `json:"favorite_count"`
}

type ItemQuery struct {
	UserId string `json:"user_id"`
	Origin Origin `json:"origin"`
}

func (c *Client) GetItem(itemId string) (*Item, error) {
	getItemUrl := c.makeUrl(ItemEndpoint) + itemId
	body := ItemQuery{
		UserId: c.userId,
	}
	result := Item{}
	err := c.postWithResult(getItemUrl, body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) GetItems(query *ItemsQueryFilter) (*ItemsQueryResult, error) {
	getItemsUrl := c.makeUrl(ItemEndpoint)
	body := ItemsQueryRequest{
		//UserId:         c.userId,
		Origin:         Origin{query.Latitude, query.Longitude},
		Radius:         query.Radius,
		PageSize:       query.PageSize,
		Page:           query.Page,
		Discover:       query.Discover,
		FavoritesOnly:  query.FavoritesOnly,
		ItemCategories: query.ItemCategories,
		PickupEarliest: query.PickupEarliest,
		PickupLatest:   query.PickupLatest,
		SearchPhrase:   query.SearchPhrase,
		WithStockOnly:  query.WithStockOnly,
		HiddenOnly:     query.HiddenOnly,
		WeCareOnly:     query.WeCareOnly,
	}
	result := ItemsQueryResult{}
	err := c.postWithResult(getItemsUrl, body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
