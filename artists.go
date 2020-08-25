package spotify

import (
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

// SimpleArtist is the representation of a Spotify simplified artist object
type SimpleArtist struct {
	Name     string `json:"name"`
	ID       string `json:"id"`
	URI      string `json:"uri"`
	Endpoint string `json:"href"`
}

// Followers is the representation of a Spotify follower object
type Followers struct {
	Total uint `json:"total"`
}

// FullArtist is the representation of an Spotify artist object
type FullArtist struct {
	SimpleArtist
	Followers        Followers `json:"followers"`
	Genres           []string  `json:"genres"`
	Popularity       int       `json:"popularity"`
	MonthlyListeners int
}

// FullArtistPage is a paging object that has full artists as items
type FullArtistPage struct {
	Page
	Artist []FullArtist `json:"items"`
}

// GetArtist gets the relevant artist information
// Can possibly fetch the monthly listeners here
func (c *Client) GetArtist(id string) (FullArtist, error) {
	var artist FullArtist

	url := constructURL(nil, c.baseURL, "artists", id)
	err := c.get(url, &artist)
	if err != nil {
		return FullArtist{}, err
	}

	return artist, nil
}

// TODO: Create endpoints for the other include groups

// GetArtistAlbums gets an all the albums by an artist
// Need a function to remove duplicates here
// Need a process for specifying options
func (c *Client) GetArtistAlbums(id string, options *Options) ([]SimpleAlbum, error) {
	var albums SimpleAlbumPage

	url := constructURL(options, c.baseURL, "artists", id, "albums")
	err := c.get(url, &albums)
	if err != nil {
		return nil, err
	}

	if len(albums.Next) != 0 {
		id, offset, err := getURLParams(albums.Next)
		if err != nil {
			return nil, err
		}

		options.Offset = offset
		a, err := c.GetArtistAlbums(id, options)
		if err != nil {
			return nil, err
		}
		albums.Albums = append(albums.Albums, a...)
	}

	return albums.Albums, nil
}

// GetArtistTopTracks gets an artist's top tracks in a specified country
func (c *Client) GetArtistTopTracks(id string, options *Options) ([]FullTrack, error) {
	var t struct {
		TopTracks []FullTrack `json:"tracks"`
	}

	url := constructURL(options, c.baseURL, "artists", "id", "top-tracks")
	err := c.get(url, &t)
	if err != nil {
		return nil, err
	}
	return t.TopTracks, nil
}

// GetRelatedArtists gets the artists related to the given artist
func (c *Client) GetRelatedArtists(id string) ([]FullArtist, error) {
	var a struct {
		RelatedArtists []FullArtist `json:"artists"`
	}

	url := constructURL(nil, c.baseURL, "artists", id, "related-artists")
	err := c.get(url, &a)
	if err != nil {
		return nil, err
	}
	return a.RelatedArtists, nil
}

// GetArtistMonthlyListeners gets the number of monthly listeners of an artist
func GetArtistMonthlyListeners(id string) (int, error) {
	var numMonthlyListeners int
	var err error

	c := colly.NewCollector()
	c.OnHTML("section", func(e *colly.HTMLElement) {
		if strings.Contains(e.Text, "Monthly Listeners") {
			monthlyListeners := strings.TrimSpace(e.Text)
			ind := strings.Index(monthlyListeners, "M")
			monthlyListeners = strings.ReplaceAll(monthlyListeners[:ind], ",", "")
			numMonthlyListeners, err = strconv.Atoi(monthlyListeners)
		}
	})

	c.Visit(constructURL(nil, SpotifyBaseURL, "artist", id))
	if err != nil {
		return 0, err
	}
	return numMonthlyListeners, err
}

// GetSeveralArtists
