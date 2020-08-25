package spotify

// SimpleTrack is the representation of a Spotify simple track object
type SimpleTrack struct {
	Name        string         `json:"name"`
	ID          string         `json:"id"`
	TrackNumber int            `json:"track_number"`
	Artists     []SimpleArtist `json:"artists"`
	Endpoint    string         `json:"href"`
	PreviewURL  string         `json:"preview_url"`
	URI         string         `json:"uri"`
}

// FullTrack is the representation of a Spotify full track object
type FullTrack struct {
	SimpleTrack
	Album      SimpleAlbum `json:"album"`
	Popularity int         `json:"popularity"`
}

// SimpleTrackPage is a paging object that has simple tracks as items
type SimpleTrackPage struct {
	Page
	Tracks []SimpleTrack `json:"items"`
}

// FullTrackPage is a paging object that has full tracks as items
type FullTrackPage struct {
	Page
	Tracks []FullTrack `json:"items"`
}

// GetTrack gets a track
func (c *Client) GetTrack(id string, options *Options) (FullTrack, error) {
	var track FullTrack

	url := constructURL(options, c.baseURL, "tracks", id)
	err := c.get(url, &track)
	if err != nil {
		return FullTrack{}, err
	}
	return track, nil
}

// GetSeveralTracks
