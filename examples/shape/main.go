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

	shapeId := "1_10002005"

	shape, err := client.Shape.Get(ctx, shapeId)
	if err != nil {
		log.Fatalf("Error fetching shape: %v", err)
	}

	fmt.Println(shape.Data.JSON.RawJSON())
}
