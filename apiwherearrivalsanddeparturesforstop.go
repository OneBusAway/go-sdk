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

// APIWhereArrivalsAndDeparturesForStopService contains methods and other services
// that help with interacting with the OneBusAway API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAPIWhereArrivalsAndDeparturesForStopService] method instead.
type APIWhereArrivalsAndDeparturesForStopService struct {
	Options []option.RequestOption
}

// NewAPIWhereArrivalsAndDeparturesForStopService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewAPIWhereArrivalsAndDeparturesForStopService(opts ...option.RequestOption) (r *APIWhereArrivalsAndDeparturesForStopService) {
	r = &APIWhereArrivalsAndDeparturesForStopService{}
	r.Options = opts
	return
}

// arrival-and-departure-for-stop
func (r *APIWhereArrivalsAndDeparturesForStopService) List(ctx context.Context, stopID string, opts ...option.RequestOption) (res *APIWhereArrivalsAndDeparturesForStopListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if stopID == "" {
		err = errors.New("missing required stopID parameter")
		return
	}
	path := fmt.Sprintf("api/where/arrivals-and-departures-for-stop/stopID.json")
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type APIWhereArrivalsAndDeparturesForStopListResponse struct {
	Data APIWhereArrivalsAndDeparturesForStopListResponseData `json:"data"`
	JSON apiWhereArrivalsAndDeparturesForStopListResponseJSON `json:"-"`
	shared.ResponseWrapper
}

// apiWhereArrivalsAndDeparturesForStopListResponseJSON contains the JSON metadata
// for the struct [APIWhereArrivalsAndDeparturesForStopListResponse]
type apiWhereArrivalsAndDeparturesForStopListResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIWhereArrivalsAndDeparturesForStopListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiWhereArrivalsAndDeparturesForStopListResponseJSON) RawJSON() string {
	return r.raw
}

type APIWhereArrivalsAndDeparturesForStopListResponseData struct {
	Entry APIWhereArrivalsAndDeparturesForStopListResponseDataEntry `json:"entry"`
	JSON  apiWhereArrivalsAndDeparturesForStopListResponseDataJSON  `json:"-"`
}

// apiWhereArrivalsAndDeparturesForStopListResponseDataJSON contains the JSON
// metadata for the struct [APIWhereArrivalsAndDeparturesForStopListResponseData]
type apiWhereArrivalsAndDeparturesForStopListResponseDataJSON struct {
	Entry       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIWhereArrivalsAndDeparturesForStopListResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiWhereArrivalsAndDeparturesForStopListResponseDataJSON) RawJSON() string {
	return r.raw
}

type APIWhereArrivalsAndDeparturesForStopListResponseDataEntry struct {
	ArrivalsAndDepartures []APIWhereArrivalsAndDeparturesForStopListResponseDataEntryArrivalsAndDeparture `json:"arrivalsAndDepartures"`
	References            APIWhereArrivalsAndDeparturesForStopListResponseDataEntryReferences             `json:"references"`
	JSON                  apiWhereArrivalsAndDeparturesForStopListResponseDataEntryJSON                   `json:"-"`
}

// apiWhereArrivalsAndDeparturesForStopListResponseDataEntryJSON contains the JSON
// metadata for the struct
// [APIWhereArrivalsAndDeparturesForStopListResponseDataEntry]
type apiWhereArrivalsAndDeparturesForStopListResponseDataEntryJSON struct {
	ArrivalsAndDepartures apijson.Field
	References            apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *APIWhereArrivalsAndDeparturesForStopListResponseDataEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiWhereArrivalsAndDeparturesForStopListResponseDataEntryJSON) RawJSON() string {
	return r.raw
}

