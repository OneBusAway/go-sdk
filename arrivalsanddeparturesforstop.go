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
	"github.com/stainless-sdks/open-transit-go/shared"
)

// ArrivalsAndDeparturesForStopService contains methods and other services that
// help with interacting with the OneBusAway API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewArrivalsAndDeparturesForStopService] method instead.
type ArrivalsAndDeparturesForStopService struct {
	Options []option.RequestOption
}

// NewArrivalsAndDeparturesForStopService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewArrivalsAndDeparturesForStopService(opts ...option.RequestOption) (r *ArrivalsAndDeparturesForStopService) {
	r = &ArrivalsAndDeparturesForStopService{}
	r.Options = opts
	return
}

// arrival-and-departure-for-stop
func (r *ArrivalsAndDeparturesForStopService) List(ctx context.Context, stopID string, opts ...option.RequestOption) (res *ArrivalsAndDeparturesForStopListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if stopID == "" {
		err = errors.New("missing required stopID parameter")
		return
	}
	path := fmt.Sprintf("api/where/arrivals-and-departures-for-stop/stopID.json")
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ArrivalsAndDeparturesForStopListResponse struct {
	Data ArrivalsAndDeparturesForStopListResponseData `json:"data"`
	JSON arrivalsAndDeparturesForStopListResponseJSON `json:"-"`
	shared.ResponseWrapper
}

// arrivalsAndDeparturesForStopListResponseJSON contains the JSON metadata for the
// struct [ArrivalsAndDeparturesForStopListResponse]
type arrivalsAndDeparturesForStopListResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ArrivalsAndDeparturesForStopListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r arrivalsAndDeparturesForStopListResponseJSON) RawJSON() string {
	return r.raw
}

type ArrivalsAndDeparturesForStopListResponseData struct {
	Entry ArrivalsAndDeparturesForStopListResponseDataEntry `json:"entry"`
	JSON  arrivalsAndDeparturesForStopListResponseDataJSON  `json:"-"`
}

// arrivalsAndDeparturesForStopListResponseDataJSON contains the JSON metadata for
// the struct [ArrivalsAndDeparturesForStopListResponseData]
type arrivalsAndDeparturesForStopListResponseDataJSON struct {
	Entry       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ArrivalsAndDeparturesForStopListResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r arrivalsAndDeparturesForStopListResponseDataJSON) RawJSON() string {
	return r.raw
}

type ArrivalsAndDeparturesForStopListResponseDataEntry struct {
	ArrivalsAndDepartures []ArrivalsAndDeparturesForStopListResponseDataEntryArrivalsAndDeparture `json:"arrivalsAndDepartures"`
	References            ArrivalsAndDeparturesForStopListResponseDataEntryReferences             `json:"references"`
	JSON                  arrivalsAndDeparturesForStopListResponseDataEntryJSON                   `json:"-"`
}

// arrivalsAndDeparturesForStopListResponseDataEntryJSON contains the JSON metadata
// for the struct [ArrivalsAndDeparturesForStopListResponseDataEntry]
type arrivalsAndDeparturesForStopListResponseDataEntryJSON struct {
	ArrivalsAndDepartures apijson.Field
	References            apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *ArrivalsAndDeparturesForStopListResponseDataEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r arrivalsAndDeparturesForStopListResponseDataEntryJSON) RawJSON() string {
	return r.raw
}

type ArrivalsAndDeparturesForStopListResponseDataEntryArrivalsAndDeparture struct {
	ActualTrack                string                                                                           `json:"actualTrack"`
	ArrivalEnabled             bool                                                                             `json:"arrivalEnabled"`
	BlockTripSequence          int64                                                                            `json:"blockTripSequence"`
	DepartureEnabled           bool                                                                             `json:"departureEnabled"`
	DistanceFromStop           float64                                                                          `json:"distanceFromStop"`
	Frequency                  string                                                                           `json:"frequency"`
	HistoricalOccupancy        string                                                                           `json:"historicalOccupancy"`
	LastUpdateTime             int64                                                                            `json:"lastUpdateTime"`
	NumberOfStopsAway          int64                                                                            `json:"numberOfStopsAway"`
	OccupancyStatus            string                                                                           `json:"occupancyStatus"`
	Predicted                  bool                                                                             `json:"predicted"`
	PredictedArrivalInterval   string                                                                           `json:"predictedArrivalInterval"`
	PredictedArrivalTime       int64                                                                            `json:"predictedArrivalTime"`
	PredictedDepartureInterval string                                                                           `json:"predictedDepartureInterval"`
	PredictedDepartureTime     int64                                                                            `json:"predictedDepartureTime"`
	PredictedOccupancy         string                                                                           `json:"predictedOccupancy"`
	RouteID                    string                                                                           `json:"routeId"`
	RouteLongName              string                                                                           `json:"routeLongName"`
	RouteShortName             string                                                                           `json:"routeShortName"`
	ScheduledArrivalInterval   string                                                                           `json:"scheduledArrivalInterval"`
	ScheduledArrivalTime       int64                                                                            `json:"scheduledArrivalTime"`
	ScheduledDepartureInterval string                                                                           `json:"scheduledDepartureInterval"`
	ScheduledDepartureTime     int64                                                                            `json:"scheduledDepartureTime"`
	ScheduledTrack             string                                                                           `json:"scheduledTrack"`
	ServiceDate                int64                                                                            `json:"serviceDate"`
	SituationIDs               []string                                                                         `json:"situationIds"`
	Status                     string                                                                           `json:"status"`
	StopID                     string                                                                           `json:"stopId"`
	StopSequence               int64                                                                            `json:"stopSequence"`
	TotalStopsInTrip           int64                                                                            `json:"totalStopsInTrip"`
	TripHeadsign               string                                                                           `json:"tripHeadsign"`
	TripID                     string                                                                           `json:"tripId"`
	TripStatus                 ArrivalsAndDeparturesForStopListResponseDataEntryArrivalsAndDeparturesTripStatus `json:"tripStatus"`
	VehicleID                  string                                                                           `json:"vehicleId"`
	JSON                       arrivalsAndDeparturesForStopListResponseDataEntryArrivalsAndDepartureJSON        `json:"-"`
}

