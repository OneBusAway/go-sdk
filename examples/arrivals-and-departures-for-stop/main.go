package main

import (
	"context"
	"fmt"
	"log"
	"time"

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

	// Define the query parameters
	stopID := "1_75403" // Replace with actual stop ID

	minutesBefore := int64(5) // include vehicles having arrived or departed in the previous n minutes (default=5)
	minutesAfter := int64(35) // include vehicles arriving or departing in the next n minutes (default=35)

	// Retrieve arrival and departure information
	response, err := client.ArrivalAndDeparture.List(ctx, stopID, onebusaway.ArrivalAndDepartureListParams{
		MinutesAfter:  onebusaway.F(minutesAfter),
		MinutesBefore: onebusaway.F(minutesBefore),
	})
	if err != nil {
		log.Fatalf("Error retrieving arrival and departure: %v", err)
	}

	for _, ad := range response.Data.Entry.ArrivalsAndDepartures {
		fmt.Printf("Route: %s\n", ad.RouteShortName)
		fmt.Printf("Trip Headsign: %s\n", ad.TripHeadsign)
		fmt.Printf("Predicted Arrival Time: %s\n", time.Unix(ad.PredictedArrivalTime/1000, 0))
		fmt.Printf("Scheduled Arrival Time: %s\n", time.Unix(ad.ScheduledArrivalTime/1000, 0))
		fmt.Printf("Vehicle ID: %s\n", ad.VehicleID)
		fmt.Println()
	}
}
