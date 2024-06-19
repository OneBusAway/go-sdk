// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package onebusaway

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/open-transit-go/internal/apijson"
	"github.com/stainless-sdks/open-transit-go/internal/requestconfig"
	"github.com/stainless-sdks/open-transit-go/option"
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

// arrivals-and-departures-for-stop
func (r *ArrivalsAndDeparturesForStopService) Get(ctx context.Context, stopID string, opts ...option.RequestOption) (res *ArrivalsAndDeparturesForStopGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if stopID == "" {
		err = errors.New("missing required stopID parameter")
		return
	}
	path := fmt.Sprintf("api/where/arrivals-and-departures-for-stop/%s", stopID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ArrivalsAndDeparturesForStopGetResponse struct {
	Code        int64                                       `json:"code,required"`
	CurrentTime int64                                       `json:"currentTime,required"`
	Text        string                                      `json:"text,required"`
	Version     int64                                       `json:"version,required"`
	Data        ArrivalsAndDeparturesForStopGetResponseData `json:"data"`
	JSON        arrivalsAndDeparturesForStopGetResponseJSON `json:"-"`
}

// arrivalsAndDeparturesForStopGetResponseJSON contains the JSON metadata for the
// struct [ArrivalsAndDeparturesForStopGetResponse]
type arrivalsAndDeparturesForStopGetResponseJSON struct {
	Code        apijson.Field
	CurrentTime apijson.Field
	Text        apijson.Field
	Version     apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ArrivalsAndDeparturesForStopGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r arrivalsAndDeparturesForStopGetResponseJSON) RawJSON() string {
	return r.raw
}

type ArrivalsAndDeparturesForStopGetResponseData struct {
	Entry ArrivalsAndDeparturesForStopGetResponseDataEntry `json:"entry"`
	JSON  arrivalsAndDeparturesForStopGetResponseDataJSON  `json:"-"`
}

// arrivalsAndDeparturesForStopGetResponseDataJSON contains the JSON metadata for
// the struct [ArrivalsAndDeparturesForStopGetResponseData]
type arrivalsAndDeparturesForStopGetResponseDataJSON struct {
	Entry       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ArrivalsAndDeparturesForStopGetResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r arrivalsAndDeparturesForStopGetResponseDataJSON) RawJSON() string {
	return r.raw
}

type ArrivalsAndDeparturesForStopGetResponseDataEntry struct {
	ArrivalsAndDepartures []ArrivalsAndDeparturesForStopGetResponseDataEntryArrivalsAndDeparture `json:"arrivalsAndDepartures"`
	References            ArrivalsAndDeparturesForStopGetResponseDataEntryReferences             `json:"references"`
	JSON                  arrivalsAndDeparturesForStopGetResponseDataEntryJSON                   `json:"-"`
}

// arrivalsAndDeparturesForStopGetResponseDataEntryJSON contains the JSON metadata
// for the struct [ArrivalsAndDeparturesForStopGetResponseDataEntry]
type arrivalsAndDeparturesForStopGetResponseDataEntryJSON struct {
	ArrivalsAndDepartures apijson.Field
	References            apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *ArrivalsAndDeparturesForStopGetResponseDataEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r arrivalsAndDeparturesForStopGetResponseDataEntryJSON) RawJSON() string {
	return r.raw
}

type ArrivalsAndDeparturesForStopGetResponseDataEntryArrivalsAndDeparture struct {
	ActualTrack                string                                                                          `json:"actualTrack"`
	ArrivalEnabled             bool                                                                            `json:"arrivalEnabled"`
	BlockTripSequence          int64                                                                           `json:"blockTripSequence"`
	DepartureEnabled           bool                                                                            `json:"departureEnabled"`
	DistanceFromStop           float64                                                                         `json:"distanceFromStop"`
	Frequency                  string                                                                          `json:"frequency"`
	HistoricalOccupancy        string                                                                          `json:"historicalOccupancy"`
	LastUpdateTime             int64                                                                           `json:"lastUpdateTime"`
	NumberOfStopsAway          int64                                                                           `json:"numberOfStopsAway"`
	OccupancyStatus            string                                                                          `json:"occupancyStatus"`
	Predicted                  bool                                                                            `json:"predicted"`
	PredictedArrivalInterval   string                                                                          `json:"predictedArrivalInterval"`
	PredictedArrivalTime       int64                                                                           `json:"predictedArrivalTime"`
	PredictedDepartureInterval string                                                                          `json:"predictedDepartureInterval"`
	PredictedDepartureTime     int64                                                                           `json:"predictedDepartureTime"`
	PredictedOccupancy         string                                                                          `json:"predictedOccupancy"`
	RouteID                    string                                                                          `json:"routeId"`
	RouteLongName              string                                                                          `json:"routeLongName"`
	RouteShortName             string                                                                          `json:"routeShortName"`
	ScheduledArrivalInterval   string                                                                          `json:"scheduledArrivalInterval"`
	ScheduledArrivalTime       int64                                                                           `json:"scheduledArrivalTime"`
	ScheduledDepartureInterval string                                                                          `json:"scheduledDepartureInterval"`
	ScheduledDepartureTime     int64                                                                           `json:"scheduledDepartureTime"`
	ScheduledTrack             string                                                                          `json:"scheduledTrack"`
	ServiceDate                int64                                                                           `json:"serviceDate"`
	SituationIDs               []string                                                                        `json:"situationIds"`
	Status                     string                                                                          `json:"status"`
	StopID                     string                                                                          `json:"stopId"`
	StopSequence               int64                                                                           `json:"stopSequence"`
	TotalStopsInTrip           int64                                                                           `json:"totalStopsInTrip"`
	TripHeadsign               string                                                                          `json:"tripHeadsign"`
	TripID                     string                                                                          `json:"tripId"`
	TripStatus                 ArrivalsAndDeparturesForStopGetResponseDataEntryArrivalsAndDeparturesTripStatus `json:"tripStatus"`
	VehicleID                  string                                                                          `json:"vehicleId"`
	JSON                       arrivalsAndDeparturesForStopGetResponseDataEntryArrivalsAndDepartureJSON        `json:"-"`
}

// arrivalsAndDeparturesForStopGetResponseDataEntryArrivalsAndDepartureJSON
// contains the JSON metadata for the struct
// [ArrivalsAndDeparturesForStopGetResponseDataEntryArrivalsAndDeparture]
type arrivalsAndDeparturesForStopGetResponseDataEntryArrivalsAndDepartureJSON struct {
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

func (r *ArrivalsAndDeparturesForStopGetResponseDataEntryArrivalsAndDeparture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r arrivalsAndDeparturesForStopGetResponseDataEntryArrivalsAndDepartureJSON) RawJSON() string {
	return r.raw
}

type ArrivalsAndDeparturesForStopGetResponseDataEntryArrivalsAndDeparturesTripStatus struct {
	ActiveTripID               string                                                                                           `json:"activeTripId"`
	BlockTripSequence          int64                                                                                            `json:"blockTripSequence"`
	ClosestStop                string                                                                                           `json:"closestStop"`
	ClosestStopTimeOffset      int64                                                                                            `json:"closestStopTimeOffset"`
	DistanceAlongTrip          float64                                                                                          `json:"distanceAlongTrip"`
	Frequency                  string                                                                                           `json:"frequency"`
	LastKnownDistanceAlongTrip float64                                                                                          `json:"lastKnownDistanceAlongTrip"`
	LastKnownLocation          ArrivalsAndDeparturesForStopGetResponseDataEntryArrivalsAndDeparturesTripStatusLastKnownLocation `json:"lastKnownLocation"`
	LastKnownOrientation       float64                                                                                          `json:"lastKnownOrientation"`
	LastLocationUpdateTime     int64                                                                                            `json:"lastLocationUpdateTime"`
	LastUpdateTime             int64                                                                                            `json:"lastUpdateTime"`
	NextStop                   string                                                                                           `json:"nextStop"`
	NextStopTimeOffset         int64                                                                                            `json:"nextStopTimeOffset"`
	OccupancyCapacity          int64                                                                                            `json:"occupancyCapacity"`
	OccupancyCount             int64                                                                                            `json:"occupancyCount"`
	OccupancyStatus            string                                                                                           `json:"occupancyStatus"`
	Orientation                float64                                                                                          `json:"orientation"`
	Phase                      string                                                                                           `json:"phase"`
	Position                   ArrivalsAndDeparturesForStopGetResponseDataEntryArrivalsAndDeparturesTripStatusPosition          `json:"position"`
	Predicted                  bool                                                                                             `json:"predicted"`
	ScheduledDistanceAlongTrip float64                                                                                          `json:"scheduledDistanceAlongTrip"`
	ScheduleDeviation          int64                                                                                            `json:"scheduleDeviation"`
	ServiceDate                int64                                                                                            `json:"serviceDate"`
	SituationIDs               []string                                                                                         `json:"situationIds"`
	Status                     string                                                                                           `json:"status"`
	TotalDistanceAlongTrip     float64                                                                                          `json:"totalDistanceAlongTrip"`
	VehicleID                  string                                                                                           `json:"vehicleId"`
	JSON                       arrivalsAndDeparturesForStopGetResponseDataEntryArrivalsAndDeparturesTripStatusJSON              `json:"-"`
}

// arrivalsAndDeparturesForStopGetResponseDataEntryArrivalsAndDeparturesTripStatusJSON
// contains the JSON metadata for the struct
// [ArrivalsAndDeparturesForStopGetResponseDataEntryArrivalsAndDeparturesTripStatus]
type arrivalsAndDeparturesForStopGetResponseDataEntryArrivalsAndDeparturesTripStatusJSON struct {
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

func (r *ArrivalsAndDeparturesForStopGetResponseDataEntryArrivalsAndDeparturesTripStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r arrivalsAndDeparturesForStopGetResponseDataEntryArrivalsAndDeparturesTripStatusJSON) RawJSON() string {
	return r.raw
}

type ArrivalsAndDeparturesForStopGetResponseDataEntryArrivalsAndDeparturesTripStatusLastKnownLocation struct {
	Lat  float64                                                                                              `json:"lat"`
	Lon  float64                                                                                              `json:"lon"`
	JSON arrivalsAndDeparturesForStopGetResponseDataEntryArrivalsAndDeparturesTripStatusLastKnownLocationJSON `json:"-"`
}

// arrivalsAndDeparturesForStopGetResponseDataEntryArrivalsAndDeparturesTripStatusLastKnownLocationJSON
// contains the JSON metadata for the struct
// [ArrivalsAndDeparturesForStopGetResponseDataEntryArrivalsAndDeparturesTripStatusLastKnownLocation]
type arrivalsAndDeparturesForStopGetResponseDataEntryArrivalsAndDeparturesTripStatusLastKnownLocationJSON struct {
	Lat         apijson.Field
	Lon         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ArrivalsAndDeparturesForStopGetResponseDataEntryArrivalsAndDeparturesTripStatusLastKnownLocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r arrivalsAndDeparturesForStopGetResponseDataEntryArrivalsAndDeparturesTripStatusLastKnownLocationJSON) RawJSON() string {
	return r.raw
}

type ArrivalsAndDeparturesForStopGetResponseDataEntryArrivalsAndDeparturesTripStatusPosition struct {
	Lat  float64                                                                                     `json:"lat"`
	Lon  float64                                                                                     `json:"lon"`
	JSON arrivalsAndDeparturesForStopGetResponseDataEntryArrivalsAndDeparturesTripStatusPositionJSON `json:"-"`
}

// arrivalsAndDeparturesForStopGetResponseDataEntryArrivalsAndDeparturesTripStatusPositionJSON
// contains the JSON metadata for the struct
// [ArrivalsAndDeparturesForStopGetResponseDataEntryArrivalsAndDeparturesTripStatusPosition]
type arrivalsAndDeparturesForStopGetResponseDataEntryArrivalsAndDeparturesTripStatusPositionJSON struct {
	Lat         apijson.Field
	Lon         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ArrivalsAndDeparturesForStopGetResponseDataEntryArrivalsAndDeparturesTripStatusPosition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r arrivalsAndDeparturesForStopGetResponseDataEntryArrivalsAndDeparturesTripStatusPositionJSON) RawJSON() string {
	return r.raw
}

type ArrivalsAndDeparturesForStopGetResponseDataEntryReferences struct {
	Agencies   []ArrivalsAndDeparturesForStopGetResponseDataEntryReferencesAgency `json:"agencies"`
	Routes     []ArrivalsAndDeparturesForStopGetResponseDataEntryReferencesRoute  `json:"routes"`
	Situations []interface{}                                                      `json:"situations"`
	Stops      []ArrivalsAndDeparturesForStopGetResponseDataEntryReferencesStop   `json:"stops"`
	StopTimes  []interface{}                                                      `json:"stopTimes"`
	Trips      []ArrivalsAndDeparturesForStopGetResponseDataEntryReferencesTrip   `json:"trips"`
	JSON       arrivalsAndDeparturesForStopGetResponseDataEntryReferencesJSON     `json:"-"`
}

// arrivalsAndDeparturesForStopGetResponseDataEntryReferencesJSON contains the JSON
// metadata for the struct
// [ArrivalsAndDeparturesForStopGetResponseDataEntryReferences]
type arrivalsAndDeparturesForStopGetResponseDataEntryReferencesJSON struct {
	Agencies    apijson.Field
	Routes      apijson.Field
	Situations  apijson.Field
	Stops       apijson.Field
	StopTimes   apijson.Field
	Trips       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ArrivalsAndDeparturesForStopGetResponseDataEntryReferences) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r arrivalsAndDeparturesForStopGetResponseDataEntryReferencesJSON) RawJSON() string {
	return r.raw
}

type ArrivalsAndDeparturesForStopGetResponseDataEntryReferencesAgency struct {
	ID             string                                                               `json:"id,required"`
	Name           string                                                               `json:"name,required"`
	Timezone       string                                                               `json:"timezone,required"`
	URL            string                                                               `json:"url,required"`
	Disclaimer     string                                                               `json:"disclaimer"`
	Email          string                                                               `json:"email"`
	FareURL        string                                                               `json:"fareUrl"`
	Lang           string                                                               `json:"lang"`
	Phone          string                                                               `json:"phone"`
	PrivateService bool                                                                 `json:"privateService"`
	JSON           arrivalsAndDeparturesForStopGetResponseDataEntryReferencesAgencyJSON `json:"-"`
}

// arrivalsAndDeparturesForStopGetResponseDataEntryReferencesAgencyJSON contains
// the JSON metadata for the struct
// [ArrivalsAndDeparturesForStopGetResponseDataEntryReferencesAgency]
type arrivalsAndDeparturesForStopGetResponseDataEntryReferencesAgencyJSON struct {
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

func (r *ArrivalsAndDeparturesForStopGetResponseDataEntryReferencesAgency) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r arrivalsAndDeparturesForStopGetResponseDataEntryReferencesAgencyJSON) RawJSON() string {
	return r.raw
}

type ArrivalsAndDeparturesForStopGetResponseDataEntryReferencesRoute struct {
	ID                string                                                              `json:"id"`
	AgencyID          string                                                              `json:"agencyId"`
	Color             string                                                              `json:"color"`
	Description       string                                                              `json:"description"`
	LongName          string                                                              `json:"longName"`
	NullSafeShortName string                                                              `json:"nullSafeShortName"`
	ShortName         string                                                              `json:"shortName"`
	TextColor         string                                                              `json:"textColor"`
	Type              int64                                                               `json:"type"`
	URL               string                                                              `json:"url"`
	JSON              arrivalsAndDeparturesForStopGetResponseDataEntryReferencesRouteJSON `json:"-"`
}

// arrivalsAndDeparturesForStopGetResponseDataEntryReferencesRouteJSON contains the
// JSON metadata for the struct
// [ArrivalsAndDeparturesForStopGetResponseDataEntryReferencesRoute]
type arrivalsAndDeparturesForStopGetResponseDataEntryReferencesRouteJSON struct {
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

func (r *ArrivalsAndDeparturesForStopGetResponseDataEntryReferencesRoute) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r arrivalsAndDeparturesForStopGetResponseDataEntryReferencesRouteJSON) RawJSON() string {
	return r.raw
}

type ArrivalsAndDeparturesForStopGetResponseDataEntryReferencesStop struct {
	ID                 string                                                             `json:"id,required"`
	Code               string                                                             `json:"code,required"`
	Lat                float64                                                            `json:"lat,required"`
	Lon                float64                                                            `json:"lon,required"`
	Name               string                                                             `json:"name,required"`
	Direction          string                                                             `json:"direction"`
	LocationType       int64                                                              `json:"locationType"`
	Parent             string                                                             `json:"parent"`
	RouteIDs           []string                                                           `json:"routeIds"`
	StaticRouteIDs     []string                                                           `json:"staticRouteIds"`
	WheelchairBoarding string                                                             `json:"wheelchairBoarding"`
	JSON               arrivalsAndDeparturesForStopGetResponseDataEntryReferencesStopJSON `json:"-"`
}

// arrivalsAndDeparturesForStopGetResponseDataEntryReferencesStopJSON contains the
// JSON metadata for the struct
// [ArrivalsAndDeparturesForStopGetResponseDataEntryReferencesStop]
type arrivalsAndDeparturesForStopGetResponseDataEntryReferencesStopJSON struct {
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

func (r *ArrivalsAndDeparturesForStopGetResponseDataEntryReferencesStop) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r arrivalsAndDeparturesForStopGetResponseDataEntryReferencesStopJSON) RawJSON() string {
	return r.raw
}

type ArrivalsAndDeparturesForStopGetResponseDataEntryReferencesTrip struct {
	ID             string                                                             `json:"id,required"`
	RouteID        string                                                             `json:"routeId,required"`
	BlockID        string                                                             `json:"blockId"`
	DirectionID    string                                                             `json:"directionId"`
	PeakOffpeak    int64                                                              `json:"peakOffpeak"`
	RouteShortName string                                                             `json:"routeShortName"`
	ServiceID      string                                                             `json:"serviceId"`
	ShapeID        string                                                             `json:"shapeId"`
	TimeZone       string                                                             `json:"timeZone"`
	TripHeadsign   string                                                             `json:"tripHeadsign"`
	TripShortName  string                                                             `json:"tripShortName"`
	JSON           arrivalsAndDeparturesForStopGetResponseDataEntryReferencesTripJSON `json:"-"`
}

// arrivalsAndDeparturesForStopGetResponseDataEntryReferencesTripJSON contains the
// JSON metadata for the struct
// [ArrivalsAndDeparturesForStopGetResponseDataEntryReferencesTrip]
type arrivalsAndDeparturesForStopGetResponseDataEntryReferencesTripJSON struct {
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

func (r *ArrivalsAndDeparturesForStopGetResponseDataEntryReferencesTrip) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r arrivalsAndDeparturesForStopGetResponseDataEntryReferencesTripJSON) RawJSON() string {
	return r.raw
}
