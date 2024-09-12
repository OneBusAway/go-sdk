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

	// Define the vehicle ID
	vehicleID := "40_9801"

	// Fetch the trip information
	trip, err := client.TripForVehicle.Get(ctx, vehicleID, onebusaway.TripForVehicleGetParams{
		IncludeSchedule: onebusaway.F(true),
		IncludeTrip:     onebusaway.F(true),
		IncludeStatus:   onebusaway.F(true),
	})

	if err != nil {
		log.Fatalf("Error fetching trip: %v", err)
	}

	fmt.Print(trip.Data.JSON.RawJSON())

}
