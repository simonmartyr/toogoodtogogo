package toogoodtogo

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// Client the too good to go api client it is recommended to use the New method to create a new instance
type Client struct {
	http         *http.Client
	baseURL      string
	accessToken  string
	refreshToken string
	username     string
	userId       string
	language     string
	cookie       string
	userAgent    string
	verbose      bool
}

type Credentials struct {
	AccessToken  string
	RefreshToken string
	Cookie       string
	UserId       string
}

const (
	BaseUrl         = "https://apptoogoodtogo.com/api/"
	DeviceType      = "ANDROID"
	DefaultLanguage = "en-GB"
	UserAgent       = "TGTG/22.5.5 Dalvik/2.1.0 (Linux; U; Android 10; SM-G935F Build/NRD90M)"
	AuthByEmail     = "auth/v3/authByEmail"
	AuthPolling     = "auth/v3/authByRequestPollingId"
	BucketEndpoint  = "discover/v1/bucket"
	RefreshEndpoint = "auth/v3/token/refresh"
	ItemEndpoint    = "item/v8/"
	PollTime        = 5
	PollAttempts    = 24
)

type ClientOption func(client *Client)

// WithBaseUrl replace the base URL to be used
func WithBaseUrl(url string) ClientOption {
	return func(client *Client) {
		client.baseURL = url
	}
}

// WithLanguage specify the particular language you wish to use.
// The default is set to en-GB
func WithLanguage(languageCode string) ClientOption {
	return func(client *Client) {
		client.language = languageCode
	}
}

func WithVerbose() ClientOption {
	return func(client *Client) {
		client.verbose = true
	}
}

// WithUsername specify your username (the email of the account) to be used for authentication.
func WithUsername(username string) ClientOption {
	return func(client *Client) {
		client.username = username
	}
}

func WithAuth(userId string, accessToken string, refreshToken string, cookie string) ClientOption {
	return func(client *Client) {
		client.userId = userId
		client.accessToken = accessToken
		client.refreshToken = refreshToken
		client.cookie = cookie
	}
}

// New Create a new instance of the too good to go api client
func New(client *http.Client, opts ...ClientOption) *Client {
	c := &Client{
		http:      client,
		baseURL:   BaseUrl,
		language:  DefaultLanguage,
		userAgent: UserAgent,
	}

	for _, opt := range opts {
		opt(c)
	}

	return c
}

func (c *Client) GetCredentials() *Credentials {
	return &Credentials{
		AccessToken:  c.accessToken,
		RefreshToken: c.refreshToken,
		Cookie:       c.cookie,
		UserId:       c.userId,
	}
}

// Logout assuming there is a valid session this endpoint will request the termination of the session and
// clear the token from the client.
func (c *Client) Logout() error {
	logoutUrl := c.baseURL + "/user/logout"
	err := c.postWithResult(logoutUrl, nil, nil)
	if err != nil {
		return err
	}
	c.accessToken = ""
	c.refreshToken = ""
	c.cookie = ""
	return nil
}

// IsAuthenticated convince method to verify the client has an auth token.
func (c *Client) IsAuthenticated() bool {
	if c == nil {
		return false
	}
	return c.accessToken != ""
}

func (c *Client) parseError(resp *http.Response) error {
	gzipBody, gzipErr := gzip.NewReader(resp.Body)
	if gzipErr != nil {
		return gzipErr
	}
	body, err := io.ReadAll(gzipBody)
	if err != nil {
		return err
	}

	if len(body) == 0 {
		return fmt.Errorf("tgtg-api: produced an error with code %d: %s empty error", resp.StatusCode, http.StatusText(resp.StatusCode))
	}

	return fmt.Errorf("tgtg-api: produced an [%d] response with an unprocessable error payload [%s]", resp.StatusCode, body)
}

func (c *Client) get(url string, result interface{}) error {
	request, requestErr := http.NewRequest("GET", url, nil)
	if requestErr != nil {
		return requestErr
	}
	c.configureHeaders(request)
	res, httpErr := c.http.Do(request)
	if httpErr != nil {
		return httpErr
	}
	return c.processResponse(res, result)
}

func (c *Client) post(url string, body any) (*http.Response, error) {
	jsonBody, jsonReqErr := json.Marshal(body)
	if jsonReqErr != nil {
		return nil, jsonReqErr
	}
	request, requestErr := http.NewRequest("POST", url, bytes.NewReader(jsonBody))
	if requestErr != nil {
		return nil, requestErr
	}
	c.configureHeaders(request)
	res, httpErr := c.http.Do(request)
	if httpErr != nil {
		return nil, httpErr
	}
	return res, nil
}

func (c *Client) postWithResult(url string, body any, result interface{}) error {
	res, postErr := c.post(url, body)
	if postErr != nil {
		return postErr
	}
	return c.processResponse(res, result)
}

// processResponse if result specified decompress the gzip payload and decode the json.
func (c *Client) processResponse(res *http.Response, result interface{}) error {
	if !(res.StatusCode >= 200 && res.StatusCode <= 299) {
		return c.parseError(res)
	}
	if result == nil {
		return nil
	}
	defer res.Body.Close()
	gzipReader, gzipErr := gzip.NewReader(res.Body)
	if gzipErr != nil {
		return gzipErr
	}
	if c.verbose {
		buf := &bytes.Buffer{}
		tee := io.TeeReader(gzipReader, buf)
		readme, _ := io.ReadAll(tee)
		log.Print(string(readme))
		jsonErr := json.NewDecoder(buf).Decode(result)
		if jsonErr != nil {
			return jsonErr
		}
		return nil
	}
	jsonErr := json.NewDecoder(gzipReader).Decode(result)
	if jsonErr != nil {
		return jsonErr
	}
	return nil
}

func (c *Client) makeUrl(path string) string {
	return c.baseURL + path
}

func (c *Client) configureHeaders(request *http.Request) {
	request.Header.Set("accept", "application/json")
	request.Header.Set("Accept-Encoding", "gzip")
	request.Header.Set("accept-language", c.language)
	request.Header.Set("content-type", "application/json; charset=utf-8")
	request.Header.Set("user-agent", c.userAgent)
	if c.cookie != "" {
		request.Header.Set("Cookie", c.cookie)
	}
	if c.accessToken != "" {
		request.Header.Set("authorization", fmt.Sprintf("Bearer %s", c.accessToken))
	}

}
