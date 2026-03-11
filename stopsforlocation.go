// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package onebusaway

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/OneBusAway/go-sdk/internal/apijson"
	"github.com/OneBusAway/go-sdk/internal/apiquery"
	"github.com/OneBusAway/go-sdk/internal/param"
	"github.com/OneBusAway/go-sdk/internal/requestconfig"
	"github.com/OneBusAway/go-sdk/option"
	"github.com/OneBusAway/go-sdk/shared"
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
	opts = slices.Concat(r.Options, opts)
	path := "api/where/stops-for-location.json"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type StopsForLocationListResponse struct {
	Data StopsForLocationListResponseData `json:"data" api:"required"`
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
	List          []StopsForLocationListResponseDataList `json:"list" api:"required"`
	References    shared.References                      `json:"references" api:"required"`
	LimitExceeded bool                                   `json:"limitExceeded"`
	OutOfRange    bool                                   `json:"outOfRange"`
	JSON          stopsForLocationListResponseDataJSON   `json:"-"`
}

// stopsForLocationListResponseDataJSON contains the JSON metadata for the struct
// [StopsForLocationListResponseData]
type stopsForLocationListResponseDataJSON struct {
	List          apijson.Field
	References    apijson.Field
	LimitExceeded apijson.Field
	OutOfRange    apijson.Field
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
	ID                 string                                   `json:"id" api:"required"`
	Lat                float64                                  `json:"lat" api:"required"`
	LocationType       int64                                    `json:"locationType" api:"required"`
	Lon                float64                                  `json:"lon" api:"required"`
	Name               string                                   `json:"name" api:"required"`
	Parent             string                                   `json:"parent" api:"required"`
	RouteIDs           []string                                 `json:"routeIds" api:"required"`
	StaticRouteIDs     []string                                 `json:"staticRouteIds" api:"required"`
	Code               string                                   `json:"code"`
	Direction          string                                   `json:"direction"`
	WheelchairBoarding string                                   `json:"wheelchairBoarding"`
	JSON               stopsForLocationListResponseDataListJSON `json:"-"`
}

// stopsForLocationListResponseDataListJSON contains the JSON metadata for the
// struct [StopsForLocationListResponseDataList]
type stopsForLocationListResponseDataListJSON struct {
	ID                 apijson.Field
	Lat                apijson.Field
	LocationType       apijson.Field
	Lon                apijson.Field
	Name               apijson.Field
	Parent             apijson.Field
	RouteIDs           apijson.Field
	StaticRouteIDs     apijson.Field
	Code               apijson.Field
	Direction          apijson.Field
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
	Lat param.Field[float64] `query:"lat" api:"required"`
	Lon param.Field[float64] `query:"lon" api:"required"`
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
