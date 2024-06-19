// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package opentransit

import (
	"context"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/open-transit-go/internal/apijson"
	"github.com/stainless-sdks/open-transit-go/internal/apiquery"
	"github.com/stainless-sdks/open-transit-go/internal/param"
	"github.com/stainless-sdks/open-transit-go/internal/requestconfig"
	"github.com/stainless-sdks/open-transit-go/option"
	"github.com/stainless-sdks/open-transit-go/shared"
)

// APIWhereStopsForLocationService contains methods and other services that help
// with interacting with the OneBusAway API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAPIWhereStopsForLocationService] method instead.
type APIWhereStopsForLocationService struct {
	Options []option.RequestOption
}

// NewAPIWhereStopsForLocationService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAPIWhereStopsForLocationService(opts ...option.RequestOption) (r *APIWhereStopsForLocationService) {
	r = &APIWhereStopsForLocationService{}
	r.Options = opts
	return
}

// stops-for-location
func (r *APIWhereStopsForLocationService) List(ctx context.Context, query APIWhereStopsForLocationListParams, opts ...option.RequestOption) (res *APIWhereStopsForLocationListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/where/stops-for-location.json"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type APIWhereStopsForLocationListResponse struct {
	Data APIWhereStopsForLocationListResponseData `json:"data"`
	JSON apiWhereStopsForLocationListResponseJSON `json:"-"`
	shared.ResponseWrapper
}

// apiWhereStopsForLocationListResponseJSON contains the JSON metadata for the
// struct [APIWhereStopsForLocationListResponse]
type apiWhereStopsForLocationListResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIWhereStopsForLocationListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiWhereStopsForLocationListResponseJSON) RawJSON() string {
	return r.raw
}

type APIWhereStopsForLocationListResponseData struct {
	LimitExceeded bool                                               `json:"limitExceeded"`
	List          []APIWhereStopsForLocationListResponseDataList     `json:"list"`
	References    APIWhereStopsForLocationListResponseDataReferences `json:"references"`
	JSON          apiWhereStopsForLocationListResponseDataJSON       `json:"-"`
}

// apiWhereStopsForLocationListResponseDataJSON contains the JSON metadata for the
// struct [APIWhereStopsForLocationListResponseData]
type apiWhereStopsForLocationListResponseDataJSON struct {
	LimitExceeded apijson.Field
	List          apijson.Field
	References    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *APIWhereStopsForLocationListResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiWhereStopsForLocationListResponseDataJSON) RawJSON() string {
	return r.raw
}

type APIWhereStopsForLocationListResponseDataList struct {
	ID                 string                                           `json:"id"`
	Code               string                                           `json:"code"`
	Direction          string                                           `json:"direction"`
	Lat                float64                                          `json:"lat"`
	LocationType       int64                                            `json:"locationType"`
	Lon                float64                                          `json:"lon"`
	Name               string                                           `json:"name"`
	Parent             string                                           `json:"parent"`
	RouteIDs           []string                                         `json:"routeIds"`
	StaticRouteIDs     []string                                         `json:"staticRouteIds"`
	WheelchairBoarding string                                           `json:"wheelchairBoarding"`
	JSON               apiWhereStopsForLocationListResponseDataListJSON `json:"-"`
}

// apiWhereStopsForLocationListResponseDataListJSON contains the JSON metadata for
// the struct [APIWhereStopsForLocationListResponseDataList]
type apiWhereStopsForLocationListResponseDataListJSON struct {
	ID                 apijson.Field
	Code               apijson.Field
	Direction          apijson.Field
	Lat                apijson.Field
	LocationType       apijson.Field
	Lon                apijson.Field
	Name               apijson.Field
	Parent             apijson.Field
	RouteIDs           apijson.Field
	StaticRouteIDs     apijson.Field
	WheelchairBoarding apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *APIWhereStopsForLocationListResponseDataList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiWhereStopsForLocationListResponseDataListJSON) RawJSON() string {
	return r.raw
}

type APIWhereStopsForLocationListResponseDataReferences struct {
	Agencies   []APIWhereStopsForLocationListResponseDataReferencesAgency `json:"agencies"`
	Routes     []APIWhereStopsForLocationListResponseDataReferencesRoute  `json:"routes"`
	Situations []interface{}                                              `json:"situations"`
	Stops      []APIWhereStopsForLocationListResponseDataReferencesStop   `json:"stops"`
	StopTimes  []interface{}                                              `json:"stopTimes"`
	Trips      []APIWhereStopsForLocationListResponseDataReferencesTrip   `json:"trips"`
	JSON       apiWhereStopsForLocationListResponseDataReferencesJSON     `json:"-"`
}

// apiWhereStopsForLocationListResponseDataReferencesJSON contains the JSON
// metadata for the struct [APIWhereStopsForLocationListResponseDataReferences]
type apiWhereStopsForLocationListResponseDataReferencesJSON struct {
	Agencies    apijson.Field
	Routes      apijson.Field
	Situations  apijson.Field
	Stops       apijson.Field
	StopTimes   apijson.Field
	Trips       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIWhereStopsForLocationListResponseDataReferences) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiWhereStopsForLocationListResponseDataReferencesJSON) RawJSON() string {
	return r.raw
}

