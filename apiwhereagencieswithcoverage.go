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

// APIWhereAgenciesWithCoverageService contains methods and other services that
// help with interacting with the OneBusAway API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAPIWhereAgenciesWithCoverageService] method instead.
type APIWhereAgenciesWithCoverageService struct {
	Options []option.RequestOption
}

// NewAPIWhereAgenciesWithCoverageService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAPIWhereAgenciesWithCoverageService(opts ...option.RequestOption) (r *APIWhereAgenciesWithCoverageService) {
	r = &APIWhereAgenciesWithCoverageService{}
	r.Options = opts
	return
}

// Retrieve Agencies with Coverage
func (r *APIWhereAgenciesWithCoverageService) List(ctx context.Context, opts ...option.RequestOption) (res *APIWhereAgenciesWithCoverageListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/where/agencies-with-coverage.json"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type APIWhereAgenciesWithCoverageListResponse struct {
	Data APIWhereAgenciesWithCoverageListResponseData `json:"data"`
	JSON apiWhereAgenciesWithCoverageListResponseJSON `json:"-"`
	shared.ResponseWrapper
}

// apiWhereAgenciesWithCoverageListResponseJSON contains the JSON metadata for the
// struct [APIWhereAgenciesWithCoverageListResponse]
type apiWhereAgenciesWithCoverageListResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIWhereAgenciesWithCoverageListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiWhereAgenciesWithCoverageListResponseJSON) RawJSON() string {
	return r.raw
}

type APIWhereAgenciesWithCoverageListResponseData struct {
	LimitExceeded bool                                                   `json:"limitExceeded"`
	List          []APIWhereAgenciesWithCoverageListResponseDataList     `json:"list"`
	References    APIWhereAgenciesWithCoverageListResponseDataReferences `json:"references"`
	JSON          apiWhereAgenciesWithCoverageListResponseDataJSON       `json:"-"`
}

// apiWhereAgenciesWithCoverageListResponseDataJSON contains the JSON metadata for
// the struct [APIWhereAgenciesWithCoverageListResponseData]
type apiWhereAgenciesWithCoverageListResponseDataJSON struct {
	LimitExceeded apijson.Field
	List          apijson.Field
	References    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *APIWhereAgenciesWithCoverageListResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiWhereAgenciesWithCoverageListResponseDataJSON) RawJSON() string {
	return r.raw
}

type APIWhereAgenciesWithCoverageListResponseDataList struct {
	AgencyID string                                               `json:"agencyId,required"`
	Lat      float64                                              `json:"lat,required"`
	LatSpan  float64                                              `json:"latSpan,required"`
	Lon      float64                                              `json:"lon,required"`
	LonSpan  float64                                              `json:"lonSpan,required"`
	JSON     apiWhereAgenciesWithCoverageListResponseDataListJSON `json:"-"`
}

// apiWhereAgenciesWithCoverageListResponseDataListJSON contains the JSON metadata
// for the struct [APIWhereAgenciesWithCoverageListResponseDataList]
type apiWhereAgenciesWithCoverageListResponseDataListJSON struct {
	AgencyID    apijson.Field
	Lat         apijson.Field
	LatSpan     apijson.Field
	Lon         apijson.Field
	LonSpan     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIWhereAgenciesWithCoverageListResponseDataList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiWhereAgenciesWithCoverageListResponseDataListJSON) RawJSON() string {
	return r.raw
}

type APIWhereAgenciesWithCoverageListResponseDataReferences struct {
	Agencies   []APIWhereAgenciesWithCoverageListResponseDataReferencesAgency `json:"agencies"`
	Routes     []APIWhereAgenciesWithCoverageListResponseDataReferencesRoute  `json:"routes"`
	Situations []interface{}                                                  `json:"situations"`
	Stops      []APIWhereAgenciesWithCoverageListResponseDataReferencesStop   `json:"stops"`
	StopTimes  []interface{}                                                  `json:"stopTimes"`
	Trips      []APIWhereAgenciesWithCoverageListResponseDataReferencesTrip   `json:"trips"`
	JSON       apiWhereAgenciesWithCoverageListResponseDataReferencesJSON     `json:"-"`
}

// apiWhereAgenciesWithCoverageListResponseDataReferencesJSON contains the JSON
// metadata for the struct [APIWhereAgenciesWithCoverageListResponseDataReferences]
type apiWhereAgenciesWithCoverageListResponseDataReferencesJSON struct {
	Agencies    apijson.Field
	Routes      apijson.Field
	Situations  apijson.Field
	Stops       apijson.Field
	StopTimes   apijson.Field
	Trips       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIWhereAgenciesWithCoverageListResponseDataReferences) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiWhereAgenciesWithCoverageListResponseDataReferencesJSON) RawJSON() string {
	return r.raw
}

