package spotify

import (
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

// ChartTrack is a track object that represents a track on the spotify charts
type ChartTrack struct {
	Artist   string
	Name     string
	Position int
	Streams  int
}

// GetSpotifyChartTracksRegional gets the tracks on the Spotify Top 200 chart
// https://spotifycharts.com/regional/{region}/{frequency}/{date}
// Create a utility function for separating out featuring artists
func GetSpotifyChartTracksRegional(region string, frequency string, date string) ([]ChartTrack, error) {
	var tracks []ChartTrack
	var err error

	c := colly.NewCollector()
	c.OnHTML("tbody", func(e *colly.HTMLElement) {
		e.ForEach("tr", func(_ int, e *colly.HTMLElement) {
			track := ChartTrack{}
			trackInfo := strings.Split(e.ChildText("td.chart-table-track"), " by ")
			artist := strings.TrimSpace(trackInfo[1])
			name := strings.TrimSpace(trackInfo[0])
			p := strings.ReplaceAll(e.ChildText("td.chart-table-position"), ",", "")
			position, posErr := strconv.Atoi(p)
			if posErr != nil {
				err = posErr
			}
			s := strings.ReplaceAll(e.ChildText("td.chart-table-streams"), ",", "")
			streams, streamErr := strconv.Atoi(s)
			if streamErr != nil {
				err = streamErr
			}

			track.Artist = artist
			track.Name = name
			track.Position = position
			track.Streams = streams
			tracks = append(tracks, track)
		})
	})
	url := constructURL(nil, ChartBaseURL, "regional", region, frequency, date)
	c.Visit(url)
	if err != nil {
		return nil, err
	}
	return tracks, nil
}

// GetSpotifyChartTracksViral gets the tracks on the Spotify Viral 50 chart
// https://spotifycharts.com/viral/{region}/{frequency}/{date}
// Create a utility function for separating out featuring artists
func GetSpotifyChartTracksViral(region string, frequency string, date string) ([]ChartTrack, error) {
	var tracks []ChartTrack
	var err error

	c := colly.NewCollector()
	c.OnHTML("tbody", func(e *colly.HTMLElement) {
		e.ForEach("tr", func(_ int, e *colly.HTMLElement) {
			track := ChartTrack{}
			trackInfo := strings.Split(e.ChildText("td.chart-table-track"), " by ")
			artist := strings.TrimSpace(trackInfo[1])
			name := strings.TrimSpace(trackInfo[0])
			p := strings.ReplaceAll(e.ChildText("td.chart-table-position"), ",", "")
			position, posErr := strconv.Atoi(p)
			if posErr != nil {
				err = posErr
			}

			track.Artist = artist
			track.Name = name
			track.Position = position
			tracks = append(tracks, track)
		})
	})
	url := constructURL(nil, ChartBaseURL, "viral", region, frequency, date)
	c.Visit(url)
	if err != nil {
		return nil, err
	}
	return tracks, nil
}
