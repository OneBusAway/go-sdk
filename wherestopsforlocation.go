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
)

// WhereStopsForLocationService contains methods and other services that help with
// interacting with the open-transit API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWhereStopsForLocationService] method instead.
type WhereStopsForLocationService struct {
	Options []option.RequestOption
}

// NewWhereStopsForLocationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewWhereStopsForLocationService(opts ...option.RequestOption) (r *WhereStopsForLocationService) {
	r = &WhereStopsForLocationService{}
	r.Options = opts
	return
}

// stops-for-location
func (r *WhereStopsForLocationService) List(ctx context.Context, query WhereStopsForLocationListParams, opts ...option.RequestOption) (res *WhereStopsForLocationListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/where/stops-for-location.json"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type WhereStopsForLocationListResponse struct {
	Code        int64                                 `json:"code,required"`
	CurrentTime int64                                 `json:"currentTime,required"`
	Text        string                                `json:"text,required"`
	Version     int64                                 `json:"version,required"`
	Data        WhereStopsForLocationListResponseData `json:"data"`
	JSON        whereStopsForLocationListResponseJSON `json:"-"`
}

// whereStopsForLocationListResponseJSON contains the JSON metadata for the struct
// [WhereStopsForLocationListResponse]
type whereStopsForLocationListResponseJSON struct {
	Code        apijson.Field
	CurrentTime apijson.Field
	Text        apijson.Field
	Version     apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WhereStopsForLocationListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r whereStopsForLocationListResponseJSON) RawJSON() string {
	return r.raw
}

type WhereStopsForLocationListResponseData struct {
	LimitExceeded bool                                            `json:"limitExceeded"`
	List          []WhereStopsForLocationListResponseDataList     `json:"list"`
	References    WhereStopsForLocationListResponseDataReferences `json:"references"`
	JSON          whereStopsForLocationListResponseDataJSON       `json:"-"`
}

// whereStopsForLocationListResponseDataJSON contains the JSON metadata for the
// struct [WhereStopsForLocationListResponseData]
type whereStopsForLocationListResponseDataJSON struct {
	LimitExceeded apijson.Field
	List          apijson.Field
	References    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *WhereStopsForLocationListResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r whereStopsForLocationListResponseDataJSON) RawJSON() string {
	return r.raw
}

type WhereStopsForLocationListResponseDataList struct {
	ID                 string                                        `json:"id"`
	Code               string                                        `json:"code"`
	Direction          string                                        `json:"direction"`
	Lat                float64                                       `json:"lat"`
	LocationType       int64                                         `json:"locationType"`
	Lon                float64                                       `json:"lon"`
	Name               string                                        `json:"name"`
	Parent             string                                        `json:"parent"`
	RouteIDs           []string                                      `json:"routeIds"`
	StaticRouteIDs     []string                                      `json:"staticRouteIds"`
	WheelchairBoarding string                                        `json:"wheelchairBoarding"`
	JSON               whereStopsForLocationListResponseDataListJSON `json:"-"`
}

// whereStopsForLocationListResponseDataListJSON contains the JSON metadata for the
// struct [WhereStopsForLocationListResponseDataList]
type whereStopsForLocationListResponseDataListJSON struct {
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

func (r *WhereStopsForLocationListResponseDataList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r whereStopsForLocationListResponseDataListJSON) RawJSON() string {
	return r.raw
}

type WhereStopsForLocationListResponseDataReferences struct {
	Agencies   []WhereStopsForLocationListResponseDataReferencesAgency `json:"agencies"`
	Routes     []WhereStopsForLocationListResponseDataReferencesRoute  `json:"routes"`
	Situations []interface{}                                           `json:"situations"`
	Stops      []WhereStopsForLocationListResponseDataReferencesStop   `json:"stops"`
	StopTimes  []interface{}                                           `json:"stopTimes"`
	Trips      []WhereStopsForLocationListResponseDataReferencesTrip   `json:"trips"`
	JSON       whereStopsForLocationListResponseDataReferencesJSON     `json:"-"`
}

// whereStopsForLocationListResponseDataReferencesJSON contains the JSON metadata
// for the struct [WhereStopsForLocationListResponseDataReferences]
type whereStopsForLocationListResponseDataReferencesJSON struct {
	Agencies    apijson.Field
	Routes      apijson.Field
	Situations  apijson.Field
	Stops       apijson.Field
	StopTimes   apijson.Field
	Trips       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WhereStopsForLocationListResponseDataReferences) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r whereStopsForLocationListResponseDataReferencesJSON) RawJSON() string {
	return r.raw
}

