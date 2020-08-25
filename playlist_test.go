package spotify

import (
	"net/http"
	"os"
	"testing"
)

func TestGetPlaylist(t *testing.T) {
	f, err := os.Open("test-data/get_playlist.txt")
	if err != nil {
		panic(err)
	}
	server := newTestServer(http.StatusOK, f)
	client := newTestClient(server.URL)
	defer server.Close()

	playlist, err := client.GetPlaylist("6iBOHQq18scMsGnLJFkMk0")
	if err != nil {
		t.Error(err)
	}

	if playlist.Name != "TAKE TIME" {
		t.Errorf("Expected playlist name TAKE TIME, got %s", playlist.Name)
	}

	if playlist.Tracks.PlaylistTracks[1].Track.Name != "THE BEACH" {
		t.Errorf("Expected second playlist track name THE BEACH, got %s",
			playlist.Tracks.PlaylistTracks[1].Track.Name)
	}

	if playlist.Tracks.PlaylistTracks[1].Track.Artists[0].Name != "Giveon" {
		t.Errorf("Expected second playlist artist name Giveon, got %s",
			playlist.Tracks.PlaylistTracks[1].Track.Artists[0].Name)
	}

	if playlist.Tracks.PlaylistTracks[1].Track.TrackNumber != 1 {
		t.Errorf("Expected second playlist track track number 1, got track number %d",
			playlist.Tracks.PlaylistTracks[1].Track.TrackNumber)
	}
}

func TestGetFeaturedPlaylists(t *testing.T) {
	f, err := os.Open("test-data/get_featured_playlists.txt")
	if err != nil {
		panic(err)
	}
	server := newTestServer(http.StatusOK, f)
	client := newTestClient(server.URL)
	defer server.Close()

	playlists, err := client.GetFeaturedPlaylists(nil)
	if err != nil {
		t.Error(err)
	}

	if playlists.Message != "Enjoy your Saturday night" {
		t.Errorf("Expected playlist message Enjoy your Saturday night, got %s", playlists.Message)
	}

	if playlists.Playlists.Playlists[0].Name != "Dance Pop" {
		t.Errorf("Expected first playlist name Dance Pop, got %s",
			playlists.Playlists.Playlists[0].Name)
	}

	if playlists.Playlists.Playlists[0].Tracks.Total != 116 {
		t.Errorf("Expected first playlist to have 116 tracks, got %d",
			playlists.Playlists.Playlists[0].Tracks.Total)
	}
}

func TestGetPlaylistTracks(t *testing.T) {
	f, err := os.Open("test-data/get_playlist_tracks.txt")
	if err != nil {
		panic(err)
	}
	server := newTestServer(http.StatusOK, f)
	client := newTestClient(server.URL)
	defer server.Close()

	tracks, err := client.GetPlaylistTracks("37i9dQZF1DWVGy1YP1ojM5", nil)
	if err != nil {
		t.Error(err)
	}

	if tracks[0].Name != "Stickin' (feat. Masego & VanJess)" {
		t.Errorf("Expected first track name Stickin' (feat. Masego & VanJess), got %s",
			tracks[0].Name)
	}

	if tracks[0].TrackNumber != 1 {
		t.Errorf("Expected first track number 1, got track number %d",
			tracks[0].TrackNumber)
	}
}
