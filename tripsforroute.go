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

// TripsForRouteService contains methods and other services that help with
// interacting with the onebusaway-sdk API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTripsForRouteService] method instead.
type TripsForRouteService struct {
	Options []option.RequestOption
}

// NewTripsForRouteService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTripsForRouteService(opts ...option.RequestOption) (r *TripsForRouteService) {
	r = &TripsForRouteService{}
	r.Options = opts
	return
}

// Search for active trips for a specific route.
func (r *TripsForRouteService) List(ctx context.Context, routeID string, query TripsForRouteListParams, opts ...option.RequestOption) (res *TripsForRouteListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if routeID == "" {
		err = errors.New("missing required routeID parameter")
		return
	}
	path := fmt.Sprintf("api/where/trips-for-route/routeID.json")
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type TripsForRouteListResponse struct {
	Data TripsForRouteListResponseData `json:"data,required"`
	JSON tripsForRouteListResponseJSON `json:"-"`
	shared.ResponseWrapper
}

// tripsForRouteListResponseJSON contains the JSON metadata for the struct
// [TripsForRouteListResponse]
type tripsForRouteListResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TripsForRouteListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tripsForRouteListResponseJSON) RawJSON() string {
	return r.raw
}

type TripsForRouteListResponseData struct {
	LimitExceeded bool                                `json:"limitExceeded,required"`
	List          []TripsForRouteListResponseDataList `json:"list,required"`
	References    shared.References                   `json:"references,required"`
	JSON          tripsForRouteListResponseDataJSON   `json:"-"`
}

// tripsForRouteListResponseDataJSON contains the JSON metadata for the struct
// [TripsForRouteListResponseData]
type tripsForRouteListResponseDataJSON struct {
	LimitExceeded apijson.Field
	List          apijson.Field
	References    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *TripsForRouteListResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tripsForRouteListResponseDataJSON) RawJSON() string {
	return r.raw
}

type TripsForRouteListResponseDataList struct {
	Schedule     TripsForRouteListResponseDataListSchedule `json:"schedule,required"`
	Status       TripsForRouteListResponseDataListStatus   `json:"status,required"`
	TripID       string                                    `json:"tripId,required"`
	Frequency    string                                    `json:"frequency,nullable"`
	ServiceDate  int64                                     `json:"serviceDate"`
	SituationIDs []string                                  `json:"situationIds"`
	JSON         tripsForRouteListResponseDataListJSON     `json:"-"`
}

// tripsForRouteListResponseDataListJSON contains the JSON metadata for the struct
// [TripsForRouteListResponseDataList]
type tripsForRouteListResponseDataListJSON struct {
	Schedule     apijson.Field
	Status       apijson.Field
	TripID       apijson.Field
	Frequency    apijson.Field
	ServiceDate  apijson.Field
	SituationIDs apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *TripsForRouteListResponseDataList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tripsForRouteListResponseDataListJSON) RawJSON() string {
	return r.raw
}

type TripsForRouteListResponseDataListSchedule struct {
	NextTripID     string                                              `json:"nextTripId,required"`
	PreviousTripID string                                              `json:"previousTripId,required"`
	StopTimes      []TripsForRouteListResponseDataListScheduleStopTime `json:"stopTimes,required"`
	TimeZone       string                                              `json:"timeZone,required"`
	Frequency      string                                              `json:"frequency,nullable"`
	JSON           tripsForRouteListResponseDataListScheduleJSON       `json:"-"`
}

// tripsForRouteListResponseDataListScheduleJSON contains the JSON metadata for the
// struct [TripsForRouteListResponseDataListSchedule]
type tripsForRouteListResponseDataListScheduleJSON struct {
	NextTripID     apijson.Field
	PreviousTripID apijson.Field
	StopTimes      apijson.Field
	TimeZone       apijson.Field
	Frequency      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *TripsForRouteListResponseDataListSchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tripsForRouteListResponseDataListScheduleJSON) RawJSON() string {
	return r.raw
}

type TripsForRouteListResponseDataListScheduleStopTime struct {
	ArrivalTime         int64                                                 `json:"arrivalTime"`
	DepartureTime       int64                                                 `json:"departureTime"`
	DistanceAlongTrip   float64                                               `json:"distanceAlongTrip"`
	HistoricalOccupancy string                                                `json:"historicalOccupancy"`
	StopHeadsign        string                                                `json:"stopHeadsign"`
	StopID              string                                                `json:"stopId"`
	JSON                tripsForRouteListResponseDataListScheduleStopTimeJSON `json:"-"`
}

