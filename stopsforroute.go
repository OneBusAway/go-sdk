// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package onebusaway

import (
	"context"
	"errors"
	"fmt"
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

// StopsForRouteService contains methods and other services that help with
// interacting with the onebusaway-sdk API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewStopsForRouteService] method instead.
type StopsForRouteService struct {
	Options []option.RequestOption
}

// NewStopsForRouteService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewStopsForRouteService(opts ...option.RequestOption) (r *StopsForRouteService) {
	r = &StopsForRouteService{}
	r.Options = opts
	return
}

// Get stops for a specific route
func (r *StopsForRouteService) List(ctx context.Context, routeID string, query StopsForRouteListParams, opts ...option.RequestOption) (res *StopsForRouteListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if routeID == "" {
		err = errors.New("missing required routeID parameter")
		return
	}
	path := fmt.Sprintf("api/where/stops-for-route/%s.json", routeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type StopsForRouteListResponse struct {
	Data StopsForRouteListResponseData `json:"data" api:"required"`
	JSON stopsForRouteListResponseJSON `json:"-"`
	shared.ResponseWrapper
}

// stopsForRouteListResponseJSON contains the JSON metadata for the struct
// [StopsForRouteListResponse]
type stopsForRouteListResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StopsForRouteListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stopsForRouteListResponseJSON) RawJSON() string {
	return r.raw
}

type StopsForRouteListResponseData struct {
	Data StopsForRouteListResponseDataData `json:"data" api:"required"`
	JSON stopsForRouteListResponseDataJSON `json:"-"`
}

// stopsForRouteListResponseDataJSON contains the JSON metadata for the struct
// [StopsForRouteListResponseData]
type stopsForRouteListResponseDataJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StopsForRouteListResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stopsForRouteListResponseDataJSON) RawJSON() string {
	return r.raw
}

type StopsForRouteListResponseDataData struct {
	Entry      StopsForRouteListResponseDataDataEntry `json:"entry" api:"required"`
	References shared.References                      `json:"references" api:"required"`
	JSON       stopsForRouteListResponseDataDataJSON  `json:"-"`
}

// stopsForRouteListResponseDataDataJSON contains the JSON metadata for the struct
// [StopsForRouteListResponseDataData]
type stopsForRouteListResponseDataDataJSON struct {
	Entry       apijson.Field
	References  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StopsForRouteListResponseDataData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stopsForRouteListResponseDataDataJSON) RawJSON() string {
	return r.raw
}

type StopsForRouteListResponseDataDataEntry struct {
	Polylines     []StopsForRouteListResponseDataDataEntryPolyline     `json:"polylines"`
	RouteID       string                                               `json:"routeId"`
	StopGroupings []StopsForRouteListResponseDataDataEntryStopGrouping `json:"stopGroupings"`
	StopIDs       []string                                             `json:"stopIds"`
	JSON          stopsForRouteListResponseDataDataEntryJSON           `json:"-"`
}

// stopsForRouteListResponseDataDataEntryJSON contains the JSON metadata for the
// struct [StopsForRouteListResponseDataDataEntry]
type stopsForRouteListResponseDataDataEntryJSON struct {
	Polylines     apijson.Field
	RouteID       apijson.Field
	StopGroupings apijson.Field
	StopIDs       apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *StopsForRouteListResponseDataDataEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stopsForRouteListResponseDataDataEntryJSON) RawJSON() string {
	return r.raw
}

type StopsForRouteListResponseDataDataEntryPolyline struct {
	Length int64                                              `json:"length"`
	Levels string                                             `json:"levels"`
	Points string                                             `json:"points"`
	JSON   stopsForRouteListResponseDataDataEntryPolylineJSON `json:"-"`
}

// stopsForRouteListResponseDataDataEntryPolylineJSON contains the JSON metadata
// for the struct [StopsForRouteListResponseDataDataEntryPolyline]
type stopsForRouteListResponseDataDataEntryPolylineJSON struct {
	Length      apijson.Field
	Levels      apijson.Field
	Points      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StopsForRouteListResponseDataDataEntryPolyline) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stopsForRouteListResponseDataDataEntryPolylineJSON) RawJSON() string {
	return r.raw
}

type StopsForRouteListResponseDataDataEntryStopGrouping struct {
	ID        string                                                        `json:"id"`
	Name      StopsForRouteListResponseDataDataEntryStopGroupingsName       `json:"name"`
	Polylines []StopsForRouteListResponseDataDataEntryStopGroupingsPolyline `json:"polylines"`
	StopIDs   []string                                                      `json:"stopIds"`
	JSON      stopsForRouteListResponseDataDataEntryStopGroupingJSON        `json:"-"`
}

// stopsForRouteListResponseDataDataEntryStopGroupingJSON contains the JSON
// metadata for the struct [StopsForRouteListResponseDataDataEntryStopGrouping]
type stopsForRouteListResponseDataDataEntryStopGroupingJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	Polylines   apijson.Field
	StopIDs     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StopsForRouteListResponseDataDataEntryStopGrouping) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stopsForRouteListResponseDataDataEntryStopGroupingJSON) RawJSON() string {
	return r.raw
}

type StopsForRouteListResponseDataDataEntryStopGroupingsName struct {
	Name  string                                                      `json:"name"`
	Names []string                                                    `json:"names"`
	Type  string                                                      `json:"type"`
	JSON  stopsForRouteListResponseDataDataEntryStopGroupingsNameJSON `json:"-"`
}

// stopsForRouteListResponseDataDataEntryStopGroupingsNameJSON contains the JSON
// metadata for the struct
// [StopsForRouteListResponseDataDataEntryStopGroupingsName]
type stopsForRouteListResponseDataDataEntryStopGroupingsNameJSON struct {
	Name        apijson.Field
	Names       apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StopsForRouteListResponseDataDataEntryStopGroupingsName) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stopsForRouteListResponseDataDataEntryStopGroupingsNameJSON) RawJSON() string {
	return r.raw
}

type StopsForRouteListResponseDataDataEntryStopGroupingsPolyline struct {
	Length int64                                                           `json:"length"`
	Levels string                                                          `json:"levels"`
	Points string                                                          `json:"points"`
	JSON   stopsForRouteListResponseDataDataEntryStopGroupingsPolylineJSON `json:"-"`
}

// stopsForRouteListResponseDataDataEntryStopGroupingsPolylineJSON contains the
// JSON metadata for the struct
// [StopsForRouteListResponseDataDataEntryStopGroupingsPolyline]
type stopsForRouteListResponseDataDataEntryStopGroupingsPolylineJSON struct {
	Length      apijson.Field
	Levels      apijson.Field
	Points      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StopsForRouteListResponseDataDataEntryStopGroupingsPolyline) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stopsForRouteListResponseDataDataEntryStopGroupingsPolylineJSON) RawJSON() string {
	return r.raw
}

type StopsForRouteListParams struct {
	// Include polyline elements in the response (default true)
	IncludePolylines param.Field[bool] `query:"includePolylines"`
	// Specify service date (YYYY-MM-DD or epoch) (default today)
	Time param.Field[string] `query:"time"`
}

// URLQuery serializes [StopsForRouteListParams]'s query parameters as
// `url.Values`.
func (r StopsForRouteListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
