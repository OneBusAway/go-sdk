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

	// Define the route ID
	routeId := "1_100223"

	// Fetch the schedule for the route
	schedule, err := client.ScheduleForRoute.Get(ctx, routeId, onebusaway.ScheduleForRouteGetParams{
		Date: onebusaway.F(time.Now()),
	})

	if err != nil {
		log.Fatalf("Error fetching schedule: %v", err)
	}

	fmt.Println(schedule.Data.JSON.RawJSON())
}
