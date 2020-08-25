package spotify

import (
	"net/http"
	"os"
	"testing"
)

func TestSearchArtist(t *testing.T) {
	f, err := os.Open("test-data/search_artist.txt")
	if err != nil {
		panic(err)
	}
	server := newTestServer(http.StatusOK, f)
	client := newTestClient(server.URL)
	defer server.Close()

	artist, err := client.SearchArtist("Drake")
	if err != nil {
		t.Error(err)
	}

	if artist.Name != "Drake" {
		t.Errorf("Expected artist name Drake, got %s", artist.Name)
	}

	if artist.Popularity != 100 {
		t.Errorf("Expected popularity 100, got %d", artist.Popularity)
	}

	if artist.Followers.Total != 47381748 {
		t.Errorf("Expected total followers 47381748, %d", artist.Followers.Total)
	}
}