type WhereStopsForLocationListResponseDataReferencesAgency struct {
	ID             string                                                    `json:"id,required"`
	Name           string                                                    `json:"name,required"`
	Timezone       string                                                    `json:"timezone,required"`
	URL            string                                                    `json:"url,required"`
	Disclaimer     string                                                    `json:"disclaimer"`
	Email          string                                                    `json:"email"`
	FareURL        string                                                    `json:"fareUrl"`
	Lang           string                                                    `json:"lang"`
	Phone          string                                                    `json:"phone"`
	PrivateService bool                                                      `json:"privateService"`
	JSON           whereStopsForLocationListResponseDataReferencesAgencyJSON `json:"-"`
}

// whereStopsForLocationListResponseDataReferencesAgencyJSON contains the JSON
// metadata for the struct [WhereStopsForLocationListResponseDataReferencesAgency]
type whereStopsForLocationListResponseDataReferencesAgencyJSON struct {
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

func (r *WhereStopsForLocationListResponseDataReferencesAgency) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r whereStopsForLocationListResponseDataReferencesAgencyJSON) RawJSON() string {
	return r.raw
}

type WhereStopsForLocationListResponseDataReferencesRoute struct {
	ID                string                                                   `json:"id"`
	AgencyID          string                                                   `json:"agencyId"`
	Color             string                                                   `json:"color"`
	Description       string                                                   `json:"description"`
	LongName          string                                                   `json:"longName"`
	NullSafeShortName string                                                   `json:"nullSafeShortName"`
	ShortName         string                                                   `json:"shortName"`
	TextColor         string                                                   `json:"textColor"`
	Type              int64                                                    `json:"type"`
	URL               string                                                   `json:"url"`
	JSON              whereStopsForLocationListResponseDataReferencesRouteJSON `json:"-"`
}

// whereStopsForLocationListResponseDataReferencesRouteJSON contains the JSON
// metadata for the struct [WhereStopsForLocationListResponseDataReferencesRoute]
type whereStopsForLocationListResponseDataReferencesRouteJSON struct {
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

func (r *WhereStopsForLocationListResponseDataReferencesRoute) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r whereStopsForLocationListResponseDataReferencesRouteJSON) RawJSON() string {
	return r.raw
}

type WhereStopsForLocationListResponseDataReferencesStop struct {
	ID                 string                                                  `json:"id,required"`
	Code               string                                                  `json:"code,required"`
	Lat                float64                                                 `json:"lat,required"`
	Lon                float64                                                 `json:"lon,required"`
	Name               string                                                  `json:"name,required"`
	Direction          string                                                  `json:"direction"`
	LocationType       int64                                                   `json:"locationType"`
	Parent             string                                                  `json:"parent"`
	RouteIDs           []string                                                `json:"routeIds"`
	StaticRouteIDs     []string                                                `json:"staticRouteIds"`
	WheelchairBoarding string                                                  `json:"wheelchairBoarding"`
	JSON               whereStopsForLocationListResponseDataReferencesStopJSON `json:"-"`
}

// whereStopsForLocationListResponseDataReferencesStopJSON contains the JSON
// metadata for the struct [WhereStopsForLocationListResponseDataReferencesStop]
type whereStopsForLocationListResponseDataReferencesStopJSON struct {
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

func (r *WhereStopsForLocationListResponseDataReferencesStop) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r whereStopsForLocationListResponseDataReferencesStopJSON) RawJSON() string {
	return r.raw
}

type WhereStopsForLocationListResponseDataReferencesTrip struct {
	ID             string                                                  `json:"id,required"`
	RouteID        string                                                  `json:"routeId,required"`
	BlockID        string                                                  `json:"blockId"`
	DirectionID    string                                                  `json:"directionId"`
	PeakOffpeak    int64                                                   `json:"peakOffpeak"`
	RouteShortName string                                                  `json:"routeShortName"`
	ServiceID      string                                                  `json:"serviceId"`
	ShapeID        string                                                  `json:"shapeId"`
	TimeZone       string                                                  `json:"timeZone"`
	TripHeadsign   string                                                  `json:"tripHeadsign"`
	TripShortName  string                                                  `json:"tripShortName"`
	JSON           whereStopsForLocationListResponseDataReferencesTripJSON `json:"-"`
}

// whereStopsForLocationListResponseDataReferencesTripJSON contains the JSON
// metadata for the struct [WhereStopsForLocationListResponseDataReferencesTrip]
type whereStopsForLocationListResponseDataReferencesTripJSON struct {
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

func (r *WhereStopsForLocationListResponseDataReferencesTrip) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r whereStopsForLocationListResponseDataReferencesTripJSON) RawJSON() string {
	return r.raw
}

type WhereStopsForLocationListParams struct {
	Key param.Field[string]  `query:"key,required"`
	Lat param.Field[float64] `query:"lat"`
	Lon param.Field[float64] `query:"lon"`
}

// URLQuery serializes [WhereStopsForLocationListParams]'s query parameters as
// `url.Values`.
func (r WhereStopsForLocationListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
