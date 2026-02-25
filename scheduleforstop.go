// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package onebusaway

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/OneBusAway/go-sdk/internal/apijson"
	"github.com/OneBusAway/go-sdk/internal/apiquery"
	"github.com/OneBusAway/go-sdk/internal/param"
	"github.com/OneBusAway/go-sdk/internal/requestconfig"
	"github.com/OneBusAway/go-sdk/option"
	"github.com/OneBusAway/go-sdk/shared"
)

// ScheduleForStopService contains methods and other services that help with
// interacting with the onebusaway-sdk API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewScheduleForStopService] method instead.
type ScheduleForStopService struct {
	Options []option.RequestOption
}

// NewScheduleForStopService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewScheduleForStopService(opts ...option.RequestOption) (r *ScheduleForStopService) {
	r = &ScheduleForStopService{}
	r.Options = opts
	return
}

// Get schedule for a specific stop
func (r *ScheduleForStopService) Get(ctx context.Context, stopID string, query ScheduleForStopGetParams, opts ...option.RequestOption) (res *ScheduleForStopGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if stopID == "" {
		err = errors.New("missing required stopID parameter")
		return
	}
	path := fmt.Sprintf("api/where/schedule-for-stop/%s.json", stopID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type ScheduleForStopGetResponse struct {
	Data ScheduleForStopGetResponseData `json:"data" api:"required"`
	JSON scheduleForStopGetResponseJSON `json:"-"`
	shared.ResponseWrapper
}

// scheduleForStopGetResponseJSON contains the JSON metadata for the struct
// [ScheduleForStopGetResponse]
type scheduleForStopGetResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScheduleForStopGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scheduleForStopGetResponseJSON) RawJSON() string {
	return r.raw
}

type ScheduleForStopGetResponseData struct {
	Entry      ScheduleForStopGetResponseDataEntry `json:"entry" api:"required"`
	References shared.References                   `json:"references" api:"required"`
	JSON       scheduleForStopGetResponseDataJSON  `json:"-"`
}

// scheduleForStopGetResponseDataJSON contains the JSON metadata for the struct
// [ScheduleForStopGetResponseData]
type scheduleForStopGetResponseDataJSON struct {
	Entry       apijson.Field
	References  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScheduleForStopGetResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scheduleForStopGetResponseDataJSON) RawJSON() string {
	return r.raw
}

type ScheduleForStopGetResponseDataEntry struct {
	Date               int64                                                  `json:"date" api:"required"`
	StopID             string                                                 `json:"stopId" api:"required"`
	StopRouteSchedules []ScheduleForStopGetResponseDataEntryStopRouteSchedule `json:"stopRouteSchedules" api:"required"`
	JSON               scheduleForStopGetResponseDataEntryJSON                `json:"-"`
}

// scheduleForStopGetResponseDataEntryJSON contains the JSON metadata for the
// struct [ScheduleForStopGetResponseDataEntry]
type scheduleForStopGetResponseDataEntryJSON struct {
	Date               apijson.Field
	StopID             apijson.Field
	StopRouteSchedules apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ScheduleForStopGetResponseDataEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scheduleForStopGetResponseDataEntryJSON) RawJSON() string {
	return r.raw
}

type ScheduleForStopGetResponseDataEntryStopRouteSchedule struct {
	RouteID                     string                                                                            `json:"routeId" api:"required"`
	StopRouteDirectionSchedules []ScheduleForStopGetResponseDataEntryStopRouteSchedulesStopRouteDirectionSchedule `json:"stopRouteDirectionSchedules" api:"required"`
	JSON                        scheduleForStopGetResponseDataEntryStopRouteScheduleJSON                          `json:"-"`
}

// scheduleForStopGetResponseDataEntryStopRouteScheduleJSON contains the JSON
// metadata for the struct [ScheduleForStopGetResponseDataEntryStopRouteSchedule]
type scheduleForStopGetResponseDataEntryStopRouteScheduleJSON struct {
	RouteID                     apijson.Field
	StopRouteDirectionSchedules apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *ScheduleForStopGetResponseDataEntryStopRouteSchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scheduleForStopGetResponseDataEntryStopRouteScheduleJSON) RawJSON() string {
	return r.raw
}

type ScheduleForStopGetResponseDataEntryStopRouteSchedulesStopRouteDirectionSchedule struct {
	ScheduleStopTimes   []ScheduleForStopGetResponseDataEntryStopRouteSchedulesStopRouteDirectionSchedulesScheduleStopTime  `json:"scheduleStopTimes" api:"required"`
	TripHeadsign        string                                                                                              `json:"tripHeadsign" api:"required"`
	ScheduleFrequencies []ScheduleForStopGetResponseDataEntryStopRouteSchedulesStopRouteDirectionSchedulesScheduleFrequency `json:"scheduleFrequencies"`
	JSON                scheduleForStopGetResponseDataEntryStopRouteSchedulesStopRouteDirectionScheduleJSON                 `json:"-"`
}

// scheduleForStopGetResponseDataEntryStopRouteSchedulesStopRouteDirectionScheduleJSON
// contains the JSON metadata for the struct
// [ScheduleForStopGetResponseDataEntryStopRouteSchedulesStopRouteDirectionSchedule]
type scheduleForStopGetResponseDataEntryStopRouteSchedulesStopRouteDirectionScheduleJSON struct {
	ScheduleStopTimes   apijson.Field
	TripHeadsign        apijson.Field
	ScheduleFrequencies apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *ScheduleForStopGetResponseDataEntryStopRouteSchedulesStopRouteDirectionSchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scheduleForStopGetResponseDataEntryStopRouteSchedulesStopRouteDirectionScheduleJSON) RawJSON() string {
	return r.raw
}

type ScheduleForStopGetResponseDataEntryStopRouteSchedulesStopRouteDirectionSchedulesScheduleStopTime struct {
	ArrivalEnabled   bool                                                                                                 `json:"arrivalEnabled" api:"required"`
	ArrivalTime      int64                                                                                                `json:"arrivalTime" api:"required"`
	DepartureEnabled bool                                                                                                 `json:"departureEnabled" api:"required"`
	DepartureTime    int64                                                                                                `json:"departureTime" api:"required"`
	ServiceID        string                                                                                               `json:"serviceId" api:"required"`
	TripID           string                                                                                               `json:"tripId" api:"required"`
	StopHeadsign     string                                                                                               `json:"stopHeadsign"`
	JSON             scheduleForStopGetResponseDataEntryStopRouteSchedulesStopRouteDirectionSchedulesScheduleStopTimeJSON `json:"-"`
}

// scheduleForStopGetResponseDataEntryStopRouteSchedulesStopRouteDirectionSchedulesScheduleStopTimeJSON
// contains the JSON metadata for the struct
// [ScheduleForStopGetResponseDataEntryStopRouteSchedulesStopRouteDirectionSchedulesScheduleStopTime]
type scheduleForStopGetResponseDataEntryStopRouteSchedulesStopRouteDirectionSchedulesScheduleStopTimeJSON struct {
	ArrivalEnabled   apijson.Field
	ArrivalTime      apijson.Field
	DepartureEnabled apijson.Field
	DepartureTime    apijson.Field
	ServiceID        apijson.Field
	TripID           apijson.Field
	StopHeadsign     apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ScheduleForStopGetResponseDataEntryStopRouteSchedulesStopRouteDirectionSchedulesScheduleStopTime) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scheduleForStopGetResponseDataEntryStopRouteSchedulesStopRouteDirectionSchedulesScheduleStopTimeJSON) RawJSON() string {
	return r.raw
}

