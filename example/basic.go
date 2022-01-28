package main

import (
	"github.com/mager/go-opensea/opensea"
)

func main() {
	// Create a new client
	client := opensea.NewOpenSeaClient("") // Your API key goes here

	// Get all assets
	address := "0x3b417FaeE9d2ff636701100891DC2755b5321Cc3" // Jay-Z's wallet address
	assets, err := client.GetAssets(address)
	if err != nil {
		// TODO: Handle error
	}

	// Print the assets
	for _, asset := range assets {
		client.Log.Info(asset.Name)
	}
}
