// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package onebusaway

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/OneBusAway/go-sdk/internal/apiquery"
	"github.com/OneBusAway/go-sdk/internal/param"
	"github.com/OneBusAway/go-sdk/internal/requestconfig"
	"github.com/OneBusAway/go-sdk/option"
	"github.com/OneBusAway/go-sdk/shared"
)

// ReportProblemWithTripService contains methods and other services that help with
// interacting with the onebusaway-sdk API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewReportProblemWithTripService] method instead.
type ReportProblemWithTripService struct {
	Options []option.RequestOption
}

// NewReportProblemWithTripService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewReportProblemWithTripService(opts ...option.RequestOption) (r *ReportProblemWithTripService) {
	r = &ReportProblemWithTripService{}
	r.Options = opts
	return
}

// Submit a user-generated problem report for a particular trip.
func (r *ReportProblemWithTripService) Get(ctx context.Context, tripID string, query ReportProblemWithTripGetParams, opts ...option.RequestOption) (res *shared.ResponseWrapper, err error) {
	opts = append(r.Options[:], opts...)
	if tripID == "" {
		err = errors.New("missing required tripID parameter")
		return
	}
	path := fmt.Sprintf("api/where/report-problem-with-trip/%s.json", tripID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type ReportProblemWithTripGetParams struct {
	// A string code identifying the nature of the problem
	Code param.Field[ReportProblemWithTripGetParamsCode] `query:"code"`
	// The service date of the trip
	ServiceDate param.Field[int64] `query:"serviceDate"`
	// A stop ID indicating where the user is experiencing the problem
	StopID param.Field[string] `query:"stopID"`
	// Additional comment text supplied by the user describing the problem
	UserComment param.Field[string] `query:"userComment"`
	// The reporting user’s current latitude
	UserLat param.Field[float64] `query:"userLat"`
	// The reporting user’s location accuracy, in meters
	UserLocationAccuracy param.Field[float64] `query:"userLocationAccuracy"`
	// The reporting user’s current longitude
	UserLon param.Field[float64] `query:"userLon"`
	// Indicator if the user is on the transit vehicle experiencing the problem
	UserOnVehicle param.Field[bool] `query:"userOnVehicle"`
	// The vehicle number, as reported by the user
	UserVehicleNumber param.Field[string] `query:"userVehicleNumber"`
	// The vehicle actively serving the trip
	VehicleID param.Field[string] `query:"vehicleID"`
}

// URLQuery serializes [ReportProblemWithTripGetParams]'s query parameters as
// `url.Values`.
func (r ReportProblemWithTripGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// A string code identifying the nature of the problem
type ReportProblemWithTripGetParamsCode string

const (
	ReportProblemWithTripGetParamsCodeVehicleNeverCame       ReportProblemWithTripGetParamsCode = "vehicle_never_came"
	ReportProblemWithTripGetParamsCodeVehicleCameEarly       ReportProblemWithTripGetParamsCode = "vehicle_came_early"
	ReportProblemWithTripGetParamsCodeVehicleCameLate        ReportProblemWithTripGetParamsCode = "vehicle_came_late"
	ReportProblemWithTripGetParamsCodeWrongHeadsign          ReportProblemWithTripGetParamsCode = "wrong_headsign"
	ReportProblemWithTripGetParamsCodeVehicleDoesNotStopHere ReportProblemWithTripGetParamsCode = "vehicle_does_not_stop_here"
	ReportProblemWithTripGetParamsCodeOther                  ReportProblemWithTripGetParamsCode = "other"
)

func (r ReportProblemWithTripGetParamsCode) IsKnown() bool {
	switch r {
	case ReportProblemWithTripGetParamsCodeVehicleNeverCame, ReportProblemWithTripGetParamsCodeVehicleCameEarly, ReportProblemWithTripGetParamsCodeVehicleCameLate, ReportProblemWithTripGetParamsCodeWrongHeadsign, ReportProblemWithTripGetParamsCodeVehicleDoesNotStopHere, ReportProblemWithTripGetParamsCodeOther:
		return true
	}
	return false
}