type ScheduleForStopGetResponseDataEntryStopRouteSchedulesStopRouteDirectionSchedulesScheduleFrequency struct {
	EndTime     int64                                                                                                 `json:"endTime" api:"required"`
	Headway     int64                                                                                                 `json:"headway" api:"required"`
	ServiceDate int64                                                                                                 `json:"serviceDate" api:"required"`
	ServiceID   string                                                                                                `json:"serviceId" api:"required"`
	StartTime   int64                                                                                                 `json:"startTime" api:"required"`
	TripID      string                                                                                                `json:"tripId" api:"required"`
	JSON        scheduleForStopGetResponseDataEntryStopRouteSchedulesStopRouteDirectionSchedulesScheduleFrequencyJSON `json:"-"`
}

// scheduleForStopGetResponseDataEntryStopRouteSchedulesStopRouteDirectionSchedulesScheduleFrequencyJSON
// contains the JSON metadata for the struct
// [ScheduleForStopGetResponseDataEntryStopRouteSchedulesStopRouteDirectionSchedulesScheduleFrequency]
type scheduleForStopGetResponseDataEntryStopRouteSchedulesStopRouteDirectionSchedulesScheduleFrequencyJSON struct {
	EndTime     apijson.Field
	Headway     apijson.Field
	ServiceDate apijson.Field
	ServiceID   apijson.Field
	StartTime   apijson.Field
	TripID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScheduleForStopGetResponseDataEntryStopRouteSchedulesStopRouteDirectionSchedulesScheduleFrequency) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scheduleForStopGetResponseDataEntryStopRouteSchedulesStopRouteDirectionSchedulesScheduleFrequencyJSON) RawJSON() string {
	return r.raw
}

type ScheduleForStopGetParams struct {
	// The date for which you want to request a schedule in the format YYYY-MM-DD
	// (optional, defaults to the current date)
	Date param.Field[time.Time] `query:"date" format:"date"`
}

// URLQuery serializes [ScheduleForStopGetParams]'s query parameters as
// `url.Values`.
func (r ScheduleForStopGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
