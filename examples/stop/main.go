package main

import (
	"context"
	"fmt"
	"log"

	onebusaway "github.com/OneBusAway/go-sdk"
	"github.com/OneBusAway/go-sdk/option"
)

func main() {

	// Create a new instance of the OneBusAway SDK with the settings
	client := onebusaway.NewClient(
		option.WithAPIKey("TEST"),
		option.WithBaseURL("https://api.pugetsound.onebusaway.org/"),
	)

	ctx := context.Background()

	// Define the stop ID
	stopID := "1_75403"

	// Fetch the stop information
	stop, err := client.Stop.Get(ctx, stopID)
	if err != nil {
		log.Fatalf("Error fetching stop: %v", err)
	}

	fmt.Print(stop.Data.Entry.JSON.RawJSON())
}
