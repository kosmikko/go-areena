package areena

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/kosmikko/go-areena/responses"
	"github.com/kosmikko/go-areena/results"
)

// Client manages communication with the Yle API.
type Client struct {
	// HTTP client used to communicate with the API.
	httpClient *http.Client

	// Base URL for API requests, should have a trailing slash.
	BaseURL *url.URL

	// User agent
	UserAgent string

	// Config for api keys etc.
	cfg *Config
}

// NewClient Client constructor
func NewClient(cfg *Config) (client *Client, err error) {
	httpClient := cfg.HTTPClient
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	baseURL, err := url.Parse(cfg.APIBaseURL)
	if err != nil {
		return
	}
	client = &Client{
		httpClient: httpClient,
		BaseURL:    baseURL,
		cfg:        cfg,
	}
	return
}

// NewRequest creates an API request. Areena API is GET only.
func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	u := c.BaseURL.ResolveReference(rel)

	q := u.Query()
	q.Set("app_id", c.cfg.YleAppID)
	q.Set("app_key", c.cfg.YleAppKey)
	u.RawQuery = q.Encode()
	req, err := http.NewRequest(method, u.String(), nil)
	if err != nil {
		return nil, err
	}
	if c.cfg.Debug {
		log.Printf("req to %s", u.String())
	}
	if c.UserAgent != "" {
		req.Header.Add("User-Agent", c.UserAgent)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	return req, nil
}

// Do sends an API request
func (c *Client) Do(req *http.Request, v interface{}) (err error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return
	}

	if resp.StatusCode == 404 {
		err = ErrNotFound
		return
	}
	if resp.StatusCode != 200 {
		err = fmt.Errorf("Failed to read from API, got status %d", resp.StatusCode)
		return
	}
	err = json.NewDecoder(resp.Body).Decode(v)
	resp.Body.Close()
	return
}

// Program fetch details of program
func (c *Client) Program(programID string) (program *responses.Program, err error) {
	req, err := c.NewRequest("GET", fmt.Sprintf("programs/items/%s.json", programID), nil)
	if err != nil {
		return
	}
	err = c.Do(req, &program)
	return
}

// Playout fetch playout details
func (c *Client) Playout(programID string, mediaID string, protocol string) (playout *responses.Playout, err error) {
	playoutsURL := fmt.Sprintf("media/playouts.json?program_id=%s&protocol=%s&media_id=%s", programID, protocol, mediaID)
	req, err := c.NewRequest("GET", playoutsURL, nil)
	if err != nil {
		return
	}
	err = c.Do(req, &playout)
	return
}

// ProgramDetails fetch & parse program details for given programID
func (c *Client) ProgramDetails(programID string) (details *results.ProgramDetails, err error) {
	program, err := c.Program(programID)
	if err != nil {
		return
	}
	mediaID := program.MediaID()
	if mediaID == "" {
		err = ErrMediaIDMissing
		return
	}
	playout, err := c.Playout(programID, mediaID, "HLS")
	if err != nil {
		return
	}
	encURL := playout.EncodedURL()
	hlsurl, err := DecryptURL(c.cfg.YleSecret, encURL)
	if err != nil {
		return
	}
	title := program.Title()
	details = &results.ProgramDetails{
		HLSURL:      hlsurl,
		MediaID:     mediaID,
		Title:       title,
		Slug:        Slugify(title),
		SubtitleURL: playout.SubtitleURL("fi"),
	}
	return
}
