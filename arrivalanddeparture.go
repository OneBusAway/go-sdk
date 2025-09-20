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

// ArrivalAndDepartureService contains methods and other services that help with
// interacting with the onebusaway-sdk API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewArrivalAndDepartureService] method instead.
type ArrivalAndDepartureService struct {
	Options []option.RequestOption
}

// NewArrivalAndDepartureService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewArrivalAndDepartureService(opts ...option.RequestOption) (r *ArrivalAndDepartureService) {
	r = &ArrivalAndDepartureService{}
	r.Options = opts
	return
}

// arrival-and-departure-for-stop
func (r *ArrivalAndDepartureService) Get(ctx context.Context, stopID string, query ArrivalAndDepartureGetParams, opts ...option.RequestOption) (res *ArrivalAndDepartureGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if stopID == "" {
		err = errors.New("missing required stopID parameter")
		return
	}
	path := fmt.Sprintf("api/where/arrival-and-departure-for-stop/%s.json", stopID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// arrivals-and-departures-for-stop
func (r *ArrivalAndDepartureService) List(ctx context.Context, stopID string, query ArrivalAndDepartureListParams, opts ...option.RequestOption) (res *ArrivalAndDepartureListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if stopID == "" {
		err = errors.New("missing required stopID parameter")
		return
	}
	path := fmt.Sprintf("api/where/arrivals-and-departures-for-stop/%s.json", stopID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type ArrivalAndDepartureGetResponse struct {
	Data ArrivalAndDepartureGetResponseData `json:"data,required"`
	JSON arrivalAndDepartureGetResponseJSON `json:"-"`
	shared.ResponseWrapper
}

// arrivalAndDepartureGetResponseJSON contains the JSON metadata for the struct
// [ArrivalAndDepartureGetResponse]
type arrivalAndDepartureGetResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ArrivalAndDepartureGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r arrivalAndDepartureGetResponseJSON) RawJSON() string {
	return r.raw
}

type ArrivalAndDepartureGetResponseData struct {
	Entry      ArrivalAndDepartureGetResponseDataEntry `json:"entry,required"`
	References shared.References                       `json:"references,required"`
	JSON       arrivalAndDepartureGetResponseDataJSON  `json:"-"`
}

// arrivalAndDepartureGetResponseDataJSON contains the JSON metadata for the struct
// [ArrivalAndDepartureGetResponseData]
type arrivalAndDepartureGetResponseDataJSON struct {
	Entry       apijson.Field
	References  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ArrivalAndDepartureGetResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r arrivalAndDepartureGetResponseDataJSON) RawJSON() string {
	return r.raw
}

type ArrivalAndDepartureGetResponseDataEntry struct {
	// Indicates if riders can arrive on this transit vehicle.
	ArrivalEnabled bool `json:"arrivalEnabled,required"`
	// Index of this arrival’s trip into the sequence of trips for the active block.
	BlockTripSequence int64 `json:"blockTripSequence,required"`
	// Indicates if riders can depart from this transit vehicle.
	DepartureEnabled bool `json:"departureEnabled,required"`
	// Number of stops between the arriving transit vehicle and the current stop
	// (excluding the current stop).
	NumberOfStopsAway int64 `json:"numberOfStopsAway,required"`
	// Predicted arrival time, in milliseconds since Unix epoch (zero if no real-time
	// available).
	PredictedArrivalTime int64 `json:"predictedArrivalTime,required"`
	// Predicted departure time, in milliseconds since Unix epoch (zero if no real-time
	// available).
	PredictedDepartureTime int64 `json:"predictedDepartureTime,required"`
	// The ID of the route for the arriving vehicle.
	RouteID string `json:"routeId,required"`
	// Scheduled arrival time, in milliseconds since Unix epoch.
	ScheduledArrivalTime int64 `json:"scheduledArrivalTime,required"`
	// Scheduled departure time, in milliseconds since Unix epoch.
	ScheduledDepartureTime int64 `json:"scheduledDepartureTime,required"`
	// Time, in milliseconds since the Unix epoch, of midnight for the start of the
	// service date for the trip.
	ServiceDate int64 `json:"serviceDate,required"`
	// The ID of the stop the vehicle is arriving at.
	StopID string `json:"stopId,required"`
	// Index of the stop into the sequence of stops that make up the trip for this
	// arrival.
	StopSequence int64 `json:"stopSequence,required"`
	// Total number of stops visited on the trip for this arrival.
	TotalStopsInTrip int64 `json:"totalStopsInTrip,required"`
	// Optional trip headsign that potentially overrides the trip headsign in the
	// referenced trip element.
	TripHeadsign string `json:"tripHeadsign,required"`
	// The ID of the trip for the arriving vehicle.
	TripID string `json:"tripId,required"`
	// ID of the transit vehicle serving this trip.
	VehicleID string `json:"vehicleId,required"`
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
	TripStatus ArrivalAndDepartureGetResponseDataEntryTripStatus `json:"tripStatus"`
	JSON       arrivalAndDepartureGetResponseDataEntryJSON       `json:"-"`
}

// arrivalAndDepartureGetResponseDataEntryJSON contains the JSON metadata for the
// struct [ArrivalAndDepartureGetResponseDataEntry]
type arrivalAndDepartureGetResponseDataEntryJSON struct {
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

func (r *ArrivalAndDepartureGetResponseDataEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r arrivalAndDepartureGetResponseDataEntryJSON) RawJSON() string {
	return r.raw
}

// Trip-specific status for the arriving transit vehicle.
type ArrivalAndDepartureGetResponseDataEntryTripStatus struct {
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
	LastKnownLocation ArrivalAndDepartureGetResponseDataEntryTripStatusLastKnownLocation `json:"lastKnownLocation"`
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
	Position ArrivalAndDepartureGetResponseDataEntryTripStatusPosition `json:"position"`
	// Distance, in meters, the transit vehicle is scheduled to have progressed along
	// the active trip.
	ScheduledDistanceAlongTrip float64 `json:"scheduledDistanceAlongTrip"`
	// References to situation elements (if any) applicable to this trip.
	SituationIDs []string `json:"situationIds"`
	// ID of the transit vehicle currently serving the trip.
	VehicleID string                                                `json:"vehicleId"`
	JSON      arrivalAndDepartureGetResponseDataEntryTripStatusJSON `json:"-"`
}

// arrivalAndDepartureGetResponseDataEntryTripStatusJSON contains the JSON metadata
// for the struct [ArrivalAndDepartureGetResponseDataEntryTripStatus]
type arrivalAndDepartureGetResponseDataEntryTripStatusJSON struct {
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

func (r *ArrivalAndDepartureGetResponseDataEntryTripStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r arrivalAndDepartureGetResponseDataEntryTripStatusJSON) RawJSON() string {
	return r.raw
}

// Last known location of the transit vehicle.
type ArrivalAndDepartureGetResponseDataEntryTripStatusLastKnownLocation struct {
	// Latitude of the last known location of the transit vehicle.
	Lat float64 `json:"lat"`
	// Longitude of the last known location of the transit vehicle.
	Lon  float64                                                                `json:"lon"`
	JSON arrivalAndDepartureGetResponseDataEntryTripStatusLastKnownLocationJSON `json:"-"`
}

// arrivalAndDepartureGetResponseDataEntryTripStatusLastKnownLocationJSON contains
// the JSON metadata for the struct
// [ArrivalAndDepartureGetResponseDataEntryTripStatusLastKnownLocation]
type arrivalAndDepartureGetResponseDataEntryTripStatusLastKnownLocationJSON struct {
	Lat         apijson.Field
	Lon         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ArrivalAndDepartureGetResponseDataEntryTripStatusLastKnownLocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r arrivalAndDepartureGetResponseDataEntryTripStatusLastKnownLocationJSON) RawJSON() string {
	return r.raw
}

// Current position of the transit vehicle.
type ArrivalAndDepartureGetResponseDataEntryTripStatusPosition struct {
	// Latitude of the current position of the transit vehicle.
	Lat float64 `json:"lat"`
	// Longitude of the current position of the transit vehicle.
	Lon  float64                                                       `json:"lon"`
	JSON arrivalAndDepartureGetResponseDataEntryTripStatusPositionJSON `json:"-"`
}

// arrivalAndDepartureGetResponseDataEntryTripStatusPositionJSON contains the JSON
// metadata for the struct
// [ArrivalAndDepartureGetResponseDataEntryTripStatusPosition]
type arrivalAndDepartureGetResponseDataEntryTripStatusPositionJSON struct {
	Lat         apijson.Field
	Lon         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ArrivalAndDepartureGetResponseDataEntryTripStatusPosition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r arrivalAndDepartureGetResponseDataEntryTripStatusPositionJSON) RawJSON() string {
	return r.raw
}

type ArrivalAndDepartureListResponse struct {
	Data ArrivalAndDepartureListResponseData `json:"data,required"`
	JSON arrivalAndDepartureListResponseJSON `json:"-"`
	shared.ResponseWrapper
}

// arrivalAndDepartureListResponseJSON contains the JSON metadata for the struct
// [ArrivalAndDepartureListResponse]
type arrivalAndDepartureListResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ArrivalAndDepartureListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r arrivalAndDepartureListResponseJSON) RawJSON() string {
	return r.raw
}

type ArrivalAndDepartureListResponseData struct {
	Entry      ArrivalAndDepartureListResponseDataEntry `json:"entry,required"`
	References shared.References                        `json:"references,required"`
	JSON       arrivalAndDepartureListResponseDataJSON  `json:"-"`
}

// arrivalAndDepartureListResponseDataJSON contains the JSON metadata for the
// struct [ArrivalAndDepartureListResponseData]
type arrivalAndDepartureListResponseDataJSON struct {
	Entry       apijson.Field
	References  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ArrivalAndDepartureListResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r arrivalAndDepartureListResponseDataJSON) RawJSON() string {
	return r.raw
}

type ArrivalAndDepartureListResponseDataEntry struct {
	ArrivalsAndDepartures []ArrivalAndDepartureListResponseDataEntryArrivalsAndDeparture `json:"arrivalsAndDepartures,required"`
	JSON                  arrivalAndDepartureListResponseDataEntryJSON                   `json:"-"`
}

// arrivalAndDepartureListResponseDataEntryJSON contains the JSON metadata for the
// struct [ArrivalAndDepartureListResponseDataEntry]
type arrivalAndDepartureListResponseDataEntryJSON struct {
	ArrivalsAndDepartures apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *ArrivalAndDepartureListResponseDataEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r arrivalAndDepartureListResponseDataEntryJSON) RawJSON() string {
	return r.raw
}

type ArrivalAndDepartureListResponseDataEntryArrivalsAndDeparture struct {
	// Indicates if riders can arrive on this transit vehicle.
	ArrivalEnabled bool `json:"arrivalEnabled,required"`
	// Index of this arrival’s trip into the sequence of trips for the active block.
	BlockTripSequence int64 `json:"blockTripSequence,required"`
	// Indicates if riders can depart from this transit vehicle.
	DepartureEnabled bool `json:"departureEnabled,required"`
	// Number of stops between the arriving transit vehicle and the current stop
	// (excluding the current stop).
	NumberOfStopsAway int64 `json:"numberOfStopsAway,required"`
	// Predicted arrival time, in milliseconds since Unix epoch (zero if no real-time
	// available).
	PredictedArrivalTime int64 `json:"predictedArrivalTime,required"`
	// Predicted departure time, in milliseconds since Unix epoch (zero if no real-time
	// available).
	PredictedDepartureTime int64 `json:"predictedDepartureTime,required"`
	// The ID of the route for the arriving vehicle.
	RouteID string `json:"routeId,required"`
	// Scheduled arrival time, in milliseconds since Unix epoch.
	ScheduledArrivalTime int64 `json:"scheduledArrivalTime,required"`
	// Scheduled departure time, in milliseconds since Unix epoch.
	ScheduledDepartureTime int64 `json:"scheduledDepartureTime,required"`
	// Time, in milliseconds since the Unix epoch, of midnight for the start of the
	// service date for the trip.
	ServiceDate int64 `json:"serviceDate,required"`
	// The ID of the stop the vehicle is arriving at.
	StopID string `json:"stopId,required"`
	// Index of the stop into the sequence of stops that make up the trip for this
	// arrival.
	StopSequence int64 `json:"stopSequence,required"`
	// Total number of stops visited on the trip for this arrival.
	TotalStopsInTrip int64 `json:"totalStopsInTrip,required"`
	// Optional trip headsign that potentially overrides the trip headsign in the
	// referenced trip element.
	TripHeadsign string `json:"tripHeadsign,required"`
	// The ID of the trip for the arriving vehicle.
	TripID string `json:"tripId,required"`
	// ID of the transit vehicle serving this trip.
	VehicleID string `json:"vehicleId,required"`
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
	TripStatus ArrivalAndDepartureListResponseDataEntryArrivalsAndDeparturesTripStatus `json:"tripStatus"`
	JSON       arrivalAndDepartureListResponseDataEntryArrivalsAndDepartureJSON        `json:"-"`
}

// arrivalAndDepartureListResponseDataEntryArrivalsAndDepartureJSON contains the
// JSON metadata for the struct
// [ArrivalAndDepartureListResponseDataEntryArrivalsAndDeparture]
type arrivalAndDepartureListResponseDataEntryArrivalsAndDepartureJSON struct {
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

func (r *ArrivalAndDepartureListResponseDataEntryArrivalsAndDeparture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r arrivalAndDepartureListResponseDataEntryArrivalsAndDepartureJSON) RawJSON() string {
	return r.raw
}

// Trip-specific status for the arriving transit vehicle.
type ArrivalAndDepartureListResponseDataEntryArrivalsAndDeparturesTripStatus struct {
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
	LastKnownLocation ArrivalAndDepartureListResponseDataEntryArrivalsAndDeparturesTripStatusLastKnownLocation `json:"lastKnownLocation"`
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
	Position ArrivalAndDepartureListResponseDataEntryArrivalsAndDeparturesTripStatusPosition `json:"position"`
	// Distance, in meters, the transit vehicle is scheduled to have progressed along
	// the active trip.
	ScheduledDistanceAlongTrip float64 `json:"scheduledDistanceAlongTrip"`
	// References to situation elements (if any) applicable to this trip.
	SituationIDs []string `json:"situationIds"`
	// ID of the transit vehicle currently serving the trip.
	VehicleID string                                                                      `json:"vehicleId"`
	JSON      arrivalAndDepartureListResponseDataEntryArrivalsAndDeparturesTripStatusJSON `json:"-"`
}

// arrivalAndDepartureListResponseDataEntryArrivalsAndDeparturesTripStatusJSON
// contains the JSON metadata for the struct
// [ArrivalAndDepartureListResponseDataEntryArrivalsAndDeparturesTripStatus]
type arrivalAndDepartureListResponseDataEntryArrivalsAndDeparturesTripStatusJSON struct {
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

func (r *ArrivalAndDepartureListResponseDataEntryArrivalsAndDeparturesTripStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r arrivalAndDepartureListResponseDataEntryArrivalsAndDeparturesTripStatusJSON) RawJSON() string {
	return r.raw
}

// Last known location of the transit vehicle.
type ArrivalAndDepartureListResponseDataEntryArrivalsAndDeparturesTripStatusLastKnownLocation struct {
	// Latitude of the last known location of the transit vehicle.
	Lat float64 `json:"lat"`
	// Longitude of the last known location of the transit vehicle.
	Lon  float64                                                                                      `json:"lon"`
	JSON arrivalAndDepartureListResponseDataEntryArrivalsAndDeparturesTripStatusLastKnownLocationJSON `json:"-"`
}

// arrivalAndDepartureListResponseDataEntryArrivalsAndDeparturesTripStatusLastKnownLocationJSON
// contains the JSON metadata for the struct
// [ArrivalAndDepartureListResponseDataEntryArrivalsAndDeparturesTripStatusLastKnownLocation]
type arrivalAndDepartureListResponseDataEntryArrivalsAndDeparturesTripStatusLastKnownLocationJSON struct {
	Lat         apijson.Field
	Lon         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ArrivalAndDepartureListResponseDataEntryArrivalsAndDeparturesTripStatusLastKnownLocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r arrivalAndDepartureListResponseDataEntryArrivalsAndDeparturesTripStatusLastKnownLocationJSON) RawJSON() string {
	return r.raw
}

// Current position of the transit vehicle.
type ArrivalAndDepartureListResponseDataEntryArrivalsAndDeparturesTripStatusPosition struct {
	// Latitude of the current position of the transit vehicle.
	Lat float64 `json:"lat"`
	// Longitude of the current position of the transit vehicle.
	Lon  float64                                                                             `json:"lon"`
	JSON arrivalAndDepartureListResponseDataEntryArrivalsAndDeparturesTripStatusPositionJSON `json:"-"`
}

// arrivalAndDepartureListResponseDataEntryArrivalsAndDeparturesTripStatusPositionJSON
// contains the JSON metadata for the struct
// [ArrivalAndDepartureListResponseDataEntryArrivalsAndDeparturesTripStatusPosition]
type arrivalAndDepartureListResponseDataEntryArrivalsAndDeparturesTripStatusPositionJSON struct {
	Lat         apijson.Field
	Lon         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ArrivalAndDepartureListResponseDataEntryArrivalsAndDeparturesTripStatusPosition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r arrivalAndDepartureListResponseDataEntryArrivalsAndDeparturesTripStatusPositionJSON) RawJSON() string {
	return r.raw
}

type ArrivalAndDepartureGetParams struct {
	ServiceDate  param.Field[int64]  `query:"serviceDate,required"`
	TripID       param.Field[string] `query:"tripId,required"`
	StopSequence param.Field[int64]  `query:"stopSequence"`
	Time         param.Field[int64]  `query:"time"`
	VehicleID    param.Field[string] `query:"vehicleId"`
}

// URLQuery serializes [ArrivalAndDepartureGetParams]'s query parameters as
// `url.Values`.
func (r ArrivalAndDepartureGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ArrivalAndDepartureListParams struct {
	// Include vehicles arriving or departing in the next n minutes.
	MinutesAfter param.Field[int64] `query:"minutesAfter"`
	// Include vehicles having arrived or departed in the previous n minutes.
	MinutesBefore param.Field[int64] `query:"minutesBefore"`
	// The specific time for querying the system status.
	Time param.Field[time.Time] `query:"time" format:"date-time"`
}

// URLQuery serializes [ArrivalAndDepartureListParams]'s query parameters as
// `url.Values`.
func (r ArrivalAndDepartureListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
