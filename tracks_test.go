package spotify

import (
	"net/http"
	"os"
	"testing"
)

func TestGetTrack(t *testing.T) {
	f, err := os.Open("test-data/get_track.txt")
	if err != nil {
		panic(err)
	}
	server := newTestServer(http.StatusOK, f)
	client := newTestClient(server.URL)
	defer server.Close()

	track, err := client.GetTrack("2SVIkGkmY9ZPZFniCwa7ar", nil)
	if err != nil {
		t.Error(err)
	}

	if track.Name != "you should call mum" {
		t.Errorf("Expected track name you should call mum, got %s", track.Name)
	}

	if track.Popularity != 49 {
		t.Errorf("Expected track popularity 49, got %d", track.Popularity)
	}

	if track.TrackNumber != 4 {
		t.Errorf("Expected track number 4, got %d", track.TrackNumber)
	}
}
