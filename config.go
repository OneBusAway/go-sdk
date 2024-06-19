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

// ConfigService contains methods and other services that help with interacting
// with the OneBusAway API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewConfigService] method instead.
type ConfigService struct {
	Options []option.RequestOption
}

// NewConfigService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewConfigService(opts ...option.RequestOption) (r *ConfigService) {
	r = &ConfigService{}
	r.Options = opts
	return
}

// config
func (r *ConfigService) Get(ctx context.Context, opts ...option.RequestOption) (res *ConfigGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/where/config.json"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ConfigGetResponse struct {
	Data ConfigGetResponseData `json:"data"`
	JSON configGetResponseJSON `json:"-"`
	shared.ResponseWrapper
}

// configGetResponseJSON contains the JSON metadata for the struct
// [ConfigGetResponse]
type configGetResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configGetResponseJSON) RawJSON() string {
	return r.raw
}

type ConfigGetResponseData struct {
	Entry      ConfigGetResponseDataEntry      `json:"entry"`
	References ConfigGetResponseDataReferences `json:"references"`
	JSON       configGetResponseDataJSON       `json:"-"`
}

// configGetResponseDataJSON contains the JSON metadata for the struct
// [ConfigGetResponseData]
type configGetResponseDataJSON struct {
	Entry       apijson.Field
	References  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigGetResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configGetResponseDataJSON) RawJSON() string {
	return r.raw
}

type ConfigGetResponseDataEntry struct {
	ID              string                                  `json:"id"`
	GitProperties   ConfigGetResponseDataEntryGitProperties `json:"gitProperties"`
	Name            string                                  `json:"name"`
	ServiceDateFrom string                                  `json:"serviceDateFrom"`
	ServiceDateTo   string                                  `json:"serviceDateTo"`
	JSON            configGetResponseDataEntryJSON          `json:"-"`
}

// configGetResponseDataEntryJSON contains the JSON metadata for the struct
// [ConfigGetResponseDataEntry]
type configGetResponseDataEntryJSON struct {
	ID              apijson.Field
	GitProperties   apijson.Field
	Name            apijson.Field
	ServiceDateFrom apijson.Field
	ServiceDateTo   apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ConfigGetResponseDataEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configGetResponseDataEntryJSON) RawJSON() string {
	return r.raw
}

type ConfigGetResponseDataEntryGitProperties struct {
	GitBranch                string                                      `json:"git.branch"`
	GitBuildHost             string                                      `json:"git.build.host"`
	GitBuildTime             string                                      `json:"git.build.time"`
	GitBuildUserEmail        string                                      `json:"git.build.user.email"`
	GitBuildUserName         string                                      `json:"git.build.user.name"`
	GitBuildVersion          string                                      `json:"git.build.version"`
	GitClosestTagCommitCount string                                      `json:"git.closest.tag.commit.count"`
	GitClosestTagName        string                                      `json:"git.closest.tag.name"`
	GitCommitID              string                                      `json:"git.commit.id"`
	GitCommitIDAbbrev        string                                      `json:"git.commit.id.abbrev"`
	GitCommitIDDescribe      string                                      `json:"git.commit.id.describe"`
	GitCommitIDDescribeShort string                                      `json:"git.commit.id.describe-short"`
	GitCommitMessageFull     string                                      `json:"git.commit.message.full"`
	GitCommitMessageShort    string                                      `json:"git.commit.message.short"`
	GitCommitTime            string                                      `json:"git.commit.time"`
	GitCommitUserEmail       string                                      `json:"git.commit.user.email"`
	GitCommitUserName        string                                      `json:"git.commit.user.name"`
	GitDirty                 string                                      `json:"git.dirty"`
	GitRemoteOriginURL       string                                      `json:"git.remote.origin.url"`
	GitTags                  string                                      `json:"git.tags"`
	JSON                     configGetResponseDataEntryGitPropertiesJSON `json:"-"`
}

// configGetResponseDataEntryGitPropertiesJSON contains the JSON metadata for the
// struct [ConfigGetResponseDataEntryGitProperties]
type configGetResponseDataEntryGitPropertiesJSON struct {
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

func (r *ConfigGetResponseDataEntryGitProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configGetResponseDataEntryGitPropertiesJSON) RawJSON() string {
	return r.raw
}

type ConfigGetResponseDataReferences struct {
	Agencies   []ConfigGetResponseDataReferencesAgency `json:"agencies"`
	Routes     []ConfigGetResponseDataReferencesRoute  `json:"routes"`
	Situations []interface{}                           `json:"situations"`
	Stops      []ConfigGetResponseDataReferencesStop   `json:"stops"`
	StopTimes  []interface{}                           `json:"stopTimes"`
	Trips      []ConfigGetResponseDataReferencesTrip   `json:"trips"`
	JSON       configGetResponseDataReferencesJSON     `json:"-"`
}

// configGetResponseDataReferencesJSON contains the JSON metadata for the struct
// [ConfigGetResponseDataReferences]
type configGetResponseDataReferencesJSON struct {
	Agencies    apijson.Field
	Routes      apijson.Field
	Situations  apijson.Field
	Stops       apijson.Field
	StopTimes   apijson.Field
	Trips       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigGetResponseDataReferences) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configGetResponseDataReferencesJSON) RawJSON() string {
	return r.raw
}

