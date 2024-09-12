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

	// Define the agency ID
	agencyID := "1"

	// Fetch the routes for the agency
	routes, err := client.RoutesForAgency.List(ctx, agencyID)

	if err != nil {
		log.Fatalf("Error fetching routes: %v", err)
	}

	fmt.Print(routes.Data.JSON.RawJSON())
}
