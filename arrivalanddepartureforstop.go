// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package opentransit

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
	Entry      ArrivalAndDepartureForStopGetResponseDataEntry      `json:"entry"`
	References ArrivalAndDepartureForStopGetResponseDataReferences `json:"references"`
	JSON       arrivalAndDepartureForStopGetResponseDataJSON       `json:"-"`
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

type ArrivalAndDepartureForStopGetResponseDataReferences struct {
	Agencies   []ArrivalAndDepartureForStopGetResponseDataReferencesAgency `json:"agencies"`
	Routes     []ArrivalAndDepartureForStopGetResponseDataReferencesRoute  `json:"routes"`
	Situations []interface{}                                               `json:"situations"`
	Stops      []ArrivalAndDepartureForStopGetResponseDataReferencesStop   `json:"stops"`
	StopTimes  []interface{}                                               `json:"stopTimes"`
	Trips      []ArrivalAndDepartureForStopGetResponseDataReferencesTrip   `json:"trips"`
	JSON       arrivalAndDepartureForStopGetResponseDataReferencesJSON     `json:"-"`
}

// arrivalAndDepartureForStopGetResponseDataReferencesJSON contains the JSON
// metadata for the struct [ArrivalAndDepartureForStopGetResponseDataReferences]
type arrivalAndDepartureForStopGetResponseDataReferencesJSON struct {
	Agencies    apijson.Field
	Routes      apijson.Field
	Situations  apijson.Field
	Stops       apijson.Field
	StopTimes   apijson.Field
	Trips       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ArrivalAndDepartureForStopGetResponseDataReferences) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r arrivalAndDepartureForStopGetResponseDataReferencesJSON) RawJSON() string {
	return r.raw
}

type ArrivalAndDepartureForStopGetResponseDataReferencesAgency struct {
	ID             string                                                        `json:"id,required"`
	Name           string                                                        `json:"name,required"`
	Timezone       string                                                        `json:"timezone,required"`
	URL            string                                                        `json:"url,required"`
	Disclaimer     string                                                        `json:"disclaimer"`
	Email          string                                                        `json:"email"`
	FareURL        string                                                        `json:"fareUrl"`
	Lang           string                                                        `json:"lang"`
	Phone          string                                                        `json:"phone"`
	PrivateService bool                                                          `json:"privateService"`
	JSON           arrivalAndDepartureForStopGetResponseDataReferencesAgencyJSON `json:"-"`
}

// arrivalAndDepartureForStopGetResponseDataReferencesAgencyJSON contains the JSON
// metadata for the struct
// [ArrivalAndDepartureForStopGetResponseDataReferencesAgency]
type arrivalAndDepartureForStopGetResponseDataReferencesAgencyJSON struct {
	ID             apijson.Field
	Name           apijson.Field
	Timezone       apijson.Field
	URL            apijson.Field
	Disclaimer     apijson.Field
	Email          apijson.Field
	FareURL        apijson.Field
	Lang           apijson.Field
	Phone          apijson.Field
	PrivateService apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ArrivalAndDepartureForStopGetResponseDataReferencesAgency) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r arrivalAndDepartureForStopGetResponseDataReferencesAgencyJSON) RawJSON() string {
	return r.raw
}

type ArrivalAndDepartureForStopGetResponseDataReferencesRoute struct {
	ID                string                                                       `json:"id"`
	AgencyID          string                                                       `json:"agencyId"`
	Color             string                                                       `json:"color"`
	Description       string                                                       `json:"description"`
	LongName          string                                                       `json:"longName"`
	NullSafeShortName string                                                       `json:"nullSafeShortName"`
	ShortName         string                                                       `json:"shortName"`
	TextColor         string                                                       `json:"textColor"`
	Type              int64                                                        `json:"type"`
	URL               string                                                       `json:"url"`
	JSON              arrivalAndDepartureForStopGetResponseDataReferencesRouteJSON `json:"-"`
}

// arrivalAndDepartureForStopGetResponseDataReferencesRouteJSON contains the JSON
// metadata for the struct
// [ArrivalAndDepartureForStopGetResponseDataReferencesRoute]
type arrivalAndDepartureForStopGetResponseDataReferencesRouteJSON struct {
	ID                apijson.Field
	AgencyID          apijson.Field
	Color             apijson.Field
	Description       apijson.Field
	LongName          apijson.Field
	NullSafeShortName apijson.Field
	ShortName         apijson.Field
	TextColor         apijson.Field
	Type              apijson.Field
	URL               apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ArrivalAndDepartureForStopGetResponseDataReferencesRoute) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r arrivalAndDepartureForStopGetResponseDataReferencesRouteJSON) RawJSON() string {
	return r.raw
}

type ArrivalAndDepartureForStopGetResponseDataReferencesStop struct {
	ID                 string                                                      `json:"id,required"`
	Code               string                                                      `json:"code,required"`
	Lat                float64                                                     `json:"lat,required"`
	Lon                float64                                                     `json:"lon,required"`
	Name               string                                                      `json:"name,required"`
	Direction          string                                                      `json:"direction"`
	LocationType       int64                                                       `json:"locationType"`
	Parent             string                                                      `json:"parent"`
	RouteIDs           []string                                                    `json:"routeIds"`
	StaticRouteIDs     []string                                                    `json:"staticRouteIds"`
	WheelchairBoarding string                                                      `json:"wheelchairBoarding"`
	JSON               arrivalAndDepartureForStopGetResponseDataReferencesStopJSON `json:"-"`
}

// arrivalAndDepartureForStopGetResponseDataReferencesStopJSON contains the JSON
// metadata for the struct
// [ArrivalAndDepartureForStopGetResponseDataReferencesStop]
type arrivalAndDepartureForStopGetResponseDataReferencesStopJSON struct {
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

func (r *ArrivalAndDepartureForStopGetResponseDataReferencesStop) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r arrivalAndDepartureForStopGetResponseDataReferencesStopJSON) RawJSON() string {
	return r.raw
}

type ArrivalAndDepartureForStopGetResponseDataReferencesTrip struct {
	ID             string                                                      `json:"id,required"`
	RouteID        string                                                      `json:"routeId,required"`
	BlockID        string                                                      `json:"blockId"`
	DirectionID    string                                                      `json:"directionId"`
	PeakOffpeak    int64                                                       `json:"peakOffpeak"`
	RouteShortName string                                                      `json:"routeShortName"`
	ServiceID      string                                                      `json:"serviceId"`
	ShapeID        string                                                      `json:"shapeId"`
	TimeZone       string                                                      `json:"timeZone"`
	TripHeadsign   string                                                      `json:"tripHeadsign"`
	TripShortName  string                                                      `json:"tripShortName"`
	JSON           arrivalAndDepartureForStopGetResponseDataReferencesTripJSON `json:"-"`
}

// arrivalAndDepartureForStopGetResponseDataReferencesTripJSON contains the JSON
// metadata for the struct
// [ArrivalAndDepartureForStopGetResponseDataReferencesTrip]
type arrivalAndDepartureForStopGetResponseDataReferencesTripJSON struct {
	ID             apijson.Field
	RouteID        apijson.Field
	BlockID        apijson.Field
	DirectionID    apijson.Field
	PeakOffpeak    apijson.Field
	RouteShortName apijson.Field
	ServiceID      apijson.Field
	ShapeID        apijson.Field
	TimeZone       apijson.Field
	TripHeadsign   apijson.Field
	TripShortName  apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ArrivalAndDepartureForStopGetResponseDataReferencesTrip) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r arrivalAndDepartureForStopGetResponseDataReferencesTripJSON) RawJSON() string {
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
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
