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
func (r *StopsForLocationService) Get(ctx context.Context, query StopsForLocationGetParams, opts ...option.RequestOption) (res *StopsForLocationGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/where/stops-for-location.json"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type StopsForLocationGetResponse struct {
	Code        int64                           `json:"code,required"`
	CurrentTime int64                           `json:"currentTime,required"`
	Text        string                          `json:"text,required"`
	Version     int64                           `json:"version,required"`
	Data        StopsForLocationGetResponseData `json:"data"`
	JSON        stopsForLocationGetResponseJSON `json:"-"`
}

// stopsForLocationGetResponseJSON contains the JSON metadata for the struct
// [StopsForLocationGetResponse]
type stopsForLocationGetResponseJSON struct {
	Code        apijson.Field
	CurrentTime apijson.Field
	Text        apijson.Field
	Version     apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StopsForLocationGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stopsForLocationGetResponseJSON) RawJSON() string {
	return r.raw
}

type StopsForLocationGetResponseData struct {
	LimitExceeded bool                                      `json:"limitExceeded"`
	List          []StopsForLocationGetResponseDataList     `json:"list"`
	References    StopsForLocationGetResponseDataReferences `json:"references"`
	JSON          stopsForLocationGetResponseDataJSON       `json:"-"`
}

// stopsForLocationGetResponseDataJSON contains the JSON metadata for the struct
// [StopsForLocationGetResponseData]
type stopsForLocationGetResponseDataJSON struct {
	LimitExceeded apijson.Field
	List          apijson.Field
	References    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *StopsForLocationGetResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stopsForLocationGetResponseDataJSON) RawJSON() string {
	return r.raw
}

type StopsForLocationGetResponseDataList struct {
	ID                 string                                  `json:"id"`
	Code               string                                  `json:"code"`
	Direction          string                                  `json:"direction"`
	Lat                float64                                 `json:"lat"`
	LocationType       int64                                   `json:"locationType"`
	Lon                float64                                 `json:"lon"`
	Name               string                                  `json:"name"`
	Parent             string                                  `json:"parent"`
	RouteIDs           []string                                `json:"routeIds"`
	StaticRouteIDs     []string                                `json:"staticRouteIds"`
	WheelchairBoarding string                                  `json:"wheelchairBoarding"`
	JSON               stopsForLocationGetResponseDataListJSON `json:"-"`
}

// stopsForLocationGetResponseDataListJSON contains the JSON metadata for the
// struct [StopsForLocationGetResponseDataList]
type stopsForLocationGetResponseDataListJSON struct {
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

func (r *StopsForLocationGetResponseDataList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stopsForLocationGetResponseDataListJSON) RawJSON() string {
	return r.raw
}

type StopsForLocationGetResponseDataReferences struct {
	Agencies   []StopsForLocationGetResponseDataReferencesAgency `json:"agencies"`
	Routes     []StopsForLocationGetResponseDataReferencesRoute  `json:"routes"`
	Situations []interface{}                                     `json:"situations"`
	Stops      []StopsForLocationGetResponseDataReferencesStop   `json:"stops"`
	StopTimes  []interface{}                                     `json:"stopTimes"`
	Trips      []StopsForLocationGetResponseDataReferencesTrip   `json:"trips"`
	JSON       stopsForLocationGetResponseDataReferencesJSON     `json:"-"`
}

// stopsForLocationGetResponseDataReferencesJSON contains the JSON metadata for the
// struct [StopsForLocationGetResponseDataReferences]
type stopsForLocationGetResponseDataReferencesJSON struct {
	Agencies    apijson.Field
	Routes      apijson.Field
	Situations  apijson.Field
	Stops       apijson.Field
	StopTimes   apijson.Field
	Trips       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StopsForLocationGetResponseDataReferences) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stopsForLocationGetResponseDataReferencesJSON) RawJSON() string {
	return r.raw
}

type StopsForLocationGetResponseDataReferencesAgency struct {
	ID             string                                              `json:"id,required"`
	Name           string                                              `json:"name,required"`
	Timezone       string                                              `json:"timezone,required"`
	URL            string                                              `json:"url,required"`
	Disclaimer     string                                              `json:"disclaimer"`
	Email          string                                              `json:"email"`
	FareURL        string                                              `json:"fareUrl"`
	Lang           string                                              `json:"lang"`
	Phone          string                                              `json:"phone"`
	PrivateService bool                                                `json:"privateService"`
	JSON           stopsForLocationGetResponseDataReferencesAgencyJSON `json:"-"`
}

