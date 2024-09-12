// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package onebusaway

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/OneBusAway/go-sdk/internal/apijson"
	"github.com/OneBusAway/go-sdk/internal/requestconfig"
	"github.com/OneBusAway/go-sdk/option"
	"github.com/OneBusAway/go-sdk/shared"
)

// StopService contains methods and other services that help with interacting with
// the onebusaway-sdk API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewStopService] method instead.
type StopService struct {
	Options []option.RequestOption
}

// NewStopService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewStopService(opts ...option.RequestOption) (r *StopService) {
	r = &StopService{}
	r.Options = opts
	return
}

// Get details of a specific stop
func (r *StopService) Get(ctx context.Context, stopID string, opts ...option.RequestOption) (res *StopGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if stopID == "" {
		err = errors.New("missing required stopID parameter")
		return
	}
	path := fmt.Sprintf("api/where/stop/%s.json", stopID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type StopGetResponse struct {
	Data StopGetResponseData `json:"data,required"`
	JSON stopGetResponseJSON `json:"-"`
	shared.ResponseWrapper
}

// stopGetResponseJSON contains the JSON metadata for the struct [StopGetResponse]
type stopGetResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StopGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stopGetResponseJSON) RawJSON() string {
	return r.raw
}

type StopGetResponseData struct {
	Entry      StopGetResponseDataEntry `json:"entry,required"`
	References shared.References        `json:"references,required"`
	JSON       stopGetResponseDataJSON  `json:"-"`
}

// stopGetResponseDataJSON contains the JSON metadata for the struct
// [StopGetResponseData]
type stopGetResponseDataJSON struct {
	Entry       apijson.Field
	References  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StopGetResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stopGetResponseDataJSON) RawJSON() string {
	return r.raw
}

type StopGetResponseDataEntry struct {
	ID                 string                       `json:"id,required"`
	Code               string                       `json:"code,required"`
	Lat                float64                      `json:"lat,required"`
	Lon                float64                      `json:"lon,required"`
	Name               string                       `json:"name,required"`
	Direction          string                       `json:"direction"`
	LocationType       int64                        `json:"locationType"`
	Parent             string                       `json:"parent"`
	RouteIDs           []string                     `json:"routeIds"`
	StaticRouteIDs     []string                     `json:"staticRouteIds"`
	WheelchairBoarding string                       `json:"wheelchairBoarding"`
	JSON               stopGetResponseDataEntryJSON `json:"-"`
}

// stopGetResponseDataEntryJSON contains the JSON metadata for the struct
// [StopGetResponseDataEntry]
type stopGetResponseDataEntryJSON struct {
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

func (r *StopGetResponseDataEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stopGetResponseDataEntryJSON) RawJSON() string {
	return r.raw
}
