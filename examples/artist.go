package main

import (
	"context"
	"fmt"

	"jollofspotify/spotify"
)

func main() {
	client := spotify.NewClientCredentialsAuth().NewClient(context.Background())
	artist, err := client.GetArtist("3TVXtAsR1Inumwj472S9r4")
	if err != nil {
		fmt.Println("Error")
		return
	}

	fmt.Println("Name:", artist.Name)
	fmt.Println("Number of followers:", artist.Followers.Total)
	fmt.Println("Genres:", artist.Genres)
	fmt.Println("Popularity:", artist.Popularity)
}
