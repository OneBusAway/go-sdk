// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package onebusaway

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/OneBusAway/go-sdk/internal/apijson"
	"github.com/OneBusAway/go-sdk/internal/apiquery"
	"github.com/OneBusAway/go-sdk/internal/param"
	"github.com/OneBusAway/go-sdk/internal/requestconfig"
	"github.com/OneBusAway/go-sdk/option"
	"github.com/OneBusAway/go-sdk/shared"
)

// ArrivalsAndDeparturesForLocationService contains methods and other services that
// help with interacting with the onebusaway-sdk API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewArrivalsAndDeparturesForLocationService] method instead.
type ArrivalsAndDeparturesForLocationService struct {
	Options []option.RequestOption
}

// NewArrivalsAndDeparturesForLocationService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewArrivalsAndDeparturesForLocationService(opts ...option.RequestOption) (r *ArrivalsAndDeparturesForLocationService) {
	r = &ArrivalsAndDeparturesForLocationService{}
	r.Options = opts
	return
}

// Returns real-time arrival and departure data for stops within a bounding box or
// radius centered on a specific location.
func (r *ArrivalsAndDeparturesForLocationService) List(ctx context.Context, query ArrivalsAndDeparturesForLocationListParams, opts ...option.RequestOption) (res *ArrivalsAndDeparturesForLocationListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/where/arrivals-and-departures-for-location.json"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type ArrivalsAndDeparturesForLocationListResponse struct {
	Data ArrivalsAndDeparturesForLocationListResponseData `json:"data" api:"required"`
	JSON arrivalsAndDeparturesForLocationListResponseJSON `json:"-"`
	shared.ResponseWrapper
}

// arrivalsAndDeparturesForLocationListResponseJSON contains the JSON metadata for
// the struct [ArrivalsAndDeparturesForLocationListResponse]
type arrivalsAndDeparturesForLocationListResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ArrivalsAndDeparturesForLocationListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r arrivalsAndDeparturesForLocationListResponseJSON) RawJSON() string {
	return r.raw
}

type ArrivalsAndDeparturesForLocationListResponseData struct {
	Entry      ArrivalsAndDeparturesForLocationListResponseDataEntry `json:"entry" api:"required"`
	References shared.References                                     `json:"references" api:"required"`
	JSON       arrivalsAndDeparturesForLocationListResponseDataJSON  `json:"-"`
}

// arrivalsAndDeparturesForLocationListResponseDataJSON contains the JSON metadata
// for the struct [ArrivalsAndDeparturesForLocationListResponseData]
type arrivalsAndDeparturesForLocationListResponseDataJSON struct {
	Entry       apijson.Field
	References  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ArrivalsAndDeparturesForLocationListResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r arrivalsAndDeparturesForLocationListResponseDataJSON) RawJSON() string {
	return r.raw
}

type ArrivalsAndDeparturesForLocationListResponseDataEntry struct {
	ArrivalsAndDepartures []ArrivalsAndDeparturesForLocationListResponseDataEntryArrivalsAndDeparture `json:"arrivalsAndDepartures" api:"required"`
	LimitExceeded         bool                                                                        `json:"limitExceeded" api:"required"`
	NearbyStopIDs         []ArrivalsAndDeparturesForLocationListResponseDataEntryNearbyStopID         `json:"nearbyStopIds" api:"required"`
	StopIDs               []string                                                                    `json:"stopIds" api:"required"`
	SituationIDs          []string                                                                    `json:"situationIds"`
	JSON                  arrivalsAndDeparturesForLocationListResponseDataEntryJSON                   `json:"-"`
}

// arrivalsAndDeparturesForLocationListResponseDataEntryJSON contains the JSON
// metadata for the struct [ArrivalsAndDeparturesForLocationListResponseDataEntry]
type arrivalsAndDeparturesForLocationListResponseDataEntryJSON struct {
	ArrivalsAndDepartures apijson.Field
	LimitExceeded         apijson.Field
	NearbyStopIDs         apijson.Field
	StopIDs               apijson.Field
	SituationIDs          apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *ArrivalsAndDeparturesForLocationListResponseDataEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r arrivalsAndDeparturesForLocationListResponseDataEntryJSON) RawJSON() string {
	return r.raw
}

