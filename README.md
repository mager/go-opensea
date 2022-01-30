# go-opensea

[![Go Reference](https://pkg.go.dev/badge/github.com/mager/go-opensea.svg)](https://pkg.go.dev/github.com/mager/go-opensea)

A Go client for the OpenSea API.

## Endpoints supported

### v1

_https://docs.opensea.io/reference/api-overview_

- [x] Retrieving assets (GetAssets)
- [ ] Retrieving events (GetEvents)
- [ ] Retrieving collections (GetCollections)
- [ ] Retrieving bundles (GetBundles)
- [ ] Retrieving a single asset (GetAsset)
- [ ] Retrieving a single contract (GetContract)
- [x] Retrieving a single collection (GetCollection)
- [ ] Retrieving collection stats (GetCollectionStats)


## Example usage

```go
package main

import (
	"github.com/mager/go-opensea/opensea"
)

func main() {
	// Create a new client
	client := opensea.NewOpenSeaClient("")

	// Get all assets
	addy := "0x3b417FaeE9d2ff636701100891DC2755b5321Cc3" // Jay-Z's wallet address
	assets, err := client.GetAssets(addy)
	if err != nil {
		// TODO: Handle error
	}

	// Print the assets
	for _, asset := range assets {
		client.Log.Info(asset.Name)
	}
}
```

