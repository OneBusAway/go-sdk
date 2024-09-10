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

// TripDetailService contains methods and other services that help with interacting
// with the onebusaway-sdk API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTripDetailService] method instead.
type TripDetailService struct {
	Options []option.RequestOption
}

// NewTripDetailService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTripDetailService(opts ...option.RequestOption) (r *TripDetailService) {
	r = &TripDetailService{}
	r.Options = opts
	return
}

// Retrieve Trip Details
func (r *TripDetailService) Get(ctx context.Context, tripID string, query TripDetailGetParams, opts ...option.RequestOption) (res *TripDetailGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if tripID == "" {
		err = errors.New("missing required tripID parameter")
		return
	}
	path := fmt.Sprintf("api/where/trip-details/tripID.json")
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type TripDetailGetResponse struct {
	Data TripDetailGetResponseData `json:"data,required"`
	JSON tripDetailGetResponseJSON `json:"-"`
	shared.ResponseWrapper
}

// tripDetailGetResponseJSON contains the JSON metadata for the struct
// [TripDetailGetResponse]
type tripDetailGetResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TripDetailGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tripDetailGetResponseJSON) RawJSON() string {
	return r.raw
}

type TripDetailGetResponseData struct {
	Entry      TripDetailGetResponseDataEntry `json:"entry,required"`
	References shared.References              `json:"references,required"`
	JSON       tripDetailGetResponseDataJSON  `json:"-"`
}

// tripDetailGetResponseDataJSON contains the JSON metadata for the struct
// [TripDetailGetResponseData]
type tripDetailGetResponseDataJSON struct {
	Entry       apijson.Field
	References  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TripDetailGetResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tripDetailGetResponseDataJSON) RawJSON() string {
	return r.raw
}

type TripDetailGetResponseDataEntry struct {
	TripID       string                                 `json:"tripId,required"`
	Frequency    string                                 `json:"frequency,nullable"`
	Schedule     TripDetailGetResponseDataEntrySchedule `json:"schedule"`
	ServiceDate  int64                                  `json:"serviceDate"`
	SituationIDs []string                               `json:"situationIds"`
	Status       TripDetailGetResponseDataEntryStatus   `json:"status"`
	JSON         tripDetailGetResponseDataEntryJSON     `json:"-"`
}

// tripDetailGetResponseDataEntryJSON contains the JSON metadata for the struct
// [TripDetailGetResponseDataEntry]
type tripDetailGetResponseDataEntryJSON struct {
	TripID       apijson.Field
	Frequency    apijson.Field
	Schedule     apijson.Field
	ServiceDate  apijson.Field
	SituationIDs apijson.Field
	Status       apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *TripDetailGetResponseDataEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tripDetailGetResponseDataEntryJSON) RawJSON() string {
	return r.raw
}

type TripDetailGetResponseDataEntrySchedule struct {
	NextTripID     string                                           `json:"nextTripId,required"`
	PreviousTripID string                                           `json:"previousTripId,required"`
	StopTimes      []TripDetailGetResponseDataEntryScheduleStopTime `json:"stopTimes,required"`
	TimeZone       string                                           `json:"timeZone,required"`
	Frequency      string                                           `json:"frequency,nullable"`
	JSON           tripDetailGetResponseDataEntryScheduleJSON       `json:"-"`
}

// tripDetailGetResponseDataEntryScheduleJSON contains the JSON metadata for the
// struct [TripDetailGetResponseDataEntrySchedule]
type tripDetailGetResponseDataEntryScheduleJSON struct {
	NextTripID     apijson.Field
	PreviousTripID apijson.Field
	StopTimes      apijson.Field
	TimeZone       apijson.Field
	Frequency      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *TripDetailGetResponseDataEntrySchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tripDetailGetResponseDataEntryScheduleJSON) RawJSON() string {
	return r.raw
}

type TripDetailGetResponseDataEntryScheduleStopTime struct {
	ArrivalTime         int64                                              `json:"arrivalTime"`
	DepartureTime       int64                                              `json:"departureTime"`
	DistanceAlongTrip   float64                                            `json:"distanceAlongTrip"`
	HistoricalOccupancy string                                             `json:"historicalOccupancy"`
	StopHeadsign        string                                             `json:"stopHeadsign"`
	StopID              string                                             `json:"stopId"`
	JSON                tripDetailGetResponseDataEntryScheduleStopTimeJSON `json:"-"`
}

