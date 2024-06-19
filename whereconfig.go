// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package opentransit

import (
	"context"
	"net/http"

	"github.com/stainless-sdks/open-transit-go/internal/apijson"
	"github.com/stainless-sdks/open-transit-go/internal/requestconfig"
	"github.com/stainless-sdks/open-transit-go/option"
)

// WhereConfigService contains methods and other services that help with
// interacting with the open-transit API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWhereConfigService] method instead.
type WhereConfigService struct {
	Options []option.RequestOption
}

// NewWhereConfigService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWhereConfigService(opts ...option.RequestOption) (r *WhereConfigService) {
	r = &WhereConfigService{}
	r.Options = opts
	return
}

// config
func (r *WhereConfigService) Get(ctx context.Context, opts ...option.RequestOption) (res *WhereConfigGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/where/config.json"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type WhereConfigGetResponse struct {
	Code        int64                      `json:"code,required"`
	CurrentTime int64                      `json:"currentTime,required"`
	Text        string                     `json:"text,required"`
	Version     int64                      `json:"version,required"`
	Data        WhereConfigGetResponseData `json:"data"`
	JSON        whereConfigGetResponseJSON `json:"-"`
}

// whereConfigGetResponseJSON contains the JSON metadata for the struct
// [WhereConfigGetResponse]
type whereConfigGetResponseJSON struct {
	Code        apijson.Field
	CurrentTime apijson.Field
	Text        apijson.Field
	Version     apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WhereConfigGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r whereConfigGetResponseJSON) RawJSON() string {
	return r.raw
}

type WhereConfigGetResponseData struct {
	Entry      WhereConfigGetResponseDataEntry      `json:"entry"`
	References WhereConfigGetResponseDataReferences `json:"references"`
	JSON       whereConfigGetResponseDataJSON       `json:"-"`
}

// whereConfigGetResponseDataJSON contains the JSON metadata for the struct
// [WhereConfigGetResponseData]
type whereConfigGetResponseDataJSON struct {
	Entry       apijson.Field
	References  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WhereConfigGetResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r whereConfigGetResponseDataJSON) RawJSON() string {
	return r.raw
}

type WhereConfigGetResponseDataEntry struct {
	ID              string                                       `json:"id"`
	GitProperties   WhereConfigGetResponseDataEntryGitProperties `json:"gitProperties"`
	Name            string                                       `json:"name"`
	ServiceDateFrom string                                       `json:"serviceDateFrom"`
	ServiceDateTo   string                                       `json:"serviceDateTo"`
	JSON            whereConfigGetResponseDataEntryJSON          `json:"-"`
}

// whereConfigGetResponseDataEntryJSON contains the JSON metadata for the struct
// [WhereConfigGetResponseDataEntry]
type whereConfigGetResponseDataEntryJSON struct {
	ID              apijson.Field
	GitProperties   apijson.Field
	Name            apijson.Field
	ServiceDateFrom apijson.Field
	ServiceDateTo   apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *WhereConfigGetResponseDataEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r whereConfigGetResponseDataEntryJSON) RawJSON() string {
	return r.raw
}

type WhereConfigGetResponseDataEntryGitProperties struct {
	GitBranch                string                                           `json:"git.branch"`
	GitBuildHost             string                                           `json:"git.build.host"`
	GitBuildTime             string                                           `json:"git.build.time"`
	GitBuildUserEmail        string                                           `json:"git.build.user.email"`
	GitBuildUserName         string                                           `json:"git.build.user.name"`
	GitBuildVersion          string                                           `json:"git.build.version"`
	GitClosestTagCommitCount string                                           `json:"git.closest.tag.commit.count"`
	GitClosestTagName        string                                           `json:"git.closest.tag.name"`
	GitCommitID              string                                           `json:"git.commit.id"`
	GitCommitIDAbbrev        string                                           `json:"git.commit.id.abbrev"`
	GitCommitIDDescribe      string                                           `json:"git.commit.id.describe"`
	GitCommitIDDescribeShort string                                           `json:"git.commit.id.describe-short"`
	GitCommitMessageFull     string                                           `json:"git.commit.message.full"`
	GitCommitMessageShort    string                                           `json:"git.commit.message.short"`
	GitCommitTime            string                                           `json:"git.commit.time"`
	GitCommitUserEmail       string                                           `json:"git.commit.user.email"`
	GitCommitUserName        string                                           `json:"git.commit.user.name"`
	GitDirty                 string                                           `json:"git.dirty"`
	GitRemoteOriginURL       string                                           `json:"git.remote.origin.url"`
	GitTags                  string                                           `json:"git.tags"`
	JSON                     whereConfigGetResponseDataEntryGitPropertiesJSON `json:"-"`
}

