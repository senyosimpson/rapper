package spotify

// Copyright is the representation of a Spotify copyright object
type Copyright struct {
	Text string `json:"text"`
	Type string `json:"type"`
}

// SimpleAlbum is the representation of a Spotify simplified album object
type SimpleAlbum struct {
	Name     string         `json:"name"`
	ID       string         `json:"id"`
	Artists  []SimpleArtist `json:"artists"`
	Endpoint string         `json:"href"`
	URI      string         `json:"uri"`
}

// FullAlbum is the representation of a Spotify full album object
type FullAlbum struct {
	SimpleAlbum
	Copyrights  []Copyright     `json:"copyrights"`
	Genres      []string        `json:"genres"`
	RecordLabel string          `json:"label"`
	Popularity  int             `json:"popularity"`
	Tracks      SimpleTrackPage `json:"tracks"`
}

// SimpleAlbumPage is a paging object that has simple albums as items
type SimpleAlbumPage struct {
	Page
	Albums []SimpleAlbum `json:"items"`
}

// FullAlbumPage is a paging object that has full albums as items
type FullAlbumPage struct {
	Page
	Albums []FullAlbum `json:"items"`
}

// GetAlbum gets an album
func (c *Client) GetAlbum(id string, options *Options) (FullAlbum, error) {
	var album FullAlbum

	url := constructURL(options, c.baseURL, "albums", id)
	err := c.get(url, &album)
	if err != nil {
		return FullAlbum{}, err
	}

	// recurse through all pages and fetch the all tracks for the album
	if len(album.Tracks.Next) != 0 {
		// parse the url to get the id and offset
		id, offset, err := getURLParams(album.Tracks.Next)
		if err != nil {
			return FullAlbum{}, err
		}

		options.Offset = offset
		t, err := c.GetAlbumTracks(id, options)
		if err != nil {
			return FullAlbum{}, err
		}
		album.Tracks.Tracks = append(album.Tracks.Tracks, t...)

	}
	return album, nil
}

// GetAlbumTracks gets the album tracks
func (c *Client) GetAlbumTracks(id string, options *Options) ([]SimpleTrack, error) {
	var tracks SimpleTrackPage

	url := constructURL(options, c.baseURL, "albums", id, "tracks")
	err := c.get(url, &tracks)
	if err != nil {
		return nil, err
	}

	// recurse through all pages and fetch the tracks
	if len(tracks.Next) != 0 {
		// parse the url to get the id and offset
		id, offset, err := getURLParams(tracks.Next)
		if err != nil {
			return nil, err
		}

		options.Offset = offset
		t, err := c.GetAlbumTracks(id, options)
		if err != nil {
			return nil, err
		}
		tracks.Tracks = append(tracks.Tracks, t...)
	}

	return tracks.Tracks, nil
}

// GetSeveralAlbums