// arrivalsAndDeparturesForStopListResponseDataEntryArrivalsAndDepartureJSON
// contains the JSON metadata for the struct
// [ArrivalsAndDeparturesForStopListResponseDataEntryArrivalsAndDeparture]
type arrivalsAndDeparturesForStopListResponseDataEntryArrivalsAndDepartureJSON struct {
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

func (r *ArrivalsAndDeparturesForStopListResponseDataEntryArrivalsAndDeparture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r arrivalsAndDeparturesForStopListResponseDataEntryArrivalsAndDepartureJSON) RawJSON() string {
	return r.raw
}

type ArrivalsAndDeparturesForStopListResponseDataEntryArrivalsAndDeparturesTripStatus struct {
	ActiveTripID               string                                                                                            `json:"activeTripId"`
	BlockTripSequence          int64                                                                                             `json:"blockTripSequence"`
	ClosestStop                string                                                                                            `json:"closestStop"`
	ClosestStopTimeOffset      int64                                                                                             `json:"closestStopTimeOffset"`
	DistanceAlongTrip          float64                                                                                           `json:"distanceAlongTrip"`
	Frequency                  string                                                                                            `json:"frequency"`
	LastKnownDistanceAlongTrip float64                                                                                           `json:"lastKnownDistanceAlongTrip"`
	LastKnownLocation          ArrivalsAndDeparturesForStopListResponseDataEntryArrivalsAndDeparturesTripStatusLastKnownLocation `json:"lastKnownLocation"`
	LastKnownOrientation       float64                                                                                           `json:"lastKnownOrientation"`
	LastLocationUpdateTime     int64                                                                                             `json:"lastLocationUpdateTime"`
	LastUpdateTime             int64                                                                                             `json:"lastUpdateTime"`
	NextStop                   string                                                                                            `json:"nextStop"`
	NextStopTimeOffset         int64                                                                                             `json:"nextStopTimeOffset"`
	OccupancyCapacity          int64                                                                                             `json:"occupancyCapacity"`
	OccupancyCount             int64                                                                                             `json:"occupancyCount"`
	OccupancyStatus            string                                                                                            `json:"occupancyStatus"`
	Orientation                float64                                                                                           `json:"orientation"`
	Phase                      string                                                                                            `json:"phase"`
	Position                   ArrivalsAndDeparturesForStopListResponseDataEntryArrivalsAndDeparturesTripStatusPosition          `json:"position"`
	Predicted                  bool                                                                                              `json:"predicted"`
	ScheduledDistanceAlongTrip float64                                                                                           `json:"scheduledDistanceAlongTrip"`
	ScheduleDeviation          int64                                                                                             `json:"scheduleDeviation"`
	ServiceDate                int64                                                                                             `json:"serviceDate"`
	SituationIDs               []string                                                                                          `json:"situationIds"`
	Status                     string                                                                                            `json:"status"`
	TotalDistanceAlongTrip     float64                                                                                           `json:"totalDistanceAlongTrip"`
	VehicleID                  string                                                                                            `json:"vehicleId"`
	JSON                       arrivalsAndDeparturesForStopListResponseDataEntryArrivalsAndDeparturesTripStatusJSON              `json:"-"`
}

// arrivalsAndDeparturesForStopListResponseDataEntryArrivalsAndDeparturesTripStatusJSON
// contains the JSON metadata for the struct
// [ArrivalsAndDeparturesForStopListResponseDataEntryArrivalsAndDeparturesTripStatus]
type arrivalsAndDeparturesForStopListResponseDataEntryArrivalsAndDeparturesTripStatusJSON struct {
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

func (r *ArrivalsAndDeparturesForStopListResponseDataEntryArrivalsAndDeparturesTripStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r arrivalsAndDeparturesForStopListResponseDataEntryArrivalsAndDeparturesTripStatusJSON) RawJSON() string {
	return r.raw
}

type ArrivalsAndDeparturesForStopListResponseDataEntryArrivalsAndDeparturesTripStatusLastKnownLocation struct {
	Lat  float64                                                                                               `json:"lat"`
	Lon  float64                                                                                               `json:"lon"`
	JSON arrivalsAndDeparturesForStopListResponseDataEntryArrivalsAndDeparturesTripStatusLastKnownLocationJSON `json:"-"`
}

// arrivalsAndDeparturesForStopListResponseDataEntryArrivalsAndDeparturesTripStatusLastKnownLocationJSON
// contains the JSON metadata for the struct
// [ArrivalsAndDeparturesForStopListResponseDataEntryArrivalsAndDeparturesTripStatusLastKnownLocation]
type arrivalsAndDeparturesForStopListResponseDataEntryArrivalsAndDeparturesTripStatusLastKnownLocationJSON struct {
	Lat         apijson.Field
	Lon         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ArrivalsAndDeparturesForStopListResponseDataEntryArrivalsAndDeparturesTripStatusLastKnownLocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r arrivalsAndDeparturesForStopListResponseDataEntryArrivalsAndDeparturesTripStatusLastKnownLocationJSON) RawJSON() string {
	return r.raw
}

type ArrivalsAndDeparturesForStopListResponseDataEntryArrivalsAndDeparturesTripStatusPosition struct {
	Lat  float64                                                                                      `json:"lat"`
	Lon  float64                                                                                      `json:"lon"`
	JSON arrivalsAndDeparturesForStopListResponseDataEntryArrivalsAndDeparturesTripStatusPositionJSON `json:"-"`
}

// arrivalsAndDeparturesForStopListResponseDataEntryArrivalsAndDeparturesTripStatusPositionJSON
// contains the JSON metadata for the struct
// [ArrivalsAndDeparturesForStopListResponseDataEntryArrivalsAndDeparturesTripStatusPosition]
type arrivalsAndDeparturesForStopListResponseDataEntryArrivalsAndDeparturesTripStatusPositionJSON struct {
	Lat         apijson.Field
	Lon         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ArrivalsAndDeparturesForStopListResponseDataEntryArrivalsAndDeparturesTripStatusPosition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r arrivalsAndDeparturesForStopListResponseDataEntryArrivalsAndDeparturesTripStatusPositionJSON) RawJSON() string {
	return r.raw
}

type ArrivalsAndDeparturesForStopListResponseDataEntryReferences struct {
	Agencies   []ArrivalsAndDeparturesForStopListResponseDataEntryReferencesAgency `json:"agencies"`
	Routes     []ArrivalsAndDeparturesForStopListResponseDataEntryReferencesRoute  `json:"routes"`
	Situations []interface{}                                                       `json:"situations"`
	Stops      []ArrivalsAndDeparturesForStopListResponseDataEntryReferencesStop   `json:"stops"`
	StopTimes  []interface{}                                                       `json:"stopTimes"`
	Trips      []ArrivalsAndDeparturesForStopListResponseDataEntryReferencesTrip   `json:"trips"`
	JSON       arrivalsAndDeparturesForStopListResponseDataEntryReferencesJSON     `json:"-"`
}

// arrivalsAndDeparturesForStopListResponseDataEntryReferencesJSON contains the
// JSON metadata for the struct
// [ArrivalsAndDeparturesForStopListResponseDataEntryReferences]
type arrivalsAndDeparturesForStopListResponseDataEntryReferencesJSON struct {
	Agencies    apijson.Field
	Routes      apijson.Field
	Situations  apijson.Field
	Stops       apijson.Field
	StopTimes   apijson.Field
	Trips       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ArrivalsAndDeparturesForStopListResponseDataEntryReferences) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r arrivalsAndDeparturesForStopListResponseDataEntryReferencesJSON) RawJSON() string {
	return r.raw
}

type ArrivalsAndDeparturesForStopListResponseDataEntryReferencesAgency struct {
	ID             string                                                                `json:"id,required"`
	Name           string                                                                `json:"name,required"`
	Timezone       string                                                                `json:"timezone,required"`
	URL            string                                                                `json:"url,required"`
	Disclaimer     string                                                                `json:"disclaimer"`
	Email          string                                                                `json:"email"`
	FareURL        string                                                                `json:"fareUrl"`
	Lang           string                                                                `json:"lang"`
	Phone          string                                                                `json:"phone"`
	PrivateService bool                                                                  `json:"privateService"`
	JSON           arrivalsAndDeparturesForStopListResponseDataEntryReferencesAgencyJSON `json:"-"`
}

// arrivalsAndDeparturesForStopListResponseDataEntryReferencesAgencyJSON contains
// the JSON metadata for the struct
// [ArrivalsAndDeparturesForStopListResponseDataEntryReferencesAgency]
type arrivalsAndDeparturesForStopListResponseDataEntryReferencesAgencyJSON struct {
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

func (r *ArrivalsAndDeparturesForStopListResponseDataEntryReferencesAgency) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r arrivalsAndDeparturesForStopListResponseDataEntryReferencesAgencyJSON) RawJSON() string {
	return r.raw
}

type ArrivalsAndDeparturesForStopListResponseDataEntryReferencesRoute struct {
	ID                string                                                               `json:"id"`
	AgencyID          string                                                               `json:"agencyId"`
	Color             string                                                               `json:"color"`
	Description       string                                                               `json:"description"`
	LongName          string                                                               `json:"longName"`
	NullSafeShortName string                                                               `json:"nullSafeShortName"`
	ShortName         string                                                               `json:"shortName"`
	TextColor         string                                                               `json:"textColor"`
	Type              int64                                                                `json:"type"`
	URL               string                                                               `json:"url"`
	JSON              arrivalsAndDeparturesForStopListResponseDataEntryReferencesRouteJSON `json:"-"`
}

// arrivalsAndDeparturesForStopListResponseDataEntryReferencesRouteJSON contains
// the JSON metadata for the struct
// [ArrivalsAndDeparturesForStopListResponseDataEntryReferencesRoute]
type arrivalsAndDeparturesForStopListResponseDataEntryReferencesRouteJSON struct {
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

func (r *ArrivalsAndDeparturesForStopListResponseDataEntryReferencesRoute) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r arrivalsAndDeparturesForStopListResponseDataEntryReferencesRouteJSON) RawJSON() string {
	return r.raw
}

type ArrivalsAndDeparturesForStopListResponseDataEntryReferencesStop struct {
	ID                 string                                                              `json:"id,required"`
	Code               string                                                              `json:"code,required"`
	Lat                float64                                                             `json:"lat,required"`
	Lon                float64                                                             `json:"lon,required"`
	Name               string                                                              `json:"name,required"`
	Direction          string                                                              `json:"direction"`
	LocationType       int64                                                               `json:"locationType"`
	Parent             string                                                              `json:"parent"`
	RouteIDs           []string                                                            `json:"routeIds"`
	StaticRouteIDs     []string                                                            `json:"staticRouteIds"`
	WheelchairBoarding string                                                              `json:"wheelchairBoarding"`
	JSON               arrivalsAndDeparturesForStopListResponseDataEntryReferencesStopJSON `json:"-"`
}

// arrivalsAndDeparturesForStopListResponseDataEntryReferencesStopJSON contains the
// JSON metadata for the struct
// [ArrivalsAndDeparturesForStopListResponseDataEntryReferencesStop]
type arrivalsAndDeparturesForStopListResponseDataEntryReferencesStopJSON struct {
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

func (r *ArrivalsAndDeparturesForStopListResponseDataEntryReferencesStop) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r arrivalsAndDeparturesForStopListResponseDataEntryReferencesStopJSON) RawJSON() string {
	return r.raw
}

type ArrivalsAndDeparturesForStopListResponseDataEntryReferencesTrip struct {
	ID             string                                                              `json:"id,required"`
	RouteID        string                                                              `json:"routeId,required"`
	BlockID        string                                                              `json:"blockId"`
	DirectionID    string                                                              `json:"directionId"`
	PeakOffpeak    int64                                                               `json:"peakOffpeak"`
	RouteShortName string                                                              `json:"routeShortName"`
	ServiceID      string                                                              `json:"serviceId"`
	ShapeID        string                                                              `json:"shapeId"`
	TimeZone       string                                                              `json:"timeZone"`
	TripHeadsign   string                                                              `json:"tripHeadsign"`
	TripShortName  string                                                              `json:"tripShortName"`
	JSON           arrivalsAndDeparturesForStopListResponseDataEntryReferencesTripJSON `json:"-"`
}

// arrivalsAndDeparturesForStopListResponseDataEntryReferencesTripJSON contains the
// JSON metadata for the struct
// [ArrivalsAndDeparturesForStopListResponseDataEntryReferencesTrip]
type arrivalsAndDeparturesForStopListResponseDataEntryReferencesTripJSON struct {
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

func (r *ArrivalsAndDeparturesForStopListResponseDataEntryReferencesTrip) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r arrivalsAndDeparturesForStopListResponseDataEntryReferencesTripJSON) RawJSON() string {
	return r.raw
}
