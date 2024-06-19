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

// APIWhereCurrentTimeService contains methods and other services that help with
// interacting with the OneBusAway API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAPIWhereCurrentTimeService] method instead.
type APIWhereCurrentTimeService struct {
	Options []option.RequestOption
}

// NewAPIWhereCurrentTimeService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAPIWhereCurrentTimeService(opts ...option.RequestOption) (r *APIWhereCurrentTimeService) {
	r = &APIWhereCurrentTimeService{}
	r.Options = opts
	return
}

// current-time
func (r *APIWhereCurrentTimeService) Get(ctx context.Context, opts ...option.RequestOption) (res *APIWhereCurrentTimeGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/where/current-time.json"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type APIWhereCurrentTimeGetResponse struct {
	Data APIWhereCurrentTimeGetResponseData `json:"data"`
	JSON apiWhereCurrentTimeGetResponseJSON `json:"-"`
	shared.ResponseWrapper
}

// apiWhereCurrentTimeGetResponseJSON contains the JSON metadata for the struct
// [APIWhereCurrentTimeGetResponse]
type apiWhereCurrentTimeGetResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIWhereCurrentTimeGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiWhereCurrentTimeGetResponseJSON) RawJSON() string {
	return r.raw
}

type APIWhereCurrentTimeGetResponseData struct {
	Entry      APIWhereCurrentTimeGetResponseDataEntry      `json:"entry"`
	References APIWhereCurrentTimeGetResponseDataReferences `json:"references"`
	JSON       apiWhereCurrentTimeGetResponseDataJSON       `json:"-"`
}

// apiWhereCurrentTimeGetResponseDataJSON contains the JSON metadata for the struct
// [APIWhereCurrentTimeGetResponseData]
type apiWhereCurrentTimeGetResponseDataJSON struct {
	Entry       apijson.Field
	References  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIWhereCurrentTimeGetResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiWhereCurrentTimeGetResponseDataJSON) RawJSON() string {
	return r.raw
}

type APIWhereCurrentTimeGetResponseDataEntry struct {
	ReadableTime string                                      `json:"readableTime"`
	Time         int64                                       `json:"time"`
	JSON         apiWhereCurrentTimeGetResponseDataEntryJSON `json:"-"`
}

// apiWhereCurrentTimeGetResponseDataEntryJSON contains the JSON metadata for the
// struct [APIWhereCurrentTimeGetResponseDataEntry]
type apiWhereCurrentTimeGetResponseDataEntryJSON struct {
	ReadableTime apijson.Field
	Time         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *APIWhereCurrentTimeGetResponseDataEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiWhereCurrentTimeGetResponseDataEntryJSON) RawJSON() string {
	return r.raw
}

type APIWhereCurrentTimeGetResponseDataReferences struct {
	Agencies   []APIWhereCurrentTimeGetResponseDataReferencesAgency `json:"agencies"`
	Routes     []APIWhereCurrentTimeGetResponseDataReferencesRoute  `json:"routes"`
	Situations []interface{}                                        `json:"situations"`
	Stops      []APIWhereCurrentTimeGetResponseDataReferencesStop   `json:"stops"`
	StopTimes  []interface{}                                        `json:"stopTimes"`
	Trips      []APIWhereCurrentTimeGetResponseDataReferencesTrip   `json:"trips"`
	JSON       apiWhereCurrentTimeGetResponseDataReferencesJSON     `json:"-"`
}

// apiWhereCurrentTimeGetResponseDataReferencesJSON contains the JSON metadata for
// the struct [APIWhereCurrentTimeGetResponseDataReferences]
type apiWhereCurrentTimeGetResponseDataReferencesJSON struct {
	Agencies    apijson.Field
	Routes      apijson.Field
	Situations  apijson.Field
	Stops       apijson.Field
	StopTimes   apijson.Field
	Trips       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIWhereCurrentTimeGetResponseDataReferences) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiWhereCurrentTimeGetResponseDataReferencesJSON) RawJSON() string {
	return r.raw
}

type APIWhereCurrentTimeGetResponseDataReferencesAgency struct {
	ID             string                                                 `json:"id,required"`
	Name           string                                                 `json:"name,required"`
	Timezone       string                                                 `json:"timezone,required"`
	URL            string                                                 `json:"url,required"`
	Disclaimer     string                                                 `json:"disclaimer"`
	Email          string                                                 `json:"email"`
	FareURL        string                                                 `json:"fareUrl"`
	Lang           string                                                 `json:"lang"`
	Phone          string                                                 `json:"phone"`
	PrivateService bool                                                   `json:"privateService"`
	JSON           apiWhereCurrentTimeGetResponseDataReferencesAgencyJSON `json:"-"`
}

