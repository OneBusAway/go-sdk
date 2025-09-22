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

// VehiclesForAgencyService contains methods and other services that help with
// interacting with the onebusaway-sdk API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewVehiclesForAgencyService] method instead.
type VehiclesForAgencyService struct {
	Options []option.RequestOption
}

// NewVehiclesForAgencyService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewVehiclesForAgencyService(opts ...option.RequestOption) (r *VehiclesForAgencyService) {
	r = &VehiclesForAgencyService{}
	r.Options = opts
	return
}

// Get vehicles for a specific agency
func (r *VehiclesForAgencyService) List(ctx context.Context, agencyID string, query VehiclesForAgencyListParams, opts ...option.RequestOption) (res *VehiclesForAgencyListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if agencyID == "" {
		err = errors.New("missing required agencyID parameter")
		return
	}
	path := fmt.Sprintf("api/where/vehicles-for-agency/%s.json", agencyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type VehiclesForAgencyListResponse struct {
	Data VehiclesForAgencyListResponseData `json:"data,required"`
	JSON vehiclesForAgencyListResponseJSON `json:"-"`
	shared.ResponseWrapper
}

// vehiclesForAgencyListResponseJSON contains the JSON metadata for the struct
// [VehiclesForAgencyListResponse]
type vehiclesForAgencyListResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VehiclesForAgencyListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r vehiclesForAgencyListResponseJSON) RawJSON() string {
	return r.raw
}

type VehiclesForAgencyListResponseData struct {
	LimitExceeded bool                                    `json:"limitExceeded,required"`
	List          []VehiclesForAgencyListResponseDataList `json:"list,required"`
	References    shared.References                       `json:"references,required"`
	JSON          vehiclesForAgencyListResponseDataJSON   `json:"-"`
}

// vehiclesForAgencyListResponseDataJSON contains the JSON metadata for the struct
// [VehiclesForAgencyListResponseData]
type vehiclesForAgencyListResponseDataJSON struct {
	LimitExceeded apijson.Field
	List          apijson.Field
	References    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *VehiclesForAgencyListResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r vehiclesForAgencyListResponseDataJSON) RawJSON() string {
	return r.raw
}

type VehiclesForAgencyListResponseDataList struct {
	LastLocationUpdateTime int64                                           `json:"lastLocationUpdateTime,required"`
	LastUpdateTime         int64                                           `json:"lastUpdateTime,required"`
	Location               VehiclesForAgencyListResponseDataListLocation   `json:"location,required"`
	TripID                 string                                          `json:"tripId,required"`
	TripStatus             VehiclesForAgencyListResponseDataListTripStatus `json:"tripStatus,required"`
	VehicleID              string                                          `json:"vehicleId,required"`
	OccupancyCapacity      int64                                           `json:"occupancyCapacity"`
	OccupancyCount         int64                                           `json:"occupancyCount"`
	OccupancyStatus        string                                          `json:"occupancyStatus"`
	Phase                  string                                          `json:"phase"`
	Status                 string                                          `json:"status"`
	JSON                   vehiclesForAgencyListResponseDataListJSON       `json:"-"`
}

// vehiclesForAgencyListResponseDataListJSON contains the JSON metadata for the
// struct [VehiclesForAgencyListResponseDataList]
type vehiclesForAgencyListResponseDataListJSON struct {
	LastLocationUpdateTime apijson.Field
	LastUpdateTime         apijson.Field
	Location               apijson.Field
	TripID                 apijson.Field
	TripStatus             apijson.Field
	VehicleID              apijson.Field
	OccupancyCapacity      apijson.Field
	OccupancyCount         apijson.Field
	OccupancyStatus        apijson.Field
	Phase                  apijson.Field
	Status                 apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *VehiclesForAgencyListResponseDataList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r vehiclesForAgencyListResponseDataListJSON) RawJSON() string {
	return r.raw
}

type VehiclesForAgencyListResponseDataListLocation struct {
	Lat  float64                                           `json:"lat"`
	Lon  float64                                           `json:"lon"`
	JSON vehiclesForAgencyListResponseDataListLocationJSON `json:"-"`
}

// vehiclesForAgencyListResponseDataListLocationJSON contains the JSON metadata for
// the struct [VehiclesForAgencyListResponseDataListLocation]
type vehiclesForAgencyListResponseDataListLocationJSON struct {
	Lat         apijson.Field
	Lon         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VehiclesForAgencyListResponseDataListLocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r vehiclesForAgencyListResponseDataListLocationJSON) RawJSON() string {
	return r.raw
}

type VehiclesForAgencyListResponseDataListTripStatus struct {
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
	LastKnownLocation VehiclesForAgencyListResponseDataListTripStatusLastKnownLocation `json:"lastKnownLocation"`
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
	Position VehiclesForAgencyListResponseDataListTripStatusPosition `json:"position"`
	// Distance, in meters, the transit vehicle is scheduled to have progressed along
	// the active trip.
	ScheduledDistanceAlongTrip float64 `json:"scheduledDistanceAlongTrip"`
	// References to situation elements (if any) applicable to this trip.
	SituationIDs []string `json:"situationIds"`
	// ID of the transit vehicle currently serving the trip.
	VehicleID string                                              `json:"vehicleId"`
	JSON      vehiclesForAgencyListResponseDataListTripStatusJSON `json:"-"`
}

// vehiclesForAgencyListResponseDataListTripStatusJSON contains the JSON metadata
// for the struct [VehiclesForAgencyListResponseDataListTripStatus]
type vehiclesForAgencyListResponseDataListTripStatusJSON struct {
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

func (r *VehiclesForAgencyListResponseDataListTripStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r vehiclesForAgencyListResponseDataListTripStatusJSON) RawJSON() string {
	return r.raw
}

// Last known location of the transit vehicle.
type VehiclesForAgencyListResponseDataListTripStatusLastKnownLocation struct {
	// Latitude of the last known location of the transit vehicle.
	Lat float64 `json:"lat"`
	// Longitude of the last known location of the transit vehicle.
	Lon  float64                                                              `json:"lon"`
	JSON vehiclesForAgencyListResponseDataListTripStatusLastKnownLocationJSON `json:"-"`
}

// vehiclesForAgencyListResponseDataListTripStatusLastKnownLocationJSON contains
// the JSON metadata for the struct
// [VehiclesForAgencyListResponseDataListTripStatusLastKnownLocation]
type vehiclesForAgencyListResponseDataListTripStatusLastKnownLocationJSON struct {
	Lat         apijson.Field
	Lon         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VehiclesForAgencyListResponseDataListTripStatusLastKnownLocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r vehiclesForAgencyListResponseDataListTripStatusLastKnownLocationJSON) RawJSON() string {
	return r.raw
}

// Current position of the transit vehicle.
type VehiclesForAgencyListResponseDataListTripStatusPosition struct {
	// Latitude of the current position of the transit vehicle.
	Lat float64 `json:"lat"`
	// Longitude of the current position of the transit vehicle.
	Lon  float64                                                     `json:"lon"`
	JSON vehiclesForAgencyListResponseDataListTripStatusPositionJSON `json:"-"`
}

// vehiclesForAgencyListResponseDataListTripStatusPositionJSON contains the JSON
// metadata for the struct
// [VehiclesForAgencyListResponseDataListTripStatusPosition]
type vehiclesForAgencyListResponseDataListTripStatusPositionJSON struct {
	Lat         apijson.Field
	Lon         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VehiclesForAgencyListResponseDataListTripStatusPosition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r vehiclesForAgencyListResponseDataListTripStatusPositionJSON) RawJSON() string {
	return r.raw
}

type VehiclesForAgencyListParams struct {
	// Specific time for querying the status (timestamp format)
	Time param.Field[string] `query:"time"`
}

// URLQuery serializes [VehiclesForAgencyListParams]'s query parameters as
// `url.Values`.
func (r VehiclesForAgencyListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
