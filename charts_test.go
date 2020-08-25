package spotify

import "testing"

func TestGetSpotifyChartTracksRegional(t *testing.T) {
	chartTracks, err := GetSpotifyChartTracksRegional("global", "daily", "latest")
	if err != nil {
		t.Error(err)
	}

	chartTrack := chartTracks[0]
	if chartTrack.Position != 1 {
		t.Errorf("Expected position 1, got %d", chartTrack.Position)
	}
}

func TestGetSpotifyChartsTracksViral(t *testing.T) {
	chartTracks, err := GetSpotifyChartTracksViral("global", "daily", "latest")
	if err != nil {
		t.Error(err)
	}

	chartTrack := chartTracks[0]
	if chartTrack.Position != 1 {
		t.Errorf("Expected position 1, got %d", chartTrack.Position)
	}
}

func TestGetSpotifyChartsTracksRegionalWeekly(t *testing.T) { 
	chartTracks, err := GetSpotifyChartTracksRegional("global", "weekly", "latest")
	if err != nil {
		t.Error(err)
	}

	chartTrack := chartTracks[0]
	if chartTrack.Position != 1 {
		t.Errorf("Expected position 1, got %d", chartTrack.Position)
	}
}

func TestGetSpotifyChartsTracksViralWeekly(t *testing.T) { 
	chartTracks, err := GetSpotifyChartTracksViral("global", "weekly", "latest")
	if err != nil {
		t.Error(err)
	}

	chartTrack := chartTracks[0]
	if chartTrack.Position != 1 {
		t.Errorf("Expected position 1, got %d", chartTrack.Position)
	}
}
func TestGetSpotifyChartsTracksRegionalDate(t *testing.T) { 
	chartTracks, err := GetSpotifyChartTracksRegional("global", "daily", "2020-07-08")
	if err != nil {
		t.Error(err)
	}

	chartTrack := chartTracks[0]
	if chartTrack.Position != 1 {
		t.Errorf("Expected position 1, got %d", chartTrack.Position)
	}
}
func TestGetSpotifyChartsTracksViralDate(t *testing.T) { 
	chartTracks, err := GetSpotifyChartTracksViral("global", "daily", "2020-07-08")
	if err != nil {
		t.Error(err)
	}

	chartTrack := chartTracks[0]
	if chartTrack.Position != 1 {
		t.Errorf("Expected position 1, got %d", chartTrack.Position)
	}
}