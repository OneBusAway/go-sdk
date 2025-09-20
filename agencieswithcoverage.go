// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package onebusaway

import (
	"context"
	"net/http"
	"slices"

	"github.com/OneBusAway/go-sdk/internal/apijson"
	"github.com/OneBusAway/go-sdk/internal/requestconfig"
	"github.com/OneBusAway/go-sdk/option"
	"github.com/OneBusAway/go-sdk/shared"
)

// AgenciesWithCoverageService contains methods and other services that help with
// interacting with the onebusaway-sdk API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAgenciesWithCoverageService] method instead.
type AgenciesWithCoverageService struct {
	Options []option.RequestOption
}

// NewAgenciesWithCoverageService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAgenciesWithCoverageService(opts ...option.RequestOption) (r *AgenciesWithCoverageService) {
	r = &AgenciesWithCoverageService{}
	r.Options = opts
	return
}

// Returns a list of all transit agencies currently supported by OneBusAway along
// with the center of their coverage area.
func (r *AgenciesWithCoverageService) List(ctx context.Context, opts ...option.RequestOption) (res *AgenciesWithCoverageListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/where/agencies-with-coverage.json"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AgenciesWithCoverageListResponse struct {
	Data AgenciesWithCoverageListResponseData `json:"data,required"`
	JSON agenciesWithCoverageListResponseJSON `json:"-"`
	shared.ResponseWrapper
}

// agenciesWithCoverageListResponseJSON contains the JSON metadata for the struct
// [AgenciesWithCoverageListResponse]
type agenciesWithCoverageListResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AgenciesWithCoverageListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r agenciesWithCoverageListResponseJSON) RawJSON() string {
	return r.raw
}

type AgenciesWithCoverageListResponseData struct {
	LimitExceeded bool                                       `json:"limitExceeded,required"`
	List          []AgenciesWithCoverageListResponseDataList `json:"list,required"`
	References    shared.References                          `json:"references,required"`
	JSON          agenciesWithCoverageListResponseDataJSON   `json:"-"`
}

// agenciesWithCoverageListResponseDataJSON contains the JSON metadata for the
// struct [AgenciesWithCoverageListResponseData]
type agenciesWithCoverageListResponseDataJSON struct {
	LimitExceeded apijson.Field
	List          apijson.Field
	References    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AgenciesWithCoverageListResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r agenciesWithCoverageListResponseDataJSON) RawJSON() string {
	return r.raw
}

type AgenciesWithCoverageListResponseDataList struct {
	AgencyID string                                       `json:"agencyId,required"`
	Lat      float64                                      `json:"lat,required"`
	LatSpan  float64                                      `json:"latSpan,required"`
	Lon      float64                                      `json:"lon,required"`
	LonSpan  float64                                      `json:"lonSpan,required"`
	JSON     agenciesWithCoverageListResponseDataListJSON `json:"-"`
}

// agenciesWithCoverageListResponseDataListJSON contains the JSON metadata for the
// struct [AgenciesWithCoverageListResponseDataList]
type agenciesWithCoverageListResponseDataListJSON struct {
	AgencyID    apijson.Field
	Lat         apijson.Field
	LatSpan     apijson.Field
	Lon         apijson.Field
	LonSpan     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AgenciesWithCoverageListResponseDataList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r agenciesWithCoverageListResponseDataListJSON) RawJSON() string {
	return r.raw
}
