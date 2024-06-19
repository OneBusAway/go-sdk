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

// APIWhereConfigService contains methods and other services that help with
// interacting with the OneBusAway API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAPIWhereConfigService] method instead.
type APIWhereConfigService struct {
	Options []option.RequestOption
}

// NewAPIWhereConfigService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAPIWhereConfigService(opts ...option.RequestOption) (r *APIWhereConfigService) {
	r = &APIWhereConfigService{}
	r.Options = opts
	return
}

// config
func (r *APIWhereConfigService) Get(ctx context.Context, opts ...option.RequestOption) (res *APIWhereConfigGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/where/config.json"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type APIWhereConfigGetResponse struct {
	Data APIWhereConfigGetResponseData `json:"data"`
	JSON apiWhereConfigGetResponseJSON `json:"-"`
	shared.ResponseWrapper
}

// apiWhereConfigGetResponseJSON contains the JSON metadata for the struct
// [APIWhereConfigGetResponse]
type apiWhereConfigGetResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIWhereConfigGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiWhereConfigGetResponseJSON) RawJSON() string {
	return r.raw
}

type APIWhereConfigGetResponseData struct {
	Entry      APIWhereConfigGetResponseDataEntry      `json:"entry"`
	References APIWhereConfigGetResponseDataReferences `json:"references"`
	JSON       apiWhereConfigGetResponseDataJSON       `json:"-"`
}

// apiWhereConfigGetResponseDataJSON contains the JSON metadata for the struct
// [APIWhereConfigGetResponseData]
type apiWhereConfigGetResponseDataJSON struct {
	Entry       apijson.Field
	References  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIWhereConfigGetResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiWhereConfigGetResponseDataJSON) RawJSON() string {
	return r.raw
}

type APIWhereConfigGetResponseDataEntry struct {
	ID              string                                          `json:"id"`
	GitProperties   APIWhereConfigGetResponseDataEntryGitProperties `json:"gitProperties"`
	Name            string                                          `json:"name"`
	ServiceDateFrom string                                          `json:"serviceDateFrom"`
	ServiceDateTo   string                                          `json:"serviceDateTo"`
	JSON            apiWhereConfigGetResponseDataEntryJSON          `json:"-"`
}

// apiWhereConfigGetResponseDataEntryJSON contains the JSON metadata for the struct
// [APIWhereConfigGetResponseDataEntry]
type apiWhereConfigGetResponseDataEntryJSON struct {
	ID              apijson.Field
	GitProperties   apijson.Field
	Name            apijson.Field
	ServiceDateFrom apijson.Field
	ServiceDateTo   apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *APIWhereConfigGetResponseDataEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiWhereConfigGetResponseDataEntryJSON) RawJSON() string {
	return r.raw
}

type APIWhereConfigGetResponseDataEntryGitProperties struct {
	GitBranch                string                                              `json:"git.branch"`
	GitBuildHost             string                                              `json:"git.build.host"`
	GitBuildTime             string                                              `json:"git.build.time"`
	GitBuildUserEmail        string                                              `json:"git.build.user.email"`
	GitBuildUserName         string                                              `json:"git.build.user.name"`
	GitBuildVersion          string                                              `json:"git.build.version"`
	GitClosestTagCommitCount string                                              `json:"git.closest.tag.commit.count"`
	GitClosestTagName        string                                              `json:"git.closest.tag.name"`
	GitCommitID              string                                              `json:"git.commit.id"`
	GitCommitIDAbbrev        string                                              `json:"git.commit.id.abbrev"`
	GitCommitIDDescribe      string                                              `json:"git.commit.id.describe"`
	GitCommitIDDescribeShort string                                              `json:"git.commit.id.describe-short"`
	GitCommitMessageFull     string                                              `json:"git.commit.message.full"`
	GitCommitMessageShort    string                                              `json:"git.commit.message.short"`
	GitCommitTime            string                                              `json:"git.commit.time"`
	GitCommitUserEmail       string                                              `json:"git.commit.user.email"`
	GitCommitUserName        string                                              `json:"git.commit.user.name"`
	GitDirty                 string                                              `json:"git.dirty"`
	GitRemoteOriginURL       string                                              `json:"git.remote.origin.url"`
	GitTags                  string                                              `json:"git.tags"`
	JSON                     apiWhereConfigGetResponseDataEntryGitPropertiesJSON `json:"-"`
}

// apiWhereConfigGetResponseDataEntryGitPropertiesJSON contains the JSON metadata
// for the struct [APIWhereConfigGetResponseDataEntryGitProperties]
type apiWhereConfigGetResponseDataEntryGitPropertiesJSON struct {
	GitBranch                apijson.Field
	GitBuildHost             apijson.Field
	GitBuildTime             apijson.Field
	GitBuildUserEmail        apijson.Field
	GitBuildUserName         apijson.Field
	GitBuildVersion          apijson.Field
	GitClosestTagCommitCount apijson.Field
	GitClosestTagName        apijson.Field
	GitCommitID              apijson.Field
	GitCommitIDAbbrev        apijson.Field
	GitCommitIDDescribe      apijson.Field
	GitCommitIDDescribeShort apijson.Field
	GitCommitMessageFull     apijson.Field
	GitCommitMessageShort    apijson.Field
	GitCommitTime            apijson.Field
	GitCommitUserEmail       apijson.Field
	GitCommitUserName        apijson.Field
	GitDirty                 apijson.Field
	GitRemoteOriginURL       apijson.Field
	GitTags                  apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *APIWhereConfigGetResponseDataEntryGitProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiWhereConfigGetResponseDataEntryGitPropertiesJSON) RawJSON() string {
	return r.raw
}