type APIWhereArrivalsAndDeparturesForStopListResponseDataEntryArrivalsAndDeparture struct {
	ActualTrack                string                                                                                   `json:"actualTrack"`
	ArrivalEnabled             bool                                                                                     `json:"arrivalEnabled"`
	BlockTripSequence          int64                                                                                    `json:"blockTripSequence"`
	DepartureEnabled           bool                                                                                     `json:"departureEnabled"`
	DistanceFromStop           float64                                                                                  `json:"distanceFromStop"`
	Frequency                  string                                                                                   `json:"frequency"`
	HistoricalOccupancy        string                                                                                   `json:"historicalOccupancy"`
	LastUpdateTime             int64                                                                                    `json:"lastUpdateTime"`
	NumberOfStopsAway          int64                                                                                    `json:"numberOfStopsAway"`
	OccupancyStatus            string                                                                                   `json:"occupancyStatus"`
	Predicted                  bool                                                                                     `json:"predicted"`
	PredictedArrivalInterval   string                                                                                   `json:"predictedArrivalInterval"`
	PredictedArrivalTime       int64                                                                                    `json:"predictedArrivalTime"`
	PredictedDepartureInterval string                                                                                   `json:"predictedDepartureInterval"`
	PredictedDepartureTime     int64                                                                                    `json:"predictedDepartureTime"`
	PredictedOccupancy         string                                                                                   `json:"predictedOccupancy"`
	RouteID                    string                                                                                   `json:"routeId"`
	RouteLongName              string                                                                                   `json:"routeLongName"`
	RouteShortName             string                                                                                   `json:"routeShortName"`
	ScheduledArrivalInterval   string                                                                                   `json:"scheduledArrivalInterval"`
	ScheduledArrivalTime       int64                                                                                    `json:"scheduledArrivalTime"`
	ScheduledDepartureInterval string                                                                                   `json:"scheduledDepartureInterval"`
	ScheduledDepartureTime     int64                                                                                    `json:"scheduledDepartureTime"`
	ScheduledTrack             string                                                                                   `json:"scheduledTrack"`
	ServiceDate                int64                                                                                    `json:"serviceDate"`
	SituationIDs               []string                                                                                 `json:"situationIds"`
	Status                     string                                                                                   `json:"status"`
	StopID                     string                                                                                   `json:"stopId"`
	StopSequence               int64                                                                                    `json:"stopSequence"`
	TotalStopsInTrip           int64                                                                                    `json:"totalStopsInTrip"`
	TripHeadsign               string                                                                                   `json:"tripHeadsign"`
	TripID                     string                                                                                   `json:"tripId"`
	TripStatus                 APIWhereArrivalsAndDeparturesForStopListResponseDataEntryArrivalsAndDeparturesTripStatus `json:"tripStatus"`
	VehicleID                  string                                                                                   `json:"vehicleId"`
	JSON                       apiWhereArrivalsAndDeparturesForStopListResponseDataEntryArrivalsAndDepartureJSON        `json:"-"`
}