// tripDetailGetResponseDataEntryScheduleStopTimeJSON contains the JSON metadata
// for the struct [TripDetailGetResponseDataEntryScheduleStopTime]
type tripDetailGetResponseDataEntryScheduleStopTimeJSON struct {
	ArrivalTime         apijson.Field
	DepartureTime       apijson.Field
	DistanceAlongTrip   apijson.Field
	HistoricalOccupancy apijson.Field
	StopHeadsign        apijson.Field
	StopID              apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *TripDetailGetResponseDataEntryScheduleStopTime) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tripDetailGetResponseDataEntryScheduleStopTimeJSON) RawJSON() string {
	return r.raw
}

type TripDetailGetResponseDataEntryStatus struct {
	// Trip ID of the trip the vehicle is actively serving.
	ActiveTripID string `json:"activeTripId,required"`
	// Index of the active trip into the sequence of trips for the active block.
	BlockTripSequence int64 `json:"blockTripSequence,required"`
	// ID of the closest stop to the current location of the transit vehicle.
	ClosestStop string `json:"closestStop,required"`
	// Distance, in meters, the transit vehicle has progressed along the active trip.
	DistanceAlongTrip float64 `json:"distanceAlongTrip,required"`
	// Last known distance along the trip received in real-time from the transit
	// vehicle.
	LastKnownDistanceAlongTrip float64 `json:"lastKnownDistanceAlongTrip,required"`
	// Timestamp of the last known real-time location update from the transit vehicle.
	LastLocationUpdateTime int64 `json:"lastLocationUpdateTime,required"`
	// Timestamp of the last known real-time update from the transit vehicle.
	LastUpdateTime int64 `json:"lastUpdateTime,required"`
	// Capacity of the transit vehicle in terms of occupancy.
	OccupancyCapacity int64 `json:"occupancyCapacity,required"`
	// Current count of occupants in the transit vehicle.
	OccupancyCount int64 `json:"occupancyCount,required"`
	// Current occupancy status of the transit vehicle.
	OccupancyStatus string `json:"occupancyStatus,required"`
	// Current journey phase of the trip.
	Phase string `json:"phase,required"`
	// Indicates if real-time arrival info is available for this trip.
	Predicted bool `json:"predicted,required"`
	// Deviation from the schedule in seconds (positive for late, negative for early).
	ScheduleDeviation int64 `json:"scheduleDeviation,required"`
	// Time, in milliseconds since the Unix epoch, of midnight for the start of the
	// service date for the trip.
	ServiceDate int64 `json:"serviceDate,required"`
	// Current status modifiers for the trip.
	Status string `json:"status,required"`
	// Total length of the trip, in meters.
	TotalDistanceAlongTrip float64 `json:"totalDistanceAlongTrip,required"`
	// Time offset from the closest stop to the current position of the transit vehicle
	// (in seconds).
	ClosestStopTimeOffset int64 `json:"closestStopTimeOffset"`
	// Information about frequency-based scheduling, if applicable to the trip.
	Frequency string `json:"frequency"`
	// Last known location of the transit vehicle.
	LastKnownLocation TripDetailGetResponseDataEntryStatusLastKnownLocation `json:"lastKnownLocation"`
	// Last known orientation value received in real-time from the transit vehicle.
	LastKnownOrientation float64 `json:"lastKnownOrientation"`
	// ID of the next stop the transit vehicle is scheduled to arrive at.
	NextStop string `json:"nextStop"`
	// Time offset from the next stop to the current position of the transit vehicle
	// (in seconds).
	NextStopTimeOffset int64 `json:"nextStopTimeOffset"`
	// Orientation of the transit vehicle, represented as an angle in degrees.
	Orientation float64 `json:"orientation"`
	// Current position of the transit vehicle.
	Position TripDetailGetResponseDataEntryStatusPosition `json:"position"`
	// Distance, in meters, the transit vehicle is scheduled to have progressed along
	// the active trip.
	ScheduledDistanceAlongTrip float64 `json:"scheduledDistanceAlongTrip"`
	// References to situation elements (if any) applicable to this trip.
	SituationIDs []string `json:"situationIds"`
	// ID of the transit vehicle currently serving the trip.
	VehicleID string                                   `json:"vehicleId"`
	JSON      tripDetailGetResponseDataEntryStatusJSON `json:"-"`
}

