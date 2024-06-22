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

// Retrieve Agencies with Coverage
func (r *AgenciesWithCoverageService) Get(ctx context.Context, opts ...option.RequestOption) (res *AgenciesWithCoverageGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/where/agencies-with-coverage.json"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AgenciesWithCoverageGetResponse struct {
	Data AgenciesWithCoverageGetResponseData `json:"data"`
	JSON agenciesWithCoverageGetResponseJSON `json:"-"`
	shared.ResponseWrapper
}

// agenciesWithCoverageGetResponseJSON contains the JSON metadata for the struct
// [AgenciesWithCoverageGetResponse]
type agenciesWithCoverageGetResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AgenciesWithCoverageGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r agenciesWithCoverageGetResponseJSON) RawJSON() string {
	return r.raw
}

type AgenciesWithCoverageGetResponseData struct {
	LimitExceeded bool                                      `json:"limitExceeded"`
	List          []AgenciesWithCoverageGetResponseDataList `json:"list"`
	References    shared.References                         `json:"references"`
	JSON          agenciesWithCoverageGetResponseDataJSON   `json:"-"`
}

// agenciesWithCoverageGetResponseDataJSON contains the JSON metadata for the
// struct [AgenciesWithCoverageGetResponseData]
type agenciesWithCoverageGetResponseDataJSON struct {
	LimitExceeded apijson.Field
	List          apijson.Field
	References    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AgenciesWithCoverageGetResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r agenciesWithCoverageGetResponseDataJSON) RawJSON() string {
	return r.raw
}

type AgenciesWithCoverageGetResponseDataList struct {
	AgencyID string                                      `json:"agencyId,required"`
	Lat      float64                                     `json:"lat,required"`
	LatSpan  float64                                     `json:"latSpan,required"`
	Lon      float64                                     `json:"lon,required"`
	LonSpan  float64                                     `json:"lonSpan,required"`
	JSON     agenciesWithCoverageGetResponseDataListJSON `json:"-"`
}

// agenciesWithCoverageGetResponseDataListJSON contains the JSON metadata for the
// struct [AgenciesWithCoverageGetResponseDataList]
type agenciesWithCoverageGetResponseDataListJSON struct {
	AgencyID    apijson.Field
	Lat         apijson.Field
	LatSpan     apijson.Field
	Lon         apijson.Field
	LonSpan     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AgenciesWithCoverageGetResponseDataList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r agenciesWithCoverageGetResponseDataListJSON) RawJSON() string {
	return r.raw
}
