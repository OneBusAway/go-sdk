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

	// Define the trip ID
	tripID := "1_100000"

	reportProblemWithTrip, err := client.ReportProblemWithTrip.Get(ctx, tripID, onebusaway.ReportProblemWithTripGetParams{
		UserComment:       onebusaway.F("This is a test report"),
		ServiceDate:       onebusaway.F(time.Now().Unix()),
		VehicleID:         onebusaway.F("1000"),
		UserLat:           onebusaway.F(47.6062),
		UserLon:           onebusaway.F(-122.3321),
		UserOnVehicle:     onebusaway.F(true),
		UserVehicleNumber: onebusaway.F("1234"),
		StopID:            onebusaway.F("1_100000"),
	})

	if err != nil {
		log.Fatalf("Error reporting problem with trip: %v", err)
	}

	fmt.Println(reportProblemWithTrip)
}
