package toogoodtogo

import (
	"fmt"
	"net/http"
)

type FavoritesQuery struct {
	Origin Origin       `json:"origin"`
	Radius int          `json:"radius"`
	UserId string       `json:"user_id"`
	Paging Paging       `json:"paging"`
	Bucket BucketFilter `json:"bucket"`
}

type FavoriteRequest struct {
	IsFavorite bool `json:"is_favorite"`
}

const FavoritesFillerType = "Favorites"

// GetFavorites retrieve a list of the stores which are favorite.
func (c *Client) GetFavorites(latitude float64, longitude float64, radius int, page int, size int) (*MobileBucket, error) {
	favoriteUrl := c.makeUrl(BucketEndpoint)
	body := FavoritesQuery{
		Origin: Origin{
			Latitude:  latitude,
			Longitude: longitude,
		},
		Radius: radius,
		UserId: c.userId,
		Paging: Paging{
			Page: page,
			Size: size,
		},
		Bucket: BucketFilter{FillerType: FavoritesFillerType},
	}
	result := MobileBucket{}
	postErr := c.postWithResult(favoriteUrl, body, &result)
	if postErr != nil {
		return nil, postErr
	}
	return &result, nil
}

// SetFavorite set or remove the favorite from an item
func (c *Client) SetFavorite(itemId string, isFavorite bool) error {
	setFavoriteUrl := c.makeUrl(ItemEndpoint) + itemId + "/setFavorite"
	body := FavoriteRequest{IsFavorite: isFavorite}
	res, err := c.post(setFavoriteUrl, body)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("tgtg-api: error when favoriting")
	}
	return nil
}
