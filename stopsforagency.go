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
	LimitExceeded bool                             `json:"limitExceeded,required"`
	List          []StopsForAgencyListResponseList `json:"list,required"`
	References    shared.References                `json:"references,required"`
	OutOfRange    bool                             `json:"outOfRange"`
	JSON          stopsForAgencyListResponseJSON   `json:"-"`
	shared.ResponseWrapper
}

// stopsForAgencyListResponseJSON contains the JSON metadata for the struct
// [StopsForAgencyListResponse]
type stopsForAgencyListResponseJSON struct {
	LimitExceeded apijson.Field
	List          apijson.Field
	References    apijson.Field
	OutOfRange    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *StopsForAgencyListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stopsForAgencyListResponseJSON) RawJSON() string {
	return r.raw
}

type StopsForAgencyListResponseList struct {
	ID                 string                             `json:"id,required"`
	Lat                float64                            `json:"lat,required"`
	LocationType       int64                              `json:"locationType,required"`
	Lon                float64                            `json:"lon,required"`
	Name               string                             `json:"name,required"`
	Parent             string                             `json:"parent,required"`
	RouteIDs           []string                           `json:"routeIds,required"`
	StaticRouteIDs     []string                           `json:"staticRouteIds,required"`
	Code               string                             `json:"code"`
	Direction          string                             `json:"direction"`
	WheelchairBoarding string                             `json:"wheelchairBoarding"`
	JSON               stopsForAgencyListResponseListJSON `json:"-"`
}

// stopsForAgencyListResponseListJSON contains the JSON metadata for the struct
// [StopsForAgencyListResponseList]
type stopsForAgencyListResponseListJSON struct {
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

func (r *StopsForAgencyListResponseList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stopsForAgencyListResponseListJSON) RawJSON() string {
	return r.raw
}
