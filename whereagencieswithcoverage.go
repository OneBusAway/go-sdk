// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package opentransit

import (
	"context"
	"net/http"

	"github.com/stainless-sdks/open-transit-go/internal/apijson"
	"github.com/stainless-sdks/open-transit-go/internal/requestconfig"
	"github.com/stainless-sdks/open-transit-go/option"
)

// WhereAgenciesWithCoverageService contains methods and other services that help
// with interacting with the open-transit API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWhereAgenciesWithCoverageService] method instead.
type WhereAgenciesWithCoverageService struct {
	Options []option.RequestOption
}

// NewWhereAgenciesWithCoverageService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewWhereAgenciesWithCoverageService(opts ...option.RequestOption) (r *WhereAgenciesWithCoverageService) {
	r = &WhereAgenciesWithCoverageService{}
	r.Options = opts
	return
}

// Retrieve Agencies with Coverage
func (r *WhereAgenciesWithCoverageService) List(ctx context.Context, opts ...option.RequestOption) (res *WhereAgenciesWithCoverageListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/where/agencies-with-coverage.json"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type WhereAgenciesWithCoverageListResponse struct {
	Code        int64                                     `json:"code,required"`
	CurrentTime int64                                     `json:"currentTime,required"`
	Text        string                                    `json:"text,required"`
	Version     int64                                     `json:"version,required"`
	Data        WhereAgenciesWithCoverageListResponseData `json:"data"`
	JSON        whereAgenciesWithCoverageListResponseJSON `json:"-"`
}

// whereAgenciesWithCoverageListResponseJSON contains the JSON metadata for the
// struct [WhereAgenciesWithCoverageListResponse]
type whereAgenciesWithCoverageListResponseJSON struct {
	Code        apijson.Field
	CurrentTime apijson.Field
	Text        apijson.Field
	Version     apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WhereAgenciesWithCoverageListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r whereAgenciesWithCoverageListResponseJSON) RawJSON() string {
	return r.raw
}

type WhereAgenciesWithCoverageListResponseData struct {
	LimitExceeded bool                                                `json:"limitExceeded"`
	List          []WhereAgenciesWithCoverageListResponseDataList     `json:"list"`
	References    WhereAgenciesWithCoverageListResponseDataReferences `json:"references"`
	JSON          whereAgenciesWithCoverageListResponseDataJSON       `json:"-"`
}

// whereAgenciesWithCoverageListResponseDataJSON contains the JSON metadata for the
// struct [WhereAgenciesWithCoverageListResponseData]
type whereAgenciesWithCoverageListResponseDataJSON struct {
	LimitExceeded apijson.Field
	List          apijson.Field
	References    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *WhereAgenciesWithCoverageListResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r whereAgenciesWithCoverageListResponseDataJSON) RawJSON() string {
	return r.raw
}

type WhereAgenciesWithCoverageListResponseDataList struct {
	AgencyID string                                            `json:"agencyId,required"`
	Lat      float64                                           `json:"lat,required"`
	LatSpan  float64                                           `json:"latSpan,required"`
	Lon      float64                                           `json:"lon,required"`
	LonSpan  float64                                           `json:"lonSpan,required"`
	JSON     whereAgenciesWithCoverageListResponseDataListJSON `json:"-"`
}

// whereAgenciesWithCoverageListResponseDataListJSON contains the JSON metadata for
// the struct [WhereAgenciesWithCoverageListResponseDataList]
type whereAgenciesWithCoverageListResponseDataListJSON struct {
	AgencyID    apijson.Field
	Lat         apijson.Field
	LatSpan     apijson.Field
	Lon         apijson.Field
	LonSpan     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WhereAgenciesWithCoverageListResponseDataList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r whereAgenciesWithCoverageListResponseDataListJSON) RawJSON() string {
	return r.raw
}

type WhereAgenciesWithCoverageListResponseDataReferences struct {
	Agencies   []WhereAgenciesWithCoverageListResponseDataReferencesAgency `json:"agencies"`
	Routes     []WhereAgenciesWithCoverageListResponseDataReferencesRoute  `json:"routes"`
	Situations []interface{}                                               `json:"situations"`
	Stops      []WhereAgenciesWithCoverageListResponseDataReferencesStop   `json:"stops"`
	StopTimes  []interface{}                                               `json:"stopTimes"`
	Trips      []WhereAgenciesWithCoverageListResponseDataReferencesTrip   `json:"trips"`
	JSON       whereAgenciesWithCoverageListResponseDataReferencesJSON     `json:"-"`
}

// whereAgenciesWithCoverageListResponseDataReferencesJSON contains the JSON
// metadata for the struct [WhereAgenciesWithCoverageListResponseDataReferences]
type whereAgenciesWithCoverageListResponseDataReferencesJSON struct {
	Agencies    apijson.Field
	Routes      apijson.Field
	Situations  apijson.Field
	Stops       apijson.Field
	StopTimes   apijson.Field
	Trips       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WhereAgenciesWithCoverageListResponseDataReferences) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r whereAgenciesWithCoverageListResponseDataReferencesJSON) RawJSON() string {
	return r.raw
}

