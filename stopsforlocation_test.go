// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package tempopentransit_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/TEMP_open-transit-go"
	"github.com/stainless-sdks/TEMP_open-transit-go/internal/testutil"
	"github.com/stainless-sdks/TEMP_open-transit-go/option"
)

func TestStopsForLocationListWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := tempopentransit.NewClient(
		option.WithBaseURL(baseURL),
	)
	err := client.StopsForLocation.List(context.TODO(), tempopentransit.StopsForLocationListParams{
		Key: tempopentransit.F("string"),
		Lat: tempopentransit.F(0.000000),
		Lon: tempopentransit.F(0.000000),
	})
	if err != nil {
		var apierr *tempopentransit.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