// whereConfigGetResponseDataEntryGitPropertiesJSON contains the JSON metadata for
// the struct [WhereConfigGetResponseDataEntryGitProperties]
type whereConfigGetResponseDataEntryGitPropertiesJSON struct {
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

func (r *WhereConfigGetResponseDataEntryGitProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r whereConfigGetResponseDataEntryGitPropertiesJSON) RawJSON() string {
	return r.raw
}

type WhereConfigGetResponseDataReferences struct {
	Agencies   []WhereConfigGetResponseDataReferencesAgency `json:"agencies"`
	Routes     []WhereConfigGetResponseDataReferencesRoute  `json:"routes"`
	Situations []interface{}                                `json:"situations"`
	Stops      []WhereConfigGetResponseDataReferencesStop   `json:"stops"`
	StopTimes  []interface{}                                `json:"stopTimes"`
	Trips      []WhereConfigGetResponseDataReferencesTrip   `json:"trips"`
	JSON       whereConfigGetResponseDataReferencesJSON     `json:"-"`
}

// whereConfigGetResponseDataReferencesJSON contains the JSON metadata for the
// struct [WhereConfigGetResponseDataReferences]
type whereConfigGetResponseDataReferencesJSON struct {
	Agencies    apijson.Field
	Routes      apijson.Field
	Situations  apijson.Field
	Stops       apijson.Field
	StopTimes   apijson.Field
	Trips       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WhereConfigGetResponseDataReferences) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r whereConfigGetResponseDataReferencesJSON) RawJSON() string {
	return r.raw
}

type WhereConfigGetResponseDataReferencesAgency struct {
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
	JSON           whereConfigGetResponseDataReferencesAgencyJSON `json:"-"`
}

// whereConfigGetResponseDataReferencesAgencyJSON contains the JSON metadata for
// the struct [WhereConfigGetResponseDataReferencesAgency]
type whereConfigGetResponseDataReferencesAgencyJSON struct {
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

func (r *WhereConfigGetResponseDataReferencesAgency) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r whereConfigGetResponseDataReferencesAgencyJSON) RawJSON() string {
	return r.raw
}

type WhereConfigGetResponseDataReferencesRoute struct {
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
	JSON              whereConfigGetResponseDataReferencesRouteJSON `json:"-"`
}

// whereConfigGetResponseDataReferencesRouteJSON contains the JSON metadata for the
// struct [WhereConfigGetResponseDataReferencesRoute]
type whereConfigGetResponseDataReferencesRouteJSON struct {
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

func (r *WhereConfigGetResponseDataReferencesRoute) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r whereConfigGetResponseDataReferencesRouteJSON) RawJSON() string {
	return r.raw
}

type WhereConfigGetResponseDataReferencesStop struct {
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
	JSON               whereConfigGetResponseDataReferencesStopJSON `json:"-"`
}

// whereConfigGetResponseDataReferencesStopJSON contains the JSON metadata for the
// struct [WhereConfigGetResponseDataReferencesStop]
type whereConfigGetResponseDataReferencesStopJSON struct {
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

func (r *WhereConfigGetResponseDataReferencesStop) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r whereConfigGetResponseDataReferencesStopJSON) RawJSON() string {
	return r.raw
}

type WhereConfigGetResponseDataReferencesTrip struct {
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
	JSON           whereConfigGetResponseDataReferencesTripJSON `json:"-"`
}

// whereConfigGetResponseDataReferencesTripJSON contains the JSON metadata for the
// struct [WhereConfigGetResponseDataReferencesTrip]
type whereConfigGetResponseDataReferencesTripJSON struct {
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

func (r *WhereConfigGetResponseDataReferencesTrip) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r whereConfigGetResponseDataReferencesTripJSON) RawJSON() string {
	return r.raw
}
