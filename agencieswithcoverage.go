// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package opentransit

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
func (r *AgenciesWithCoverageService) List(ctx context.Context, opts ...option.RequestOption) (res *AgenciesWithCoverageListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/where/agencies-with-coverage.json"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AgenciesWithCoverageListResponse struct {
	Data AgenciesWithCoverageListResponseData `json:"data"`
	JSON agenciesWithCoverageListResponseJSON `json:"-"`
	shared.ResponseWrapper
}

// agenciesWithCoverageListResponseJSON contains the JSON metadata for the struct
// [AgenciesWithCoverageListResponse]
type agenciesWithCoverageListResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AgenciesWithCoverageListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r agenciesWithCoverageListResponseJSON) RawJSON() string {
	return r.raw
}

type AgenciesWithCoverageListResponseData struct {
	LimitExceeded bool                                           `json:"limitExceeded"`
	List          []AgenciesWithCoverageListResponseDataList     `json:"list"`
	References    AgenciesWithCoverageListResponseDataReferences `json:"references"`
	JSON          agenciesWithCoverageListResponseDataJSON       `json:"-"`
}

// agenciesWithCoverageListResponseDataJSON contains the JSON metadata for the
// struct [AgenciesWithCoverageListResponseData]
type agenciesWithCoverageListResponseDataJSON struct {
	LimitExceeded apijson.Field
	List          apijson.Field
	References    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AgenciesWithCoverageListResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r agenciesWithCoverageListResponseDataJSON) RawJSON() string {
	return r.raw
}

type AgenciesWithCoverageListResponseDataList struct {
	AgencyID string                                       `json:"agencyId,required"`
	Lat      float64                                      `json:"lat,required"`
	LatSpan  float64                                      `json:"latSpan,required"`
	Lon      float64                                      `json:"lon,required"`
	LonSpan  float64                                      `json:"lonSpan,required"`
	JSON     agenciesWithCoverageListResponseDataListJSON `json:"-"`
}

// agenciesWithCoverageListResponseDataListJSON contains the JSON metadata for the
// struct [AgenciesWithCoverageListResponseDataList]
type agenciesWithCoverageListResponseDataListJSON struct {
	AgencyID    apijson.Field
	Lat         apijson.Field
	LatSpan     apijson.Field
	Lon         apijson.Field
	LonSpan     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AgenciesWithCoverageListResponseDataList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r agenciesWithCoverageListResponseDataListJSON) RawJSON() string {
	return r.raw
}

type AgenciesWithCoverageListResponseDataReferences struct {
	Agencies   []AgenciesWithCoverageListResponseDataReferencesAgency `json:"agencies"`
	Routes     []AgenciesWithCoverageListResponseDataReferencesRoute  `json:"routes"`
	Situations []interface{}                                          `json:"situations"`
	Stops      []AgenciesWithCoverageListResponseDataReferencesStop   `json:"stops"`
	StopTimes  []interface{}                                          `json:"stopTimes"`
	Trips      []AgenciesWithCoverageListResponseDataReferencesTrip   `json:"trips"`
	JSON       agenciesWithCoverageListResponseDataReferencesJSON     `json:"-"`
}

// agenciesWithCoverageListResponseDataReferencesJSON contains the JSON metadata
// for the struct [AgenciesWithCoverageListResponseDataReferences]
type agenciesWithCoverageListResponseDataReferencesJSON struct {
	Agencies    apijson.Field
	Routes      apijson.Field
	Situations  apijson.Field
	Stops       apijson.Field
	StopTimes   apijson.Field
	Trips       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AgenciesWithCoverageListResponseDataReferences) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r agenciesWithCoverageListResponseDataReferencesJSON) RawJSON() string {
	return r.raw
}

type AgenciesWithCoverageListResponseDataReferencesAgency struct {
	ID             string                                                   `json:"id,required"`
	Name           string                                                   `json:"name,required"`
	Timezone       string                                                   `json:"timezone,required"`
	URL            string                                                   `json:"url,required"`
	Disclaimer     string                                                   `json:"disclaimer"`
	Email          string                                                   `json:"email"`
	FareURL        string                                                   `json:"fareUrl"`
	Lang           string                                                   `json:"lang"`
	Phone          string                                                   `json:"phone"`
	PrivateService bool                                                     `json:"privateService"`
	JSON           agenciesWithCoverageListResponseDataReferencesAgencyJSON `json:"-"`
}

