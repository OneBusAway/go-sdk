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

// ArrivalAndDepartureForStopService contains methods and other services that help
// with interacting with the OneBusAway API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewArrivalAndDepartureForStopService] method instead.
type ArrivalAndDepartureForStopService struct {
	Options []option.RequestOption
}

// NewArrivalAndDepartureForStopService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewArrivalAndDepartureForStopService(opts ...option.RequestOption) (r *ArrivalAndDepartureForStopService) {
	r = &ArrivalAndDepartureForStopService{}
	r.Options = opts
	return
}

// arrival-and-departure-for-stop
func (r *ArrivalAndDepartureForStopService) Get(ctx context.Context, stopID string, query ArrivalAndDepartureForStopGetParams, opts ...option.RequestOption) (res *ArrivalAndDepartureForStopGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if stopID == "" {
		err = errors.New("missing required stopID parameter")
		return
	}
	path := fmt.Sprintf("api/where/arrival-and-departure-for-stop/stopID.json")
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type ArrivalAndDepartureForStopGetResponse struct {
	Data ArrivalAndDepartureForStopGetResponseData `json:"data"`
	JSON arrivalAndDepartureForStopGetResponseJSON `json:"-"`
	shared.ResponseWrapper
}

// arrivalAndDepartureForStopGetResponseJSON contains the JSON metadata for the
// struct [ArrivalAndDepartureForStopGetResponse]
type arrivalAndDepartureForStopGetResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ArrivalAndDepartureForStopGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r arrivalAndDepartureForStopGetResponseJSON) RawJSON() string {
	return r.raw
}

type ArrivalAndDepartureForStopGetResponseData struct {
	Entry      ArrivalAndDepartureForStopGetResponseDataEntry `json:"entry"`
	References shared.References                              `json:"references"`
	JSON       arrivalAndDepartureForStopGetResponseDataJSON  `json:"-"`
}

// arrivalAndDepartureForStopGetResponseDataJSON contains the JSON metadata for the
// struct [ArrivalAndDepartureForStopGetResponseData]
type arrivalAndDepartureForStopGetResponseDataJSON struct {
	Entry       apijson.Field
	References  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ArrivalAndDepartureForStopGetResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r arrivalAndDepartureForStopGetResponseDataJSON) RawJSON() string {
	return r.raw
}

type ArrivalAndDepartureForStopGetResponseDataEntry struct {
	ActualTrack                string                                                   `json:"actualTrack"`
	ArrivalEnabled             bool                                                     `json:"arrivalEnabled"`
	BlockTripSequence          int64                                                    `json:"blockTripSequence"`
	DepartureEnabled           bool                                                     `json:"departureEnabled"`
	DistanceFromStop           float64                                                  `json:"distanceFromStop"`
	Frequency                  string                                                   `json:"frequency"`
	HistoricalOccupancy        string                                                   `json:"historicalOccupancy"`
	LastUpdateTime             int64                                                    `json:"lastUpdateTime"`
	NumberOfStopsAway          int64                                                    `json:"numberOfStopsAway"`
	OccupancyStatus            string                                                   `json:"occupancyStatus"`
	Predicted                  bool                                                     `json:"predicted"`
	PredictedArrivalInterval   string                                                   `json:"predictedArrivalInterval"`
	PredictedArrivalTime       int64                                                    `json:"predictedArrivalTime"`
	PredictedDepartureInterval string                                                   `json:"predictedDepartureInterval"`
	PredictedDepartureTime     int64                                                    `json:"predictedDepartureTime"`
	PredictedOccupancy         string                                                   `json:"predictedOccupancy"`
	RouteID                    string                                                   `json:"routeId"`
	RouteLongName              string                                                   `json:"routeLongName"`
	RouteShortName             string                                                   `json:"routeShortName"`
	ScheduledArrivalInterval   string                                                   `json:"scheduledArrivalInterval"`
	ScheduledArrivalTime       int64                                                    `json:"scheduledArrivalTime"`
	ScheduledDepartureInterval string                                                   `json:"scheduledDepartureInterval"`
	ScheduledDepartureTime     int64                                                    `json:"scheduledDepartureTime"`
	ScheduledTrack             string                                                   `json:"scheduledTrack"`
	ServiceDate                int64                                                    `json:"serviceDate"`
	SituationIDs               []string                                                 `json:"situationIds"`
	Status                     string                                                   `json:"status"`
	StopID                     string                                                   `json:"stopId"`
	StopSequence               int64                                                    `json:"stopSequence"`
	TotalStopsInTrip           int64                                                    `json:"totalStopsInTrip"`
	TripHeadsign               string                                                   `json:"tripHeadsign"`
	TripID                     string                                                   `json:"tripId"`
	TripStatus                 ArrivalAndDepartureForStopGetResponseDataEntryTripStatus `json:"tripStatus"`
	VehicleID                  string                                                   `json:"vehicleId"`
	JSON                       arrivalAndDepartureForStopGetResponseDataEntryJSON       `json:"-"`
}

