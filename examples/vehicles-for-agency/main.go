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

	// Define the agency ID
	agencyID := "1"

	// Fetch the vehicles for the specified agency
	vehicles, err := client.VehiclesForAgency.List(ctx, agencyID, onebusaway.VehiclesForAgencyListParams{})

	if err != nil {
		log.Fatalf("Error fetching vehicles for agency: %v", err)
	}

	fmt.Print(vehicles.Data.JSON.RawJSON())
}