// apiWhereArrivalsAndDeparturesForStopListResponseDataEntryArrivalsAndDepartureJSON
// contains the JSON metadata for the struct
// [APIWhereArrivalsAndDeparturesForStopListResponseDataEntryArrivalsAndDeparture]
type apiWhereArrivalsAndDeparturesForStopListResponseDataEntryArrivalsAndDepartureJSON struct {
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

func (r *APIWhereArrivalsAndDeparturesForStopListResponseDataEntryArrivalsAndDeparture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiWhereArrivalsAndDeparturesForStopListResponseDataEntryArrivalsAndDepartureJSON) RawJSON() string {
	return r.raw
}

type APIWhereArrivalsAndDeparturesForStopListResponseDataEntryArrivalsAndDeparturesTripStatus struct {
	ActiveTripID               string                                                                                                    `json:"activeTripId"`
	BlockTripSequence          int64                                                                                                     `json:"blockTripSequence"`
	ClosestStop                string                                                                                                    `json:"closestStop"`
	ClosestStopTimeOffset      int64                                                                                                     `json:"closestStopTimeOffset"`
	DistanceAlongTrip          float64                                                                                                   `json:"distanceAlongTrip"`
	Frequency                  string                                                                                                    `json:"frequency"`
	LastKnownDistanceAlongTrip float64                                                                                                   `json:"lastKnownDistanceAlongTrip"`
	LastKnownLocation          APIWhereArrivalsAndDeparturesForStopListResponseDataEntryArrivalsAndDeparturesTripStatusLastKnownLocation `json:"lastKnownLocation"`
	LastKnownOrientation       float64                                                                                                   `json:"lastKnownOrientation"`
	LastLocationUpdateTime     int64                                                                                                     `json:"lastLocationUpdateTime"`
	LastUpdateTime             int64                                                                                                     `json:"lastUpdateTime"`
	NextStop                   string                                                                                                    `json:"nextStop"`
	NextStopTimeOffset         int64                                                                                                     `json:"nextStopTimeOffset"`
	OccupancyCapacity          int64                                                                                                     `json:"occupancyCapacity"`
	OccupancyCount             int64                                                                                                     `json:"occupancyCount"`
	OccupancyStatus            string                                                                                                    `json:"occupancyStatus"`
	Orientation                float64                                                                                                   `json:"orientation"`
	Phase                      string                                                                                                    `json:"phase"`
	Position                   APIWhereArrivalsAndDeparturesForStopListResponseDataEntryArrivalsAndDeparturesTripStatusPosition          `json:"position"`
	Predicted                  bool                                                                                                      `json:"predicted"`
	ScheduledDistanceAlongTrip float64                                                                                                   `json:"scheduledDistanceAlongTrip"`
	ScheduleDeviation          int64                                                                                                     `json:"scheduleDeviation"`
	ServiceDate                int64                                                                                                     `json:"serviceDate"`
	SituationIDs               []string                                                                                                  `json:"situationIds"`
	Status                     string                                                                                                    `json:"status"`
	TotalDistanceAlongTrip     float64                                                                                                   `json:"totalDistanceAlongTrip"`
	VehicleID                  string                                                                                                    `json:"vehicleId"`
	JSON                       apiWhereArrivalsAndDeparturesForStopListResponseDataEntryArrivalsAndDeparturesTripStatusJSON              `json:"-"`
}

// apiWhereArrivalsAndDeparturesForStopListResponseDataEntryArrivalsAndDeparturesTripStatusJSON
// contains the JSON metadata for the struct
// [APIWhereArrivalsAndDeparturesForStopListResponseDataEntryArrivalsAndDeparturesTripStatus]
type apiWhereArrivalsAndDeparturesForStopListResponseDataEntryArrivalsAndDeparturesTripStatusJSON struct {
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

func (r *APIWhereArrivalsAndDeparturesForStopListResponseDataEntryArrivalsAndDeparturesTripStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiWhereArrivalsAndDeparturesForStopListResponseDataEntryArrivalsAndDeparturesTripStatusJSON) RawJSON() string {
	return r.raw
}

type APIWhereArrivalsAndDeparturesForStopListResponseDataEntryArrivalsAndDeparturesTripStatusLastKnownLocation struct {
	Lat  float64                                                                                                       `json:"lat"`
	Lon  float64                                                                                                       `json:"lon"`
	JSON apiWhereArrivalsAndDeparturesForStopListResponseDataEntryArrivalsAndDeparturesTripStatusLastKnownLocationJSON `json:"-"`
}

// apiWhereArrivalsAndDeparturesForStopListResponseDataEntryArrivalsAndDeparturesTripStatusLastKnownLocationJSON
// contains the JSON metadata for the struct
// [APIWhereArrivalsAndDeparturesForStopListResponseDataEntryArrivalsAndDeparturesTripStatusLastKnownLocation]
type apiWhereArrivalsAndDeparturesForStopListResponseDataEntryArrivalsAndDeparturesTripStatusLastKnownLocationJSON struct {
	Lat         apijson.Field
	Lon         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIWhereArrivalsAndDeparturesForStopListResponseDataEntryArrivalsAndDeparturesTripStatusLastKnownLocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiWhereArrivalsAndDeparturesForStopListResponseDataEntryArrivalsAndDeparturesTripStatusLastKnownLocationJSON) RawJSON() string {
	return r.raw
}

type APIWhereArrivalsAndDeparturesForStopListResponseDataEntryArrivalsAndDeparturesTripStatusPosition struct {
	Lat  float64                                                                                              `json:"lat"`
	Lon  float64                                                                                              `json:"lon"`
	JSON apiWhereArrivalsAndDeparturesForStopListResponseDataEntryArrivalsAndDeparturesTripStatusPositionJSON `json:"-"`
}

// apiWhereArrivalsAndDeparturesForStopListResponseDataEntryArrivalsAndDeparturesTripStatusPositionJSON
// contains the JSON metadata for the struct
// [APIWhereArrivalsAndDeparturesForStopListResponseDataEntryArrivalsAndDeparturesTripStatusPosition]
type apiWhereArrivalsAndDeparturesForStopListResponseDataEntryArrivalsAndDeparturesTripStatusPositionJSON struct {
	Lat         apijson.Field
	Lon         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIWhereArrivalsAndDeparturesForStopListResponseDataEntryArrivalsAndDeparturesTripStatusPosition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiWhereArrivalsAndDeparturesForStopListResponseDataEntryArrivalsAndDeparturesTripStatusPositionJSON) RawJSON() string {
	return r.raw
}

type APIWhereArrivalsAndDeparturesForStopListResponseDataEntryReferences struct {
	Agencies   []APIWhereArrivalsAndDeparturesForStopListResponseDataEntryReferencesAgency `json:"agencies"`
	Routes     []APIWhereArrivalsAndDeparturesForStopListResponseDataEntryReferencesRoute  `json:"routes"`
	Situations []interface{}                                                               `json:"situations"`
	Stops      []APIWhereArrivalsAndDeparturesForStopListResponseDataEntryReferencesStop   `json:"stops"`
	StopTimes  []interface{}                                                               `json:"stopTimes"`
	Trips      []APIWhereArrivalsAndDeparturesForStopListResponseDataEntryReferencesTrip   `json:"trips"`
	JSON       apiWhereArrivalsAndDeparturesForStopListResponseDataEntryReferencesJSON     `json:"-"`
}

// apiWhereArrivalsAndDeparturesForStopListResponseDataEntryReferencesJSON contains
// the JSON metadata for the struct
// [APIWhereArrivalsAndDeparturesForStopListResponseDataEntryReferences]
type apiWhereArrivalsAndDeparturesForStopListResponseDataEntryReferencesJSON struct {
	Agencies    apijson.Field
	Routes      apijson.Field
	Situations  apijson.Field
	Stops       apijson.Field
	StopTimes   apijson.Field
	Trips       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIWhereArrivalsAndDeparturesForStopListResponseDataEntryReferences) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiWhereArrivalsAndDeparturesForStopListResponseDataEntryReferencesJSON) RawJSON() string {
	return r.raw
}

type APIWhereArrivalsAndDeparturesForStopListResponseDataEntryReferencesAgency struct {
	ID             string                                                                        `json:"id,required"`
	Name           string                                                                        `json:"name,required"`
	Timezone       string                                                                        `json:"timezone,required"`
	URL            string                                                                        `json:"url,required"`
	Disclaimer     string                                                                        `json:"disclaimer"`
	Email          string                                                                        `json:"email"`
	FareURL        string                                                                        `json:"fareUrl"`
	Lang           string                                                                        `json:"lang"`
	Phone          string                                                                        `json:"phone"`
	PrivateService bool                                                                          `json:"privateService"`
	JSON           apiWhereArrivalsAndDeparturesForStopListResponseDataEntryReferencesAgencyJSON `json:"-"`
}

// apiWhereArrivalsAndDeparturesForStopListResponseDataEntryReferencesAgencyJSON
// contains the JSON metadata for the struct
// [APIWhereArrivalsAndDeparturesForStopListResponseDataEntryReferencesAgency]
type apiWhereArrivalsAndDeparturesForStopListResponseDataEntryReferencesAgencyJSON struct {
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

func (r *APIWhereArrivalsAndDeparturesForStopListResponseDataEntryReferencesAgency) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiWhereArrivalsAndDeparturesForStopListResponseDataEntryReferencesAgencyJSON) RawJSON() string {
	return r.raw
}

type APIWhereArrivalsAndDeparturesForStopListResponseDataEntryReferencesRoute struct {
	ID                string                                                                       `json:"id"`
	AgencyID          string                                                                       `json:"agencyId"`
	Color             string                                                                       `json:"color"`
	Description       string                                                                       `json:"description"`
	LongName          string                                                                       `json:"longName"`
	NullSafeShortName string                                                                       `json:"nullSafeShortName"`
	ShortName         string                                                                       `json:"shortName"`
	TextColor         string                                                                       `json:"textColor"`
	Type              int64                                                                        `json:"type"`
	URL               string                                                                       `json:"url"`
	JSON              apiWhereArrivalsAndDeparturesForStopListResponseDataEntryReferencesRouteJSON `json:"-"`
}

// apiWhereArrivalsAndDeparturesForStopListResponseDataEntryReferencesRouteJSON
// contains the JSON metadata for the struct
// [APIWhereArrivalsAndDeparturesForStopListResponseDataEntryReferencesRoute]
type apiWhereArrivalsAndDeparturesForStopListResponseDataEntryReferencesRouteJSON struct {
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

func (r *APIWhereArrivalsAndDeparturesForStopListResponseDataEntryReferencesRoute) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiWhereArrivalsAndDeparturesForStopListResponseDataEntryReferencesRouteJSON) RawJSON() string {
	return r.raw
}

type APIWhereArrivalsAndDeparturesForStopListResponseDataEntryReferencesStop struct {
	ID                 string                                                                      `json:"id,required"`
	Code               string                                                                      `json:"code,required"`
	Lat                float64                                                                     `json:"lat,required"`
	Lon                float64                                                                     `json:"lon,required"`
	Name               string                                                                      `json:"name,required"`
	Direction          string                                                                      `json:"direction"`
	LocationType       int64                                                                       `json:"locationType"`
	Parent             string                                                                      `json:"parent"`
	RouteIDs           []string                                                                    `json:"routeIds"`
	StaticRouteIDs     []string                                                                    `json:"staticRouteIds"`
	WheelchairBoarding string                                                                      `json:"wheelchairBoarding"`
	JSON               apiWhereArrivalsAndDeparturesForStopListResponseDataEntryReferencesStopJSON `json:"-"`
}

// apiWhereArrivalsAndDeparturesForStopListResponseDataEntryReferencesStopJSON
// contains the JSON metadata for the struct
// [APIWhereArrivalsAndDeparturesForStopListResponseDataEntryReferencesStop]
type apiWhereArrivalsAndDeparturesForStopListResponseDataEntryReferencesStopJSON struct {
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

func (r *APIWhereArrivalsAndDeparturesForStopListResponseDataEntryReferencesStop) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiWhereArrivalsAndDeparturesForStopListResponseDataEntryReferencesStopJSON) RawJSON() string {
	return r.raw
}

type APIWhereArrivalsAndDeparturesForStopListResponseDataEntryReferencesTrip struct {
	ID             string                                                                      `json:"id,required"`
	RouteID        string                                                                      `json:"routeId,required"`
	BlockID        string                                                                      `json:"blockId"`
	DirectionID    string                                                                      `json:"directionId"`
	PeakOffpeak    int64                                                                       `json:"peakOffpeak"`
	RouteShortName string                                                                      `json:"routeShortName"`
	ServiceID      string                                                                      `json:"serviceId"`
	ShapeID        string                                                                      `json:"shapeId"`
	TimeZone       string                                                                      `json:"timeZone"`
	TripHeadsign   string                                                                      `json:"tripHeadsign"`
	TripShortName  string                                                                      `json:"tripShortName"`
	JSON           apiWhereArrivalsAndDeparturesForStopListResponseDataEntryReferencesTripJSON `json:"-"`
}

// apiWhereArrivalsAndDeparturesForStopListResponseDataEntryReferencesTripJSON
// contains the JSON metadata for the struct
// [APIWhereArrivalsAndDeparturesForStopListResponseDataEntryReferencesTrip]
type apiWhereArrivalsAndDeparturesForStopListResponseDataEntryReferencesTripJSON struct {
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

func (r *APIWhereArrivalsAndDeparturesForStopListResponseDataEntryReferencesTrip) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiWhereArrivalsAndDeparturesForStopListResponseDataEntryReferencesTripJSON) RawJSON() string {
	return r.raw
}
