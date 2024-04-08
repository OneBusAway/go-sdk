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

// CurrentTimeService contains methods and other services that help with
// interacting with the open-transit API. Note, unlike clients, this service does
// not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the [NewCurrentTimeService]
// method instead.
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
func (r *CurrentTimeService) Get(ctx context.Context, query CurrentTimeGetParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "api/where/current-time.json"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, nil, opts...)
	return
}

type CurrentTimeGetParams struct {
	Key param.Field[string] `query:"key"`
}

// URLQuery serializes [CurrentTimeGetParams]'s query parameters as `url.Values`.
func (r CurrentTimeGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
