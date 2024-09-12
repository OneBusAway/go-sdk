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

// TripService contains methods and other services that help with interacting with
// the onebusaway-sdk API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTripService] method instead.
type TripService struct {
	Options []option.RequestOption
}

// NewTripService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewTripService(opts ...option.RequestOption) (r *TripService) {
	r = &TripService{}
	r.Options = opts
	return
}

// Get details of a specific trip
func (r *TripService) Get(ctx context.Context, tripID string, opts ...option.RequestOption) (res *TripGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if tripID == "" {
		err = errors.New("missing required tripID parameter")
		return
	}
	path := fmt.Sprintf("api/where/trip/%s.json", tripID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type TripGetResponse struct {
	Data TripGetResponseData `json:"data,required"`
	JSON tripGetResponseJSON `json:"-"`
	shared.ResponseWrapper
}

// tripGetResponseJSON contains the JSON metadata for the struct [TripGetResponse]
type tripGetResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TripGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tripGetResponseJSON) RawJSON() string {
	return r.raw
}

type TripGetResponseData struct {
	Entry      TripGetResponseDataEntry `json:"entry,required"`
	References shared.References        `json:"references,required"`
	JSON       tripGetResponseDataJSON  `json:"-"`
}

// tripGetResponseDataJSON contains the JSON metadata for the struct
// [TripGetResponseData]
type tripGetResponseDataJSON struct {
	Entry       apijson.Field
	References  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TripGetResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tripGetResponseDataJSON) RawJSON() string {
	return r.raw
}

type TripGetResponseDataEntry struct {
	ID             string                       `json:"id,required"`
	RouteID        string                       `json:"routeId,required"`
	ServiceID      string                       `json:"serviceId,required"`
	BlockID        string                       `json:"blockId"`
	DirectionID    string                       `json:"directionId"`
	PeakOffpeak    int64                        `json:"peakOffpeak"`
	RouteShortName string                       `json:"routeShortName"`
	ShapeID        string                       `json:"shapeId"`
	TimeZone       string                       `json:"timeZone"`
	TripHeadsign   string                       `json:"tripHeadsign"`
	TripShortName  string                       `json:"tripShortName"`
	JSON           tripGetResponseDataEntryJSON `json:"-"`
}

// tripGetResponseDataEntryJSON contains the JSON metadata for the struct
// [TripGetResponseDataEntry]
type tripGetResponseDataEntryJSON struct {
	ID             apijson.Field
	RouteID        apijson.Field
	ServiceID      apijson.Field
	BlockID        apijson.Field
	DirectionID    apijson.Field
	PeakOffpeak    apijson.Field
	RouteShortName apijson.Field
	ShapeID        apijson.Field
	TimeZone       apijson.Field
	TripHeadsign   apijson.Field
	TripShortName  apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *TripGetResponseDataEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tripGetResponseDataEntryJSON) RawJSON() string {
	return r.raw
}
