// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package tempopentransit_test

import (
	"context"
	"os"
	"testing"

	"github.com/stainless-sdks/TEMP_open-transit-go"
	"github.com/stainless-sdks/TEMP_open-transit-go/internal/testutil"
	"github.com/stainless-sdks/TEMP_open-transit-go/option"
)

func TestUsage(t *testing.T) {
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
	_, err := client.AgenciesWithCoverage.List(context.TODO(), tempopentransit.AgenciesWithCoverageListParams{})
	if err != nil {
		t.Error(err)
	}
}
