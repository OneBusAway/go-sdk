// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package onebusaway_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/OneBusAway/go-sdk"
	"github.com/OneBusAway/go-sdk/internal/testutil"
	"github.com/OneBusAway/go-sdk/option"
)

func TestArrivalAndDepartureGetWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := onebusaway.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.ArrivalAndDeparture.Get(
		context.TODO(),
		"1_75403",
		onebusaway.ArrivalAndDepartureGetParams{
			ServiceDate:  onebusaway.F(int64(0)),
			TripID:       onebusaway.F("tripId"),
			StopSequence: onebusaway.F(int64(0)),
			Time:         onebusaway.F(int64(0)),
			VehicleID:    onebusaway.F("vehicleId"),
		},
	)
	if err != nil {
		var apierr *onebusaway.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestArrivalAndDepartureListWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := onebusaway.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.ArrivalAndDeparture.List(
		context.TODO(),
		"1_75403",
		onebusaway.ArrivalAndDepartureListParams{
			MinutesAfter:  onebusaway.F(int64(0)),
			MinutesBefore: onebusaway.F(int64(0)),
			Time:          onebusaway.F(time.Now()),
		},
	)
	if err != nil {
		var apierr *onebusaway.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
