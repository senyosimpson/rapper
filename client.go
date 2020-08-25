package spotify

import (
	"encoding/json"
	"net/http"
)

// Client is a struct that contains a http client and the base url to Spotify Web API
type Client struct {
	client  *http.Client
	baseURL string
}

func (c *Client) get(url string, result interface{}) error {
	resp, err := c.client.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	json.NewDecoder(resp.Body).Decode(result)
	return nil
}

// Options is a struct that contains common optional parameters when calling the API
type Options struct {
	Country      string
	Limit        int
	Offset       int
	Market       string
	Locale       string
	IncludeGroup string
	Type         string
	Query        string
}
