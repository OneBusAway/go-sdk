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
	)

	ctx := context.Background()

	agencies, err := client.AgenciesWithCoverage.List(ctx)
	if err != nil {
		log.Fatalf("Error fetching agencies: %v", err)
	}

	for _, agency := range agencies.Data.List {
		fmt.Println(agency.JSON.RawJSON())
	}
}
