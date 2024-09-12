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

	currentTime, err := client.CurrentTime.Get(ctx)

	if err != nil {
		log.Fatalf("Error fetching current time: %v", err)
	}

	fmt.Println(currentTime.Data.JSON.RawJSON())
}
