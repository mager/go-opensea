# go-opensea

[![Go Reference](https://pkg.go.dev/badge/github.com/mager/go-opensea.svg)](https://pkg.go.dev/github.com/mager/go-opensea)

A Go client for the OpenSea API.

## Endpoints supported

- [x] Get Assets

## Example usage

```go
package main

import (
	"fmt"

	"github.com/mager/go-opensea"
)

func main() {
	// Create a new client
	client := opensea.NewClient("YOUR_API_KEY")

	// Get all assets
	assets, err := client.GetAssets("0x064DcA21b1377D1655AC3CA3e95282D9494E5611")
	if err != nil {
		// TODO: Handle error
	}

	// Print the assets
	for _, asset := range assets {
		fmt.Println(asset.Name)
	}
}
```