type APIWhereStopsForLocationListResponseDataReferencesAgency struct {
	ID             string                                                       `json:"id,required"`
	Name           string                                                       `json:"name,required"`
	Timezone       string                                                       `json:"timezone,required"`
	URL            string                                                       `json:"url,required"`
	Disclaimer     string                                                       `json:"disclaimer"`
	Email          string                                                       `json:"email"`
	FareURL        string                                                       `json:"fareUrl"`
	Lang           string                                                       `json:"lang"`
	Phone          string                                                       `json:"phone"`
	PrivateService bool                                                         `json:"privateService"`
	JSON           apiWhereStopsForLocationListResponseDataReferencesAgencyJSON `json:"-"`
}

// apiWhereStopsForLocationListResponseDataReferencesAgencyJSON contains the JSON
// metadata for the struct
// [APIWhereStopsForLocationListResponseDataReferencesAgency]
type apiWhereStopsForLocationListResponseDataReferencesAgencyJSON struct {
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

func (r *APIWhereStopsForLocationListResponseDataReferencesAgency) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiWhereStopsForLocationListResponseDataReferencesAgencyJSON) RawJSON() string {
	return r.raw
}

type APIWhereStopsForLocationListResponseDataReferencesRoute struct {
	ID                string                                                      `json:"id"`
	AgencyID          string                                                      `json:"agencyId"`
	Color             string                                                      `json:"color"`
	Description       string                                                      `json:"description"`
	LongName          string                                                      `json:"longName"`
	NullSafeShortName string                                                      `json:"nullSafeShortName"`
	ShortName         string                                                      `json:"shortName"`
	TextColor         string                                                      `json:"textColor"`
	Type              int64                                                       `json:"type"`
	URL               string                                                      `json:"url"`
	JSON              apiWhereStopsForLocationListResponseDataReferencesRouteJSON `json:"-"`
}

// apiWhereStopsForLocationListResponseDataReferencesRouteJSON contains the JSON
// metadata for the struct
// [APIWhereStopsForLocationListResponseDataReferencesRoute]
type apiWhereStopsForLocationListResponseDataReferencesRouteJSON struct {
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

func (r *APIWhereStopsForLocationListResponseDataReferencesRoute) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiWhereStopsForLocationListResponseDataReferencesRouteJSON) RawJSON() string {
	return r.raw
}

type APIWhereStopsForLocationListResponseDataReferencesStop struct {
	ID                 string                                                     `json:"id,required"`
	Code               string                                                     `json:"code,required"`
	Lat                float64                                                    `json:"lat,required"`
	Lon                float64                                                    `json:"lon,required"`
	Name               string                                                     `json:"name,required"`
	Direction          string                                                     `json:"direction"`
	LocationType       int64                                                      `json:"locationType"`
	Parent             string                                                     `json:"parent"`
	RouteIDs           []string                                                   `json:"routeIds"`
	StaticRouteIDs     []string                                                   `json:"staticRouteIds"`
	WheelchairBoarding string                                                     `json:"wheelchairBoarding"`
	JSON               apiWhereStopsForLocationListResponseDataReferencesStopJSON `json:"-"`
}

// apiWhereStopsForLocationListResponseDataReferencesStopJSON contains the JSON
// metadata for the struct [APIWhereStopsForLocationListResponseDataReferencesStop]
type apiWhereStopsForLocationListResponseDataReferencesStopJSON struct {
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

func (r *APIWhereStopsForLocationListResponseDataReferencesStop) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiWhereStopsForLocationListResponseDataReferencesStopJSON) RawJSON() string {
	return r.raw
}

type APIWhereStopsForLocationListResponseDataReferencesTrip struct {
	ID             string                                                     `json:"id,required"`
	RouteID        string                                                     `json:"routeId,required"`
	BlockID        string                                                     `json:"blockId"`
	DirectionID    string                                                     `json:"directionId"`
	PeakOffpeak    int64                                                      `json:"peakOffpeak"`
	RouteShortName string                                                     `json:"routeShortName"`
	ServiceID      string                                                     `json:"serviceId"`
	ShapeID        string                                                     `json:"shapeId"`
	TimeZone       string                                                     `json:"timeZone"`
	TripHeadsign   string                                                     `json:"tripHeadsign"`
	TripShortName  string                                                     `json:"tripShortName"`
	JSON           apiWhereStopsForLocationListResponseDataReferencesTripJSON `json:"-"`
}

// apiWhereStopsForLocationListResponseDataReferencesTripJSON contains the JSON
// metadata for the struct [APIWhereStopsForLocationListResponseDataReferencesTrip]
type apiWhereStopsForLocationListResponseDataReferencesTripJSON struct {
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

func (r *APIWhereStopsForLocationListResponseDataReferencesTrip) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiWhereStopsForLocationListResponseDataReferencesTripJSON) RawJSON() string {
	return r.raw
}

type APIWhereStopsForLocationListParams struct {
	Key param.Field[string]  `query:"key,required"`
	Lat param.Field[float64] `query:"lat"`
	Lon param.Field[float64] `query:"lon"`
}

// URLQuery serializes [APIWhereStopsForLocationListParams]'s query parameters as
// `url.Values`.
func (r APIWhereStopsForLocationListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
