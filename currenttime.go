// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package onebusaway

import (
	"context"
	"net/http"

	"github.com/stainless-sdks/open-transit-go/internal/apijson"
	"github.com/stainless-sdks/open-transit-go/internal/requestconfig"
	"github.com/stainless-sdks/open-transit-go/option"
	"github.com/stainless-sdks/open-transit-go/shared"
)

// CurrentTimeService contains methods and other services that help with
// interacting with the onebusaway-sdk API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCurrentTimeService] method instead.
type CurrentTimeService struct {
	Options []option.RequestOption
}

// NewCurrentTimeService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCurrentTimeService(opts ...option.RequestOption) (r *CurrentTimeService) {
	r = &CurrentTimeService{}
	r.Options = opts
	return
}

// current-time
func (r *CurrentTimeService) Get(ctx context.Context, opts ...option.RequestOption) (res *CurrentTimeGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/where/current-time.json"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type CurrentTimeGetResponse struct {
	Data CurrentTimeGetResponseData `json:"data"`
	JSON currentTimeGetResponseJSON `json:"-"`
	shared.ResponseWrapper
}

// currentTimeGetResponseJSON contains the JSON metadata for the struct
// [CurrentTimeGetResponse]
type currentTimeGetResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CurrentTimeGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r currentTimeGetResponseJSON) RawJSON() string {
	return r.raw
}

type CurrentTimeGetResponseData struct {
	Entry      CurrentTimeGetResponseDataEntry `json:"entry"`
	References shared.References               `json:"references"`
	JSON       currentTimeGetResponseDataJSON  `json:"-"`
}

// currentTimeGetResponseDataJSON contains the JSON metadata for the struct
// [CurrentTimeGetResponseData]
type currentTimeGetResponseDataJSON struct {
	Entry       apijson.Field
	References  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CurrentTimeGetResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r currentTimeGetResponseDataJSON) RawJSON() string {
	return r.raw
}

type CurrentTimeGetResponseDataEntry struct {
	ReadableTime string                              `json:"readableTime"`
	Time         int64                               `json:"time"`
	JSON         currentTimeGetResponseDataEntryJSON `json:"-"`
}

// currentTimeGetResponseDataEntryJSON contains the JSON metadata for the struct
// [CurrentTimeGetResponseDataEntry]
type currentTimeGetResponseDataEntryJSON struct {
	ReadableTime apijson.Field
	Time         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *CurrentTimeGetResponseDataEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r currentTimeGetResponseDataEntryJSON) RawJSON() string {
	return r.raw
}
