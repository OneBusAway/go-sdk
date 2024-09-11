package main

import (
	"context"
	"fmt"
	"log"
	"time"

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

	// Define the query parameters
	stopID := "1_75403"
	tripID := "1_604670535"
	serviceDate := time.Unix(1810918000, 0) // Convert Unix timestamp to time.Time

	// Retrieve arrival and departure information
	response, err := client.ArrivalAndDeparture.Get(ctx, stopID, onebusaway.ArrivalAndDepartureGetParams{
		TripID:      onebusaway.F(tripID),
		ServiceDate: onebusaway.F(serviceDate.Unix()),
	})
	if err != nil {
		log.Fatalf("Error retrieving arrival and departure: %v", err)
	}

	fmt.Printf("%+v\n", response)
}
