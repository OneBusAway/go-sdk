// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package onebusaway_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/open-transit-go"
	"github.com/stainless-sdks/open-transit-go/internal/testutil"
	"github.com/stainless-sdks/open-transit-go/option"
)

func TestTripsForLocationListWithOptionalParams(t *testing.T) {
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
	_, err := client.TripsForLocation.List(context.TODO(), onebusaway.TripsForLocationListParams{
		Lat:             onebusaway.F(0.000000),
		LatSpan:         onebusaway.F(0.000000),
		Lon:             onebusaway.F(0.000000),
		LonSpan:         onebusaway.F(0.000000),
		IncludeSchedule: onebusaway.F(true),
		IncludeTrip:     onebusaway.F(true),
		Time:            onebusaway.F(int64(0)),
	})
	if err != nil {
		var apierr *onebusaway.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}