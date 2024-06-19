// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package onebusaway

import (
	"context"
	"net/http"

	"github.com/stainless-sdks/open-transit-go/internal/apijson"
	"github.com/stainless-sdks/open-transit-go/internal/requestconfig"
	"github.com/stainless-sdks/open-transit-go/option"
	"github.com/stainless-sdks/open-transit-go/shared"
)

// AgenciesWithCoverageService contains methods and other services that help with
// interacting with the OneBusAway API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAgenciesWithCoverageService] method instead.
type AgenciesWithCoverageService struct {
	Options []option.RequestOption
}

// NewAgenciesWithCoverageService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAgenciesWithCoverageService(opts ...option.RequestOption) (r *AgenciesWithCoverageService) {
	r = &AgenciesWithCoverageService{}
	r.Options = opts
	return
}

// Retrieve Agencies with Coverage
func (r *AgenciesWithCoverageService) Get(ctx context.Context, opts ...option.RequestOption) (res *AgenciesWithCoverageGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/where/agencies-with-coverage.json"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AgenciesWithCoverageGetResponse struct {
	Data AgenciesWithCoverageGetResponseData `json:"data"`
	JSON agenciesWithCoverageGetResponseJSON `json:"-"`
	shared.ResponseWrapper
}

// agenciesWithCoverageGetResponseJSON contains the JSON metadata for the struct
// [AgenciesWithCoverageGetResponse]
type agenciesWithCoverageGetResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AgenciesWithCoverageGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r agenciesWithCoverageGetResponseJSON) RawJSON() string {
	return r.raw
}

type AgenciesWithCoverageGetResponseData struct {
	LimitExceeded bool                                          `json:"limitExceeded"`
	List          []AgenciesWithCoverageGetResponseDataList     `json:"list"`
	References    AgenciesWithCoverageGetResponseDataReferences `json:"references"`
	JSON          agenciesWithCoverageGetResponseDataJSON       `json:"-"`
}

// agenciesWithCoverageGetResponseDataJSON contains the JSON metadata for the
// struct [AgenciesWithCoverageGetResponseData]
type agenciesWithCoverageGetResponseDataJSON struct {
	LimitExceeded apijson.Field
	List          apijson.Field
	References    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AgenciesWithCoverageGetResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r agenciesWithCoverageGetResponseDataJSON) RawJSON() string {
	return r.raw
}

type AgenciesWithCoverageGetResponseDataList struct {
	AgencyID string                                      `json:"agencyId,required"`
	Lat      float64                                     `json:"lat,required"`
	LatSpan  float64                                     `json:"latSpan,required"`
	Lon      float64                                     `json:"lon,required"`
	LonSpan  float64                                     `json:"lonSpan,required"`
	JSON     agenciesWithCoverageGetResponseDataListJSON `json:"-"`
}

// agenciesWithCoverageGetResponseDataListJSON contains the JSON metadata for the
// struct [AgenciesWithCoverageGetResponseDataList]
type agenciesWithCoverageGetResponseDataListJSON struct {
	AgencyID    apijson.Field
	Lat         apijson.Field
	LatSpan     apijson.Field
	Lon         apijson.Field
	LonSpan     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AgenciesWithCoverageGetResponseDataList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r agenciesWithCoverageGetResponseDataListJSON) RawJSON() string {
	return r.raw
}

type AgenciesWithCoverageGetResponseDataReferences struct {
	Agencies   []AgenciesWithCoverageGetResponseDataReferencesAgency `json:"agencies"`
	Routes     []AgenciesWithCoverageGetResponseDataReferencesRoute  `json:"routes"`
	Situations []interface{}                                         `json:"situations"`
	Stops      []AgenciesWithCoverageGetResponseDataReferencesStop   `json:"stops"`
	StopTimes  []interface{}                                         `json:"stopTimes"`
	Trips      []AgenciesWithCoverageGetResponseDataReferencesTrip   `json:"trips"`
	JSON       agenciesWithCoverageGetResponseDataReferencesJSON     `json:"-"`
}

// agenciesWithCoverageGetResponseDataReferencesJSON contains the JSON metadata for
// the struct [AgenciesWithCoverageGetResponseDataReferences]
type agenciesWithCoverageGetResponseDataReferencesJSON struct {
	Agencies    apijson.Field
	Routes      apijson.Field
	Situations  apijson.Field
	Stops       apijson.Field
	StopTimes   apijson.Field
	Trips       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AgenciesWithCoverageGetResponseDataReferences) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r agenciesWithCoverageGetResponseDataReferencesJSON) RawJSON() string {
	return r.raw
}

