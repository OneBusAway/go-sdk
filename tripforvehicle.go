// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package onebusaway

import (
	"context"
	"errors"
	"fmt"
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

// TripForVehicleService contains methods and other services that help with
// interacting with the onebusaway-sdk API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTripForVehicleService] method instead.
type TripForVehicleService struct {
	Options []option.RequestOption
}

// NewTripForVehicleService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTripForVehicleService(opts ...option.RequestOption) (r *TripForVehicleService) {
	r = &TripForVehicleService{}
	r.Options = opts
	return
}

// Retrieve trip for a specific vehicle
func (r *TripForVehicleService) Get(ctx context.Context, vehicleID string, query TripForVehicleGetParams, opts ...option.RequestOption) (res *TripForVehicleGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if vehicleID == "" {
		err = errors.New("missing required vehicleID parameter")
		return
	}
	path := fmt.Sprintf("api/where/trip-for-vehicle/%s.json", vehicleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type TripForVehicleGetResponse struct {
	Data TripForVehicleGetResponseData `json:"data,required"`
	JSON tripForVehicleGetResponseJSON `json:"-"`
	shared.ResponseWrapper
}

// tripForVehicleGetResponseJSON contains the JSON metadata for the struct
// [TripForVehicleGetResponse]
type tripForVehicleGetResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TripForVehicleGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tripForVehicleGetResponseJSON) RawJSON() string {
	return r.raw
}

type TripForVehicleGetResponseData struct {
	Entry      TripForVehicleGetResponseDataEntry `json:"entry,required"`
	References shared.References                  `json:"references,required"`
	JSON       tripForVehicleGetResponseDataJSON  `json:"-"`
}

// tripForVehicleGetResponseDataJSON contains the JSON metadata for the struct
// [TripForVehicleGetResponseData]
type tripForVehicleGetResponseDataJSON struct {
	Entry       apijson.Field
	References  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TripForVehicleGetResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tripForVehicleGetResponseDataJSON) RawJSON() string {
	return r.raw
}

type TripForVehicleGetResponseDataEntry struct {
	TripID       string                                     `json:"tripId,required"`
	Frequency    string                                     `json:"frequency,nullable"`
	Schedule     TripForVehicleGetResponseDataEntrySchedule `json:"schedule"`
	ServiceDate  int64                                      `json:"serviceDate"`
	SituationIDs []string                                   `json:"situationIds"`
	Status       TripForVehicleGetResponseDataEntryStatus   `json:"status"`
	JSON         tripForVehicleGetResponseDataEntryJSON     `json:"-"`
}

// tripForVehicleGetResponseDataEntryJSON contains the JSON metadata for the struct
// [TripForVehicleGetResponseDataEntry]
type tripForVehicleGetResponseDataEntryJSON struct {
	TripID       apijson.Field
	Frequency    apijson.Field
	Schedule     apijson.Field
	ServiceDate  apijson.Field
	SituationIDs apijson.Field
	Status       apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *TripForVehicleGetResponseDataEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tripForVehicleGetResponseDataEntryJSON) RawJSON() string {
	return r.raw
}

type TripForVehicleGetResponseDataEntrySchedule struct {
	NextTripID     string                                               `json:"nextTripId,required"`
	PreviousTripID string                                               `json:"previousTripId,required"`
	StopTimes      []TripForVehicleGetResponseDataEntryScheduleStopTime `json:"stopTimes,required"`
	TimeZone       string                                               `json:"timeZone,required"`
	Frequency      string                                               `json:"frequency,nullable"`
	JSON           tripForVehicleGetResponseDataEntryScheduleJSON       `json:"-"`
}

// tripForVehicleGetResponseDataEntryScheduleJSON contains the JSON metadata for
// the struct [TripForVehicleGetResponseDataEntrySchedule]
type tripForVehicleGetResponseDataEntryScheduleJSON struct {
	NextTripID     apijson.Field
	PreviousTripID apijson.Field
	StopTimes      apijson.Field
	TimeZone       apijson.Field
	Frequency      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *TripForVehicleGetResponseDataEntrySchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tripForVehicleGetResponseDataEntryScheduleJSON) RawJSON() string {
	return r.raw
}

type TripForVehicleGetResponseDataEntryScheduleStopTime struct {
	ArrivalTime         int64                                                  `json:"arrivalTime"`
	DepartureTime       int64                                                  `json:"departureTime"`
	DistanceAlongTrip   float64                                                `json:"distanceAlongTrip"`
	HistoricalOccupancy string                                                 `json:"historicalOccupancy"`
	StopHeadsign        string                                                 `json:"stopHeadsign"`
	StopID              string                                                 `json:"stopId"`
	JSON                tripForVehicleGetResponseDataEntryScheduleStopTimeJSON `json:"-"`
}

// tripForVehicleGetResponseDataEntryScheduleStopTimeJSON contains the JSON
// metadata for the struct [TripForVehicleGetResponseDataEntryScheduleStopTime]
type tripForVehicleGetResponseDataEntryScheduleStopTimeJSON struct {
	ArrivalTime         apijson.Field
	DepartureTime       apijson.Field
	DistanceAlongTrip   apijson.Field
	HistoricalOccupancy apijson.Field
	StopHeadsign        apijson.Field
	StopID              apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *TripForVehicleGetResponseDataEntryScheduleStopTime) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tripForVehicleGetResponseDataEntryScheduleStopTimeJSON) RawJSON() string {
	return r.raw
}

