// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package onebusaway

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

// StopsForLocationService contains methods and other services that help with
// interacting with the OneBusAway API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewStopsForLocationService] method instead.
type StopsForLocationService struct {
	Options []option.RequestOption
}

// NewStopsForLocationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewStopsForLocationService(opts ...option.RequestOption) (r *StopsForLocationService) {
	r = &StopsForLocationService{}
	r.Options = opts
	return
}

// stops-for-location
func (r *StopsForLocationService) List(ctx context.Context, query StopsForLocationListParams, opts ...option.RequestOption) (res *StopsForLocationListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/where/stops-for-location.json"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type StopsForLocationListResponse struct {
	Data StopsForLocationListResponseData `json:"data"`
	JSON stopsForLocationListResponseJSON `json:"-"`
	shared.ResponseWrapper
}

// stopsForLocationListResponseJSON contains the JSON metadata for the struct
// [StopsForLocationListResponse]
type stopsForLocationListResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StopsForLocationListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stopsForLocationListResponseJSON) RawJSON() string {
	return r.raw
}

type StopsForLocationListResponseData struct {
	LimitExceeded bool                                       `json:"limitExceeded"`
	List          []StopsForLocationListResponseDataList     `json:"list"`
	References    StopsForLocationListResponseDataReferences `json:"references"`
	JSON          stopsForLocationListResponseDataJSON       `json:"-"`
}

// stopsForLocationListResponseDataJSON contains the JSON metadata for the struct
// [StopsForLocationListResponseData]
type stopsForLocationListResponseDataJSON struct {
	LimitExceeded apijson.Field
	List          apijson.Field
	References    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *StopsForLocationListResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stopsForLocationListResponseDataJSON) RawJSON() string {
	return r.raw
}

type StopsForLocationListResponseDataList struct {
	ID                 string                                   `json:"id"`
	Code               string                                   `json:"code"`
	Direction          string                                   `json:"direction"`
	Lat                float64                                  `json:"lat"`
	LocationType       int64                                    `json:"locationType"`
	Lon                float64                                  `json:"lon"`
	Name               string                                   `json:"name"`
	Parent             string                                   `json:"parent"`
	RouteIDs           []string                                 `json:"routeIds"`
	StaticRouteIDs     []string                                 `json:"staticRouteIds"`
	WheelchairBoarding string                                   `json:"wheelchairBoarding"`
	JSON               stopsForLocationListResponseDataListJSON `json:"-"`
}

// stopsForLocationListResponseDataListJSON contains the JSON metadata for the
// struct [StopsForLocationListResponseDataList]
type stopsForLocationListResponseDataListJSON struct {
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

func (r *StopsForLocationListResponseDataList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stopsForLocationListResponseDataListJSON) RawJSON() string {
	return r.raw
}

type StopsForLocationListResponseDataReferences struct {
	Agencies   []StopsForLocationListResponseDataReferencesAgency `json:"agencies"`
	Routes     []StopsForLocationListResponseDataReferencesRoute  `json:"routes"`
	Situations []interface{}                                      `json:"situations"`
	Stops      []StopsForLocationListResponseDataReferencesStop   `json:"stops"`
	StopTimes  []interface{}                                      `json:"stopTimes"`
	Trips      []StopsForLocationListResponseDataReferencesTrip   `json:"trips"`
	JSON       stopsForLocationListResponseDataReferencesJSON     `json:"-"`
}

// stopsForLocationListResponseDataReferencesJSON contains the JSON metadata for
// the struct [StopsForLocationListResponseDataReferences]
type stopsForLocationListResponseDataReferencesJSON struct {
	Agencies    apijson.Field
	Routes      apijson.Field
	Situations  apijson.Field
	Stops       apijson.Field
	StopTimes   apijson.Field
	Trips       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StopsForLocationListResponseDataReferences) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stopsForLocationListResponseDataReferencesJSON) RawJSON() string {
	return r.raw
}

type StopsForLocationListResponseDataReferencesAgency struct {
	ID             string                                               `json:"id,required"`
	Name           string                                               `json:"name,required"`
	Timezone       string                                               `json:"timezone,required"`
	URL            string                                               `json:"url,required"`
	Disclaimer     string                                               `json:"disclaimer"`
	Email          string                                               `json:"email"`
	FareURL        string                                               `json:"fareUrl"`
	Lang           string                                               `json:"lang"`
	Phone          string                                               `json:"phone"`
	PrivateService bool                                                 `json:"privateService"`
	JSON           stopsForLocationListResponseDataReferencesAgencyJSON `json:"-"`
}