type ArrivalsAndDeparturesForLocationListResponseDataEntryArrivalsAndDeparture struct {
	// Indicates if riders can arrive on this transit vehicle.
	ArrivalEnabled bool `json:"arrivalEnabled" api:"required"`
	// Index of this arrival’s trip into the sequence of trips for the active block.
	BlockTripSequence int64 `json:"blockTripSequence" api:"required"`
	// Indicates if riders can depart from this transit vehicle.
	DepartureEnabled bool `json:"departureEnabled" api:"required"`
	// Number of stops between the arriving transit vehicle and the current stop
	// (excluding the current stop).
	NumberOfStopsAway int64 `json:"numberOfStopsAway" api:"required"`
	// Predicted arrival time, in milliseconds since Unix epoch (zero if no real-time
	// available).
	PredictedArrivalTime int64 `json:"predictedArrivalTime" api:"required"`
	// Predicted departure time, in milliseconds since Unix epoch (zero if no real-time
	// available).
	PredictedDepartureTime int64 `json:"predictedDepartureTime" api:"required"`
	// The ID of the route for the arriving vehicle.
	RouteID string `json:"routeId" api:"required"`
	// Scheduled arrival time, in milliseconds since Unix epoch.
	ScheduledArrivalTime int64 `json:"scheduledArrivalTime" api:"required"`
	// Scheduled departure time, in milliseconds since Unix epoch.
	ScheduledDepartureTime int64 `json:"scheduledDepartureTime" api:"required"`
	// Time, in milliseconds since the Unix epoch, of midnight for the start of the
	// service date for the trip.
	ServiceDate int64 `json:"serviceDate" api:"required"`
	// The ID of the stop the vehicle is arriving at.
	StopID string `json:"stopId" api:"required"`
	// Index of the stop into the sequence of stops that make up the trip for this
	// arrival.
	StopSequence int64 `json:"stopSequence" api:"required"`
	// Total number of stops visited on the trip for this arrival.
	TotalStopsInTrip int64 `json:"totalStopsInTrip" api:"required"`
	// Optional trip headsign that potentially overrides the trip headsign in the
	// referenced trip element.
	TripHeadsign string `json:"tripHeadsign" api:"required"`
	// The ID of the trip for the arriving vehicle.
	TripID string `json:"tripId" api:"required"`
	// ID of the transit vehicle serving this trip.
	VehicleID string `json:"vehicleId" api:"required"`
	// The actual track information of the arriving transit vehicle.
	ActualTrack string `json:"actualTrack"`
	// Distance of the arriving transit vehicle from the stop, in meters.
	DistanceFromStop float64 `json:"distanceFromStop"`
	// Information about frequency-based scheduling, if applicable to the trip.
	Frequency string `json:"frequency"`
	// Historical occupancy information of the transit vehicle.
	HistoricalOccupancy string `json:"historicalOccupancy"`
	// Timestamp of the last update time for this arrival.
	LastUpdateTime int64 `json:"lastUpdateTime"`
	// Current occupancy status of the transit vehicle.
	OccupancyStatus string `json:"occupancyStatus"`
	// Indicates if real-time arrival info is available for this trip.
	Predicted bool `json:"predicted"`
	// Interval for predicted arrival time, if available.
	PredictedArrivalInterval string `json:"predictedArrivalInterval"`
	// Interval for predicted departure time, if available.
	PredictedDepartureInterval string `json:"predictedDepartureInterval"`
	// Predicted occupancy status of the transit vehicle.
	PredictedOccupancy string `json:"predictedOccupancy"`
	// Optional route long name that potentially overrides the route long name in the
	// referenced route element.
	RouteLongName string `json:"routeLongName"`
	// Optional route short name that potentially overrides the route short name in the
	// referenced route element.
	RouteShortName string `json:"routeShortName"`
	// Interval for scheduled arrival time.
	ScheduledArrivalInterval string `json:"scheduledArrivalInterval"`
	// Interval for scheduled departure time.
	ScheduledDepartureInterval string `json:"scheduledDepartureInterval"`
	// Scheduled track information of the arriving transit vehicle.
	ScheduledTrack string `json:"scheduledTrack"`
	// References to situation elements (if any) applicable to this arrival.
	SituationIDs []string `json:"situationIds"`
	// Current status of the arrival.
	Status string `json:"status"`
	// Trip-specific status for the arriving transit vehicle.
	TripStatus ArrivalsAndDeparturesForLocationListResponseDataEntryArrivalsAndDeparturesTripStatus `json:"tripStatus"`
	JSON       arrivalsAndDeparturesForLocationListResponseDataEntryArrivalsAndDepartureJSON        `json:"-"`
}

