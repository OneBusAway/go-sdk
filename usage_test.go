// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package onebusaway_test

import (
	"context"
	"os"
	"testing"

	"github.com/OneBusAway/go-sdk"
	"github.com/OneBusAway/go-sdk/internal/testutil"
	"github.com/OneBusAway/go-sdk/option"
)

func TestUsage(t *testing.T) {
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
	currentTime, err := client.CurrentTime.Get(context.TODO())
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%+v\n", currentTime)
}
