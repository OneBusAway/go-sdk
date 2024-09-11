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

	stopID := "1_100000"

	reportProblemWithStop, err := client.ReportProblemWithStop.Get(ctx, stopID, onebusaway.ReportProblemWithStopGetParams{
		UserComment: onebusaway.F("This is a test report"),
		UserLat:     onebusaway.F(47.6062),
		UserLon:     onebusaway.F(-122.3321),
	})

	if err != nil {
		log.Fatalf("Error reporting problem with stop: %v", err)
	}

	fmt.Println(reportProblemWithStop)
}
