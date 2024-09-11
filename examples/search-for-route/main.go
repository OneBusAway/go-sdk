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

	// Define the search input
	searchInput := "crysta"

	// Fetch the routes
	routes, err := client.SearchForRoute.List(ctx, onebusaway.SearchForRouteListParams{
		Input: onebusaway.F(searchInput),
	})

	if err != nil {
		log.Fatalf("Error fetching routes: %v", err)
	}

	fmt.Print(routes.Data)

}
