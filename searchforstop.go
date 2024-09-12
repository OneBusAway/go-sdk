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

// SearchForStopService contains methods and other services that help with
// interacting with the onebusaway-sdk API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSearchForStopService] method instead.
type SearchForStopService struct {
	Options []option.RequestOption
}

// NewSearchForStopService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSearchForStopService(opts ...option.RequestOption) (r *SearchForStopService) {
	r = &SearchForStopService{}
	r.Options = opts
	return
}

// Search for a stop based on its name.
func (r *SearchForStopService) List(ctx context.Context, query SearchForStopListParams, opts ...option.RequestOption) (res *SearchForStopListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/where/search/stop.json"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type SearchForStopListResponse struct {
	Data SearchForStopListResponseData `json:"data"`
	JSON searchForStopListResponseJSON `json:"-"`
	shared.ResponseWrapper
}

// searchForStopListResponseJSON contains the JSON metadata for the struct
// [SearchForStopListResponse]
type searchForStopListResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SearchForStopListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r searchForStopListResponseJSON) RawJSON() string {
	return r.raw
}

type SearchForStopListResponseData struct {
	LimitExceeded bool                                `json:"limitExceeded,required"`
	List          []SearchForStopListResponseDataList `json:"list,required"`
	OutOfRange    bool                                `json:"outOfRange,required"`
	References    shared.References                   `json:"references,required"`
	JSON          searchForStopListResponseDataJSON   `json:"-"`
}

// searchForStopListResponseDataJSON contains the JSON metadata for the struct
// [SearchForStopListResponseData]
type searchForStopListResponseDataJSON struct {
	LimitExceeded apijson.Field
	List          apijson.Field
	OutOfRange    apijson.Field
	References    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *SearchForStopListResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r searchForStopListResponseDataJSON) RawJSON() string {
	return r.raw
}

type SearchForStopListResponseDataList struct {
	ID                 string                                `json:"id,required"`
	Code               string                                `json:"code,required"`
	Lat                float64                               `json:"lat,required"`
	Lon                float64                               `json:"lon,required"`
	Name               string                                `json:"name,required"`
	Direction          string                                `json:"direction"`
	LocationType       int64                                 `json:"locationType"`
	Parent             string                                `json:"parent"`
	RouteIDs           []string                              `json:"routeIds"`
	StaticRouteIDs     []string                              `json:"staticRouteIds"`
	WheelchairBoarding string                                `json:"wheelchairBoarding"`
	JSON               searchForStopListResponseDataListJSON `json:"-"`
}

// searchForStopListResponseDataListJSON contains the JSON metadata for the struct
// [SearchForStopListResponseDataList]
type searchForStopListResponseDataListJSON struct {
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

func (r *SearchForStopListResponseDataList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r searchForStopListResponseDataListJSON) RawJSON() string {
	return r.raw
}

type SearchForStopListParams struct {
	// The string to search for.
	Input param.Field[string] `query:"input,required"`
	// The max number of results to return. Defaults to 20.
	MaxCount param.Field[int64] `query:"maxCount"`
}

// URLQuery serializes [SearchForStopListParams]'s query parameters as
// `url.Values`.
func (r SearchForStopListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
