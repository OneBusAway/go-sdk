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

// CurrentTimeService contains methods and other services that help with
// interacting with the OneBusAway API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCurrentTimeService] method instead.
type CurrentTimeService struct {
	Options []option.RequestOption
}

// NewCurrentTimeService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCurrentTimeService(opts ...option.RequestOption) (r *CurrentTimeService) {
	r = &CurrentTimeService{}
	r.Options = opts
	return
}

// current-time
func (r *CurrentTimeService) Get(ctx context.Context, opts ...option.RequestOption) (res *CurrentTimeGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/where/current-time.json"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type CurrentTimeGetResponse struct {
	Data CurrentTimeGetResponseData `json:"data"`
	JSON currentTimeGetResponseJSON `json:"-"`
	shared.ResponseWrapper
}

// currentTimeGetResponseJSON contains the JSON metadata for the struct
// [CurrentTimeGetResponse]
type currentTimeGetResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CurrentTimeGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r currentTimeGetResponseJSON) RawJSON() string {
	return r.raw
}

type CurrentTimeGetResponseData struct {
	Entry      CurrentTimeGetResponseDataEntry      `json:"entry"`
	References CurrentTimeGetResponseDataReferences `json:"references"`
	JSON       currentTimeGetResponseDataJSON       `json:"-"`
}

// currentTimeGetResponseDataJSON contains the JSON metadata for the struct
// [CurrentTimeGetResponseData]
type currentTimeGetResponseDataJSON struct {
	Entry       apijson.Field
	References  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CurrentTimeGetResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r currentTimeGetResponseDataJSON) RawJSON() string {
	return r.raw
}

type CurrentTimeGetResponseDataEntry struct {
	ReadableTime string                              `json:"readableTime"`
	Time         int64                               `json:"time"`
	JSON         currentTimeGetResponseDataEntryJSON `json:"-"`
}

// currentTimeGetResponseDataEntryJSON contains the JSON metadata for the struct
// [CurrentTimeGetResponseDataEntry]
type currentTimeGetResponseDataEntryJSON struct {
	ReadableTime apijson.Field
	Time         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *CurrentTimeGetResponseDataEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r currentTimeGetResponseDataEntryJSON) RawJSON() string {
	return r.raw
}

type CurrentTimeGetResponseDataReferences struct {
	Agencies   []CurrentTimeGetResponseDataReferencesAgency `json:"agencies"`
	Routes     []CurrentTimeGetResponseDataReferencesRoute  `json:"routes"`
	Situations []interface{}                                `json:"situations"`
	Stops      []CurrentTimeGetResponseDataReferencesStop   `json:"stops"`
	StopTimes  []interface{}                                `json:"stopTimes"`
	Trips      []CurrentTimeGetResponseDataReferencesTrip   `json:"trips"`
	JSON       currentTimeGetResponseDataReferencesJSON     `json:"-"`
}

// currentTimeGetResponseDataReferencesJSON contains the JSON metadata for the
// struct [CurrentTimeGetResponseDataReferences]
type currentTimeGetResponseDataReferencesJSON struct {
	Agencies    apijson.Field
	Routes      apijson.Field
	Situations  apijson.Field
	Stops       apijson.Field
	StopTimes   apijson.Field
	Trips       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CurrentTimeGetResponseDataReferences) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r currentTimeGetResponseDataReferencesJSON) RawJSON() string {
	return r.raw
}

type CurrentTimeGetResponseDataReferencesAgency struct {
	ID             string                                         `json:"id,required"`
	Name           string                                         `json:"name,required"`
	Timezone       string                                         `json:"timezone,required"`
	URL            string                                         `json:"url,required"`
	Disclaimer     string                                         `json:"disclaimer"`
	Email          string                                         `json:"email"`
	FareURL        string                                         `json:"fareUrl"`
	Lang           string                                         `json:"lang"`
	Phone          string                                         `json:"phone"`
	PrivateService bool                                           `json:"privateService"`
	JSON           currentTimeGetResponseDataReferencesAgencyJSON `json:"-"`
}

