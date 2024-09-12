// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package onebusaway

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

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
	opts = append(r.Options[:], opts...)
	if routeID == "" {
		err = errors.New("missing required routeID parameter")
		return
	}
	path := fmt.Sprintf("api/where/stops-for-route/%s.json", routeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type StopsForRouteListResponse struct {
	Data StopsForRouteListResponseData `json:"data,required"`
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
	Entry      StopsForRouteListResponseDataEntry `json:"entry,required"`
	References shared.References                  `json:"references,required"`
	JSON       stopsForRouteListResponseDataJSON  `json:"-"`
}

// stopsForRouteListResponseDataJSON contains the JSON metadata for the struct
// [StopsForRouteListResponseData]
type stopsForRouteListResponseDataJSON struct {
	Entry       apijson.Field
	References  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StopsForRouteListResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stopsForRouteListResponseDataJSON) RawJSON() string {
	return r.raw
}

type StopsForRouteListResponseDataEntry struct {
	Polylines     []StopsForRouteListResponseDataEntryPolyline     `json:"polylines"`
	RouteID       string                                           `json:"routeId"`
	StopGroupings []StopsForRouteListResponseDataEntryStopGrouping `json:"stopGroupings"`
	StopIDs       []string                                         `json:"stopIds"`
	JSON          stopsForRouteListResponseDataEntryJSON           `json:"-"`
}

// stopsForRouteListResponseDataEntryJSON contains the JSON metadata for the struct
// [StopsForRouteListResponseDataEntry]
type stopsForRouteListResponseDataEntryJSON struct {
	Polylines     apijson.Field
	RouteID       apijson.Field
	StopGroupings apijson.Field
	StopIDs       apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *StopsForRouteListResponseDataEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stopsForRouteListResponseDataEntryJSON) RawJSON() string {
	return r.raw
}

type StopsForRouteListResponseDataEntryPolyline struct {
	Length int64                                          `json:"length"`
	Levels string                                         `json:"levels"`
	Points string                                         `json:"points"`
	JSON   stopsForRouteListResponseDataEntryPolylineJSON `json:"-"`
}

// stopsForRouteListResponseDataEntryPolylineJSON contains the JSON metadata for
// the struct [StopsForRouteListResponseDataEntryPolyline]
type stopsForRouteListResponseDataEntryPolylineJSON struct {
	Length      apijson.Field
	Levels      apijson.Field
	Points      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StopsForRouteListResponseDataEntryPolyline) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stopsForRouteListResponseDataEntryPolylineJSON) RawJSON() string {
	return r.raw
}

type StopsForRouteListResponseDataEntryStopGrouping struct {
	ID        string                                                    `json:"id"`
	Name      StopsForRouteListResponseDataEntryStopGroupingsName       `json:"name"`
	Polylines []StopsForRouteListResponseDataEntryStopGroupingsPolyline `json:"polylines"`
	StopIDs   []string                                                  `json:"stopIds"`
	JSON      stopsForRouteListResponseDataEntryStopGroupingJSON        `json:"-"`
}

// stopsForRouteListResponseDataEntryStopGroupingJSON contains the JSON metadata
// for the struct [StopsForRouteListResponseDataEntryStopGrouping]
type stopsForRouteListResponseDataEntryStopGroupingJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	Polylines   apijson.Field
	StopIDs     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StopsForRouteListResponseDataEntryStopGrouping) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stopsForRouteListResponseDataEntryStopGroupingJSON) RawJSON() string {
	return r.raw
}

type StopsForRouteListResponseDataEntryStopGroupingsName struct {
	Name  string                                                  `json:"name"`
	Names []string                                                `json:"names"`
	Type  string                                                  `json:"type"`
	JSON  stopsForRouteListResponseDataEntryStopGroupingsNameJSON `json:"-"`
}

// stopsForRouteListResponseDataEntryStopGroupingsNameJSON contains the JSON
// metadata for the struct [StopsForRouteListResponseDataEntryStopGroupingsName]
type stopsForRouteListResponseDataEntryStopGroupingsNameJSON struct {
	Name        apijson.Field
	Names       apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StopsForRouteListResponseDataEntryStopGroupingsName) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stopsForRouteListResponseDataEntryStopGroupingsNameJSON) RawJSON() string {
	return r.raw
}

type StopsForRouteListResponseDataEntryStopGroupingsPolyline struct {
	Length int64                                                       `json:"length"`
	Levels string                                                      `json:"levels"`
	Points string                                                      `json:"points"`
	JSON   stopsForRouteListResponseDataEntryStopGroupingsPolylineJSON `json:"-"`
}

// stopsForRouteListResponseDataEntryStopGroupingsPolylineJSON contains the JSON
// metadata for the struct
// [StopsForRouteListResponseDataEntryStopGroupingsPolyline]
type stopsForRouteListResponseDataEntryStopGroupingsPolylineJSON struct {
	Length      apijson.Field
	Levels      apijson.Field
	Points      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StopsForRouteListResponseDataEntryStopGroupingsPolyline) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stopsForRouteListResponseDataEntryStopGroupingsPolylineJSON) RawJSON() string {
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
