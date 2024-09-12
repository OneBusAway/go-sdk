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
	)

	ctx := context.Background()

	// Define the trip ID
	tripID := "40_608344966"

	// Fetch the trip details
	tripDetails, err := client.TripDetails.Get(ctx, tripID, onebusaway.TripDetailGetParams{})

	if err != nil {
		log.Fatalf("Error fetching trip details: %v", err)
	}

	fmt.Print(tripDetails.Data.JSON.RawJSON())

}
