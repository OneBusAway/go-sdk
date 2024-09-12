// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package onebusaway

import (
	"context"
	"net/http"
	"net/url"

	"github.com/OneBusAway/go-sdk/internal/apijson"
	"github.com/OneBusAway/go-sdk/internal/apiquery"
	"github.com/OneBusAway/go-sdk/internal/param"
	"github.com/OneBusAway/go-sdk/internal/requestconfig"
	"github.com/OneBusAway/go-sdk/option"
	"github.com/OneBusAway/go-sdk/shared"
)

// TripsForLocationService contains methods and other services that help with
// interacting with the onebusaway-sdk API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTripsForLocationService] method instead.
type TripsForLocationService struct {
	Options []option.RequestOption
}

// NewTripsForLocationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewTripsForLocationService(opts ...option.RequestOption) (r *TripsForLocationService) {
	r = &TripsForLocationService{}
	r.Options = opts
	return
}

// Retrieve trips for a given location
func (r *TripsForLocationService) List(ctx context.Context, query TripsForLocationListParams, opts ...option.RequestOption) (res *TripsForLocationListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/where/trips-for-location.json"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type TripsForLocationListResponse struct {
	Data TripsForLocationListResponseData `json:"data,required"`
	JSON tripsForLocationListResponseJSON `json:"-"`
	shared.ResponseWrapper
}

// tripsForLocationListResponseJSON contains the JSON metadata for the struct
// [TripsForLocationListResponse]
type tripsForLocationListResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TripsForLocationListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tripsForLocationListResponseJSON) RawJSON() string {
	return r.raw
}

type TripsForLocationListResponseData struct {
	// Indicates if the limit of trips has been exceeded
	LimitExceeded bool                                   `json:"limitExceeded,required"`
	List          []TripsForLocationListResponseDataList `json:"list,required"`
	References    shared.References                      `json:"references,required"`
	// Indicates if the search location is out of range
	OutOfRange bool                                 `json:"outOfRange"`
	JSON       tripsForLocationListResponseDataJSON `json:"-"`
}

// tripsForLocationListResponseDataJSON contains the JSON metadata for the struct
// [TripsForLocationListResponseData]
type tripsForLocationListResponseDataJSON struct {
	LimitExceeded apijson.Field
	List          apijson.Field
	References    apijson.Field
	OutOfRange    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *TripsForLocationListResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tripsForLocationListResponseDataJSON) RawJSON() string {
	return r.raw
}

type TripsForLocationListResponseDataList struct {
	Schedule     TripsForLocationListResponseDataListSchedule `json:"schedule,required"`
	Status       TripsForLocationListResponseDataListStatus   `json:"status,required"`
	TripID       string                                       `json:"tripId,required"`
	Frequency    string                                       `json:"frequency,nullable"`
	ServiceDate  int64                                        `json:"serviceDate"`
	SituationIDs []string                                     `json:"situationIds"`
	JSON         tripsForLocationListResponseDataListJSON     `json:"-"`
}

// tripsForLocationListResponseDataListJSON contains the JSON metadata for the
// struct [TripsForLocationListResponseDataList]
type tripsForLocationListResponseDataListJSON struct {
	Schedule     apijson.Field
	Status       apijson.Field
	TripID       apijson.Field
	Frequency    apijson.Field
	ServiceDate  apijson.Field
	SituationIDs apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *TripsForLocationListResponseDataList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tripsForLocationListResponseDataListJSON) RawJSON() string {
	return r.raw
}

type TripsForLocationListResponseDataListSchedule struct {
	NextTripID     string                                                 `json:"nextTripId,required"`
	PreviousTripID string                                                 `json:"previousTripId,required"`
	StopTimes      []TripsForLocationListResponseDataListScheduleStopTime `json:"stopTimes,required"`
	TimeZone       string                                                 `json:"timeZone,required"`
	Frequency      string                                                 `json:"frequency,nullable"`
	JSON           tripsForLocationListResponseDataListScheduleJSON       `json:"-"`
}

// tripsForLocationListResponseDataListScheduleJSON contains the JSON metadata for
// the struct [TripsForLocationListResponseDataListSchedule]
type tripsForLocationListResponseDataListScheduleJSON struct {
	NextTripID     apijson.Field
	PreviousTripID apijson.Field
	StopTimes      apijson.Field
	TimeZone       apijson.Field
	Frequency      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *TripsForLocationListResponseDataListSchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tripsForLocationListResponseDataListScheduleJSON) RawJSON() string {
	return r.raw
}