// arrivalAndDepartureForStopGetResponseDataEntryJSON contains the JSON metadata
// for the struct [ArrivalAndDepartureForStopGetResponseDataEntry]
type arrivalAndDepartureForStopGetResponseDataEntryJSON struct {
	ActualTrack                apijson.Field
	ArrivalEnabled             apijson.Field
	BlockTripSequence          apijson.Field
	DepartureEnabled           apijson.Field
	DistanceFromStop           apijson.Field
	Frequency                  apijson.Field
	HistoricalOccupancy        apijson.Field
	LastUpdateTime             apijson.Field
	NumberOfStopsAway          apijson.Field
	OccupancyStatus            apijson.Field
	Predicted                  apijson.Field
	PredictedArrivalInterval   apijson.Field
	PredictedArrivalTime       apijson.Field
	PredictedDepartureInterval apijson.Field
	PredictedDepartureTime     apijson.Field
	PredictedOccupancy         apijson.Field
	RouteID                    apijson.Field
	RouteLongName              apijson.Field
	RouteShortName             apijson.Field
	ScheduledArrivalInterval   apijson.Field
	ScheduledArrivalTime       apijson.Field
	ScheduledDepartureInterval apijson.Field
	ScheduledDepartureTime     apijson.Field
	ScheduledTrack             apijson.Field
	ServiceDate                apijson.Field
	SituationIDs               apijson.Field
	Status                     apijson.Field
	StopID                     apijson.Field
	StopSequence               apijson.Field
	TotalStopsInTrip           apijson.Field
	TripHeadsign               apijson.Field
	TripID                     apijson.Field
	TripStatus                 apijson.Field
	VehicleID                  apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *ArrivalAndDepartureForStopGetResponseDataEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r arrivalAndDepartureForStopGetResponseDataEntryJSON) RawJSON() string {
	return r.raw
}

type ArrivalAndDepartureForStopGetResponseDataEntryTripStatus struct {
	ActiveTripID               string                                                                    `json:"activeTripId"`
	BlockTripSequence          int64                                                                     `json:"blockTripSequence"`
	ClosestStop                string                                                                    `json:"closestStop"`
	ClosestStopTimeOffset      int64                                                                     `json:"closestStopTimeOffset"`
	DistanceAlongTrip          float64                                                                   `json:"distanceAlongTrip"`
	Frequency                  string                                                                    `json:"frequency"`
	LastKnownDistanceAlongTrip float64                                                                   `json:"lastKnownDistanceAlongTrip"`
	LastKnownLocation          ArrivalAndDepartureForStopGetResponseDataEntryTripStatusLastKnownLocation `json:"lastKnownLocation"`
	LastKnownOrientation       float64                                                                   `json:"lastKnownOrientation"`
	LastLocationUpdateTime     int64                                                                     `json:"lastLocationUpdateTime"`
	LastUpdateTime             int64                                                                     `json:"lastUpdateTime"`
	NextStop                   string                                                                    `json:"nextStop"`
	NextStopTimeOffset         int64                                                                     `json:"nextStopTimeOffset"`
	OccupancyCapacity          int64                                                                     `json:"occupancyCapacity"`
	OccupancyCount             int64                                                                     `json:"occupancyCount"`
	OccupancyStatus            string                                                                    `json:"occupancyStatus"`
	Orientation                float64                                                                   `json:"orientation"`
	Phase                      string                                                                    `json:"phase"`
	Position                   ArrivalAndDepartureForStopGetResponseDataEntryTripStatusPosition          `json:"position"`
	Predicted                  bool                                                                      `json:"predicted"`
	ScheduledDistanceAlongTrip float64                                                                   `json:"scheduledDistanceAlongTrip"`
	ScheduleDeviation          int64                                                                     `json:"scheduleDeviation"`
	ServiceDate                int64                                                                     `json:"serviceDate"`
	SituationIDs               []string                                                                  `json:"situationIds"`
	Status                     string                                                                    `json:"status"`
	TotalDistanceAlongTrip     float64                                                                   `json:"totalDistanceAlongTrip"`
	VehicleID                  string                                                                    `json:"vehicleId"`
	JSON                       arrivalAndDepartureForStopGetResponseDataEntryTripStatusJSON              `json:"-"`
}

