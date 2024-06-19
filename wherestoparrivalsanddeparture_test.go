// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package opentransit_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/open-transit-go"
	"github.com/stainless-sdks/open-transit-go/internal/testutil"
	"github.com/stainless-sdks/open-transit-go/option"
)

func TestWhereStopArrivalsAndDepartureList(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := opentransit.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Where.Stop.ArrivalsAndDepartures.List(context.TODO(), "1_75403")
	if err != nil {
		var apierr *opentransit.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
