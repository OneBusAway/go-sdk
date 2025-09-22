// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package onebusaway

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/OneBusAway/go-sdk/internal/apiquery"
	"github.com/OneBusAway/go-sdk/internal/param"
	"github.com/OneBusAway/go-sdk/internal/requestconfig"
	"github.com/OneBusAway/go-sdk/option"
	"github.com/OneBusAway/go-sdk/shared"
)

// ReportProblemWithStopService contains methods and other services that help with
// interacting with the onebusaway-sdk API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewReportProblemWithStopService] method instead.
type ReportProblemWithStopService struct {
	Options []option.RequestOption
}

// NewReportProblemWithStopService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewReportProblemWithStopService(opts ...option.RequestOption) (r *ReportProblemWithStopService) {
	r = &ReportProblemWithStopService{}
	r.Options = opts
	return
}

// Submit a user-generated problem report for a stop
func (r *ReportProblemWithStopService) Get(ctx context.Context, stopID string, query ReportProblemWithStopGetParams, opts ...option.RequestOption) (res *shared.ResponseWrapper, err error) {
	opts = slices.Concat(r.Options, opts)
	if stopID == "" {
		err = errors.New("missing required stopID parameter")
		return
	}
	path := fmt.Sprintf("api/where/report-problem-with-stop/%s.json", stopID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type ReportProblemWithStopGetParams struct {
	// A string code identifying the nature of the problem
	Code param.Field[ReportProblemWithStopGetParamsCode] `query:"code"`
	// Additional comment text supplied by the user describing the problem
	UserComment param.Field[string] `query:"userComment"`
	// The reporting user’s current latitude
	UserLat param.Field[float64] `query:"userLat"`
	// The reporting user’s location accuracy, in meters
	UserLocationAccuracy param.Field[float64] `query:"userLocationAccuracy"`
	// The reporting user’s current longitude
	UserLon param.Field[float64] `query:"userLon"`
}

// URLQuery serializes [ReportProblemWithStopGetParams]'s query parameters as
// `url.Values`.
func (r ReportProblemWithStopGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// A string code identifying the nature of the problem
type ReportProblemWithStopGetParamsCode string

const (
	ReportProblemWithStopGetParamsCodeStopNameWrong      ReportProblemWithStopGetParamsCode = "stop_name_wrong"
	ReportProblemWithStopGetParamsCodeStopNumberWrong    ReportProblemWithStopGetParamsCode = "stop_number_wrong"
	ReportProblemWithStopGetParamsCodeStopLocationWrong  ReportProblemWithStopGetParamsCode = "stop_location_wrong"
	ReportProblemWithStopGetParamsCodeRouteOrTripMissing ReportProblemWithStopGetParamsCode = "route_or_trip_missing"
	ReportProblemWithStopGetParamsCodeOther              ReportProblemWithStopGetParamsCode = "other"
)

func (r ReportProblemWithStopGetParamsCode) IsKnown() bool {
	switch r {
	case ReportProblemWithStopGetParamsCodeStopNameWrong, ReportProblemWithStopGetParamsCodeStopNumberWrong, ReportProblemWithStopGetParamsCodeStopLocationWrong, ReportProblemWithStopGetParamsCodeRouteOrTripMissing, ReportProblemWithStopGetParamsCodeOther:
		return true
	}
	return false
}