type WhereAgenciesWithCoverageListResponseDataReferencesAgency struct {
	ID             string                                                        `json:"id,required"`
	Name           string                                                        `json:"name,required"`
	Timezone       string                                                        `json:"timezone,required"`
	URL            string                                                        `json:"url,required"`
	Disclaimer     string                                                        `json:"disclaimer"`
	Email          string                                                        `json:"email"`
	FareURL        string                                                        `json:"fareUrl"`
	Lang           string                                                        `json:"lang"`
	Phone          string                                                        `json:"phone"`
	PrivateService bool                                                          `json:"privateService"`
	JSON           whereAgenciesWithCoverageListResponseDataReferencesAgencyJSON `json:"-"`
}

// whereAgenciesWithCoverageListResponseDataReferencesAgencyJSON contains the JSON
// metadata for the struct
// [WhereAgenciesWithCoverageListResponseDataReferencesAgency]
type whereAgenciesWithCoverageListResponseDataReferencesAgencyJSON struct {
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

func (r *WhereAgenciesWithCoverageListResponseDataReferencesAgency) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r whereAgenciesWithCoverageListResponseDataReferencesAgencyJSON) RawJSON() string {
	return r.raw
}

type WhereAgenciesWithCoverageListResponseDataReferencesRoute struct {
	ID                string                                                       `json:"id"`
	AgencyID          string                                                       `json:"agencyId"`
	Color             string                                                       `json:"color"`
	Description       string                                                       `json:"description"`
	LongName          string                                                       `json:"longName"`
	NullSafeShortName string                                                       `json:"nullSafeShortName"`
	ShortName         string                                                       `json:"shortName"`
	TextColor         string                                                       `json:"textColor"`
	Type              int64                                                        `json:"type"`
	URL               string                                                       `json:"url"`
	JSON              whereAgenciesWithCoverageListResponseDataReferencesRouteJSON `json:"-"`
}

// whereAgenciesWithCoverageListResponseDataReferencesRouteJSON contains the JSON
// metadata for the struct
// [WhereAgenciesWithCoverageListResponseDataReferencesRoute]
type whereAgenciesWithCoverageListResponseDataReferencesRouteJSON struct {
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

func (r *WhereAgenciesWithCoverageListResponseDataReferencesRoute) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r whereAgenciesWithCoverageListResponseDataReferencesRouteJSON) RawJSON() string {
	return r.raw
}

type WhereAgenciesWithCoverageListResponseDataReferencesStop struct {
	ID                 string                                                      `json:"id,required"`
	Code               string                                                      `json:"code,required"`
	Lat                float64                                                     `json:"lat,required"`
	Lon                float64                                                     `json:"lon,required"`
	Name               string                                                      `json:"name,required"`
	Direction          string                                                      `json:"direction"`
	LocationType       int64                                                       `json:"locationType"`
	Parent             string                                                      `json:"parent"`
	RouteIDs           []string                                                    `json:"routeIds"`
	StaticRouteIDs     []string                                                    `json:"staticRouteIds"`
	WheelchairBoarding string                                                      `json:"wheelchairBoarding"`
	JSON               whereAgenciesWithCoverageListResponseDataReferencesStopJSON `json:"-"`
}

// whereAgenciesWithCoverageListResponseDataReferencesStopJSON contains the JSON
// metadata for the struct
// [WhereAgenciesWithCoverageListResponseDataReferencesStop]
type whereAgenciesWithCoverageListResponseDataReferencesStopJSON struct {
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

func (r *WhereAgenciesWithCoverageListResponseDataReferencesStop) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r whereAgenciesWithCoverageListResponseDataReferencesStopJSON) RawJSON() string {
	return r.raw
}

type WhereAgenciesWithCoverageListResponseDataReferencesTrip struct {
	ID             string                                                      `json:"id,required"`
	RouteID        string                                                      `json:"routeId,required"`
	BlockID        string                                                      `json:"blockId"`
	DirectionID    string                                                      `json:"directionId"`
	PeakOffpeak    int64                                                       `json:"peakOffpeak"`
	RouteShortName string                                                      `json:"routeShortName"`
	ServiceID      string                                                      `json:"serviceId"`
	ShapeID        string                                                      `json:"shapeId"`
	TimeZone       string                                                      `json:"timeZone"`
	TripHeadsign   string                                                      `json:"tripHeadsign"`
	TripShortName  string                                                      `json:"tripShortName"`
	JSON           whereAgenciesWithCoverageListResponseDataReferencesTripJSON `json:"-"`
}

// whereAgenciesWithCoverageListResponseDataReferencesTripJSON contains the JSON
// metadata for the struct
// [WhereAgenciesWithCoverageListResponseDataReferencesTrip]
type whereAgenciesWithCoverageListResponseDataReferencesTripJSON struct {
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

func (r *WhereAgenciesWithCoverageListResponseDataReferencesTrip) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r whereAgenciesWithCoverageListResponseDataReferencesTripJSON) RawJSON() string {
	return r.raw
}