type ConfigGetResponseDataReferencesAgency struct {
	ID             string                                    `json:"id,required"`
	Name           string                                    `json:"name,required"`
	Timezone       string                                    `json:"timezone,required"`
	URL            string                                    `json:"url,required"`
	Disclaimer     string                                    `json:"disclaimer"`
	Email          string                                    `json:"email"`
	FareURL        string                                    `json:"fareUrl"`
	Lang           string                                    `json:"lang"`
	Phone          string                                    `json:"phone"`
	PrivateService bool                                      `json:"privateService"`
	JSON           configGetResponseDataReferencesAgencyJSON `json:"-"`
}

// configGetResponseDataReferencesAgencyJSON contains the JSON metadata for the
// struct [ConfigGetResponseDataReferencesAgency]
type configGetResponseDataReferencesAgencyJSON struct {
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

func (r *ConfigGetResponseDataReferencesAgency) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configGetResponseDataReferencesAgencyJSON) RawJSON() string {
	return r.raw
}

type ConfigGetResponseDataReferencesRoute struct {
	ID                string                                   `json:"id"`
	AgencyID          string                                   `json:"agencyId"`
	Color             string                                   `json:"color"`
	Description       string                                   `json:"description"`
	LongName          string                                   `json:"longName"`
	NullSafeShortName string                                   `json:"nullSafeShortName"`
	ShortName         string                                   `json:"shortName"`
	TextColor         string                                   `json:"textColor"`
	Type              int64                                    `json:"type"`
	URL               string                                   `json:"url"`
	JSON              configGetResponseDataReferencesRouteJSON `json:"-"`
}

// configGetResponseDataReferencesRouteJSON contains the JSON metadata for the
// struct [ConfigGetResponseDataReferencesRoute]
type configGetResponseDataReferencesRouteJSON struct {
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

func (r *ConfigGetResponseDataReferencesRoute) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configGetResponseDataReferencesRouteJSON) RawJSON() string {
	return r.raw
}

type ConfigGetResponseDataReferencesStop struct {
	ID                 string                                  `json:"id,required"`
	Code               string                                  `json:"code,required"`
	Lat                float64                                 `json:"lat,required"`
	Lon                float64                                 `json:"lon,required"`
	Name               string                                  `json:"name,required"`
	Direction          string                                  `json:"direction"`
	LocationType       int64                                   `json:"locationType"`
	Parent             string                                  `json:"parent"`
	RouteIDs           []string                                `json:"routeIds"`
	StaticRouteIDs     []string                                `json:"staticRouteIds"`
	WheelchairBoarding string                                  `json:"wheelchairBoarding"`
	JSON               configGetResponseDataReferencesStopJSON `json:"-"`
}

// configGetResponseDataReferencesStopJSON contains the JSON metadata for the
// struct [ConfigGetResponseDataReferencesStop]
type configGetResponseDataReferencesStopJSON struct {
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

func (r *ConfigGetResponseDataReferencesStop) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configGetResponseDataReferencesStopJSON) RawJSON() string {
	return r.raw
}

type ConfigGetResponseDataReferencesTrip struct {
	ID             string                                  `json:"id,required"`
	RouteID        string                                  `json:"routeId,required"`
	BlockID        string                                  `json:"blockId"`
	DirectionID    string                                  `json:"directionId"`
	PeakOffpeak    int64                                   `json:"peakOffpeak"`
	RouteShortName string                                  `json:"routeShortName"`
	ServiceID      string                                  `json:"serviceId"`
	ShapeID        string                                  `json:"shapeId"`
	TimeZone       string                                  `json:"timeZone"`
	TripHeadsign   string                                  `json:"tripHeadsign"`
	TripShortName  string                                  `json:"tripShortName"`
	JSON           configGetResponseDataReferencesTripJSON `json:"-"`
}

// configGetResponseDataReferencesTripJSON contains the JSON metadata for the
// struct [ConfigGetResponseDataReferencesTrip]
type configGetResponseDataReferencesTripJSON struct {
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

func (r *ConfigGetResponseDataReferencesTrip) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configGetResponseDataReferencesTripJSON) RawJSON() string {
	return r.raw
}
