// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package opentransit

import (
	"context"
	"net/http"

	"github.com/stainless-sdks/open-transit-go/internal/apijson"
	"github.com/stainless-sdks/open-transit-go/internal/requestconfig"
	"github.com/stainless-sdks/open-transit-go/option"
)

// WhereCurrentTimeService contains methods and other services that help with
// interacting with the open-transit API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWhereCurrentTimeService] method instead.
type WhereCurrentTimeService struct {
	Options []option.RequestOption
}

// NewWhereCurrentTimeService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewWhereCurrentTimeService(opts ...option.RequestOption) (r *WhereCurrentTimeService) {
	r = &WhereCurrentTimeService{}
	r.Options = opts
	return
}

// current-time
func (r *WhereCurrentTimeService) Get(ctx context.Context, opts ...option.RequestOption) (res *WhereCurrentTimeGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/where/current-time.json"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type WhereCurrentTimeGetResponse struct {
	Code        int64                           `json:"code,required"`
	CurrentTime int64                           `json:"currentTime,required"`
	Text        string                          `json:"text,required"`
	Version     int64                           `json:"version,required"`
	Data        WhereCurrentTimeGetResponseData `json:"data"`
	JSON        whereCurrentTimeGetResponseJSON `json:"-"`
}

// whereCurrentTimeGetResponseJSON contains the JSON metadata for the struct
// [WhereCurrentTimeGetResponse]
type whereCurrentTimeGetResponseJSON struct {
	Code        apijson.Field
	CurrentTime apijson.Field
	Text        apijson.Field
	Version     apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WhereCurrentTimeGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r whereCurrentTimeGetResponseJSON) RawJSON() string {
	return r.raw
}

type WhereCurrentTimeGetResponseData struct {
	Entry      WhereCurrentTimeGetResponseDataEntry      `json:"entry"`
	References WhereCurrentTimeGetResponseDataReferences `json:"references"`
	JSON       whereCurrentTimeGetResponseDataJSON       `json:"-"`
}

// whereCurrentTimeGetResponseDataJSON contains the JSON metadata for the struct
// [WhereCurrentTimeGetResponseData]
type whereCurrentTimeGetResponseDataJSON struct {
	Entry       apijson.Field
	References  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WhereCurrentTimeGetResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r whereCurrentTimeGetResponseDataJSON) RawJSON() string {
	return r.raw
}

type WhereCurrentTimeGetResponseDataEntry struct {
	ReadableTime string                                   `json:"readableTime"`
	Time         int64                                    `json:"time"`
	JSON         whereCurrentTimeGetResponseDataEntryJSON `json:"-"`
}

// whereCurrentTimeGetResponseDataEntryJSON contains the JSON metadata for the
// struct [WhereCurrentTimeGetResponseDataEntry]
type whereCurrentTimeGetResponseDataEntryJSON struct {
	ReadableTime apijson.Field
	Time         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *WhereCurrentTimeGetResponseDataEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r whereCurrentTimeGetResponseDataEntryJSON) RawJSON() string {
	return r.raw
}

type WhereCurrentTimeGetResponseDataReferences struct {
	Agencies   []WhereCurrentTimeGetResponseDataReferencesAgency `json:"agencies"`
	Routes     []WhereCurrentTimeGetResponseDataReferencesRoute  `json:"routes"`
	Situations []interface{}                                     `json:"situations"`
	Stops      []WhereCurrentTimeGetResponseDataReferencesStop   `json:"stops"`
	StopTimes  []interface{}                                     `json:"stopTimes"`
	Trips      []WhereCurrentTimeGetResponseDataReferencesTrip   `json:"trips"`
	JSON       whereCurrentTimeGetResponseDataReferencesJSON     `json:"-"`
}

// whereCurrentTimeGetResponseDataReferencesJSON contains the JSON metadata for the
// struct [WhereCurrentTimeGetResponseDataReferences]
type whereCurrentTimeGetResponseDataReferencesJSON struct {
	Agencies    apijson.Field
	Routes      apijson.Field
	Situations  apijson.Field
	Stops       apijson.Field
	StopTimes   apijson.Field
	Trips       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WhereCurrentTimeGetResponseDataReferences) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r whereCurrentTimeGetResponseDataReferencesJSON) RawJSON() string {
	return r.raw
}

type WhereCurrentTimeGetResponseDataReferencesAgency struct {
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
	JSON           whereCurrentTimeGetResponseDataReferencesAgencyJSON `json:"-"`
}

// whereCurrentTimeGetResponseDataReferencesAgencyJSON contains the JSON metadata
// for the struct [WhereCurrentTimeGetResponseDataReferencesAgency]
type whereCurrentTimeGetResponseDataReferencesAgencyJSON struct {
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

func (r *WhereCurrentTimeGetResponseDataReferencesAgency) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r whereCurrentTimeGetResponseDataReferencesAgencyJSON) RawJSON() string {
	return r.raw
}

type WhereCurrentTimeGetResponseDataReferencesRoute struct {
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
	JSON              whereCurrentTimeGetResponseDataReferencesRouteJSON `json:"-"`
}

// whereCurrentTimeGetResponseDataReferencesRouteJSON contains the JSON metadata
// for the struct [WhereCurrentTimeGetResponseDataReferencesRoute]
type whereCurrentTimeGetResponseDataReferencesRouteJSON struct {
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

func (r *WhereCurrentTimeGetResponseDataReferencesRoute) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r whereCurrentTimeGetResponseDataReferencesRouteJSON) RawJSON() string {
	return r.raw
}

type WhereCurrentTimeGetResponseDataReferencesStop struct {
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
	JSON               whereCurrentTimeGetResponseDataReferencesStopJSON `json:"-"`
}

// whereCurrentTimeGetResponseDataReferencesStopJSON contains the JSON metadata for
// the struct [WhereCurrentTimeGetResponseDataReferencesStop]
type whereCurrentTimeGetResponseDataReferencesStopJSON struct {
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

func (r *WhereCurrentTimeGetResponseDataReferencesStop) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r whereCurrentTimeGetResponseDataReferencesStopJSON) RawJSON() string {
	return r.raw
}

type WhereCurrentTimeGetResponseDataReferencesTrip struct {
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
	JSON           whereCurrentTimeGetResponseDataReferencesTripJSON `json:"-"`
}

// whereCurrentTimeGetResponseDataReferencesTripJSON contains the JSON metadata for
// the struct [WhereCurrentTimeGetResponseDataReferencesTrip]
type whereCurrentTimeGetResponseDataReferencesTripJSON struct {
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

func (r *WhereCurrentTimeGetResponseDataReferencesTrip) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r whereCurrentTimeGetResponseDataReferencesTripJSON) RawJSON() string {
	return r.raw
}
