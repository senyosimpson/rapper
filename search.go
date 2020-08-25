package spotify

// SearchArtist is used to search for an artist
func (c *Client) SearchArtist(name string) (FullArtist, error) {
	var artist struct {
		Page FullArtistPage `json:"artists"`
	}

	options := &Options{}
	options.Query = name
	options.Type = "artist"
	options.Limit = 1
	url := constructURL(options, c.baseURL, "search")

	err := c.get(url, &artist)
	if err != nil {
		return FullArtist{}, err
	}

	return artist.Page.Artist[0], nil
}
