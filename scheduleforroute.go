// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package onebusaway

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/OneBusAway/go-sdk/internal/apijson"
	"github.com/OneBusAway/go-sdk/internal/apiquery"
	"github.com/OneBusAway/go-sdk/internal/param"
	"github.com/OneBusAway/go-sdk/internal/requestconfig"
	"github.com/OneBusAway/go-sdk/option"
	"github.com/OneBusAway/go-sdk/shared"
)

// ScheduleForRouteService contains methods and other services that help with
// interacting with the onebusaway-sdk API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewScheduleForRouteService] method instead.
type ScheduleForRouteService struct {
	Options []option.RequestOption
}

// NewScheduleForRouteService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewScheduleForRouteService(opts ...option.RequestOption) (r *ScheduleForRouteService) {
	r = &ScheduleForRouteService{}
	r.Options = opts
	return
}

// Retrieve the full schedule for a route on a particular day
func (r *ScheduleForRouteService) Get(ctx context.Context, routeID string, query ScheduleForRouteGetParams, opts ...option.RequestOption) (res *ScheduleForRouteGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if routeID == "" {
		err = errors.New("missing required routeID parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/where/schedule-for-route/%s.json", routeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type ScheduleForRouteGetResponse struct {
	Data ScheduleForRouteGetResponseData `json:"data" api:"required"`
	JSON scheduleForRouteGetResponseJSON `json:"-"`
	shared.ResponseWrapper
}

// scheduleForRouteGetResponseJSON contains the JSON metadata for the struct
// [ScheduleForRouteGetResponse]
type scheduleForRouteGetResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScheduleForRouteGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scheduleForRouteGetResponseJSON) RawJSON() string {
	return r.raw
}

type ScheduleForRouteGetResponseData struct {
	Entry ScheduleForRouteGetResponseDataEntry `json:"entry" api:"required"`
	JSON  scheduleForRouteGetResponseDataJSON  `json:"-"`
}

// scheduleForRouteGetResponseDataJSON contains the JSON metadata for the struct
// [ScheduleForRouteGetResponseData]
type scheduleForRouteGetResponseDataJSON struct {
	Entry       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScheduleForRouteGetResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scheduleForRouteGetResponseDataJSON) RawJSON() string {
	return r.raw
}

type ScheduleForRouteGetResponseDataEntry struct {
	RouteID           string                                                 `json:"routeId" api:"required"`
	ScheduleDate      int64                                                  `json:"scheduleDate" api:"required"`
	ServiceIDs        []string                                               `json:"serviceIds" api:"required"`
	StopTripGroupings []ScheduleForRouteGetResponseDataEntryStopTripGrouping `json:"stopTripGroupings" api:"required"`
	JSON              scheduleForRouteGetResponseDataEntryJSON               `json:"-"`
}

// scheduleForRouteGetResponseDataEntryJSON contains the JSON metadata for the
// struct [ScheduleForRouteGetResponseDataEntry]
type scheduleForRouteGetResponseDataEntryJSON struct {
	RouteID           apijson.Field
	ScheduleDate      apijson.Field
	ServiceIDs        apijson.Field
	StopTripGroupings apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ScheduleForRouteGetResponseDataEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scheduleForRouteGetResponseDataEntryJSON) RawJSON() string {
	return r.raw
}

type ScheduleForRouteGetResponseDataEntryStopTripGrouping struct {
	DirectionID        string                                                                   `json:"directionId" api:"required"`
	StopIDs            []string                                                                 `json:"stopIds" api:"required"`
	TripHeadsigns      []string                                                                 `json:"tripHeadsigns" api:"required"`
	TripIDs            []string                                                                 `json:"tripIds" api:"required"`
	TripsWithStopTimes []ScheduleForRouteGetResponseDataEntryStopTripGroupingsTripsWithStopTime `json:"tripsWithStopTimes"`
	JSON               scheduleForRouteGetResponseDataEntryStopTripGroupingJSON                 `json:"-"`
}

// scheduleForRouteGetResponseDataEntryStopTripGroupingJSON contains the JSON
// metadata for the struct [ScheduleForRouteGetResponseDataEntryStopTripGrouping]
type scheduleForRouteGetResponseDataEntryStopTripGroupingJSON struct {
	DirectionID        apijson.Field
	StopIDs            apijson.Field
	TripHeadsigns      apijson.Field
	TripIDs            apijson.Field
	TripsWithStopTimes apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ScheduleForRouteGetResponseDataEntryStopTripGrouping) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scheduleForRouteGetResponseDataEntryStopTripGroupingJSON) RawJSON() string {
	return r.raw
}

type ScheduleForRouteGetResponseDataEntryStopTripGroupingsTripsWithStopTime struct {
	StopTimes []ScheduleForRouteGetResponseDataEntryStopTripGroupingsTripsWithStopTimesStopTime `json:"stopTimes" api:"required"`
	TripID    string                                                                            `json:"tripId" api:"required"`
	JSON      scheduleForRouteGetResponseDataEntryStopTripGroupingsTripsWithStopTimeJSON        `json:"-"`
}

// scheduleForRouteGetResponseDataEntryStopTripGroupingsTripsWithStopTimeJSON
// contains the JSON metadata for the struct
// [ScheduleForRouteGetResponseDataEntryStopTripGroupingsTripsWithStopTime]
type scheduleForRouteGetResponseDataEntryStopTripGroupingsTripsWithStopTimeJSON struct {
	StopTimes   apijson.Field
	TripID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScheduleForRouteGetResponseDataEntryStopTripGroupingsTripsWithStopTime) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scheduleForRouteGetResponseDataEntryStopTripGroupingsTripsWithStopTimeJSON) RawJSON() string {
	return r.raw
}

