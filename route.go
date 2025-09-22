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

// RouteService contains methods and other services that help with interacting with
// the onebusaway-sdk API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRouteService] method instead.
type RouteService struct {
	Options []option.RequestOption
}

// NewRouteService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewRouteService(opts ...option.RequestOption) (r *RouteService) {
	r = &RouteService{}
	r.Options = opts
	return
}

// Retrieve information for a specific route identified by its unique ID.
func (r *RouteService) Get(ctx context.Context, routeID string, opts ...option.RequestOption) (res *RouteGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if routeID == "" {
		err = errors.New("missing required routeID parameter")
		return
	}
	path := fmt.Sprintf("api/where/route/%s.json", routeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type RouteGetResponse struct {
	Data RouteGetResponseData `json:"data,required"`
	JSON routeGetResponseJSON `json:"-"`
	shared.ResponseWrapper
}

// routeGetResponseJSON contains the JSON metadata for the struct
// [RouteGetResponse]
type routeGetResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RouteGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routeGetResponseJSON) RawJSON() string {
	return r.raw
}

type RouteGetResponseData struct {
	Entry      RouteGetResponseDataEntry `json:"entry,required"`
	References shared.References         `json:"references,required"`
	JSON       routeGetResponseDataJSON  `json:"-"`
}

// routeGetResponseDataJSON contains the JSON metadata for the struct
// [RouteGetResponseData]
type routeGetResponseDataJSON struct {
	Entry       apijson.Field
	References  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RouteGetResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routeGetResponseDataJSON) RawJSON() string {
	return r.raw
}

type RouteGetResponseDataEntry struct {
	ID                string                        `json:"id,required"`
	AgencyID          string                        `json:"agencyId,required"`
	Type              int64                         `json:"type,required"`
	Color             string                        `json:"color"`
	Description       string                        `json:"description"`
	LongName          string                        `json:"longName"`
	NullSafeShortName string                        `json:"nullSafeShortName"`
	ShortName         string                        `json:"shortName"`
	TextColor         string                        `json:"textColor"`
	URL               string                        `json:"url"`
	JSON              routeGetResponseDataEntryJSON `json:"-"`
}

// routeGetResponseDataEntryJSON contains the JSON metadata for the struct
// [RouteGetResponseDataEntry]
type routeGetResponseDataEntryJSON struct {
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

func (r *RouteGetResponseDataEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routeGetResponseDataEntryJSON) RawJSON() string {
	return r.raw
}
