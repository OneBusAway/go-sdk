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
)

// WhereStopArrivalAndDepartureService contains methods and other services that
// help with interacting with the open-transit API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWhereStopArrivalAndDepartureService] method instead.
type WhereStopArrivalAndDepartureService struct {
	Options []option.RequestOption
}

// NewWhereStopArrivalAndDepartureService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewWhereStopArrivalAndDepartureService(opts ...option.RequestOption) (r *WhereStopArrivalAndDepartureService) {
	r = &WhereStopArrivalAndDepartureService{}
	r.Options = opts
	return
}

// arrival-and-departure-for-stop
func (r *WhereStopArrivalAndDepartureService) Get(ctx context.Context, stopID string, query WhereStopArrivalAndDepartureGetParams, opts ...option.RequestOption) (res *WhereStopArrivalAndDepartureGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if stopID == "" {
		err = errors.New("missing required stopID parameter")
		return
	}
	path := fmt.Sprintf("api/where/arrival-and-departure-for-stop/stopID.json")
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type WhereStopArrivalAndDepartureGetResponse struct {
	Code        int64                                       `json:"code,required"`
	CurrentTime int64                                       `json:"currentTime,required"`
	Text        string                                      `json:"text,required"`
	Version     int64                                       `json:"version,required"`
	Data        WhereStopArrivalAndDepartureGetResponseData `json:"data"`
	JSON        whereStopArrivalAndDepartureGetResponseJSON `json:"-"`
}

// whereStopArrivalAndDepartureGetResponseJSON contains the JSON metadata for the
// struct [WhereStopArrivalAndDepartureGetResponse]
type whereStopArrivalAndDepartureGetResponseJSON struct {
	Code        apijson.Field
	CurrentTime apijson.Field
	Text        apijson.Field
	Version     apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WhereStopArrivalAndDepartureGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r whereStopArrivalAndDepartureGetResponseJSON) RawJSON() string {
	return r.raw
}

type WhereStopArrivalAndDepartureGetResponseData struct {
	Entry      WhereStopArrivalAndDepartureGetResponseDataEntry      `json:"entry"`
	References WhereStopArrivalAndDepartureGetResponseDataReferences `json:"references"`
	JSON       whereStopArrivalAndDepartureGetResponseDataJSON       `json:"-"`
}

// whereStopArrivalAndDepartureGetResponseDataJSON contains the JSON metadata for
// the struct [WhereStopArrivalAndDepartureGetResponseData]
type whereStopArrivalAndDepartureGetResponseDataJSON struct {
	Entry       apijson.Field
	References  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WhereStopArrivalAndDepartureGetResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r whereStopArrivalAndDepartureGetResponseDataJSON) RawJSON() string {
	return r.raw
}

type WhereStopArrivalAndDepartureGetResponseDataEntry struct {
	ActualTrack                string                                                     `json:"actualTrack"`
	ArrivalEnabled             bool                                                       `json:"arrivalEnabled"`
	BlockTripSequence          int64                                                      `json:"blockTripSequence"`
	DepartureEnabled           bool                                                       `json:"departureEnabled"`
	DistanceFromStop           float64                                                    `json:"distanceFromStop"`
	Frequency                  string                                                     `json:"frequency"`
	HistoricalOccupancy        string                                                     `json:"historicalOccupancy"`
	LastUpdateTime             int64                                                      `json:"lastUpdateTime"`
	NumberOfStopsAway          int64                                                      `json:"numberOfStopsAway"`
	OccupancyStatus            string                                                     `json:"occupancyStatus"`
	Predicted                  bool                                                       `json:"predicted"`
	PredictedArrivalInterval   string                                                     `json:"predictedArrivalInterval"`
	PredictedArrivalTime       int64                                                      `json:"predictedArrivalTime"`
	PredictedDepartureInterval string                                                     `json:"predictedDepartureInterval"`
	PredictedDepartureTime     int64                                                      `json:"predictedDepartureTime"`
	PredictedOccupancy         string                                                     `json:"predictedOccupancy"`
	RouteID                    string                                                     `json:"routeId"`
	RouteLongName              string                                                     `json:"routeLongName"`
	RouteShortName             string                                                     `json:"routeShortName"`
	ScheduledArrivalInterval   string                                                     `json:"scheduledArrivalInterval"`
	ScheduledArrivalTime       int64                                                      `json:"scheduledArrivalTime"`
	ScheduledDepartureInterval string                                                     `json:"scheduledDepartureInterval"`
	ScheduledDepartureTime     int64                                                      `json:"scheduledDepartureTime"`
	ScheduledTrack             string                                                     `json:"scheduledTrack"`
	ServiceDate                int64                                                      `json:"serviceDate"`
	SituationIDs               []string                                                   `json:"situationIds"`
	Status                     string                                                     `json:"status"`
	StopID                     string                                                     `json:"stopId"`
	StopSequence               int64                                                      `json:"stopSequence"`
	TotalStopsInTrip           int64                                                      `json:"totalStopsInTrip"`
	TripHeadsign               string                                                     `json:"tripHeadsign"`
	TripID                     string                                                     `json:"tripId"`
	TripStatus                 WhereStopArrivalAndDepartureGetResponseDataEntryTripStatus `json:"tripStatus"`
	VehicleID                  string                                                     `json:"vehicleId"`
	JSON                       whereStopArrivalAndDepartureGetResponseDataEntryJSON       `json:"-"`
}

// whereStopArrivalAndDepartureGetResponseDataEntryJSON contains the JSON metadata
// for the struct [WhereStopArrivalAndDepartureGetResponseDataEntry]
type whereStopArrivalAndDepartureGetResponseDataEntryJSON struct {
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

func (r *WhereStopArrivalAndDepartureGetResponseDataEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r whereStopArrivalAndDepartureGetResponseDataEntryJSON) RawJSON() string {
	return r.raw
}

type WhereStopArrivalAndDepartureGetResponseDataEntryTripStatus struct {
	ActiveTripID               string                                                                      `json:"activeTripId"`
	BlockTripSequence          int64                                                                       `json:"blockTripSequence"`
	ClosestStop                string                                                                      `json:"closestStop"`
	ClosestStopTimeOffset      int64                                                                       `json:"closestStopTimeOffset"`
	DistanceAlongTrip          float64                                                                     `json:"distanceAlongTrip"`
	Frequency                  string                                                                      `json:"frequency"`
	LastKnownDistanceAlongTrip float64                                                                     `json:"lastKnownDistanceAlongTrip"`
	LastKnownLocation          WhereStopArrivalAndDepartureGetResponseDataEntryTripStatusLastKnownLocation `json:"lastKnownLocation"`
	LastKnownOrientation       float64                                                                     `json:"lastKnownOrientation"`
	LastLocationUpdateTime     int64                                                                       `json:"lastLocationUpdateTime"`
	LastUpdateTime             int64                                                                       `json:"lastUpdateTime"`
	NextStop                   string                                                                      `json:"nextStop"`
	NextStopTimeOffset         int64                                                                       `json:"nextStopTimeOffset"`
	OccupancyCapacity          int64                                                                       `json:"occupancyCapacity"`
	OccupancyCount             int64                                                                       `json:"occupancyCount"`
	OccupancyStatus            string                                                                      `json:"occupancyStatus"`
	Orientation                float64                                                                     `json:"orientation"`
	Phase                      string                                                                      `json:"phase"`
	Position                   WhereStopArrivalAndDepartureGetResponseDataEntryTripStatusPosition          `json:"position"`
	Predicted                  bool                                                                        `json:"predicted"`
	ScheduledDistanceAlongTrip float64                                                                     `json:"scheduledDistanceAlongTrip"`
	ScheduleDeviation          int64                                                                       `json:"scheduleDeviation"`
	ServiceDate                int64                                                                       `json:"serviceDate"`
	SituationIDs               []string                                                                    `json:"situationIds"`
	Status                     string                                                                      `json:"status"`
	TotalDistanceAlongTrip     float64                                                                     `json:"totalDistanceAlongTrip"`
	VehicleID                  string                                                                      `json:"vehicleId"`
	JSON                       whereStopArrivalAndDepartureGetResponseDataEntryTripStatusJSON              `json:"-"`
}

// whereStopArrivalAndDepartureGetResponseDataEntryTripStatusJSON contains the JSON
// metadata for the struct
// [WhereStopArrivalAndDepartureGetResponseDataEntryTripStatus]
type whereStopArrivalAndDepartureGetResponseDataEntryTripStatusJSON struct {
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

func (r *WhereStopArrivalAndDepartureGetResponseDataEntryTripStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r whereStopArrivalAndDepartureGetResponseDataEntryTripStatusJSON) RawJSON() string {
	return r.raw
}

type WhereStopArrivalAndDepartureGetResponseDataEntryTripStatusLastKnownLocation struct {
	Lat  float64                                                                         `json:"lat"`
	Lon  float64                                                                         `json:"lon"`
	JSON whereStopArrivalAndDepartureGetResponseDataEntryTripStatusLastKnownLocationJSON `json:"-"`
}

// whereStopArrivalAndDepartureGetResponseDataEntryTripStatusLastKnownLocationJSON
// contains the JSON metadata for the struct
// [WhereStopArrivalAndDepartureGetResponseDataEntryTripStatusLastKnownLocation]
type whereStopArrivalAndDepartureGetResponseDataEntryTripStatusLastKnownLocationJSON struct {
	Lat         apijson.Field
	Lon         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WhereStopArrivalAndDepartureGetResponseDataEntryTripStatusLastKnownLocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r whereStopArrivalAndDepartureGetResponseDataEntryTripStatusLastKnownLocationJSON) RawJSON() string {
	return r.raw
}

type WhereStopArrivalAndDepartureGetResponseDataEntryTripStatusPosition struct {
	Lat  float64                                                                `json:"lat"`
	Lon  float64                                                                `json:"lon"`
	JSON whereStopArrivalAndDepartureGetResponseDataEntryTripStatusPositionJSON `json:"-"`
}

// whereStopArrivalAndDepartureGetResponseDataEntryTripStatusPositionJSON contains
// the JSON metadata for the struct
// [WhereStopArrivalAndDepartureGetResponseDataEntryTripStatusPosition]
type whereStopArrivalAndDepartureGetResponseDataEntryTripStatusPositionJSON struct {
	Lat         apijson.Field
	Lon         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WhereStopArrivalAndDepartureGetResponseDataEntryTripStatusPosition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r whereStopArrivalAndDepartureGetResponseDataEntryTripStatusPositionJSON) RawJSON() string {
	return r.raw
}

type WhereStopArrivalAndDepartureGetResponseDataReferences struct {
	Agencies   []WhereStopArrivalAndDepartureGetResponseDataReferencesAgency `json:"agencies"`
	Routes     []WhereStopArrivalAndDepartureGetResponseDataReferencesRoute  `json:"routes"`
	Situations []interface{}                                                 `json:"situations"`
	Stops      []WhereStopArrivalAndDepartureGetResponseDataReferencesStop   `json:"stops"`
	StopTimes  []interface{}                                                 `json:"stopTimes"`
	Trips      []WhereStopArrivalAndDepartureGetResponseDataReferencesTrip   `json:"trips"`
	JSON       whereStopArrivalAndDepartureGetResponseDataReferencesJSON     `json:"-"`
}

// whereStopArrivalAndDepartureGetResponseDataReferencesJSON contains the JSON
// metadata for the struct [WhereStopArrivalAndDepartureGetResponseDataReferences]
type whereStopArrivalAndDepartureGetResponseDataReferencesJSON struct {
	Agencies    apijson.Field
	Routes      apijson.Field
	Situations  apijson.Field
	Stops       apijson.Field
	StopTimes   apijson.Field
	Trips       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WhereStopArrivalAndDepartureGetResponseDataReferences) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r whereStopArrivalAndDepartureGetResponseDataReferencesJSON) RawJSON() string {
	return r.raw
}

type WhereStopArrivalAndDepartureGetResponseDataReferencesAgency struct {
	ID             string                                                          `json:"id,required"`
	Name           string                                                          `json:"name,required"`
	Timezone       string                                                          `json:"timezone,required"`
	URL            string                                                          `json:"url,required"`
	Disclaimer     string                                                          `json:"disclaimer"`
	Email          string                                                          `json:"email"`
	FareURL        string                                                          `json:"fareUrl"`
	Lang           string                                                          `json:"lang"`
	Phone          string                                                          `json:"phone"`
	PrivateService bool                                                            `json:"privateService"`
	JSON           whereStopArrivalAndDepartureGetResponseDataReferencesAgencyJSON `json:"-"`
}

// whereStopArrivalAndDepartureGetResponseDataReferencesAgencyJSON contains the
// JSON metadata for the struct
// [WhereStopArrivalAndDepartureGetResponseDataReferencesAgency]
type whereStopArrivalAndDepartureGetResponseDataReferencesAgencyJSON struct {
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

func (r *WhereStopArrivalAndDepartureGetResponseDataReferencesAgency) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r whereStopArrivalAndDepartureGetResponseDataReferencesAgencyJSON) RawJSON() string {
	return r.raw
}

type WhereStopArrivalAndDepartureGetResponseDataReferencesRoute struct {
	ID                string                                                         `json:"id"`
	AgencyID          string                                                         `json:"agencyId"`
	Color             string                                                         `json:"color"`
	Description       string                                                         `json:"description"`
	LongName          string                                                         `json:"longName"`
	NullSafeShortName string                                                         `json:"nullSafeShortName"`
	ShortName         string                                                         `json:"shortName"`
	TextColor         string                                                         `json:"textColor"`
	Type              int64                                                          `json:"type"`
	URL               string                                                         `json:"url"`
	JSON              whereStopArrivalAndDepartureGetResponseDataReferencesRouteJSON `json:"-"`
}

// whereStopArrivalAndDepartureGetResponseDataReferencesRouteJSON contains the JSON
// metadata for the struct
// [WhereStopArrivalAndDepartureGetResponseDataReferencesRoute]
type whereStopArrivalAndDepartureGetResponseDataReferencesRouteJSON struct {
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

func (r *WhereStopArrivalAndDepartureGetResponseDataReferencesRoute) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r whereStopArrivalAndDepartureGetResponseDataReferencesRouteJSON) RawJSON() string {
	return r.raw
}

type WhereStopArrivalAndDepartureGetResponseDataReferencesStop struct {
	ID                 string                                                        `json:"id,required"`
	Code               string                                                        `json:"code,required"`
	Lat                float64                                                       `json:"lat,required"`
	Lon                float64                                                       `json:"lon,required"`
	Name               string                                                        `json:"name,required"`
	Direction          string                                                        `json:"direction"`
	LocationType       int64                                                         `json:"locationType"`
	Parent             string                                                        `json:"parent"`
	RouteIDs           []string                                                      `json:"routeIds"`
	StaticRouteIDs     []string                                                      `json:"staticRouteIds"`
	WheelchairBoarding string                                                        `json:"wheelchairBoarding"`
	JSON               whereStopArrivalAndDepartureGetResponseDataReferencesStopJSON `json:"-"`
}

// whereStopArrivalAndDepartureGetResponseDataReferencesStopJSON contains the JSON
// metadata for the struct
// [WhereStopArrivalAndDepartureGetResponseDataReferencesStop]
type whereStopArrivalAndDepartureGetResponseDataReferencesStopJSON struct {
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

func (r *WhereStopArrivalAndDepartureGetResponseDataReferencesStop) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r whereStopArrivalAndDepartureGetResponseDataReferencesStopJSON) RawJSON() string {
	return r.raw
}

type WhereStopArrivalAndDepartureGetResponseDataReferencesTrip struct {
	ID             string                                                        `json:"id,required"`
	RouteID        string                                                        `json:"routeId,required"`
	BlockID        string                                                        `json:"blockId"`
	DirectionID    string                                                        `json:"directionId"`
	PeakOffpeak    int64                                                         `json:"peakOffpeak"`
	RouteShortName string                                                        `json:"routeShortName"`
	ServiceID      string                                                        `json:"serviceId"`
	ShapeID        string                                                        `json:"shapeId"`
	TimeZone       string                                                        `json:"timeZone"`
	TripHeadsign   string                                                        `json:"tripHeadsign"`
	TripShortName  string                                                        `json:"tripShortName"`
	JSON           whereStopArrivalAndDepartureGetResponseDataReferencesTripJSON `json:"-"`
}

// whereStopArrivalAndDepartureGetResponseDataReferencesTripJSON contains the JSON
// metadata for the struct
// [WhereStopArrivalAndDepartureGetResponseDataReferencesTrip]
type whereStopArrivalAndDepartureGetResponseDataReferencesTripJSON struct {
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

func (r *WhereStopArrivalAndDepartureGetResponseDataReferencesTrip) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r whereStopArrivalAndDepartureGetResponseDataReferencesTripJSON) RawJSON() string {
	return r.raw
}

type WhereStopArrivalAndDepartureGetParams struct {
	ServiceDate  param.Field[int64]  `query:"serviceDate,required"`
	TripID       param.Field[string] `query:"tripId,required"`
	StopSequence param.Field[int64]  `query:"stopSequence"`
	Time         param.Field[int64]  `query:"time"`
	VehicleID    param.Field[string] `query:"vehicleId"`
}

// URLQuery serializes [WhereStopArrivalAndDepartureGetParams]'s query parameters
// as `url.Values`.
func (r WhereStopArrivalAndDepartureGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
