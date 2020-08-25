package spotify

// PlaylistTrack is a representation of a Spotify playlist track object
type PlaylistTrack struct {
	Local bool      `json:"is_local"`
	Track FullTrack `json:"track"`
}

// PlaylistTrackPage is a paging object that has playlist tracks as items
type PlaylistTrackPage struct {
	Page
	PlaylistTracks []PlaylistTrack `json:"items"`
}

// SimplePlaylist is a representation of a Spotify simple playlist object
type SimplePlaylist struct {
	Name        string            `json:"name"`
	ID          string            `json:"id"`
	Description string            `json:"description"`
	Tracks      PlaylistTrackPage `json:"tracks"`
	Endpoint    string            `json:"href"`
	URI         string            `json:"uri"`
}

// FullPlaylist is a representation of a Spotify full playlist object
type FullPlaylist struct {
	SimplePlaylist
	Followers Followers `json:"followers"`
}

// SimplePlaylistPage is a paging object that has simple tracks as items
type SimplePlaylistPage struct {
	Page
	Playlist []SimplePlaylist `json:"items"`
}

// FeaturedPlaylistTrack is a track contained within a featured playlist
type FeaturedPlaylistTrack struct {
	Endpoint string `json:"href"`
	Total    int    `json:"total"`
}

// SimpleFeaturedPlaylist is a representation of a Spotify simple playlist object returned
// by the featured playlist endpoint. The tracks object is different
type SimpleFeaturedPlaylist struct {
	Name        string                `json:"name"`
	ID          string                `json:"id"`
	Description string                `json:"description"`
	Tracks      FeaturedPlaylistTrack `json:"tracks"`
	Endpoint    string                `json:"href"`
	URI         string                `json:"uri"`
}

// FeaturedPlaylistPage is a paging object that has featured playlist tracks as items
type FeaturedPlaylistPage struct {
	Page
	Playlists []SimpleFeaturedPlaylist `json:"items"`
}

// FeaturedPlaylist is a representation of a featured playlist end point response
type FeaturedPlaylist struct {
	Message   string               `json:"message"`
	Playlists FeaturedPlaylistPage `json:"playlists"`
}

// GetPlaylist gets a playlist
func (c *Client) GetPlaylist(id string) (FullPlaylist, error) {
	var playlist FullPlaylist

	url := constructURL(nil, c.baseURL, "playlists", id)
	err := c.get(url, &playlist)
	if err != nil {
		return FullPlaylist{}, err
	}
	return playlist, nil
}

// GetPlaylistTracks gets the tracks in a playlist
func (c *Client) GetPlaylistTracks(id string, options *Options) ([]FullTrack, error) {
	var tracks PlaylistTrackPage
	var fullTracks []FullTrack

	url := constructURL(options, c.baseURL, "playlists", id, "tracks")
	err := c.get(url, &tracks)
	if err != nil {
		return nil, err
	}

	// add tracks to return
	for _, track := range tracks.PlaylistTracks {
		fullTracks = append(fullTracks, track.Track)
	}

	// recurse through all pages and fetch the all tracks for the album
	if len(tracks.Next) != 0 {
		// parse the url to get the id and offset
		id, offset, err := getURLParams(tracks.Next)
		if err != nil {
			return nil, err
		}

		options.Offset = offset
		t, err := c.GetPlaylistTracks(id, options)
		if err != nil {
			return nil, err
		}
		fullTracks = append(fullTracks, t...)
	}
	return fullTracks, nil
}

// GetFeaturedPlaylists gets featured playlists
func (c *Client) GetFeaturedPlaylists(options *Options) (FeaturedPlaylist, error) {
	var playlist FeaturedPlaylist

	url := constructURL(options, c.baseURL, "browse", "featured-playlists")
	err := c.get(url, &playlist)
	if err != nil {
		return FeaturedPlaylist{}, err
	}

	if len(playlist.Playlists.Next) != 0 {
		url := playlist.Playlists.Next
		p, err := c.getRestPlaylists(url)
		if err != nil {
			return FeaturedPlaylist{}, err
		}
		playlist.Playlists.Playlists = append(playlist.Playlists.Playlists, p...)
	}
	return playlist, nil
}

// Helper function for getting another playlist with a specified url
func (c *Client) getRestPlaylists(url string) ([]SimpleFeaturedPlaylist, error) {
	var playlist FeaturedPlaylist

	err := c.get(url, &playlist)
	if err != nil {
		return nil, err
	}

	if len(playlist.Playlists.Next) != 0 {
		url := playlist.Playlists.Next
		t, err := c.getRestPlaylists(url)
		if err != nil {
			return nil, err
		}
		playlist.Playlists.Playlists = append(playlist.Playlists.Playlists, t...)
	}
	return playlist.Playlists.Playlists, err
}