// stopsForLocationGetResponseDataReferencesAgencyJSON contains the JSON metadata
// for the struct [StopsForLocationGetResponseDataReferencesAgency]
type stopsForLocationGetResponseDataReferencesAgencyJSON struct {
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

func (r *StopsForLocationGetResponseDataReferencesAgency) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stopsForLocationGetResponseDataReferencesAgencyJSON) RawJSON() string {
	return r.raw
}

type StopsForLocationGetResponseDataReferencesRoute struct {
	ID                string                                             `json:"id"`
	AgencyID          string                                             `json:"agencyId"`
	Color             string                                             `json:"color"`
	Description       string                                             `json:"description"`
	LongName          string                                             `json:"longName"`
	NullSafeShortName string                                             `json:"nullSafeShortName"`
	ShortName         string                                             `json:"shortName"`
	TextColor         string                                             `json:"textColor"`
	Type              int64                                              `json:"type"`
	URL               string                                             `json:"url"`
	JSON              stopsForLocationGetResponseDataReferencesRouteJSON `json:"-"`
}

// stopsForLocationGetResponseDataReferencesRouteJSON contains the JSON metadata
// for the struct [StopsForLocationGetResponseDataReferencesRoute]
type stopsForLocationGetResponseDataReferencesRouteJSON struct {
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

func (r *StopsForLocationGetResponseDataReferencesRoute) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stopsForLocationGetResponseDataReferencesRouteJSON) RawJSON() string {
	return r.raw
}

type StopsForLocationGetResponseDataReferencesStop struct {
	ID                 string                                            `json:"id,required"`
	Code               string                                            `json:"code,required"`
	Lat                float64                                           `json:"lat,required"`
	Lon                float64                                           `json:"lon,required"`
	Name               string                                            `json:"name,required"`
	Direction          string                                            `json:"direction"`
	LocationType       int64                                             `json:"locationType"`
	Parent             string                                            `json:"parent"`
	RouteIDs           []string                                          `json:"routeIds"`
	StaticRouteIDs     []string                                          `json:"staticRouteIds"`
	WheelchairBoarding string                                            `json:"wheelchairBoarding"`
	JSON               stopsForLocationGetResponseDataReferencesStopJSON `json:"-"`
}

// stopsForLocationGetResponseDataReferencesStopJSON contains the JSON metadata for
// the struct [StopsForLocationGetResponseDataReferencesStop]
type stopsForLocationGetResponseDataReferencesStopJSON struct {
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

func (r *StopsForLocationGetResponseDataReferencesStop) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stopsForLocationGetResponseDataReferencesStopJSON) RawJSON() string {
	return r.raw
}

type StopsForLocationGetResponseDataReferencesTrip struct {
	ID             string                                            `json:"id,required"`
	RouteID        string                                            `json:"routeId,required"`
	BlockID        string                                            `json:"blockId"`
	DirectionID    string                                            `json:"directionId"`
	PeakOffpeak    int64                                             `json:"peakOffpeak"`
	RouteShortName string                                            `json:"routeShortName"`
	ServiceID      string                                            `json:"serviceId"`
	ShapeID        string                                            `json:"shapeId"`
	TimeZone       string                                            `json:"timeZone"`
	TripHeadsign   string                                            `json:"tripHeadsign"`
	TripShortName  string                                            `json:"tripShortName"`
	JSON           stopsForLocationGetResponseDataReferencesTripJSON `json:"-"`
}

// stopsForLocationGetResponseDataReferencesTripJSON contains the JSON metadata for
// the struct [StopsForLocationGetResponseDataReferencesTrip]
type stopsForLocationGetResponseDataReferencesTripJSON struct {
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

func (r *StopsForLocationGetResponseDataReferencesTrip) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stopsForLocationGetResponseDataReferencesTripJSON) RawJSON() string {
	return r.raw
}

type StopsForLocationGetParams struct {
	Key param.Field[string]  `query:"key,required"`
	Lat param.Field[float64] `query:"lat"`
	Lon param.Field[float64] `query:"lon"`
}

// URLQuery serializes [StopsForLocationGetParams]'s query parameters as
// `url.Values`.
func (r StopsForLocationGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
