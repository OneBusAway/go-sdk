package main

import (
	"context"
	"fmt"
	"log"

	onebusaway "github.com/OneBusAway/go-sdk"
	"github.com/OneBusAway/go-sdk/option"
)

func main() {

	// Initialize the OneBusAway client with the API key
	client := onebusaway.NewClient(
		option.WithAPIKey("TEST"),
		option.WithBaseURL("https://api.pugetsound.onebusaway.org/"),
	)

	ctx := context.Background()

	blockID := "1_7331695" // Replace with actual block ID

	block, err := client.Block.Get(ctx, blockID)

	if err != nil {
		log.Fatalf("Error fetching block: %v", err)
	}

	fmt.Print(block.Data.JSON.RawJSON())
}