type TripsForLocationListResponseDataListScheduleStopTime struct {
	ArrivalTime         int64                                                    `json:"arrivalTime"`
	DepartureTime       int64                                                    `json:"departureTime"`
	DistanceAlongTrip   float64                                                  `json:"distanceAlongTrip"`
	HistoricalOccupancy string                                                   `json:"historicalOccupancy"`
	StopHeadsign        string                                                   `json:"stopHeadsign"`
	StopID              string                                                   `json:"stopId"`
	JSON                tripsForLocationListResponseDataListScheduleStopTimeJSON `json:"-"`
}

// tripsForLocationListResponseDataListScheduleStopTimeJSON contains the JSON
// metadata for the struct [TripsForLocationListResponseDataListScheduleStopTime]
type tripsForLocationListResponseDataListScheduleStopTimeJSON struct {
	ArrivalTime         apijson.Field
	DepartureTime       apijson.Field
	DistanceAlongTrip   apijson.Field
	HistoricalOccupancy apijson.Field
	StopHeadsign        apijson.Field
	StopID              apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *TripsForLocationListResponseDataListScheduleStopTime) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tripsForLocationListResponseDataListScheduleStopTimeJSON) RawJSON() string {
	return r.raw
}

type TripsForLocationListResponseDataListStatus struct {
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
	LastKnownLocation TripsForLocationListResponseDataListStatusLastKnownLocation `json:"lastKnownLocation"`
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
	Position TripsForLocationListResponseDataListStatusPosition `json:"position"`
	// Distance, in meters, the transit vehicle is scheduled to have progressed along
	// the active trip.
	ScheduledDistanceAlongTrip float64 `json:"scheduledDistanceAlongTrip"`
	// References to situation elements (if any) applicable to this trip.
	SituationIDs []string `json:"situationIds"`
	// ID of the transit vehicle currently serving the trip.
	VehicleID string                                         `json:"vehicleId"`
	JSON      tripsForLocationListResponseDataListStatusJSON `json:"-"`
}

// tripsForLocationListResponseDataListStatusJSON contains the JSON metadata for
// the struct [TripsForLocationListResponseDataListStatus]
type tripsForLocationListResponseDataListStatusJSON struct {
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

func (r *TripsForLocationListResponseDataListStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tripsForLocationListResponseDataListStatusJSON) RawJSON() string {
	return r.raw
}

// Last known location of the transit vehicle.
type TripsForLocationListResponseDataListStatusLastKnownLocation struct {
	// Latitude of the last known location of the transit vehicle.
	Lat float64 `json:"lat"`
	// Longitude of the last known location of the transit vehicle.
	Lon  float64                                                         `json:"lon"`
	JSON tripsForLocationListResponseDataListStatusLastKnownLocationJSON `json:"-"`
}

// tripsForLocationListResponseDataListStatusLastKnownLocationJSON contains the
// JSON metadata for the struct
// [TripsForLocationListResponseDataListStatusLastKnownLocation]
type tripsForLocationListResponseDataListStatusLastKnownLocationJSON struct {
	Lat         apijson.Field
	Lon         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TripsForLocationListResponseDataListStatusLastKnownLocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tripsForLocationListResponseDataListStatusLastKnownLocationJSON) RawJSON() string {
	return r.raw
}

// Current position of the transit vehicle.
type TripsForLocationListResponseDataListStatusPosition struct {
	// Latitude of the current position of the transit vehicle.
	Lat float64 `json:"lat"`
	// Longitude of the current position of the transit vehicle.
	Lon  float64                                                `json:"lon"`
	JSON tripsForLocationListResponseDataListStatusPositionJSON `json:"-"`
}

// tripsForLocationListResponseDataListStatusPositionJSON contains the JSON
// metadata for the struct [TripsForLocationListResponseDataListStatusPosition]
type tripsForLocationListResponseDataListStatusPositionJSON struct {
	Lat         apijson.Field
	Lon         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TripsForLocationListResponseDataListStatusPosition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tripsForLocationListResponseDataListStatusPositionJSON) RawJSON() string {
	return r.raw
}

type TripsForLocationListParams struct {
	// The latitude coordinate of the search center
	Lat param.Field[float64] `query:"lat,required"`
	// Latitude span of the search bounding box
	LatSpan param.Field[float64] `query:"latSpan,required"`
	// The longitude coordinate of the search center
	Lon param.Field[float64] `query:"lon,required"`
	// Longitude span of the search bounding box
	LonSpan param.Field[float64] `query:"lonSpan,required"`
	// Whether to include full schedule elements in the tripDetails section. Defaults
	// to false.
	IncludeSchedule param.Field[bool] `query:"includeSchedule"`
	// Whether to include full trip elements in the references section. Defaults to
	// false.
	IncludeTrip param.Field[bool] `query:"includeTrip"`
	// Specific time for the query. Defaults to the current time.
	Time param.Field[int64] `query:"time"`
}

// URLQuery serializes [TripsForLocationListParams]'s query parameters as
// `url.Values`.
func (r TripsForLocationListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
