// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package onebusaway

import (
	"context"
	"net/http"
	"net/url"

	"github.com/OneBusAway/go-sdk/internal/apijson"
	"github.com/OneBusAway/go-sdk/internal/apiquery"
	"github.com/OneBusAway/go-sdk/internal/param"
	"github.com/OneBusAway/go-sdk/internal/requestconfig"
	"github.com/OneBusAway/go-sdk/option"
	"github.com/OneBusAway/go-sdk/shared"
)

// RoutesForLocationService contains methods and other services that help with
// interacting with the onebusaway-sdk API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRoutesForLocationService] method instead.
type RoutesForLocationService struct {
	Options []option.RequestOption
}

// NewRoutesForLocationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRoutesForLocationService(opts ...option.RequestOption) (r *RoutesForLocationService) {
	r = &RoutesForLocationService{}
	r.Options = opts
	return
}

// routes-for-location
func (r *RoutesForLocationService) List(ctx context.Context, query RoutesForLocationListParams, opts ...option.RequestOption) (res *RoutesForLocationListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/where/routes-for-location.json"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RoutesForLocationListResponse struct {
	Data RoutesForLocationListResponseData `json:"data,required"`
	JSON routesForLocationListResponseJSON `json:"-"`
	shared.ResponseWrapper
}

// routesForLocationListResponseJSON contains the JSON metadata for the struct
// [RoutesForLocationListResponse]
type routesForLocationListResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RoutesForLocationListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routesForLocationListResponseJSON) RawJSON() string {
	return r.raw
}

type RoutesForLocationListResponseData struct {
	LimitExceeded bool                                    `json:"limitExceeded,required"`
	List          []RoutesForLocationListResponseDataList `json:"list,required"`
	OutOfRange    bool                                    `json:"outOfRange,required"`
	References    shared.References                       `json:"references,required"`
	JSON          routesForLocationListResponseDataJSON   `json:"-"`
}

// routesForLocationListResponseDataJSON contains the JSON metadata for the struct
// [RoutesForLocationListResponseData]
type routesForLocationListResponseDataJSON struct {
	LimitExceeded apijson.Field
	List          apijson.Field
	OutOfRange    apijson.Field
	References    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *RoutesForLocationListResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routesForLocationListResponseDataJSON) RawJSON() string {
	return r.raw
}

type RoutesForLocationListResponseDataList struct {
	ID                string                                    `json:"id,required"`
	AgencyID          string                                    `json:"agencyId,required"`
	Type              int64                                     `json:"type,required"`
	Color             string                                    `json:"color"`
	Description       string                                    `json:"description"`
	LongName          string                                    `json:"longName"`
	NullSafeShortName string                                    `json:"nullSafeShortName"`
	ShortName         string                                    `json:"shortName"`
	TextColor         string                                    `json:"textColor"`
	URL               string                                    `json:"url"`
	JSON              routesForLocationListResponseDataListJSON `json:"-"`
}

// routesForLocationListResponseDataListJSON contains the JSON metadata for the
// struct [RoutesForLocationListResponseDataList]
type routesForLocationListResponseDataListJSON struct {
	ID                apijson.Field
	AgencyID          apijson.Field
	Type              apijson.Field
	Color             apijson.Field
	Description       apijson.Field
	LongName          apijson.Field
	NullSafeShortName apijson.Field
	ShortName         apijson.Field
	TextColor         apijson.Field
	URL               apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RoutesForLocationListResponseDataList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routesForLocationListResponseDataListJSON) RawJSON() string {
	return r.raw
}

type RoutesForLocationListParams struct {
	Lat     param.Field[float64] `query:"lat,required"`
	Lon     param.Field[float64] `query:"lon,required"`
	LatSpan param.Field[float64] `query:"latSpan"`
	LonSpan param.Field[float64] `query:"lonSpan"`
	Query   param.Field[string]  `query:"query"`
	Radius  param.Field[float64] `query:"radius"`
}

// URLQuery serializes [RoutesForLocationListParams]'s query parameters as
// `url.Values`.
func (r RoutesForLocationListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
