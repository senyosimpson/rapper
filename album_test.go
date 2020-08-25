package spotify

import (
	"net/http"
	"os"
	"testing"
)

// TODO: Add tests for failure cases

func TestGetAlbum(t *testing.T) {
	f, err := os.Open("test-data/get_album.txt")
	if err != nil {
		panic(err)
	}
	server := newTestServer(http.StatusOK, f)
	client := newTestClient(server.URL)
	defer server.Close()

	album, err := client.GetAlbum("2IjEb3Ob7GHEpLwSgaWJiX", nil)
	if err != nil {
		t.Fatal(err)
	}

	if album.Name != "The Healing Component" {
		t.Errorf("Expected album name The Healing Component, got %s", album.Name)
	}

	if album.Popularity != 56 {
		t.Errorf("Expected popularity score 56, got %d", album.Popularity)
	}

	artist := SimpleArtist{
		Name:     "Mick Jenkins",
		ID:       "1FvjvACFvko2Z91IvDljrx",
		URI:      "spotify:artist:1FvjvACFvko2Z91IvDljrx",
		Endpoint: "https://api.spotify.com/v1/artists/1FvjvACFvko2Z91IvDljrx",
	}
	artists := []SimpleArtist{artist}
	if album.Artists[0] != artists[0] {
		t.Error("Wrong artist identified")
	}
}

func TestGetAlbumTracks(t *testing.T) {
	f, err := os.Open("test-data/get_album_tracks.txt")
	if err != nil {
		panic(err)
	}
	server := newTestServer(http.StatusOK, f)
	client := newTestClient(server.URL)
	defer server.Close()

	tracks, err := client.GetAlbumTracks("2IjEb3Ob7GHEpLwSgaWJiX", nil)
	if err != nil {
		t.Fatal(err)
	}

	track := tracks[0]
	if track.Artists[0].Name != "6LACK" {
		t.Errorf("Expected artist name 6LACK, got %s", track.Artists[0].Name)
	}

	if track.Name != "Never Know" {
		t.Errorf("Expected track Never Know, got %s", track.Name)
	}

	if track.TrackNumber != 1 {
		t.Errorf("Expected track number 0, got %d", track.TrackNumber)
	}
}