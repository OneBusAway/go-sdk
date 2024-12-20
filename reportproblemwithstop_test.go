// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package onebusaway_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/OneBusAway/go-sdk"
	"github.com/OneBusAway/go-sdk/internal/testutil"
	"github.com/OneBusAway/go-sdk/option"
)

func TestReportProblemWithStopGetWithOptionalParams(t *testing.T) {
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
	_, err := client.ReportProblemWithStop.Get(
		context.TODO(),
		"stopID",
		onebusaway.ReportProblemWithStopGetParams{
			Code:                 onebusaway.F(onebusaway.ReportProblemWithStopGetParamsCodeStopNameWrong),
			UserComment:          onebusaway.F("userComment"),
			UserLat:              onebusaway.F(0.000000),
			UserLocationAccuracy: onebusaway.F(0.000000),
			UserLon:              onebusaway.F(0.000000),
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