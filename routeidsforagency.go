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

// RouteIDsForAgencyService contains methods and other services that help with
// interacting with the onebusaway-sdk API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRouteIDsForAgencyService] method instead.
type RouteIDsForAgencyService struct {
	Options []option.RequestOption
}

// NewRouteIDsForAgencyService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRouteIDsForAgencyService(opts ...option.RequestOption) (r *RouteIDsForAgencyService) {
	r = &RouteIDsForAgencyService{}
	r.Options = opts
	return
}

// Get route IDs for a specific agency
func (r *RouteIDsForAgencyService) List(ctx context.Context, agencyID string, opts ...option.RequestOption) (res *RouteIDsForAgencyListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if agencyID == "" {
		err = errors.New("missing required agencyID parameter")
		return
	}
	path := fmt.Sprintf("api/where/route-ids-for-agency/%s.json", agencyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type RouteIDsForAgencyListResponse struct {
	Data RouteIDsForAgencyListResponseData `json:"data,required"`
	JSON routeIDsForAgencyListResponseJSON `json:"-"`
	shared.ResponseWrapper
}

// routeIDsForAgencyListResponseJSON contains the JSON metadata for the struct
// [RouteIDsForAgencyListResponse]
type routeIDsForAgencyListResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RouteIDsForAgencyListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routeIDsForAgencyListResponseJSON) RawJSON() string {
	return r.raw
}

type RouteIDsForAgencyListResponseData struct {
	LimitExceeded bool                                  `json:"limitExceeded,required"`
	List          []string                              `json:"list,required"`
	References    shared.References                     `json:"references,required"`
	JSON          routeIDsForAgencyListResponseDataJSON `json:"-"`
}

// routeIDsForAgencyListResponseDataJSON contains the JSON metadata for the struct
// [RouteIDsForAgencyListResponseData]
type routeIDsForAgencyListResponseDataJSON struct {
	LimitExceeded apijson.Field
	List          apijson.Field
	References    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *RouteIDsForAgencyListResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routeIDsForAgencyListResponseDataJSON) RawJSON() string {
	return r.raw
}