// apiWhereCurrentTimeGetResponseDataReferencesAgencyJSON contains the JSON
// metadata for the struct [APIWhereCurrentTimeGetResponseDataReferencesAgency]
type apiWhereCurrentTimeGetResponseDataReferencesAgencyJSON struct {
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

func (r *APIWhereCurrentTimeGetResponseDataReferencesAgency) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiWhereCurrentTimeGetResponseDataReferencesAgencyJSON) RawJSON() string {
	return r.raw
}

type APIWhereCurrentTimeGetResponseDataReferencesRoute struct {
	ID                string                                                `json:"id"`
	AgencyID          string                                                `json:"agencyId"`
	Color             string                                                `json:"color"`
	Description       string                                                `json:"description"`
	LongName          string                                                `json:"longName"`
	NullSafeShortName string                                                `json:"nullSafeShortName"`
	ShortName         string                                                `json:"shortName"`
	TextColor         string                                                `json:"textColor"`
	Type              int64                                                 `json:"type"`
	URL               string                                                `json:"url"`
	JSON              apiWhereCurrentTimeGetResponseDataReferencesRouteJSON `json:"-"`
}

// apiWhereCurrentTimeGetResponseDataReferencesRouteJSON contains the JSON metadata
// for the struct [APIWhereCurrentTimeGetResponseDataReferencesRoute]
type apiWhereCurrentTimeGetResponseDataReferencesRouteJSON struct {
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

func (r *APIWhereCurrentTimeGetResponseDataReferencesRoute) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiWhereCurrentTimeGetResponseDataReferencesRouteJSON) RawJSON() string {
	return r.raw
}

type APIWhereCurrentTimeGetResponseDataReferencesStop struct {
	ID                 string                                               `json:"id,required"`
	Code               string                                               `json:"code,required"`
	Lat                float64                                              `json:"lat,required"`
	Lon                float64                                              `json:"lon,required"`
	Name               string                                               `json:"name,required"`
	Direction          string                                               `json:"direction"`
	LocationType       int64                                                `json:"locationType"`
	Parent             string                                               `json:"parent"`
	RouteIDs           []string                                             `json:"routeIds"`
	StaticRouteIDs     []string                                             `json:"staticRouteIds"`
	WheelchairBoarding string                                               `json:"wheelchairBoarding"`
	JSON               apiWhereCurrentTimeGetResponseDataReferencesStopJSON `json:"-"`
}

// apiWhereCurrentTimeGetResponseDataReferencesStopJSON contains the JSON metadata
// for the struct [APIWhereCurrentTimeGetResponseDataReferencesStop]
type apiWhereCurrentTimeGetResponseDataReferencesStopJSON struct {
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

func (r *APIWhereCurrentTimeGetResponseDataReferencesStop) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiWhereCurrentTimeGetResponseDataReferencesStopJSON) RawJSON() string {
	return r.raw
}

type APIWhereCurrentTimeGetResponseDataReferencesTrip struct {
	ID             string                                               `json:"id,required"`
	RouteID        string                                               `json:"routeId,required"`
	BlockID        string                                               `json:"blockId"`
	DirectionID    string                                               `json:"directionId"`
	PeakOffpeak    int64                                                `json:"peakOffpeak"`
	RouteShortName string                                               `json:"routeShortName"`
	ServiceID      string                                               `json:"serviceId"`
	ShapeID        string                                               `json:"shapeId"`
	TimeZone       string                                               `json:"timeZone"`
	TripHeadsign   string                                               `json:"tripHeadsign"`
	TripShortName  string                                               `json:"tripShortName"`
	JSON           apiWhereCurrentTimeGetResponseDataReferencesTripJSON `json:"-"`
}

// apiWhereCurrentTimeGetResponseDataReferencesTripJSON contains the JSON metadata
// for the struct [APIWhereCurrentTimeGetResponseDataReferencesTrip]
type apiWhereCurrentTimeGetResponseDataReferencesTripJSON struct {
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

func (r *APIWhereCurrentTimeGetResponseDataReferencesTrip) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiWhereCurrentTimeGetResponseDataReferencesTripJSON) RawJSON() string {
	return r.raw
}
