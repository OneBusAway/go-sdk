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

	routeID := "1_100224"

	route, err := client.Route.Get(ctx, routeID)

	if err != nil {
		log.Fatalf("Error fetching route: %v", err)
	}

	fmt.Print(route)

}
