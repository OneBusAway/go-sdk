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

// ConfigService contains methods and other services that help with interacting
// with the onebusaway-sdk API.
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
	Entry      ConfigGetResponseDataEntry `json:"entry"`
	References shared.References          `json:"references"`
	JSON       configGetResponseDataJSON  `json:"-"`
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