type AgenciesWithCoverageGetResponseDataReferencesAgency struct {
	ID             string                                                  `json:"id,required"`
	Name           string                                                  `json:"name,required"`
	Timezone       string                                                  `json:"timezone,required"`
	URL            string                                                  `json:"url,required"`
	Disclaimer     string                                                  `json:"disclaimer"`
	Email          string                                                  `json:"email"`
	FareURL        string                                                  `json:"fareUrl"`
	Lang           string                                                  `json:"lang"`
	Phone          string                                                  `json:"phone"`
	PrivateService bool                                                    `json:"privateService"`
	JSON           agenciesWithCoverageGetResponseDataReferencesAgencyJSON `json:"-"`
}

// agenciesWithCoverageGetResponseDataReferencesAgencyJSON contains the JSON
// metadata for the struct [AgenciesWithCoverageGetResponseDataReferencesAgency]
type agenciesWithCoverageGetResponseDataReferencesAgencyJSON struct {
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

func (r *AgenciesWithCoverageGetResponseDataReferencesAgency) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r agenciesWithCoverageGetResponseDataReferencesAgencyJSON) RawJSON() string {
	return r.raw
}

type AgenciesWithCoverageGetResponseDataReferencesRoute struct {
	ID                string                                                 `json:"id"`
	AgencyID          string                                                 `json:"agencyId"`
	Color             string                                                 `json:"color"`
	Description       string                                                 `json:"description"`
	LongName          string                                                 `json:"longName"`
	NullSafeShortName string                                                 `json:"nullSafeShortName"`
	ShortName         string                                                 `json:"shortName"`
	TextColor         string                                                 `json:"textColor"`
	Type              int64                                                  `json:"type"`
	URL               string                                                 `json:"url"`
	JSON              agenciesWithCoverageGetResponseDataReferencesRouteJSON `json:"-"`
}

// agenciesWithCoverageGetResponseDataReferencesRouteJSON contains the JSON
// metadata for the struct [AgenciesWithCoverageGetResponseDataReferencesRoute]
type agenciesWithCoverageGetResponseDataReferencesRouteJSON struct {
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

func (r *AgenciesWithCoverageGetResponseDataReferencesRoute) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r agenciesWithCoverageGetResponseDataReferencesRouteJSON) RawJSON() string {
	return r.raw
}

type AgenciesWithCoverageGetResponseDataReferencesStop struct {
	ID                 string                                                `json:"id,required"`
	Code               string                                                `json:"code,required"`
	Lat                float64                                               `json:"lat,required"`
	Lon                float64                                               `json:"lon,required"`
	Name               string                                                `json:"name,required"`
	Direction          string                                                `json:"direction"`
	LocationType       int64                                                 `json:"locationType"`
	Parent             string                                                `json:"parent"`
	RouteIDs           []string                                              `json:"routeIds"`
	StaticRouteIDs     []string                                              `json:"staticRouteIds"`
	WheelchairBoarding string                                                `json:"wheelchairBoarding"`
	JSON               agenciesWithCoverageGetResponseDataReferencesStopJSON `json:"-"`
}

// agenciesWithCoverageGetResponseDataReferencesStopJSON contains the JSON metadata
// for the struct [AgenciesWithCoverageGetResponseDataReferencesStop]
type agenciesWithCoverageGetResponseDataReferencesStopJSON struct {
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

func (r *AgenciesWithCoverageGetResponseDataReferencesStop) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r agenciesWithCoverageGetResponseDataReferencesStopJSON) RawJSON() string {
	return r.raw
}

type AgenciesWithCoverageGetResponseDataReferencesTrip struct {
	ID             string                                                `json:"id,required"`
	RouteID        string                                                `json:"routeId,required"`
	BlockID        string                                                `json:"blockId"`
	DirectionID    string                                                `json:"directionId"`
	PeakOffpeak    int64                                                 `json:"peakOffpeak"`
	RouteShortName string                                                `json:"routeShortName"`
	ServiceID      string                                                `json:"serviceId"`
	ShapeID        string                                                `json:"shapeId"`
	TimeZone       string                                                `json:"timeZone"`
	TripHeadsign   string                                                `json:"tripHeadsign"`
	TripShortName  string                                                `json:"tripShortName"`
	JSON           agenciesWithCoverageGetResponseDataReferencesTripJSON `json:"-"`
}

// agenciesWithCoverageGetResponseDataReferencesTripJSON contains the JSON metadata
// for the struct [AgenciesWithCoverageGetResponseDataReferencesTrip]
type agenciesWithCoverageGetResponseDataReferencesTripJSON struct {
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

func (r *AgenciesWithCoverageGetResponseDataReferencesTrip) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r agenciesWithCoverageGetResponseDataReferencesTripJSON) RawJSON() string {
	return r.raw
}