// arrivalsAndDeparturesForLocationListResponseDataEntryArrivalsAndDepartureJSON
// contains the JSON metadata for the struct
// [ArrivalsAndDeparturesForLocationListResponseDataEntryArrivalsAndDeparture]
type arrivalsAndDeparturesForLocationListResponseDataEntryArrivalsAndDepartureJSON struct {
	ArrivalEnabled             apijson.Field
	BlockTripSequence          apijson.Field
	DepartureEnabled           apijson.Field
	NumberOfStopsAway          apijson.Field
	PredictedArrivalTime       apijson.Field
	PredictedDepartureTime     apijson.Field
	RouteID                    apijson.Field
	ScheduledArrivalTime       apijson.Field
	ScheduledDepartureTime     apijson.Field
	ServiceDate                apijson.Field
	StopID                     apijson.Field
	StopSequence               apijson.Field
	TotalStopsInTrip           apijson.Field
	TripHeadsign               apijson.Field
	TripID                     apijson.Field
	VehicleID                  apijson.Field
	ActualTrack                apijson.Field
	DistanceFromStop           apijson.Field
	Frequency                  apijson.Field
	HistoricalOccupancy        apijson.Field
	LastUpdateTime             apijson.Field
	OccupancyStatus            apijson.Field
	Predicted                  apijson.Field
	PredictedArrivalInterval   apijson.Field
	PredictedDepartureInterval apijson.Field
	PredictedOccupancy         apijson.Field
	RouteLongName              apijson.Field
	RouteShortName             apijson.Field
	ScheduledArrivalInterval   apijson.Field
	ScheduledDepartureInterval apijson.Field
	ScheduledTrack             apijson.Field
	SituationIDs               apijson.Field
	Status                     apijson.Field
	TripStatus                 apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *ArrivalsAndDeparturesForLocationListResponseDataEntryArrivalsAndDeparture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r arrivalsAndDeparturesForLocationListResponseDataEntryArrivalsAndDepartureJSON) RawJSON() string {
	return r.raw
}