// tripsForRouteListResponseDataListScheduleStopTimeJSON contains the JSON metadata
// for the struct [TripsForRouteListResponseDataListScheduleStopTime]
type tripsForRouteListResponseDataListScheduleStopTimeJSON struct {
	ArrivalTime         apijson.Field
	DepartureTime       apijson.Field
	DistanceAlongTrip   apijson.Field
	HistoricalOccupancy apijson.Field
	StopHeadsign        apijson.Field
	StopID              apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *TripsForRouteListResponseDataListScheduleStopTime) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tripsForRouteListResponseDataListScheduleStopTimeJSON) RawJSON() string {
	return r.raw
}

type TripsForRouteListResponseDataListStatus struct {
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
	LastKnownLocation TripsForRouteListResponseDataListStatusLastKnownLocation `json:"lastKnownLocation"`
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
	Position TripsForRouteListResponseDataListStatusPosition `json:"position"`
	// Distance, in meters, the transit vehicle is scheduled to have progressed along
	// the active trip.
	ScheduledDistanceAlongTrip float64 `json:"scheduledDistanceAlongTrip"`
	// References to situation elements (if any) applicable to this trip.
	SituationIDs []string `json:"situationIds"`
	// ID of the transit vehicle currently serving the trip.
	VehicleID string                                      `json:"vehicleId"`
	JSON      tripsForRouteListResponseDataListStatusJSON `json:"-"`
}

// tripsForRouteListResponseDataListStatusJSON contains the JSON metadata for the
// struct [TripsForRouteListResponseDataListStatus]
type tripsForRouteListResponseDataListStatusJSON struct {
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

func (r *TripsForRouteListResponseDataListStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tripsForRouteListResponseDataListStatusJSON) RawJSON() string {
	return r.raw
}

// Last known location of the transit vehicle.
type TripsForRouteListResponseDataListStatusLastKnownLocation struct {
	// Latitude of the last known location of the transit vehicle.
	Lat float64 `json:"lat"`
	// Longitude of the last known location of the transit vehicle.
	Lon  float64                                                      `json:"lon"`
	JSON tripsForRouteListResponseDataListStatusLastKnownLocationJSON `json:"-"`
}

// tripsForRouteListResponseDataListStatusLastKnownLocationJSON contains the JSON
// metadata for the struct
// [TripsForRouteListResponseDataListStatusLastKnownLocation]
type tripsForRouteListResponseDataListStatusLastKnownLocationJSON struct {
	Lat         apijson.Field
	Lon         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TripsForRouteListResponseDataListStatusLastKnownLocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tripsForRouteListResponseDataListStatusLastKnownLocationJSON) RawJSON() string {
	return r.raw
}

// Current position of the transit vehicle.
type TripsForRouteListResponseDataListStatusPosition struct {
	// Latitude of the current position of the transit vehicle.
	Lat float64 `json:"lat"`
	// Longitude of the current position of the transit vehicle.
	Lon  float64                                             `json:"lon"`
	JSON tripsForRouteListResponseDataListStatusPositionJSON `json:"-"`
}

// tripsForRouteListResponseDataListStatusPositionJSON contains the JSON metadata
// for the struct [TripsForRouteListResponseDataListStatusPosition]
type tripsForRouteListResponseDataListStatusPositionJSON struct {
	Lat         apijson.Field
	Lon         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TripsForRouteListResponseDataListStatusPosition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tripsForRouteListResponseDataListStatusPositionJSON) RawJSON() string {
	return r.raw
}

type TripsForRouteListParams struct {
	// Determine whether full schedule elements are included. Defaults to false.
	IncludeSchedule param.Field[bool] `query:"includeSchedule"`
	// Determine whether full tripStatus elements with real-time information are
	// included. Defaults to false.
	IncludeStatus param.Field[bool] `query:"includeStatus"`
	// Query the system at a specific time. Useful for testing.
	Time param.Field[int64] `query:"time"`
}

// URLQuery serializes [TripsForRouteListParams]'s query parameters as
// `url.Values`.
func (r TripsForRouteListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