type APIWhereAgenciesWithCoverageListResponseDataReferencesAgency struct {
	ID             string                                                           `json:"id,required"`
	Name           string                                                           `json:"name,required"`
	Timezone       string                                                           `json:"timezone,required"`
	URL            string                                                           `json:"url,required"`
	Disclaimer     string                                                           `json:"disclaimer"`
	Email          string                                                           `json:"email"`
	FareURL        string                                                           `json:"fareUrl"`
	Lang           string                                                           `json:"lang"`
	Phone          string                                                           `json:"phone"`
	PrivateService bool                                                             `json:"privateService"`
	JSON           apiWhereAgenciesWithCoverageListResponseDataReferencesAgencyJSON `json:"-"`
}

// apiWhereAgenciesWithCoverageListResponseDataReferencesAgencyJSON contains the
// JSON metadata for the struct
// [APIWhereAgenciesWithCoverageListResponseDataReferencesAgency]
type apiWhereAgenciesWithCoverageListResponseDataReferencesAgencyJSON struct {
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

func (r *APIWhereAgenciesWithCoverageListResponseDataReferencesAgency) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiWhereAgenciesWithCoverageListResponseDataReferencesAgencyJSON) RawJSON() string {
	return r.raw
}

type APIWhereAgenciesWithCoverageListResponseDataReferencesRoute struct {
	ID                string                                                          `json:"id"`
	AgencyID          string                                                          `json:"agencyId"`
	Color             string                                                          `json:"color"`
	Description       string                                                          `json:"description"`
	LongName          string                                                          `json:"longName"`
	NullSafeShortName string                                                          `json:"nullSafeShortName"`
	ShortName         string                                                          `json:"shortName"`
	TextColor         string                                                          `json:"textColor"`
	Type              int64                                                           `json:"type"`
	URL               string                                                          `json:"url"`
	JSON              apiWhereAgenciesWithCoverageListResponseDataReferencesRouteJSON `json:"-"`
}

// apiWhereAgenciesWithCoverageListResponseDataReferencesRouteJSON contains the
// JSON metadata for the struct
// [APIWhereAgenciesWithCoverageListResponseDataReferencesRoute]
type apiWhereAgenciesWithCoverageListResponseDataReferencesRouteJSON struct {
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

func (r *APIWhereAgenciesWithCoverageListResponseDataReferencesRoute) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiWhereAgenciesWithCoverageListResponseDataReferencesRouteJSON) RawJSON() string {
	return r.raw
}

type APIWhereAgenciesWithCoverageListResponseDataReferencesStop struct {
	ID                 string                                                         `json:"id,required"`
	Code               string                                                         `json:"code,required"`
	Lat                float64                                                        `json:"lat,required"`
	Lon                float64                                                        `json:"lon,required"`
	Name               string                                                         `json:"name,required"`
	Direction          string                                                         `json:"direction"`
	LocationType       int64                                                          `json:"locationType"`
	Parent             string                                                         `json:"parent"`
	RouteIDs           []string                                                       `json:"routeIds"`
	StaticRouteIDs     []string                                                       `json:"staticRouteIds"`
	WheelchairBoarding string                                                         `json:"wheelchairBoarding"`
	JSON               apiWhereAgenciesWithCoverageListResponseDataReferencesStopJSON `json:"-"`
}

// apiWhereAgenciesWithCoverageListResponseDataReferencesStopJSON contains the JSON
// metadata for the struct
// [APIWhereAgenciesWithCoverageListResponseDataReferencesStop]
type apiWhereAgenciesWithCoverageListResponseDataReferencesStopJSON struct {
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

func (r *APIWhereAgenciesWithCoverageListResponseDataReferencesStop) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiWhereAgenciesWithCoverageListResponseDataReferencesStopJSON) RawJSON() string {
	return r.raw
}

type APIWhereAgenciesWithCoverageListResponseDataReferencesTrip struct {
	ID             string                                                         `json:"id,required"`
	RouteID        string                                                         `json:"routeId,required"`
	BlockID        string                                                         `json:"blockId"`
	DirectionID    string                                                         `json:"directionId"`
	PeakOffpeak    int64                                                          `json:"peakOffpeak"`
	RouteShortName string                                                         `json:"routeShortName"`
	ServiceID      string                                                         `json:"serviceId"`
	ShapeID        string                                                         `json:"shapeId"`
	TimeZone       string                                                         `json:"timeZone"`
	TripHeadsign   string                                                         `json:"tripHeadsign"`
	TripShortName  string                                                         `json:"tripShortName"`
	JSON           apiWhereAgenciesWithCoverageListResponseDataReferencesTripJSON `json:"-"`
}

// apiWhereAgenciesWithCoverageListResponseDataReferencesTripJSON contains the JSON
// metadata for the struct
// [APIWhereAgenciesWithCoverageListResponseDataReferencesTrip]
type apiWhereAgenciesWithCoverageListResponseDataReferencesTripJSON struct {
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

func (r *APIWhereAgenciesWithCoverageListResponseDataReferencesTrip) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiWhereAgenciesWithCoverageListResponseDataReferencesTripJSON) RawJSON() string {
	return r.raw
}
