package spotify

import (
	"net/http"
	"os"
	"testing"
)

func TestGetArtist(t *testing.T) {
	f, err := os.Open("test-data/get_artist.txt")
	if err != nil {
		panic(err)
	}

	server := newTestServer(http.StatusOK, f)
	client := newTestClient(server.URL)
	defer server.Close()

	artist, err := client.GetArtist("3TVXtAsR1Inumwj472S9r4")
	if err != nil {
		t.Fatal(err)
	}

	if artist.Name != "Drake" {
		t.Errorf("Expected artist name Drake, got %s", artist.Name)
	}
	
	if artist.Followers.Total != 47381748 {
		t.Errorf("Expected number of followers 47381748, got %d", artist.Followers.Total)
	}

	if artist.Popularity != 100 {
		t.Errorf("Expected popularity 100, got %d", artist.Popularity)
	}
}

func TestGetArtistTopTracks(t *testing.T) {
	f, err := os.Open("test-data/get_artist_top_tracks.txt")
	if err != nil {
		panic(err)
	}

	server := newTestServer(http.StatusOK, f)
	client := newTestClient(server.URL)
	defer server.Close()

	tracks, err := client.GetArtistTopTracks("3TVXtAsR1Inumwj472S9r4", nil)
	if err != nil {
		t.Fatal(err)
	}

	track := tracks[0]
	if track.Album.Name != "Dark Lane Demo Tapes" {
		t.Errorf("Expected track album Dark Lane Demo Tapes, got %s", track.Album.Name)
	}

	if track.Name != "Toosie Slide" {
		t.Errorf("Expected track name Toosie Slide, got %s", track.Name)
	}

	if track.Artists[0].Name != "Drake" {
		t.Errorf("Expected track artist name Drake, got %s", track.Artists[0].Name)
	}

	if track.Popularity != 87 {
		t.Errorf("Expected popularity 87, got %d", track.Popularity)
	}
}

func TestGetRelatedArtists(t *testing.T) {
	f, err := os.Open("test-data/get_related_artists.txt")
	if err != nil {
		panic(err)
	}

	server := newTestServer(http.StatusOK, f)
	client := newTestClient(server.URL)
	defer server.Close()

	artists, err := client.GetRelatedArtists("3TVXtAsR1Inumwj472S9r4")
	if err != nil {
		t.Fatal(err)
	}

	artist := artists[1]

	if artist.Name != "J. Cole" {
		t.Errorf("Expected related artist name J. Cole, got %s", artist.Name)
	}

	if artist.Popularity != 89 {
		t.Errorf("Expected related artist popularity 89, got %d", artist.Popularity)
	}

	if artist.Genres[0] != "conscious hip hop" {
		t.Errorf("Expected related artist genre conscious hip hop, got %s", artist.Genres[0])
	}

	if artist.Followers.Total != 11103702 {
		t.Errorf("Expected related artist number of followers 11103702, got %d", artist.Followers.Total)
	}
}

func TestGetMonthlyListeners(t *testing.T) {
	_, err := GetArtistMonthlyListeners("3TVXtAsR1Inumwj472S9r4")
	if err != nil {
		t.Error(err)
	}
}