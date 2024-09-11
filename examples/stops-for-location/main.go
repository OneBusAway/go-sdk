package main

import (
	"context"
	"fmt"
	"log"
	"strings"

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

	// Define the location parameters for the Space Needle
	params := onebusaway.StopsForLocationListParams{
		Lat:    onebusaway.F(47.6205),
		Lon:    onebusaway.F(-122.3493),
		Radius: onebusaway.F(500.0),
	}

	// Fetch stops for the location
	response, err := client.StopsForLocation.List(ctx, params)
	if err != nil {
		log.Fatalf("Error fetching stops: %v", err)
	}

	stops := response.Data.List
	references := response.Data.References

	// Create a map for easy route lookup
	referenceMap := make(map[string]onebusaway.ReferencesRoute)
	for _, route := range references.Routes {
		referenceMap[route.ID] = route
	}

	// Print information for each stop
	for _, stop := range stops {
		fmt.Printf("%s (%f, %f)\n", stop.Name, stop.Lat, stop.Lon)
		fmt.Println("  Routes:")

		for _, routeID := range stop.RouteIDs {
			route, exists := referenceMap[routeID]
			if !exists {
				continue
			}

			descriptionParts := []string{}
			if route.ShortName != "" {
				descriptionParts = append(descriptionParts, route.ShortName)
			}
			if route.Description != "" {
				descriptionParts = append(descriptionParts, route.Description)
			}
			description := strings.Join(descriptionParts, " - ")

			fmt.Printf("    %s\n", description)
		}

		fmt.Println()
	}
}