// currentTimeGetResponseDataReferencesAgencyJSON contains the JSON metadata for
// the struct [CurrentTimeGetResponseDataReferencesAgency]
type currentTimeGetResponseDataReferencesAgencyJSON struct {
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

func (r *CurrentTimeGetResponseDataReferencesAgency) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r currentTimeGetResponseDataReferencesAgencyJSON) RawJSON() string {
	return r.raw
}

type CurrentTimeGetResponseDataReferencesRoute struct {
	ID                string                                        `json:"id"`
	AgencyID          string                                        `json:"agencyId"`
	Color             string                                        `json:"color"`
	Description       string                                        `json:"description"`
	LongName          string                                        `json:"longName"`
	NullSafeShortName string                                        `json:"nullSafeShortName"`
	ShortName         string                                        `json:"shortName"`
	TextColor         string                                        `json:"textColor"`
	Type              int64                                         `json:"type"`
	URL               string                                        `json:"url"`
	JSON              currentTimeGetResponseDataReferencesRouteJSON `json:"-"`
}

// currentTimeGetResponseDataReferencesRouteJSON contains the JSON metadata for the
// struct [CurrentTimeGetResponseDataReferencesRoute]
type currentTimeGetResponseDataReferencesRouteJSON struct {
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

func (r *CurrentTimeGetResponseDataReferencesRoute) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r currentTimeGetResponseDataReferencesRouteJSON) RawJSON() string {
	return r.raw
}

type CurrentTimeGetResponseDataReferencesStop struct {
	ID                 string                                       `json:"id,required"`
	Code               string                                       `json:"code,required"`
	Lat                float64                                      `json:"lat,required"`
	Lon                float64                                      `json:"lon,required"`
	Name               string                                       `json:"name,required"`
	Direction          string                                       `json:"direction"`
	LocationType       int64                                        `json:"locationType"`
	Parent             string                                       `json:"parent"`
	RouteIDs           []string                                     `json:"routeIds"`
	StaticRouteIDs     []string                                     `json:"staticRouteIds"`
	WheelchairBoarding string                                       `json:"wheelchairBoarding"`
	JSON               currentTimeGetResponseDataReferencesStopJSON `json:"-"`
}

// currentTimeGetResponseDataReferencesStopJSON contains the JSON metadata for the
// struct [CurrentTimeGetResponseDataReferencesStop]
type currentTimeGetResponseDataReferencesStopJSON struct {
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

func (r *CurrentTimeGetResponseDataReferencesStop) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r currentTimeGetResponseDataReferencesStopJSON) RawJSON() string {
	return r.raw
}

type CurrentTimeGetResponseDataReferencesTrip struct {
	ID             string                                       `json:"id,required"`
	RouteID        string                                       `json:"routeId,required"`
	BlockID        string                                       `json:"blockId"`
	DirectionID    string                                       `json:"directionId"`
	PeakOffpeak    int64                                        `json:"peakOffpeak"`
	RouteShortName string                                       `json:"routeShortName"`
	ServiceID      string                                       `json:"serviceId"`
	ShapeID        string                                       `json:"shapeId"`
	TimeZone       string                                       `json:"timeZone"`
	TripHeadsign   string                                       `json:"tripHeadsign"`
	TripShortName  string                                       `json:"tripShortName"`
	JSON           currentTimeGetResponseDataReferencesTripJSON `json:"-"`
}

// currentTimeGetResponseDataReferencesTripJSON contains the JSON metadata for the
// struct [CurrentTimeGetResponseDataReferencesTrip]
type currentTimeGetResponseDataReferencesTripJSON struct {
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

func (r *CurrentTimeGetResponseDataReferencesTrip) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r currentTimeGetResponseDataReferencesTripJSON) RawJSON() string {
	return r.raw
}
