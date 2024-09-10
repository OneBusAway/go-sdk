// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package onebusaway

import (
	"context"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/open-transit-go/internal/apijson"
	"github.com/stainless-sdks/open-transit-go/internal/apiquery"
	"github.com/stainless-sdks/open-transit-go/internal/param"
	"github.com/stainless-sdks/open-transit-go/internal/requestconfig"
	"github.com/stainless-sdks/open-transit-go/option"
	"github.com/stainless-sdks/open-transit-go/shared"
)

// StopsForLocationService contains methods and other services that help with
// interacting with the onebusaway-sdk API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewStopsForLocationService] method instead.
type StopsForLocationService struct {
	Options []option.RequestOption
}

// NewStopsForLocationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewStopsForLocationService(opts ...option.RequestOption) (r *StopsForLocationService) {
	r = &StopsForLocationService{}
	r.Options = opts
	return
}

// stops-for-location
func (r *StopsForLocationService) List(ctx context.Context, query StopsForLocationListParams, opts ...option.RequestOption) (res *StopsForLocationListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/where/stops-for-location.json"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type StopsForLocationListResponse struct {
	Data StopsForLocationListResponseData `json:"data,required"`
	JSON stopsForLocationListResponseJSON `json:"-"`
	shared.ResponseWrapper
}

// stopsForLocationListResponseJSON contains the JSON metadata for the struct
// [StopsForLocationListResponse]
type stopsForLocationListResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StopsForLocationListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stopsForLocationListResponseJSON) RawJSON() string {
	return r.raw
}

type StopsForLocationListResponseData struct {
	LimitExceeded bool                                   `json:"limitExceeded,required"`
	List          []StopsForLocationListResponseDataList `json:"list,required"`
	References    shared.References                      `json:"references,required"`
	JSON          stopsForLocationListResponseDataJSON   `json:"-"`
}

// stopsForLocationListResponseDataJSON contains the JSON metadata for the struct
// [StopsForLocationListResponseData]
type stopsForLocationListResponseDataJSON struct {
	LimitExceeded apijson.Field
	List          apijson.Field
	References    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *StopsForLocationListResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stopsForLocationListResponseDataJSON) RawJSON() string {
	return r.raw
}

type StopsForLocationListResponseDataList struct {
	ID                 string                                   `json:"id,required"`
	Lat                float64                                  `json:"lat,required"`
	Lon                float64                                  `json:"lon,required"`
	Name               string                                   `json:"name,required"`
	Parent             string                                   `json:"parent,required"`
	RouteIDs           []string                                 `json:"routeIds,required"`
	StaticRouteIDs     []string                                 `json:"staticRouteIds,required"`
	Code               string                                   `json:"code"`
	Direction          string                                   `json:"direction"`
	LocationType       int64                                    `json:"locationType"`
	WheelchairBoarding string                                   `json:"wheelchairBoarding"`
	JSON               stopsForLocationListResponseDataListJSON `json:"-"`
}

// stopsForLocationListResponseDataListJSON contains the JSON metadata for the
// struct [StopsForLocationListResponseDataList]
type stopsForLocationListResponseDataListJSON struct {
	ID                 apijson.Field
	Lat                apijson.Field
	Lon                apijson.Field
	Name               apijson.Field
	Parent             apijson.Field
	RouteIDs           apijson.Field
	StaticRouteIDs     apijson.Field
	Code               apijson.Field
	Direction          apijson.Field
	LocationType       apijson.Field
	WheelchairBoarding apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *StopsForLocationListResponseDataList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stopsForLocationListResponseDataListJSON) RawJSON() string {
	return r.raw
}

type StopsForLocationListParams struct {
	Lat param.Field[float64] `query:"lat,required"`
	Lon param.Field[float64] `query:"lon,required"`
	// An alternative to radius to set the search bounding box (optional)
	LatSpan param.Field[float64] `query:"latSpan"`
	// An alternative to radius to set the search bounding box (optional)
	LonSpan param.Field[float64] `query:"lonSpan"`
	// A search query string to filter the results
	Query param.Field[string] `query:"query"`
	// The radius in meters to search within
	Radius param.Field[float64] `query:"radius"`
}

// URLQuery serializes [StopsForLocationListParams]'s query parameters as
// `url.Values`.
func (r StopsForLocationListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