// tripDetailGetResponseDataEntryStatusJSON contains the JSON metadata for the
// struct [TripDetailGetResponseDataEntryStatus]
type tripDetailGetResponseDataEntryStatusJSON struct {
	ActiveTripID               apijson.Field
	BlockTripSequence          apijson.Field
	ClosestStop                apijson.Field
	DistanceAlongTrip          apijson.Field
	LastKnownDistanceAlongTrip apijson.Field
	LastLocationUpdateTime     apijson.Field
	LastUpdateTime             apijson.Field
	OccupancyCapacity          apijson.Field
	OccupancyCount             apijson.Field
	OccupancyStatus            apijson.Field
	Phase                      apijson.Field
	Predicted                  apijson.Field
	ScheduleDeviation          apijson.Field
	ServiceDate                apijson.Field
	Status                     apijson.Field
	TotalDistanceAlongTrip     apijson.Field
	ClosestStopTimeOffset      apijson.Field
	Frequency                  apijson.Field
	LastKnownLocation          apijson.Field
	LastKnownOrientation       apijson.Field
	NextStop                   apijson.Field
	NextStopTimeOffset         apijson.Field
	Orientation                apijson.Field
	Position                   apijson.Field
	ScheduledDistanceAlongTrip apijson.Field
	SituationIDs               apijson.Field
	VehicleID                  apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *TripDetailGetResponseDataEntryStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tripDetailGetResponseDataEntryStatusJSON) RawJSON() string {
	return r.raw
}

// Last known location of the transit vehicle.
type TripDetailGetResponseDataEntryStatusLastKnownLocation struct {
	// Latitude of the last known location of the transit vehicle.
	Lat float64 `json:"lat"`
	// Longitude of the last known location of the transit vehicle.
	Lon  float64                                                   `json:"lon"`
	JSON tripDetailGetResponseDataEntryStatusLastKnownLocationJSON `json:"-"`
}

// tripDetailGetResponseDataEntryStatusLastKnownLocationJSON contains the JSON
// metadata for the struct [TripDetailGetResponseDataEntryStatusLastKnownLocation]
type tripDetailGetResponseDataEntryStatusLastKnownLocationJSON struct {
	Lat         apijson.Field
	Lon         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TripDetailGetResponseDataEntryStatusLastKnownLocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tripDetailGetResponseDataEntryStatusLastKnownLocationJSON) RawJSON() string {
	return r.raw
}

// Current position of the transit vehicle.
type TripDetailGetResponseDataEntryStatusPosition struct {
	// Latitude of the current position of the transit vehicle.
	Lat float64 `json:"lat"`
	// Longitude of the current position of the transit vehicle.
	Lon  float64                                          `json:"lon"`
	JSON tripDetailGetResponseDataEntryStatusPositionJSON `json:"-"`
}

// tripDetailGetResponseDataEntryStatusPositionJSON contains the JSON metadata for
// the struct [TripDetailGetResponseDataEntryStatusPosition]
type tripDetailGetResponseDataEntryStatusPositionJSON struct {
	Lat         apijson.Field
	Lon         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TripDetailGetResponseDataEntryStatusPosition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tripDetailGetResponseDataEntryStatusPositionJSON) RawJSON() string {
	return r.raw
}

type TripDetailGetParams struct {
	// Whether to include the full schedule element in the tripDetails section
	// (defaults to true).
	IncludeSchedule param.Field[bool] `query:"includeSchedule"`
	// Whether to include the full status element in the tripDetails section (defaults
	// to true).
	IncludeStatus param.Field[bool] `query:"includeStatus"`
	// Whether to include the full trip element in the references section (defaults to
	// true).
	IncludeTrip param.Field[bool] `query:"includeTrip"`
	// Service date for the trip as Unix time in milliseconds (optional).
	ServiceDate param.Field[int64] `query:"serviceDate"`
	// Time parameter to query the system at a specific time (optional).
	Time param.Field[int64] `query:"time"`
}

// URLQuery serializes [TripDetailGetParams]'s query parameters as `url.Values`.
func (r TripDetailGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