type ScheduleForRouteGetResponseDataEntryStopTripGroupingsTripsWithStopTimesStopTime struct {
	ArrivalEnabled   bool                                                                                `json:"arrivalEnabled" api:"required"`
	ArrivalTime      int64                                                                               `json:"arrivalTime" api:"required"`
	DepartureEnabled bool                                                                                `json:"departureEnabled" api:"required"`
	DepartureTime    int64                                                                               `json:"departureTime" api:"required"`
	StopID           string                                                                              `json:"stopId" api:"required"`
	TripID           string                                                                              `json:"tripId" api:"required"`
	ServiceID        string                                                                              `json:"serviceId"`
	StopHeadsign     string                                                                              `json:"stopHeadsign"`
	JSON             scheduleForRouteGetResponseDataEntryStopTripGroupingsTripsWithStopTimesStopTimeJSON `json:"-"`
}

// scheduleForRouteGetResponseDataEntryStopTripGroupingsTripsWithStopTimesStopTimeJSON
// contains the JSON metadata for the struct
// [ScheduleForRouteGetResponseDataEntryStopTripGroupingsTripsWithStopTimesStopTime]
type scheduleForRouteGetResponseDataEntryStopTripGroupingsTripsWithStopTimesStopTimeJSON struct {
	ArrivalEnabled   apijson.Field
	ArrivalTime      apijson.Field
	DepartureEnabled apijson.Field
	DepartureTime    apijson.Field
	StopID           apijson.Field
	TripID           apijson.Field
	ServiceID        apijson.Field
	StopHeadsign     apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ScheduleForRouteGetResponseDataEntryStopTripGroupingsTripsWithStopTimesStopTime) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scheduleForRouteGetResponseDataEntryStopTripGroupingsTripsWithStopTimesStopTimeJSON) RawJSON() string {
	return r.raw
}

type ScheduleForRouteGetParams struct {
	// The date for which you want to request a schedule in the format YYYY-MM-DD
	// (optional, defaults to current date)
	Date param.Field[time.Time] `query:"date" format:"date"`
}

// URLQuery serializes [ScheduleForRouteGetParams]'s query parameters as
// `url.Values`.
func (r ScheduleForRouteGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
