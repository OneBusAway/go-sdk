// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package opentransit

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/open-transit-go/internal/apijson"
	"github.com/stainless-sdks/open-transit-go/internal/requestconfig"
	"github.com/stainless-sdks/open-transit-go/option"
)

// WhereStopArrivalsAndDepartureService contains methods and other services that
// help with interacting with the open-transit API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWhereStopArrivalsAndDepartureService] method instead.
type WhereStopArrivalsAndDepartureService struct {
	Options []option.RequestOption
}

// NewWhereStopArrivalsAndDepartureService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewWhereStopArrivalsAndDepartureService(opts ...option.RequestOption) (r *WhereStopArrivalsAndDepartureService) {
	r = &WhereStopArrivalsAndDepartureService{}
	r.Options = opts
	return
}

// arrival-and-departure-for-stop
func (r *WhereStopArrivalsAndDepartureService) List(ctx context.Context, stopID string, opts ...option.RequestOption) (res *WhereStopArrivalsAndDepartureListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if stopID == "" {
		err = errors.New("missing required stopID parameter")
		return
	}
	path := fmt.Sprintf("api/where/arrivals-and-departures-for-stop/stopID.json")
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type WhereStopArrivalsAndDepartureListResponse struct {
	Code        int64                                         `json:"code,required"`
	CurrentTime int64                                         `json:"currentTime,required"`
	Text        string                                        `json:"text,required"`
	Version     int64                                         `json:"version,required"`
	Data        WhereStopArrivalsAndDepartureListResponseData `json:"data"`
	JSON        whereStopArrivalsAndDepartureListResponseJSON `json:"-"`
}

// whereStopArrivalsAndDepartureListResponseJSON contains the JSON metadata for the
// struct [WhereStopArrivalsAndDepartureListResponse]
type whereStopArrivalsAndDepartureListResponseJSON struct {
	Code        apijson.Field
	CurrentTime apijson.Field
	Text        apijson.Field
	Version     apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WhereStopArrivalsAndDepartureListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r whereStopArrivalsAndDepartureListResponseJSON) RawJSON() string {
	return r.raw
}

type WhereStopArrivalsAndDepartureListResponseData struct {
	Entry WhereStopArrivalsAndDepartureListResponseDataEntry `json:"entry"`
	JSON  whereStopArrivalsAndDepartureListResponseDataJSON  `json:"-"`
}

// whereStopArrivalsAndDepartureListResponseDataJSON contains the JSON metadata for
// the struct [WhereStopArrivalsAndDepartureListResponseData]
type whereStopArrivalsAndDepartureListResponseDataJSON struct {
	Entry       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WhereStopArrivalsAndDepartureListResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r whereStopArrivalsAndDepartureListResponseDataJSON) RawJSON() string {
	return r.raw
}

type WhereStopArrivalsAndDepartureListResponseDataEntry struct {
	ArrivalsAndDepartures []WhereStopArrivalsAndDepartureListResponseDataEntryArrivalsAndDeparture `json:"arrivalsAndDepartures"`
	References            WhereStopArrivalsAndDepartureListResponseDataEntryReferences             `json:"references"`
	JSON                  whereStopArrivalsAndDepartureListResponseDataEntryJSON                   `json:"-"`
}

// whereStopArrivalsAndDepartureListResponseDataEntryJSON contains the JSON
// metadata for the struct [WhereStopArrivalsAndDepartureListResponseDataEntry]
type whereStopArrivalsAndDepartureListResponseDataEntryJSON struct {
	ArrivalsAndDepartures apijson.Field
	References            apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *WhereStopArrivalsAndDepartureListResponseDataEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r whereStopArrivalsAndDepartureListResponseDataEntryJSON) RawJSON() string {
	return r.raw
}

type WhereStopArrivalsAndDepartureListResponseDataEntryArrivalsAndDeparture struct {
	ActualTrack                string                                                                            `json:"actualTrack"`
	ArrivalEnabled             bool                                                                              `json:"arrivalEnabled"`
	BlockTripSequence          int64                                                                             `json:"blockTripSequence"`
	DepartureEnabled           bool                                                                              `json:"departureEnabled"`
	DistanceFromStop           float64                                                                           `json:"distanceFromStop"`
	Frequency                  string                                                                            `json:"frequency"`
	HistoricalOccupancy        string                                                                            `json:"historicalOccupancy"`
	LastUpdateTime             int64                                                                             `json:"lastUpdateTime"`
	NumberOfStopsAway          int64                                                                             `json:"numberOfStopsAway"`
	OccupancyStatus            string                                                                            `json:"occupancyStatus"`
	Predicted                  bool                                                                              `json:"predicted"`
	PredictedArrivalInterval   string                                                                            `json:"predictedArrivalInterval"`
	PredictedArrivalTime       int64                                                                             `json:"predictedArrivalTime"`
	PredictedDepartureInterval string                                                                            `json:"predictedDepartureInterval"`
	PredictedDepartureTime     int64                                                                             `json:"predictedDepartureTime"`
	PredictedOccupancy         string                                                                            `json:"predictedOccupancy"`
	RouteID                    string                                                                            `json:"routeId"`
	RouteLongName              string                                                                            `json:"routeLongName"`
	RouteShortName             string                                                                            `json:"routeShortName"`
	ScheduledArrivalInterval   string                                                                            `json:"scheduledArrivalInterval"`
	ScheduledArrivalTime       int64                                                                             `json:"scheduledArrivalTime"`
	ScheduledDepartureInterval string                                                                            `json:"scheduledDepartureInterval"`
	ScheduledDepartureTime     int64                                                                             `json:"scheduledDepartureTime"`
	ScheduledTrack             string                                                                            `json:"scheduledTrack"`
	ServiceDate                int64                                                                             `json:"serviceDate"`
	SituationIDs               []string                                                                          `json:"situationIds"`
	Status                     string                                                                            `json:"status"`
	StopID                     string                                                                            `json:"stopId"`
	StopSequence               int64                                                                             `json:"stopSequence"`
	TotalStopsInTrip           int64                                                                             `json:"totalStopsInTrip"`
	TripHeadsign               string                                                                            `json:"tripHeadsign"`
	TripID                     string                                                                            `json:"tripId"`
	TripStatus                 WhereStopArrivalsAndDepartureListResponseDataEntryArrivalsAndDeparturesTripStatus `json:"tripStatus"`
	VehicleID                  string                                                                            `json:"vehicleId"`
	JSON                       whereStopArrivalsAndDepartureListResponseDataEntryArrivalsAndDepartureJSON        `json:"-"`
}

// whereStopArrivalsAndDepartureListResponseDataEntryArrivalsAndDepartureJSON
// contains the JSON metadata for the struct
// [WhereStopArrivalsAndDepartureListResponseDataEntryArrivalsAndDeparture]
type whereStopArrivalsAndDepartureListResponseDataEntryArrivalsAndDepartureJSON struct {
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

func (r *WhereStopArrivalsAndDepartureListResponseDataEntryArrivalsAndDeparture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r whereStopArrivalsAndDepartureListResponseDataEntryArrivalsAndDepartureJSON) RawJSON() string {
	return r.raw
}

type WhereStopArrivalsAndDepartureListResponseDataEntryArrivalsAndDeparturesTripStatus struct {
	ActiveTripID               string                                                                                             `json:"activeTripId"`
	BlockTripSequence          int64                                                                                              `json:"blockTripSequence"`
	ClosestStop                string                                                                                             `json:"closestStop"`
	ClosestStopTimeOffset      int64                                                                                              `json:"closestStopTimeOffset"`
	DistanceAlongTrip          float64                                                                                            `json:"distanceAlongTrip"`
	Frequency                  string                                                                                             `json:"frequency"`
	LastKnownDistanceAlongTrip float64                                                                                            `json:"lastKnownDistanceAlongTrip"`
	LastKnownLocation          WhereStopArrivalsAndDepartureListResponseDataEntryArrivalsAndDeparturesTripStatusLastKnownLocation `json:"lastKnownLocation"`
	LastKnownOrientation       float64                                                                                            `json:"lastKnownOrientation"`
	LastLocationUpdateTime     int64                                                                                              `json:"lastLocationUpdateTime"`
	LastUpdateTime             int64                                                                                              `json:"lastUpdateTime"`
	NextStop                   string                                                                                             `json:"nextStop"`
	NextStopTimeOffset         int64                                                                                              `json:"nextStopTimeOffset"`
	OccupancyCapacity          int64                                                                                              `json:"occupancyCapacity"`
	OccupancyCount             int64                                                                                              `json:"occupancyCount"`
	OccupancyStatus            string                                                                                             `json:"occupancyStatus"`
	Orientation                float64                                                                                            `json:"orientation"`
	Phase                      string                                                                                             `json:"phase"`
	Position                   WhereStopArrivalsAndDepartureListResponseDataEntryArrivalsAndDeparturesTripStatusPosition          `json:"position"`
	Predicted                  bool                                                                                               `json:"predicted"`
	ScheduledDistanceAlongTrip float64                                                                                            `json:"scheduledDistanceAlongTrip"`
	ScheduleDeviation          int64                                                                                              `json:"scheduleDeviation"`
	ServiceDate                int64                                                                                              `json:"serviceDate"`
	SituationIDs               []string                                                                                           `json:"situationIds"`
	Status                     string                                                                                             `json:"status"`
	TotalDistanceAlongTrip     float64                                                                                            `json:"totalDistanceAlongTrip"`
	VehicleID                  string                                                                                             `json:"vehicleId"`
	JSON                       whereStopArrivalsAndDepartureListResponseDataEntryArrivalsAndDeparturesTripStatusJSON              `json:"-"`
}

// whereStopArrivalsAndDepartureListResponseDataEntryArrivalsAndDeparturesTripStatusJSON
// contains the JSON metadata for the struct
// [WhereStopArrivalsAndDepartureListResponseDataEntryArrivalsAndDeparturesTripStatus]
type whereStopArrivalsAndDepartureListResponseDataEntryArrivalsAndDeparturesTripStatusJSON struct {
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

func (r *WhereStopArrivalsAndDepartureListResponseDataEntryArrivalsAndDeparturesTripStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r whereStopArrivalsAndDepartureListResponseDataEntryArrivalsAndDeparturesTripStatusJSON) RawJSON() string {
	return r.raw
}

type WhereStopArrivalsAndDepartureListResponseDataEntryArrivalsAndDeparturesTripStatusLastKnownLocation struct {
	Lat  float64                                                                                                `json:"lat"`
	Lon  float64                                                                                                `json:"lon"`
	JSON whereStopArrivalsAndDepartureListResponseDataEntryArrivalsAndDeparturesTripStatusLastKnownLocationJSON `json:"-"`
}

// whereStopArrivalsAndDepartureListResponseDataEntryArrivalsAndDeparturesTripStatusLastKnownLocationJSON
// contains the JSON metadata for the struct
// [WhereStopArrivalsAndDepartureListResponseDataEntryArrivalsAndDeparturesTripStatusLastKnownLocation]
type whereStopArrivalsAndDepartureListResponseDataEntryArrivalsAndDeparturesTripStatusLastKnownLocationJSON struct {
	Lat         apijson.Field
	Lon         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WhereStopArrivalsAndDepartureListResponseDataEntryArrivalsAndDeparturesTripStatusLastKnownLocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r whereStopArrivalsAndDepartureListResponseDataEntryArrivalsAndDeparturesTripStatusLastKnownLocationJSON) RawJSON() string {
	return r.raw
}

type WhereStopArrivalsAndDepartureListResponseDataEntryArrivalsAndDeparturesTripStatusPosition struct {
	Lat  float64                                                                                       `json:"lat"`
	Lon  float64                                                                                       `json:"lon"`
	JSON whereStopArrivalsAndDepartureListResponseDataEntryArrivalsAndDeparturesTripStatusPositionJSON `json:"-"`
}

// whereStopArrivalsAndDepartureListResponseDataEntryArrivalsAndDeparturesTripStatusPositionJSON
// contains the JSON metadata for the struct
// [WhereStopArrivalsAndDepartureListResponseDataEntryArrivalsAndDeparturesTripStatusPosition]
type whereStopArrivalsAndDepartureListResponseDataEntryArrivalsAndDeparturesTripStatusPositionJSON struct {
	Lat         apijson.Field
	Lon         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WhereStopArrivalsAndDepartureListResponseDataEntryArrivalsAndDeparturesTripStatusPosition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r whereStopArrivalsAndDepartureListResponseDataEntryArrivalsAndDeparturesTripStatusPositionJSON) RawJSON() string {
	return r.raw
}

type WhereStopArrivalsAndDepartureListResponseDataEntryReferences struct {
	Agencies   []WhereStopArrivalsAndDepartureListResponseDataEntryReferencesAgency `json:"agencies"`
	Routes     []WhereStopArrivalsAndDepartureListResponseDataEntryReferencesRoute  `json:"routes"`
	Situations []interface{}                                                        `json:"situations"`
	Stops      []WhereStopArrivalsAndDepartureListResponseDataEntryReferencesStop   `json:"stops"`
	StopTimes  []interface{}                                                        `json:"stopTimes"`
	Trips      []WhereStopArrivalsAndDepartureListResponseDataEntryReferencesTrip   `json:"trips"`
	JSON       whereStopArrivalsAndDepartureListResponseDataEntryReferencesJSON     `json:"-"`
}

// whereStopArrivalsAndDepartureListResponseDataEntryReferencesJSON contains the
// JSON metadata for the struct
// [WhereStopArrivalsAndDepartureListResponseDataEntryReferences]
type whereStopArrivalsAndDepartureListResponseDataEntryReferencesJSON struct {
	Agencies    apijson.Field
	Routes      apijson.Field
	Situations  apijson.Field
	Stops       apijson.Field
	StopTimes   apijson.Field
	Trips       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WhereStopArrivalsAndDepartureListResponseDataEntryReferences) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r whereStopArrivalsAndDepartureListResponseDataEntryReferencesJSON) RawJSON() string {
	return r.raw
}

type WhereStopArrivalsAndDepartureListResponseDataEntryReferencesAgency struct {
	ID             string                                                                 `json:"id,required"`
	Name           string                                                                 `json:"name,required"`
	Timezone       string                                                                 `json:"timezone,required"`
	URL            string                                                                 `json:"url,required"`
	Disclaimer     string                                                                 `json:"disclaimer"`
	Email          string                                                                 `json:"email"`
	FareURL        string                                                                 `json:"fareUrl"`
	Lang           string                                                                 `json:"lang"`
	Phone          string                                                                 `json:"phone"`
	PrivateService bool                                                                   `json:"privateService"`
	JSON           whereStopArrivalsAndDepartureListResponseDataEntryReferencesAgencyJSON `json:"-"`
}

// whereStopArrivalsAndDepartureListResponseDataEntryReferencesAgencyJSON contains
// the JSON metadata for the struct
// [WhereStopArrivalsAndDepartureListResponseDataEntryReferencesAgency]
type whereStopArrivalsAndDepartureListResponseDataEntryReferencesAgencyJSON struct {
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

func (r *WhereStopArrivalsAndDepartureListResponseDataEntryReferencesAgency) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r whereStopArrivalsAndDepartureListResponseDataEntryReferencesAgencyJSON) RawJSON() string {
	return r.raw
}

type WhereStopArrivalsAndDepartureListResponseDataEntryReferencesRoute struct {
	ID                string                                                                `json:"id"`
	AgencyID          string                                                                `json:"agencyId"`
	Color             string                                                                `json:"color"`
	Description       string                                                                `json:"description"`
	LongName          string                                                                `json:"longName"`
	NullSafeShortName string                                                                `json:"nullSafeShortName"`
	ShortName         string                                                                `json:"shortName"`
	TextColor         string                                                                `json:"textColor"`
	Type              int64                                                                 `json:"type"`
	URL               string                                                                `json:"url"`
	JSON              whereStopArrivalsAndDepartureListResponseDataEntryReferencesRouteJSON `json:"-"`
}

// whereStopArrivalsAndDepartureListResponseDataEntryReferencesRouteJSON contains
// the JSON metadata for the struct
// [WhereStopArrivalsAndDepartureListResponseDataEntryReferencesRoute]
type whereStopArrivalsAndDepartureListResponseDataEntryReferencesRouteJSON struct {
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

func (r *WhereStopArrivalsAndDepartureListResponseDataEntryReferencesRoute) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r whereStopArrivalsAndDepartureListResponseDataEntryReferencesRouteJSON) RawJSON() string {
	return r.raw
}

type WhereStopArrivalsAndDepartureListResponseDataEntryReferencesStop struct {
	ID                 string                                                               `json:"id,required"`
	Code               string                                                               `json:"code,required"`
	Lat                float64                                                              `json:"lat,required"`
	Lon                float64                                                              `json:"lon,required"`
	Name               string                                                               `json:"name,required"`
	Direction          string                                                               `json:"direction"`
	LocationType       int64                                                                `json:"locationType"`
	Parent             string                                                               `json:"parent"`
	RouteIDs           []string                                                             `json:"routeIds"`
	StaticRouteIDs     []string                                                             `json:"staticRouteIds"`
	WheelchairBoarding string                                                               `json:"wheelchairBoarding"`
	JSON               whereStopArrivalsAndDepartureListResponseDataEntryReferencesStopJSON `json:"-"`
}

// whereStopArrivalsAndDepartureListResponseDataEntryReferencesStopJSON contains
// the JSON metadata for the struct
// [WhereStopArrivalsAndDepartureListResponseDataEntryReferencesStop]
type whereStopArrivalsAndDepartureListResponseDataEntryReferencesStopJSON struct {
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

func (r *WhereStopArrivalsAndDepartureListResponseDataEntryReferencesStop) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r whereStopArrivalsAndDepartureListResponseDataEntryReferencesStopJSON) RawJSON() string {
	return r.raw
}

type WhereStopArrivalsAndDepartureListResponseDataEntryReferencesTrip struct {
	ID             string                                                               `json:"id,required"`
	RouteID        string                                                               `json:"routeId,required"`
	BlockID        string                                                               `json:"blockId"`
	DirectionID    string                                                               `json:"directionId"`
	PeakOffpeak    int64                                                                `json:"peakOffpeak"`
	RouteShortName string                                                               `json:"routeShortName"`
	ServiceID      string                                                               `json:"serviceId"`
	ShapeID        string                                                               `json:"shapeId"`
	TimeZone       string                                                               `json:"timeZone"`
	TripHeadsign   string                                                               `json:"tripHeadsign"`
	TripShortName  string                                                               `json:"tripShortName"`
	JSON           whereStopArrivalsAndDepartureListResponseDataEntryReferencesTripJSON `json:"-"`
}

// whereStopArrivalsAndDepartureListResponseDataEntryReferencesTripJSON contains
// the JSON metadata for the struct
// [WhereStopArrivalsAndDepartureListResponseDataEntryReferencesTrip]
type whereStopArrivalsAndDepartureListResponseDataEntryReferencesTripJSON struct {
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

func (r *WhereStopArrivalsAndDepartureListResponseDataEntryReferencesTrip) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r whereStopArrivalsAndDepartureListResponseDataEntryReferencesTripJSON) RawJSON() string {
	return r.raw
}
