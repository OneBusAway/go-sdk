// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package tempopentransit

import (
	"context"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/TEMP_open-transit-go/internal/apiquery"
	"github.com/stainless-sdks/TEMP_open-transit-go/internal/param"
	"github.com/stainless-sdks/TEMP_open-transit-go/internal/requestconfig"
	"github.com/stainless-sdks/TEMP_open-transit-go/option"
)

// StopsForLocationService contains methods and other services that help with
// interacting with the open-transit API. Note, unlike clients, this service does
// not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewStopsForLocationService] method instead.
type StopsForLocationService struct {
	Options []option.RequestOption
}

// NewStopsForLocationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewStopsForLocationService(opts ...option.RequestOption) (r *StopsForLocationService) {
	r = &StopsForLocationService{}
	r.Options = opts
	return
}

// stops-for-location
func (r *StopsForLocationService) List(ctx context.Context, query StopsForLocationListParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "api/where/stops-for-location.json"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, nil, opts...)
	return
}

type StopsForLocationListParams struct {
	Key param.Field[string]  `query:"key"`
	Lat param.Field[float64] `query:"lat"`
	Lon param.Field[float64] `query:"lon"`
}

// URLQuery serializes [StopsForLocationListParams]'s query parameters as
// `url.Values`.
func (r StopsForLocationListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
