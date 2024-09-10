// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package onebusaway

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/open-transit-go/internal/apijson"
	"github.com/stainless-sdks/open-transit-go/internal/requestconfig"
	"github.com/stainless-sdks/open-transit-go/option"
	"github.com/stainless-sdks/open-transit-go/shared"
)

// BlockService contains methods and other services that help with interacting with
// the onebusaway-sdk API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBlockService] method instead.
type BlockService struct {
	Options []option.RequestOption
}

// NewBlockService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewBlockService(opts ...option.RequestOption) (r *BlockService) {
	r = &BlockService{}
	r.Options = opts
	return
}

// Get details of a specific block by ID
func (r *BlockService) Get(ctx context.Context, blockID string, opts ...option.RequestOption) (res *BlockGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if blockID == "" {
		err = errors.New("missing required blockID parameter")
		return
	}
	path := fmt.Sprintf("api/where/block/%s.json", blockID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type BlockGetResponse struct {
	Data BlockGetResponseData `json:"data,required"`
	JSON blockGetResponseJSON `json:"-"`
	shared.ResponseWrapper
}

// blockGetResponseJSON contains the JSON metadata for the struct
// [BlockGetResponse]
type blockGetResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BlockGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r blockGetResponseJSON) RawJSON() string {
	return r.raw
}

type BlockGetResponseData struct {
	Entry      BlockGetResponseDataEntry `json:"entry,required"`
	References shared.References         `json:"references,required"`
	JSON       blockGetResponseDataJSON  `json:"-"`
}

// blockGetResponseDataJSON contains the JSON metadata for the struct
// [BlockGetResponseData]
type blockGetResponseDataJSON struct {
	Entry       apijson.Field
	References  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BlockGetResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r blockGetResponseDataJSON) RawJSON() string {
	return r.raw
}

type BlockGetResponseDataEntry struct {
	ID             string                                   `json:"id,required"`
	Configurations []BlockGetResponseDataEntryConfiguration `json:"configurations,required"`
	JSON           blockGetResponseDataEntryJSON            `json:"-"`
}

// blockGetResponseDataEntryJSON contains the JSON metadata for the struct
// [BlockGetResponseDataEntry]
type blockGetResponseDataEntryJSON struct {
	ID             apijson.Field
	Configurations apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *BlockGetResponseDataEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r blockGetResponseDataEntryJSON) RawJSON() string {
	return r.raw
}

type BlockGetResponseDataEntryConfiguration struct {
	ActiveServiceIDs   []string                                      `json:"activeServiceIds,required"`
	Trips              []BlockGetResponseDataEntryConfigurationsTrip `json:"trips,required"`
	InactiveServiceIDs []string                                      `json:"inactiveServiceIds"`
	JSON               blockGetResponseDataEntryConfigurationJSON    `json:"-"`
}

// blockGetResponseDataEntryConfigurationJSON contains the JSON metadata for the
// struct [BlockGetResponseDataEntryConfiguration]
type blockGetResponseDataEntryConfigurationJSON struct {
	ActiveServiceIDs   apijson.Field
	Trips              apijson.Field
	InactiveServiceIDs apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *BlockGetResponseDataEntryConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r blockGetResponseDataEntryConfigurationJSON) RawJSON() string {
	return r.raw
}

type BlockGetResponseDataEntryConfigurationsTrip struct {
	AccumulatedSlackTime float64                                                     `json:"accumulatedSlackTime,required"`
	BlockStopTimes       []BlockGetResponseDataEntryConfigurationsTripsBlockStopTime `json:"blockStopTimes,required"`
	DistanceAlongBlock   float64                                                     `json:"distanceAlongBlock,required"`
	TripID               string                                                      `json:"tripId,required"`
	JSON                 blockGetResponseDataEntryConfigurationsTripJSON             `json:"-"`
}

// blockGetResponseDataEntryConfigurationsTripJSON contains the JSON metadata for
// the struct [BlockGetResponseDataEntryConfigurationsTrip]
type blockGetResponseDataEntryConfigurationsTripJSON struct {
	AccumulatedSlackTime apijson.Field
	BlockStopTimes       apijson.Field
	DistanceAlongBlock   apijson.Field
	TripID               apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *BlockGetResponseDataEntryConfigurationsTrip) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r blockGetResponseDataEntryConfigurationsTripJSON) RawJSON() string {
	return r.raw
}

type BlockGetResponseDataEntryConfigurationsTripsBlockStopTime struct {
	AccumulatedSlackTime float64                                                            `json:"accumulatedSlackTime,required"`
	BlockSequence        int64                                                              `json:"blockSequence,required"`
	DistanceAlongBlock   float64                                                            `json:"distanceAlongBlock,required"`
	StopTime             BlockGetResponseDataEntryConfigurationsTripsBlockStopTimesStopTime `json:"stopTime,required"`
	JSON                 blockGetResponseDataEntryConfigurationsTripsBlockStopTimeJSON      `json:"-"`
}

// blockGetResponseDataEntryConfigurationsTripsBlockStopTimeJSON contains the JSON
// metadata for the struct
// [BlockGetResponseDataEntryConfigurationsTripsBlockStopTime]
type blockGetResponseDataEntryConfigurationsTripsBlockStopTimeJSON struct {
	AccumulatedSlackTime apijson.Field
	BlockSequence        apijson.Field
	DistanceAlongBlock   apijson.Field
	StopTime             apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *BlockGetResponseDataEntryConfigurationsTripsBlockStopTime) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r blockGetResponseDataEntryConfigurationsTripsBlockStopTimeJSON) RawJSON() string {
	return r.raw
}

type BlockGetResponseDataEntryConfigurationsTripsBlockStopTimesStopTime struct {
	ArrivalTime   int64                                                                  `json:"arrivalTime,required"`
	DepartureTime int64                                                                  `json:"departureTime,required"`
	StopID        string                                                                 `json:"stopId,required"`
	DropOffType   int64                                                                  `json:"dropOffType"`
	PickupType    int64                                                                  `json:"pickupType"`
	JSON          blockGetResponseDataEntryConfigurationsTripsBlockStopTimesStopTimeJSON `json:"-"`
}

// blockGetResponseDataEntryConfigurationsTripsBlockStopTimesStopTimeJSON contains
// the JSON metadata for the struct
// [BlockGetResponseDataEntryConfigurationsTripsBlockStopTimesStopTime]
type blockGetResponseDataEntryConfigurationsTripsBlockStopTimesStopTimeJSON struct {
	ArrivalTime   apijson.Field
	DepartureTime apijson.Field
	StopID        apijson.Field
	DropOffType   apijson.Field
	PickupType    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *BlockGetResponseDataEntryConfigurationsTripsBlockStopTimesStopTime) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r blockGetResponseDataEntryConfigurationsTripsBlockStopTimesStopTimeJSON) RawJSON() string {
	return r.raw
}
