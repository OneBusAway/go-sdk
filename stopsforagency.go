// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package onebusaway

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/OneBusAway/go-sdk/internal/apijson"
	"github.com/OneBusAway/go-sdk/internal/requestconfig"
	"github.com/OneBusAway/go-sdk/option"
	"github.com/OneBusAway/go-sdk/shared"
)

// StopsForAgencyService contains methods and other services that help with
// interacting with the onebusaway-sdk API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewStopsForAgencyService] method instead.
type StopsForAgencyService struct {
	Options []option.RequestOption
}

// NewStopsForAgencyService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewStopsForAgencyService(opts ...option.RequestOption) (r *StopsForAgencyService) {
	r = &StopsForAgencyService{}
	r.Options = opts
	return
}

// Get stops for a specific agency
func (r *StopsForAgencyService) List(ctx context.Context, agencyID string, opts ...option.RequestOption) (res *StopsForAgencyListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if agencyID == "" {
		err = errors.New("missing required agencyID parameter")
		return
	}
	path := fmt.Sprintf("api/where/stops-for-agency/%s.json", agencyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type StopsForAgencyListResponse struct {
	Data StopsForAgencyListResponseData `json:"data" api:"required"`
	JSON stopsForAgencyListResponseJSON `json:"-"`
	shared.ResponseWrapper
}

// stopsForAgencyListResponseJSON contains the JSON metadata for the struct
// [StopsForAgencyListResponse]
type stopsForAgencyListResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StopsForAgencyListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stopsForAgencyListResponseJSON) RawJSON() string {
	return r.raw
}

type StopsForAgencyListResponseData struct {
	List          []StopsForAgencyListResponseDataList `json:"list" api:"required"`
	References    shared.References                    `json:"references" api:"required"`
	LimitExceeded bool                                 `json:"limitExceeded"`
	OutOfRange    bool                                 `json:"outOfRange"`
	JSON          stopsForAgencyListResponseDataJSON   `json:"-"`
}

// stopsForAgencyListResponseDataJSON contains the JSON metadata for the struct
// [StopsForAgencyListResponseData]
type stopsForAgencyListResponseDataJSON struct {
	List          apijson.Field
	References    apijson.Field
	LimitExceeded apijson.Field
	OutOfRange    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *StopsForAgencyListResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stopsForAgencyListResponseDataJSON) RawJSON() string {
	return r.raw
}

type StopsForAgencyListResponseDataList struct {
	ID                 string                                 `json:"id" api:"required"`
	Lat                float64                                `json:"lat" api:"required"`
	LocationType       int64                                  `json:"locationType" api:"required"`
	Lon                float64                                `json:"lon" api:"required"`
	Name               string                                 `json:"name" api:"required"`
	Parent             string                                 `json:"parent" api:"required"`
	RouteIDs           []string                               `json:"routeIds" api:"required"`
	StaticRouteIDs     []string                               `json:"staticRouteIds" api:"required"`
	Code               string                                 `json:"code"`
	Direction          string                                 `json:"direction"`
	WheelchairBoarding string                                 `json:"wheelchairBoarding"`
	JSON               stopsForAgencyListResponseDataListJSON `json:"-"`
}

// stopsForAgencyListResponseDataListJSON contains the JSON metadata for the struct
// [StopsForAgencyListResponseDataList]
type stopsForAgencyListResponseDataListJSON struct {
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

func (r *StopsForAgencyListResponseDataList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stopsForAgencyListResponseDataListJSON) RawJSON() string {
	return r.raw
}
