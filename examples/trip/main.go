package main

import (
	"context"
	"fmt"
	"log"

	onebusaway "github.com/stainless-sdks/open-transit-go"
	"github.com/stainless-sdks/open-transit-go/option"
)

func main() {

	// Create a new instance of the OneBusAway SDK with the settings
	client := onebusaway.NewClient(
		option.WithAPIKey("TEST"),
		option.WithBaseURL("https://api.pugetsound.onebusaway.org/"),
	)

	ctx := context.Background()

	// Define the trip ID
	tripID := "1_100224_1"

	// Fetch the trip information
	trip, err := client.Trip.Get(ctx, tripID)

	if err != nil {
		log.Fatalf("Error fetching trip: %v", err)
	}

	fmt.Print(trip.Data.JSON.RawJSON())
}
