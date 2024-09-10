// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package onebusaway

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/open-transit-go/internal/apijson"
	"github.com/stainless-sdks/open-transit-go/internal/apiquery"
	"github.com/stainless-sdks/open-transit-go/internal/param"
	"github.com/stainless-sdks/open-transit-go/internal/requestconfig"
	"github.com/stainless-sdks/open-transit-go/option"
	"github.com/stainless-sdks/open-transit-go/shared"
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
	opts = append(r.Options[:], opts...)
	if routeID == "" {
		err = errors.New("missing required routeID parameter")
		return
	}
	path := fmt.Sprintf("api/where/schedule-for-route/routeID.json")
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type ScheduleForRouteGetResponse struct {
	Data ScheduleForRouteGetResponseData `json:"data,required"`
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
	Entry ScheduleForRouteGetResponseDataEntry `json:"entry,required"`
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
	RouteID           string                                                 `json:"routeId,required"`
	ScheduleDate      int64                                                  `json:"scheduleDate,required"`
	ServiceIDs        []string                                               `json:"serviceIds,required"`
	Stops             []ScheduleForRouteGetResponseDataEntryStop             `json:"stops,required"`
	StopTripGroupings []ScheduleForRouteGetResponseDataEntryStopTripGrouping `json:"stopTripGroupings,required"`
	Trips             []ScheduleForRouteGetResponseDataEntryTrip             `json:"trips,required"`
	JSON              scheduleForRouteGetResponseDataEntryJSON               `json:"-"`
}

// scheduleForRouteGetResponseDataEntryJSON contains the JSON metadata for the
// struct [ScheduleForRouteGetResponseDataEntry]
type scheduleForRouteGetResponseDataEntryJSON struct {
	RouteID           apijson.Field
	ScheduleDate      apijson.Field
	ServiceIDs        apijson.Field
	Stops             apijson.Field
	StopTripGroupings apijson.Field
	Trips             apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ScheduleForRouteGetResponseDataEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scheduleForRouteGetResponseDataEntryJSON) RawJSON() string {
	return r.raw
}

type ScheduleForRouteGetResponseDataEntryStop struct {
	ID                 string                                       `json:"id,required"`
	Code               string                                       `json:"code,required"`
	Lat                float64                                      `json:"lat,required"`
	Lon                float64                                      `json:"lon,required"`
	Name               string                                       `json:"name,required"`
	Direction          string                                       `json:"direction"`
	LocationType       int64                                        `json:"locationType"`
	Parent             string                                       `json:"parent"`
	RouteIDs           []string                                     `json:"routeIds"`
	StaticRouteIDs     []string                                     `json:"staticRouteIds"`
	WheelchairBoarding string                                       `json:"wheelchairBoarding"`
	JSON               scheduleForRouteGetResponseDataEntryStopJSON `json:"-"`
}

// scheduleForRouteGetResponseDataEntryStopJSON contains the JSON metadata for the
// struct [ScheduleForRouteGetResponseDataEntryStop]
type scheduleForRouteGetResponseDataEntryStopJSON struct {
	ID                 apijson.Field
	Code               apijson.Field
	Lat                apijson.Field
	Lon                apijson.Field
	Name               apijson.Field
	Direction          apijson.Field
	LocationType       apijson.Field
	Parent             apijson.Field
	RouteIDs           apijson.Field
	StaticRouteIDs     apijson.Field
	WheelchairBoarding apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ScheduleForRouteGetResponseDataEntryStop) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scheduleForRouteGetResponseDataEntryStopJSON) RawJSON() string {
	return r.raw
}

type ScheduleForRouteGetResponseDataEntryStopTripGrouping struct {
	DirectionID        string                                                                   `json:"directionId,required"`
	StopIDs            []string                                                                 `json:"stopIds,required"`
	TripHeadsigns      []string                                                                 `json:"tripHeadsigns,required"`
	TripIDs            []string                                                                 `json:"tripIds,required"`
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
	StopTimes []ScheduleForRouteGetResponseDataEntryStopTripGroupingsTripsWithStopTimesStopTime `json:"stopTimes,required"`
	TripID    string                                                                            `json:"tripId,required"`
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
	ArrivalEnabled   bool                                                                                `json:"arrivalEnabled,required"`
	ArrivalTime      int64                                                                               `json:"arrivalTime,required"`
	DepartureEnabled bool                                                                                `json:"departureEnabled,required"`
	DepartureTime    int64                                                                               `json:"departureTime,required"`
	StopID           string                                                                              `json:"stopId,required"`
	TripID           string                                                                              `json:"tripId,required"`
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

type ScheduleForRouteGetResponseDataEntryTrip struct {
	ID             string                                       `json:"id,required"`
	RouteID        string                                       `json:"routeId,required"`
	ServiceID      string                                       `json:"serviceId,required"`
	BlockID        string                                       `json:"blockId"`
	DirectionID    string                                       `json:"directionId"`
	PeakOffpeak    int64                                        `json:"peakOffpeak"`
	RouteShortName string                                       `json:"routeShortName"`
	ShapeID        string                                       `json:"shapeId"`
	TimeZone       string                                       `json:"timeZone"`
	TripHeadsign   string                                       `json:"tripHeadsign"`
	TripShortName  string                                       `json:"tripShortName"`
	JSON           scheduleForRouteGetResponseDataEntryTripJSON `json:"-"`
}

// scheduleForRouteGetResponseDataEntryTripJSON contains the JSON metadata for the
// struct [ScheduleForRouteGetResponseDataEntryTrip]
type scheduleForRouteGetResponseDataEntryTripJSON struct {
	ID             apijson.Field
	RouteID        apijson.Field
	ServiceID      apijson.Field
	BlockID        apijson.Field
	DirectionID    apijson.Field
	PeakOffpeak    apijson.Field
	RouteShortName apijson.Field
	ShapeID        apijson.Field
	TimeZone       apijson.Field
	TripHeadsign   apijson.Field
	TripShortName  apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ScheduleForRouteGetResponseDataEntryTrip) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scheduleForRouteGetResponseDataEntryTripJSON) RawJSON() string {
	return r.raw
}

type ScheduleForRouteGetParams struct {
	// The date for which you want to request a schedule in the format YYYY-MM-DD
	// (optional, defaults to current date)
	Date param.Field[string] `query:"date"`
}

// URLQuery serializes [ScheduleForRouteGetParams]'s query parameters as
// `url.Values`.
func (r ScheduleForRouteGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