// Trip-specific status for the arriving transit vehicle.
type ArrivalsAndDeparturesForLocationListResponseDataEntryArrivalsAndDeparturesTripStatus struct {
	// Trip ID of the trip the vehicle is actively serving.
	ActiveTripID string `json:"activeTripId" api:"required"`
	// Index of the active trip into the sequence of trips for the active block.
	BlockTripSequence int64 `json:"blockTripSequence" api:"required"`
	// ID of the closest stop to the current location of the transit vehicle.
	ClosestStop string `json:"closestStop" api:"required"`
	// Distance, in meters, the transit vehicle has progressed along the active trip.
	DistanceAlongTrip float64 `json:"distanceAlongTrip" api:"required"`
	// Last known distance along the trip received in real-time from the transit
	// vehicle.
	LastKnownDistanceAlongTrip float64 `json:"lastKnownDistanceAlongTrip" api:"required"`
	// Timestamp of the last known real-time location update from the transit vehicle.
	LastLocationUpdateTime int64 `json:"lastLocationUpdateTime" api:"required"`
	// Timestamp of the last known real-time update from the transit vehicle.
	LastUpdateTime int64 `json:"lastUpdateTime" api:"required"`
	// Capacity of the transit vehicle in terms of occupancy.
	OccupancyCapacity int64 `json:"occupancyCapacity" api:"required"`
	// Current count of occupants in the transit vehicle.
	OccupancyCount int64 `json:"occupancyCount" api:"required"`
	// Current occupancy status of the transit vehicle.
	OccupancyStatus string `json:"occupancyStatus" api:"required"`
	// Current journey phase of the trip.
	Phase string `json:"phase" api:"required"`
	// Indicates if real-time arrival info is available for this trip.
	Predicted bool `json:"predicted" api:"required"`
	// Deviation from the schedule in seconds (positive for late, negative for early).
	ScheduleDeviation int64 `json:"scheduleDeviation" api:"required"`
	// Time, in milliseconds since the Unix epoch, of midnight for the start of the
	// service date for the trip.
	ServiceDate int64 `json:"serviceDate" api:"required"`
	// Current status modifiers for the trip.
	Status string `json:"status" api:"required"`
	// Total length of the trip, in meters.
	TotalDistanceAlongTrip float64 `json:"totalDistanceAlongTrip" api:"required"`
	// Time offset from the closest stop to the current position of the transit vehicle
	// (in seconds).
	ClosestStopTimeOffset int64 `json:"closestStopTimeOffset"`
	// Information about frequency-based scheduling, if applicable to the trip.
	Frequency string `json:"frequency"`
	// Last known location of the transit vehicle (optional).
	LastKnownLocation ArrivalsAndDeparturesForLocationListResponseDataEntryArrivalsAndDeparturesTripStatusLastKnownLocation `json:"lastKnownLocation" api:"nullable"`
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
	Position ArrivalsAndDeparturesForLocationListResponseDataEntryArrivalsAndDeparturesTripStatusPosition `json:"position"`
	// Distance, in meters, the transit vehicle is scheduled to have progressed along
	// the active trip.
	ScheduledDistanceAlongTrip float64 `json:"scheduledDistanceAlongTrip"`
	// References to situation elements (if any) applicable to this trip.
	SituationIDs []string `json:"situationIds"`
	// ID of the transit vehicle currently serving the trip.
	VehicleID string                                                                                   `json:"vehicleId"`
	JSON      arrivalsAndDeparturesForLocationListResponseDataEntryArrivalsAndDeparturesTripStatusJSON `json:"-"`
}