type APIWhereConfigGetResponseDataReferences struct {
	Agencies   []APIWhereConfigGetResponseDataReferencesAgency `json:"agencies"`
	Routes     []APIWhereConfigGetResponseDataReferencesRoute  `json:"routes"`
	Situations []interface{}                                   `json:"situations"`
	Stops      []APIWhereConfigGetResponseDataReferencesStop   `json:"stops"`
	StopTimes  []interface{}                                   `json:"stopTimes"`
	Trips      []APIWhereConfigGetResponseDataReferencesTrip   `json:"trips"`
	JSON       apiWhereConfigGetResponseDataReferencesJSON     `json:"-"`
}

// apiWhereConfigGetResponseDataReferencesJSON contains the JSON metadata for the
// struct [APIWhereConfigGetResponseDataReferences]
type apiWhereConfigGetResponseDataReferencesJSON struct {
	Agencies    apijson.Field
	Routes      apijson.Field
	Situations  apijson.Field
	Stops       apijson.Field
	StopTimes   apijson.Field
	Trips       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIWhereConfigGetResponseDataReferences) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiWhereConfigGetResponseDataReferencesJSON) RawJSON() string {
	return r.raw
}

type APIWhereConfigGetResponseDataReferencesAgency struct {
	ID             string                                            `json:"id,required"`
	Name           string                                            `json:"name,required"`
	Timezone       string                                            `json:"timezone,required"`
	URL            string                                            `json:"url,required"`
	Disclaimer     string                                            `json:"disclaimer"`
	Email          string                                            `json:"email"`
	FareURL        string                                            `json:"fareUrl"`
	Lang           string                                            `json:"lang"`
	Phone          string                                            `json:"phone"`
	PrivateService bool                                              `json:"privateService"`
	JSON           apiWhereConfigGetResponseDataReferencesAgencyJSON `json:"-"`
}

// apiWhereConfigGetResponseDataReferencesAgencyJSON contains the JSON metadata for
// the struct [APIWhereConfigGetResponseDataReferencesAgency]
type apiWhereConfigGetResponseDataReferencesAgencyJSON struct {
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

func (r *APIWhereConfigGetResponseDataReferencesAgency) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiWhereConfigGetResponseDataReferencesAgencyJSON) RawJSON() string {
	return r.raw
}

type APIWhereConfigGetResponseDataReferencesRoute struct {
	ID                string                                           `json:"id"`
	AgencyID          string                                           `json:"agencyId"`
	Color             string                                           `json:"color"`
	Description       string                                           `json:"description"`
	LongName          string                                           `json:"longName"`
	NullSafeShortName string                                           `json:"nullSafeShortName"`
	ShortName         string                                           `json:"shortName"`
	TextColor         string                                           `json:"textColor"`
	Type              int64                                            `json:"type"`
	URL               string                                           `json:"url"`
	JSON              apiWhereConfigGetResponseDataReferencesRouteJSON `json:"-"`
}

// apiWhereConfigGetResponseDataReferencesRouteJSON contains the JSON metadata for
// the struct [APIWhereConfigGetResponseDataReferencesRoute]
type apiWhereConfigGetResponseDataReferencesRouteJSON struct {
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

func (r *APIWhereConfigGetResponseDataReferencesRoute) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiWhereConfigGetResponseDataReferencesRouteJSON) RawJSON() string {
	return r.raw
}

type APIWhereConfigGetResponseDataReferencesStop struct {
	ID                 string                                          `json:"id,required"`
	Code               string                                          `json:"code,required"`
	Lat                float64                                         `json:"lat,required"`
	Lon                float64                                         `json:"lon,required"`
	Name               string                                          `json:"name,required"`
	Direction          string                                          `json:"direction"`
	LocationType       int64                                           `json:"locationType"`
	Parent             string                                          `json:"parent"`
	RouteIDs           []string                                        `json:"routeIds"`
	StaticRouteIDs     []string                                        `json:"staticRouteIds"`
	WheelchairBoarding string                                          `json:"wheelchairBoarding"`
	JSON               apiWhereConfigGetResponseDataReferencesStopJSON `json:"-"`
}

// apiWhereConfigGetResponseDataReferencesStopJSON contains the JSON metadata for
// the struct [APIWhereConfigGetResponseDataReferencesStop]
type apiWhereConfigGetResponseDataReferencesStopJSON struct {
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

func (r *APIWhereConfigGetResponseDataReferencesStop) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiWhereConfigGetResponseDataReferencesStopJSON) RawJSON() string {
	return r.raw
}

type APIWhereConfigGetResponseDataReferencesTrip struct {
	ID             string                                          `json:"id,required"`
	RouteID        string                                          `json:"routeId,required"`
	BlockID        string                                          `json:"blockId"`
	DirectionID    string                                          `json:"directionId"`
	PeakOffpeak    int64                                           `json:"peakOffpeak"`
	RouteShortName string                                          `json:"routeShortName"`
	ServiceID      string                                          `json:"serviceId"`
	ShapeID        string                                          `json:"shapeId"`
	TimeZone       string                                          `json:"timeZone"`
	TripHeadsign   string                                          `json:"tripHeadsign"`
	TripShortName  string                                          `json:"tripShortName"`
	JSON           apiWhereConfigGetResponseDataReferencesTripJSON `json:"-"`
}

// apiWhereConfigGetResponseDataReferencesTripJSON contains the JSON metadata for
// the struct [APIWhereConfigGetResponseDataReferencesTrip]
type apiWhereConfigGetResponseDataReferencesTripJSON struct {
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

func (r *APIWhereConfigGetResponseDataReferencesTrip) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiWhereConfigGetResponseDataReferencesTripJSON) RawJSON() string {
	return r.raw
}