// agenciesWithCoverageListResponseDataReferencesAgencyJSON contains the JSON
// metadata for the struct [AgenciesWithCoverageListResponseDataReferencesAgency]
type agenciesWithCoverageListResponseDataReferencesAgencyJSON struct {
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

func (r *AgenciesWithCoverageListResponseDataReferencesAgency) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r agenciesWithCoverageListResponseDataReferencesAgencyJSON) RawJSON() string {
	return r.raw
}

type AgenciesWithCoverageListResponseDataReferencesRoute struct {
	ID                string                                                  `json:"id"`
	AgencyID          string                                                  `json:"agencyId"`
	Color             string                                                  `json:"color"`
	Description       string                                                  `json:"description"`
	LongName          string                                                  `json:"longName"`
	NullSafeShortName string                                                  `json:"nullSafeShortName"`
	ShortName         string                                                  `json:"shortName"`
	TextColor         string                                                  `json:"textColor"`
	Type              int64                                                   `json:"type"`
	URL               string                                                  `json:"url"`
	JSON              agenciesWithCoverageListResponseDataReferencesRouteJSON `json:"-"`
}

// agenciesWithCoverageListResponseDataReferencesRouteJSON contains the JSON
// metadata for the struct [AgenciesWithCoverageListResponseDataReferencesRoute]
type agenciesWithCoverageListResponseDataReferencesRouteJSON struct {
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

func (r *AgenciesWithCoverageListResponseDataReferencesRoute) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r agenciesWithCoverageListResponseDataReferencesRouteJSON) RawJSON() string {
	return r.raw
}

type AgenciesWithCoverageListResponseDataReferencesStop struct {
	ID                 string                                                 `json:"id,required"`
	Code               string                                                 `json:"code,required"`
	Lat                float64                                                `json:"lat,required"`
	Lon                float64                                                `json:"lon,required"`
	Name               string                                                 `json:"name,required"`
	Direction          string                                                 `json:"direction"`
	LocationType       int64                                                  `json:"locationType"`
	Parent             string                                                 `json:"parent"`
	RouteIDs           []string                                               `json:"routeIds"`
	StaticRouteIDs     []string                                               `json:"staticRouteIds"`
	WheelchairBoarding string                                                 `json:"wheelchairBoarding"`
	JSON               agenciesWithCoverageListResponseDataReferencesStopJSON `json:"-"`
}

// agenciesWithCoverageListResponseDataReferencesStopJSON contains the JSON
// metadata for the struct [AgenciesWithCoverageListResponseDataReferencesStop]
type agenciesWithCoverageListResponseDataReferencesStopJSON struct {
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

func (r *AgenciesWithCoverageListResponseDataReferencesStop) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r agenciesWithCoverageListResponseDataReferencesStopJSON) RawJSON() string {
	return r.raw
}

type AgenciesWithCoverageListResponseDataReferencesTrip struct {
	ID             string                                                 `json:"id,required"`
	RouteID        string                                                 `json:"routeId,required"`
	BlockID        string                                                 `json:"blockId"`
	DirectionID    string                                                 `json:"directionId"`
	PeakOffpeak    int64                                                  `json:"peakOffpeak"`
	RouteShortName string                                                 `json:"routeShortName"`
	ServiceID      string                                                 `json:"serviceId"`
	ShapeID        string                                                 `json:"shapeId"`
	TimeZone       string                                                 `json:"timeZone"`
	TripHeadsign   string                                                 `json:"tripHeadsign"`
	TripShortName  string                                                 `json:"tripShortName"`
	JSON           agenciesWithCoverageListResponseDataReferencesTripJSON `json:"-"`
}

// agenciesWithCoverageListResponseDataReferencesTripJSON contains the JSON
// metadata for the struct [AgenciesWithCoverageListResponseDataReferencesTrip]
type agenciesWithCoverageListResponseDataReferencesTripJSON struct {
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

func (r *AgenciesWithCoverageListResponseDataReferencesTrip) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r agenciesWithCoverageListResponseDataReferencesTripJSON) RawJSON() string {
	return r.raw
}
