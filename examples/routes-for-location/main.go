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

	location := onebusaway.RoutesForLocationListParams{
		Lat: onebusaway.F(47.6097),
		Lon: onebusaway.F(-122.3331),
	}

	// Fetch the routes for the location
	routes, err := client.RoutesForLocation.List(ctx, location)

	if err != nil {
		log.Fatalf("Error fetching routes: %v", err)
	}

	fmt.Print(routes.Data.JSON.RawJSON())

}
