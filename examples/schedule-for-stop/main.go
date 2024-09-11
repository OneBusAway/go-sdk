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

	// Define the stop ID
	stopId := "1_75403"

	schedule, err := client.ScheduleForStop.Get(ctx, stopId, onebusaway.ScheduleForStopGetParams{
		Date: onebusaway.F(time.Now()),
	})

	if err != nil {
		log.Fatalf("Error fetching schedule: %v", err)
	}

	fmt.Println(schedule.Data.JSON.RawJSON())
}