// stopsForLocationListResponseDataReferencesAgencyJSON contains the JSON metadata
// for the struct [StopsForLocationListResponseDataReferencesAgency]
type stopsForLocationListResponseDataReferencesAgencyJSON struct {
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

func (r *StopsForLocationListResponseDataReferencesAgency) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stopsForLocationListResponseDataReferencesAgencyJSON) RawJSON() string {
	return r.raw
}

type StopsForLocationListResponseDataReferencesRoute struct {
	ID                string                                              `json:"id"`
	AgencyID          string                                              `json:"agencyId"`
	Color             string                                              `json:"color"`
	Description       string                                              `json:"description"`
	LongName          string                                              `json:"longName"`
	NullSafeShortName string                                              `json:"nullSafeShortName"`
	ShortName         string                                              `json:"shortName"`
	TextColor         string                                              `json:"textColor"`
	Type              int64                                               `json:"type"`
	URL               string                                              `json:"url"`
	JSON              stopsForLocationListResponseDataReferencesRouteJSON `json:"-"`
}

// stopsForLocationListResponseDataReferencesRouteJSON contains the JSON metadata
// for the struct [StopsForLocationListResponseDataReferencesRoute]
type stopsForLocationListResponseDataReferencesRouteJSON struct {
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

func (r *StopsForLocationListResponseDataReferencesRoute) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stopsForLocationListResponseDataReferencesRouteJSON) RawJSON() string {
	return r.raw
}

type StopsForLocationListResponseDataReferencesStop struct {
	ID                 string                                             `json:"id,required"`
	Code               string                                             `json:"code,required"`
	Lat                float64                                            `json:"lat,required"`
	Lon                float64                                            `json:"lon,required"`
	Name               string                                             `json:"name,required"`
	Direction          string                                             `json:"direction"`
	LocationType       int64                                              `json:"locationType"`
	Parent             string                                             `json:"parent"`
	RouteIDs           []string                                           `json:"routeIds"`
	StaticRouteIDs     []string                                           `json:"staticRouteIds"`
	WheelchairBoarding string                                             `json:"wheelchairBoarding"`
	JSON               stopsForLocationListResponseDataReferencesStopJSON `json:"-"`
}

// stopsForLocationListResponseDataReferencesStopJSON contains the JSON metadata
// for the struct [StopsForLocationListResponseDataReferencesStop]
type stopsForLocationListResponseDataReferencesStopJSON struct {
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

func (r *StopsForLocationListResponseDataReferencesStop) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stopsForLocationListResponseDataReferencesStopJSON) RawJSON() string {
	return r.raw
}

type StopsForLocationListResponseDataReferencesTrip struct {
	ID             string                                             `json:"id,required"`
	RouteID        string                                             `json:"routeId,required"`
	BlockID        string                                             `json:"blockId"`
	DirectionID    string                                             `json:"directionId"`
	PeakOffpeak    int64                                              `json:"peakOffpeak"`
	RouteShortName string                                             `json:"routeShortName"`
	ServiceID      string                                             `json:"serviceId"`
	ShapeID        string                                             `json:"shapeId"`
	TimeZone       string                                             `json:"timeZone"`
	TripHeadsign   string                                             `json:"tripHeadsign"`
	TripShortName  string                                             `json:"tripShortName"`
	JSON           stopsForLocationListResponseDataReferencesTripJSON `json:"-"`
}

// stopsForLocationListResponseDataReferencesTripJSON contains the JSON metadata
// for the struct [StopsForLocationListResponseDataReferencesTrip]
type stopsForLocationListResponseDataReferencesTripJSON struct {
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

func (r *StopsForLocationListResponseDataReferencesTrip) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stopsForLocationListResponseDataReferencesTripJSON) RawJSON() string {
	return r.raw
}

type StopsForLocationListParams struct {
	Key param.Field[string]  `query:"key,required"`
	Lat param.Field[float64] `query:"lat"`
	Lon param.Field[float64] `query:"lon"`
}

// URLQuery serializes [StopsForLocationListParams]'s query parameters as
// `url.Values`.
func (r StopsForLocationListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