// arrivalAndDepartureForStopGetResponseDataEntryTripStatusJSON contains the JSON
// metadata for the struct
// [ArrivalAndDepartureForStopGetResponseDataEntryTripStatus]
type arrivalAndDepartureForStopGetResponseDataEntryTripStatusJSON struct {
	ActiveTripID               apijson.Field
	BlockTripSequence          apijson.Field
	ClosestStop                apijson.Field
	ClosestStopTimeOffset      apijson.Field
	DistanceAlongTrip          apijson.Field
	Frequency                  apijson.Field
	LastKnownDistanceAlongTrip apijson.Field
	LastKnownLocation          apijson.Field
	LastKnownOrientation       apijson.Field
	LastLocationUpdateTime     apijson.Field
	LastUpdateTime             apijson.Field
	NextStop                   apijson.Field
	NextStopTimeOffset         apijson.Field
	OccupancyCapacity          apijson.Field
	OccupancyCount             apijson.Field
	OccupancyStatus            apijson.Field
	Orientation                apijson.Field
	Phase                      apijson.Field
	Position                   apijson.Field
	Predicted                  apijson.Field
	ScheduledDistanceAlongTrip apijson.Field
	ScheduleDeviation          apijson.Field
	ServiceDate                apijson.Field
	SituationIDs               apijson.Field
	Status                     apijson.Field
	TotalDistanceAlongTrip     apijson.Field
	VehicleID                  apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *ArrivalAndDepartureForStopGetResponseDataEntryTripStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r arrivalAndDepartureForStopGetResponseDataEntryTripStatusJSON) RawJSON() string {
	return r.raw
}

type ArrivalAndDepartureForStopGetResponseDataEntryTripStatusLastKnownLocation struct {
	Lat  float64                                                                       `json:"lat"`
	Lon  float64                                                                       `json:"lon"`
	JSON arrivalAndDepartureForStopGetResponseDataEntryTripStatusLastKnownLocationJSON `json:"-"`
}

// arrivalAndDepartureForStopGetResponseDataEntryTripStatusLastKnownLocationJSON
// contains the JSON metadata for the struct
// [ArrivalAndDepartureForStopGetResponseDataEntryTripStatusLastKnownLocation]
type arrivalAndDepartureForStopGetResponseDataEntryTripStatusLastKnownLocationJSON struct {
	Lat         apijson.Field
	Lon         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ArrivalAndDepartureForStopGetResponseDataEntryTripStatusLastKnownLocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r arrivalAndDepartureForStopGetResponseDataEntryTripStatusLastKnownLocationJSON) RawJSON() string {
	return r.raw
}

type ArrivalAndDepartureForStopGetResponseDataEntryTripStatusPosition struct {
	Lat  float64                                                              `json:"lat"`
	Lon  float64                                                              `json:"lon"`
	JSON arrivalAndDepartureForStopGetResponseDataEntryTripStatusPositionJSON `json:"-"`
}

// arrivalAndDepartureForStopGetResponseDataEntryTripStatusPositionJSON contains
// the JSON metadata for the struct
// [ArrivalAndDepartureForStopGetResponseDataEntryTripStatusPosition]
type arrivalAndDepartureForStopGetResponseDataEntryTripStatusPositionJSON struct {
	Lat         apijson.Field
	Lon         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ArrivalAndDepartureForStopGetResponseDataEntryTripStatusPosition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r arrivalAndDepartureForStopGetResponseDataEntryTripStatusPositionJSON) RawJSON() string {
	return r.raw
}

type ArrivalAndDepartureForStopGetParams struct {
	ServiceDate  param.Field[int64]  `query:"serviceDate,required"`
	TripID       param.Field[string] `query:"tripId,required"`
	StopSequence param.Field[int64]  `query:"stopSequence"`
	Time         param.Field[int64]  `query:"time"`
	VehicleID    param.Field[string] `query:"vehicleId"`
}

// URLQuery serializes [ArrivalAndDepartureForStopGetParams]'s query parameters as
// `url.Values`.
func (r ArrivalAndDepartureForStopGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
