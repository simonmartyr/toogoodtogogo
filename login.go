package toogoodtogo

import (
	"compress/gzip"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type LoginInput struct {
	DeviceType string `json:"device_type"`
	Email      string `json:"email"`
}

type LoginResponse struct {
	State     string `json:"state"`
	PollingId string `json:"polling_id"`
}

type PollingRequest struct {
	DeviceType       string `json:"device_type"`
	Email            string `json:"email"`
	RequestPollingId string `json:"request_polling_id"`
}

type PollResponse struct {
	AccessToken  string      `json:"access_token"`
	RefreshToken string      `json:"refresh_token"`
	StartupData  StartupData `json:"startup_data"`
}
type StartupData struct {
	User StartupDataUser `json:"user"`
}
type StartupDataUser struct {
	UserId    string `json:"user_id"`
	Name      string `json:"name"`
	CountryId string `json:"country_id"`
	Email     string `json:"email"`
}

// Authenticate upon initial success ToGoodToGo will send the configured email a login link
// this method will poll until user has clicked the link (within a two-minute window).
// Once the polling is successful client will be enriched with the client's access tokens.
func (c *Client) Authenticate() error {
	authenticationUrl := c.makeUrl(AuthByEmail)
	body := LoginInput{
		DeviceType: DeviceType,
		Email:      c.username,
	}
	var response = LoginResponse{}
	err := c.postWithResult(authenticationUrl, body, &response)
	if err != nil {
		return err
	}
	pollErr := c.pollAuth(response.PollingId)
	if pollErr != nil {
		return pollErr
	}
	//do logic.
	return nil
}

// pollAuth begin looping for two minutes until the user accepts the auth link from the mail
func (c *Client) pollAuth(pollingId string) error {
	pollingUrl := c.makeUrl(AuthPolling)
	body := PollingRequest{
		DeviceType:       DeviceType,
		Email:            c.username,
		RequestPollingId: pollingId,
	}
	log.Println("tgtg: authentication polling started - Please accept link in your email")
	for pollCount := 0; pollCount < PollAttempts; pollCount++ {
		res, httpErr := c.post(pollingUrl, body)
		if httpErr != nil {
			return httpErr
		}
		if res.StatusCode == http.StatusOK {
			log.Println("tgtg: polling finished login successful")
			processErr := c.processSuccessfulPoll(res)
			if processErr != nil {
				return processErr
			}
			break
		}
		if res.StatusCode > 299 {
			return c.parseError(res)
		}
		log.Println("tgtg: Re-polling authentication - Please accept link in your email")
		time.Sleep(PollTime * time.Second)
	}
	return nil
}

// processSuccessfulPoll extract the AccessToken, RefreshToken, UserId and Cookie from the
// Successful authentication
func (c *Client) processSuccessfulPoll(res *http.Response) error {
	defer res.Body.Close()
	var pollRes = PollResponse{}
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
	c.userId = pollRes.StartupData.User.UserId
	c.cookie = res.Header.Get("Set-Cookie")
	return nil
}