type TripForVehicleGetResponseDataEntryStatus struct {
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
	LastKnownLocation TripForVehicleGetResponseDataEntryStatusLastKnownLocation `json:"lastKnownLocation"`
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
	Position TripForVehicleGetResponseDataEntryStatusPosition `json:"position"`
	// Distance, in meters, the transit vehicle is scheduled to have progressed along
	// the active trip.
	ScheduledDistanceAlongTrip float64 `json:"scheduledDistanceAlongTrip"`
	// References to situation elements (if any) applicable to this trip.
	SituationIDs []string `json:"situationIds"`
	// ID of the transit vehicle currently serving the trip.
	VehicleID string                                       `json:"vehicleId"`
	JSON      tripForVehicleGetResponseDataEntryStatusJSON `json:"-"`
}

// tripForVehicleGetResponseDataEntryStatusJSON contains the JSON metadata for the
// struct [TripForVehicleGetResponseDataEntryStatus]
type tripForVehicleGetResponseDataEntryStatusJSON struct {
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

func (r *TripForVehicleGetResponseDataEntryStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tripForVehicleGetResponseDataEntryStatusJSON) RawJSON() string {
	return r.raw
}

// Last known location of the transit vehicle.
type TripForVehicleGetResponseDataEntryStatusLastKnownLocation struct {
	// Latitude of the last known location of the transit vehicle.
	Lat float64 `json:"lat"`
	// Longitude of the last known location of the transit vehicle.
	Lon  float64                                                       `json:"lon"`
	JSON tripForVehicleGetResponseDataEntryStatusLastKnownLocationJSON `json:"-"`
}

// tripForVehicleGetResponseDataEntryStatusLastKnownLocationJSON contains the JSON
// metadata for the struct
// [TripForVehicleGetResponseDataEntryStatusLastKnownLocation]
type tripForVehicleGetResponseDataEntryStatusLastKnownLocationJSON struct {
	Lat         apijson.Field
	Lon         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TripForVehicleGetResponseDataEntryStatusLastKnownLocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tripForVehicleGetResponseDataEntryStatusLastKnownLocationJSON) RawJSON() string {
	return r.raw
}

// Current position of the transit vehicle.
type TripForVehicleGetResponseDataEntryStatusPosition struct {
	// Latitude of the current position of the transit vehicle.
	Lat float64 `json:"lat"`
	// Longitude of the current position of the transit vehicle.
	Lon  float64                                              `json:"lon"`
	JSON tripForVehicleGetResponseDataEntryStatusPositionJSON `json:"-"`
}

// tripForVehicleGetResponseDataEntryStatusPositionJSON contains the JSON metadata
// for the struct [TripForVehicleGetResponseDataEntryStatusPosition]
type tripForVehicleGetResponseDataEntryStatusPositionJSON struct {
	Lat         apijson.Field
	Lon         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TripForVehicleGetResponseDataEntryStatusPosition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tripForVehicleGetResponseDataEntryStatusPositionJSON) RawJSON() string {
	return r.raw
}

type TripForVehicleGetParams struct {
	// Determines whether full <schedule/> element is included in the <tripDetails/>
	// section. Defaults to false.
	IncludeSchedule param.Field[bool] `query:"includeSchedule"`
	// Determines whether the full <status/> element is included in the <tripDetails/>
	// section. Defaults to true.
	IncludeStatus param.Field[bool] `query:"includeStatus"`
	// Determines whether full <trip/> element is included in the <references/>
	// section. Defaults to false.
	IncludeTrip param.Field[bool] `query:"includeTrip"`
	// Time parameter to query the system at a specific time (optional).
	Time param.Field[int64] `query:"time"`
}

// URLQuery serializes [TripForVehicleGetParams]'s query parameters as
// `url.Values`.
func (r TripForVehicleGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
