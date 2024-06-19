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
// interacting with the OneBusAway API.
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
func (r *StopsForLocationService) Get(ctx context.Context, query StopsForLocationGetParams, opts ...option.RequestOption) (res *StopsForLocationGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/where/stops-for-location.json"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type StopsForLocationGetResponse struct {
	Data StopsForLocationGetResponseData `json:"data"`
	JSON stopsForLocationGetResponseJSON `json:"-"`
	shared.ResponseWrapper
}

// stopsForLocationGetResponseJSON contains the JSON metadata for the struct
// [StopsForLocationGetResponse]
type stopsForLocationGetResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StopsForLocationGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stopsForLocationGetResponseJSON) RawJSON() string {
	return r.raw
}

type StopsForLocationGetResponseData struct {
	LimitExceeded bool                                  `json:"limitExceeded"`
	List          []StopsForLocationGetResponseDataList `json:"list"`
	References    shared.References                     `json:"references"`
	JSON          stopsForLocationGetResponseDataJSON   `json:"-"`
}

// stopsForLocationGetResponseDataJSON contains the JSON metadata for the struct
// [StopsForLocationGetResponseData]
type stopsForLocationGetResponseDataJSON struct {
	LimitExceeded apijson.Field
	List          apijson.Field
	References    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *StopsForLocationGetResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stopsForLocationGetResponseDataJSON) RawJSON() string {
	return r.raw
}

type StopsForLocationGetResponseDataList struct {
	ID                 string                                  `json:"id"`
	Code               string                                  `json:"code"`
	Direction          string                                  `json:"direction"`
	Lat                float64                                 `json:"lat"`
	LocationType       int64                                   `json:"locationType"`
	Lon                float64                                 `json:"lon"`
	Name               string                                  `json:"name"`
	Parent             string                                  `json:"parent"`
	RouteIDs           []string                                `json:"routeIds"`
	StaticRouteIDs     []string                                `json:"staticRouteIds"`
	WheelchairBoarding string                                  `json:"wheelchairBoarding"`
	JSON               stopsForLocationGetResponseDataListJSON `json:"-"`
}

// stopsForLocationGetResponseDataListJSON contains the JSON metadata for the
// struct [StopsForLocationGetResponseDataList]
type stopsForLocationGetResponseDataListJSON struct {
	ID                 apijson.Field
	Code               apijson.Field
	Direction          apijson.Field
	Lat                apijson.Field
	LocationType       apijson.Field
	Lon                apijson.Field
	Name               apijson.Field
	Parent             apijson.Field
	RouteIDs           apijson.Field
	StaticRouteIDs     apijson.Field
	WheelchairBoarding apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *StopsForLocationGetResponseDataList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stopsForLocationGetResponseDataListJSON) RawJSON() string {
	return r.raw
}

type StopsForLocationGetParams struct {
	Key param.Field[string]  `query:"key,required"`
	Lat param.Field[float64] `query:"lat"`
	Lon param.Field[float64] `query:"lon"`
}

// URLQuery serializes [StopsForLocationGetParams]'s query parameters as
// `url.Values`.
func (r StopsForLocationGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
