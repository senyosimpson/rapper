package spotify

// Page is the base page object that is common among all Spotify paging objects
type Page struct {
	Endpoint string `json:"href"`
	Limit    int    `json:"limit"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Offset   int    `json:"offset"`
	Total    int    `json:"total"`
}