// arrivalsAndDeparturesForLocationListResponseDataEntryArrivalsAndDeparturesTripStatusJSON
// contains the JSON metadata for the struct
// [ArrivalsAndDeparturesForLocationListResponseDataEntryArrivalsAndDeparturesTripStatus]
type arrivalsAndDeparturesForLocationListResponseDataEntryArrivalsAndDeparturesTripStatusJSON struct {
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

func (r *ArrivalsAndDeparturesForLocationListResponseDataEntryArrivalsAndDeparturesTripStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r arrivalsAndDeparturesForLocationListResponseDataEntryArrivalsAndDeparturesTripStatusJSON) RawJSON() string {
	return r.raw
}

// Last known location of the transit vehicle (optional).
type ArrivalsAndDeparturesForLocationListResponseDataEntryArrivalsAndDeparturesTripStatusLastKnownLocation struct {
	// Latitude of the last known location of the transit vehicle.
	Lat float64 `json:"lat"`
	// Longitude of the last known location of the transit vehicle.
	Lon  float64                                                                                                   `json:"lon"`
	JSON arrivalsAndDeparturesForLocationListResponseDataEntryArrivalsAndDeparturesTripStatusLastKnownLocationJSON `json:"-"`
}

// arrivalsAndDeparturesForLocationListResponseDataEntryArrivalsAndDeparturesTripStatusLastKnownLocationJSON
// contains the JSON metadata for the struct
// [ArrivalsAndDeparturesForLocationListResponseDataEntryArrivalsAndDeparturesTripStatusLastKnownLocation]
type arrivalsAndDeparturesForLocationListResponseDataEntryArrivalsAndDeparturesTripStatusLastKnownLocationJSON struct {
	Lat         apijson.Field
	Lon         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ArrivalsAndDeparturesForLocationListResponseDataEntryArrivalsAndDeparturesTripStatusLastKnownLocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r arrivalsAndDeparturesForLocationListResponseDataEntryArrivalsAndDeparturesTripStatusLastKnownLocationJSON) RawJSON() string {
	return r.raw
}

// Current position of the transit vehicle.
type ArrivalsAndDeparturesForLocationListResponseDataEntryArrivalsAndDeparturesTripStatusPosition struct {
	// Latitude of the current position of the transit vehicle.
	Lat float64 `json:"lat"`
	// Longitude of the current position of the transit vehicle.
	Lon  float64                                                                                          `json:"lon"`
	JSON arrivalsAndDeparturesForLocationListResponseDataEntryArrivalsAndDeparturesTripStatusPositionJSON `json:"-"`
}

// arrivalsAndDeparturesForLocationListResponseDataEntryArrivalsAndDeparturesTripStatusPositionJSON
// contains the JSON metadata for the struct
// [ArrivalsAndDeparturesForLocationListResponseDataEntryArrivalsAndDeparturesTripStatusPosition]
type arrivalsAndDeparturesForLocationListResponseDataEntryArrivalsAndDeparturesTripStatusPositionJSON struct {
	Lat         apijson.Field
	Lon         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ArrivalsAndDeparturesForLocationListResponseDataEntryArrivalsAndDeparturesTripStatusPosition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r arrivalsAndDeparturesForLocationListResponseDataEntryArrivalsAndDeparturesTripStatusPositionJSON) RawJSON() string {
	return r.raw
}

type ArrivalsAndDeparturesForLocationListResponseDataEntryNearbyStopID struct {
	DistanceFromQuery float64                                                               `json:"distanceFromQuery"`
	StopID            string                                                                `json:"stopId"`
	JSON              arrivalsAndDeparturesForLocationListResponseDataEntryNearbyStopIDJSON `json:"-"`
}

// arrivalsAndDeparturesForLocationListResponseDataEntryNearbyStopIDJSON contains
// the JSON metadata for the struct
// [ArrivalsAndDeparturesForLocationListResponseDataEntryNearbyStopID]
type arrivalsAndDeparturesForLocationListResponseDataEntryNearbyStopIDJSON struct {
	DistanceFromQuery apijson.Field
	StopID            apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ArrivalsAndDeparturesForLocationListResponseDataEntryNearbyStopID) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r arrivalsAndDeparturesForLocationListResponseDataEntryNearbyStopIDJSON) RawJSON() string {
	return r.raw
}

type ArrivalsAndDeparturesForLocationListParams struct {
	// The latitude coordinate of the search center.
	Lat param.Field[float64] `query:"lat" api:"required"`
	// The longitude coordinate of the search center.
	Lon param.Field[float64] `query:"lon" api:"required"`
	// If true, returns a 404 Not Found error instead of an empty result.
	EmptyReturnsNotFound param.Field[bool] `query:"emptyReturnsNotFound"`
	// Sets the latitude limits of the search bounding box.
	LatSpan param.Field[float64] `query:"latSpan"`
	// Sets the longitude limits of the search bounding box.
	LonSpan param.Field[float64] `query:"lonSpan"`
	// The max size of the list of nearby stops and arrivals to return. Defaults to
	// 250, max 1000.
	MaxCount param.Field[int64] `query:"maxCount"`
	// Include arrivals and departures this many minutes after the query time.
	MinutesAfter param.Field[int64] `query:"minutesAfter"`
	// Include arrivals and departures this many minutes before the query time.
	MinutesBefore param.Field[int64] `query:"minutesBefore"`
	// The search radius in meters.
	Radius param.Field[float64] `query:"radius"`
	// Optional list of GTFS routeTypes to filter by (comma delimited) e.g. "1,2,3".
	RouteType param.Field[string] `query:"routeType"`
	// By default, returns the status right now. Can be queried at a specific time
	// (milliseconds since epoch) for testing.
	Time param.Field[int64] `query:"time"`
}

// URLQuery serializes [ArrivalsAndDeparturesForLocationListParams]'s query
// parameters as `url.Values`.
func (r ArrivalsAndDeparturesForLocationListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
