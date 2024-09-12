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

	// Define the route ID
	routeID := "1_100224"

	// Fetch the trips for the specified route
	trips, err := client.TripsForRoute.List(ctx, routeID, onebusaway.TripsForRouteListParams{})

	if err != nil {
		log.Fatalf("Error fetching trips for route: %v", err)
	}

	fmt.Print(trips.Data.List)
}
