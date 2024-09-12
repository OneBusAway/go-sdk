package main

import (
	"context"
	"fmt"
	"log"

	onebusaway "github.com/OneBusAway/go-sdk"
	"github.com/OneBusAway/go-sdk/option"
)

func main() {

	// Initialize the OneBusAway client with the API key
	client := onebusaway.NewClient(
		option.WithAPIKey("TEST"),
		option.WithBaseURL("https://api.pugetsound.onebusaway.org/"),
	)

	ctx := context.Background()
	agencies, err := client.AgenciesWithCoverage.List(ctx)
	if err != nil {
		log.Fatalf("Error fetching agencies: %v", err)
	}

	fmt.Print(agencies.Data.List[0].JSON.RawJSON())
}
