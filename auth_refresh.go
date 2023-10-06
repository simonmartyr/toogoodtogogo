package toogoodtogo

import (
	"compress/gzip"
	"encoding/json"
	"net/http"
)

type RefreshRequest struct {
	RefreshToken string `json:"refresh_token"`
}

type RefreshResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func (c *Client) RefreshAuthentication() error {
	refreshUrl := c.makeUrl(RefreshEndpoint)
	body := RefreshRequest{RefreshToken: c.refreshToken}
	res, httpErr := c.post(refreshUrl, body)
	if httpErr != nil {
		return httpErr
	}
	if res.StatusCode != http.StatusOK {
		return c.parseError(res)
	}
	return c.processSuccessfulRefresh(res)
}

// processSuccessfulRefresh extract the AccessToken, RefreshToken and Cookie from the
// Successful refresh
func (c *Client) processSuccessfulRefresh(res *http.Response) error {
	defer res.Body.Close()
	var pollRes = RefreshResponse{}
	gzipReader, gzipErr := gzip.NewReader(res.Body)
	if gzipErr != nil {
		return gzipErr
	}
	jsonErr := json.NewDecoder(gzipReader).Decode(&pollRes)
	if jsonErr != nil {
		return jsonErr
	}
	c.accessToken = pollRes.AccessToken
	c.refreshToken = pollRes.RefreshToken
	c.cookie = res.Header.Get("Set-Cookie")
	return nil
}
