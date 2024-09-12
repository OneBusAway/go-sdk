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

	agencyID := "1"

	stops, err := client.StopIDsForAgency.List(ctx, agencyID)
	if err != nil {
		log.Fatalf("Error fetching stops: %v", err)
	}

	for _, stop := range stops.Data.List {
		fmt.Println(stop)
	}
}
