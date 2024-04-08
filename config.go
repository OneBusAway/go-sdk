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

// ConfigService contains methods and other services that help with interacting
// with the open-transit API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewConfigService] method instead.
type ConfigService struct {
	Options []option.RequestOption
}

// NewConfigService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewConfigService(opts ...option.RequestOption) (r *ConfigService) {
	r = &ConfigService{}
	r.Options = opts
	return
}

// config
func (r *ConfigService) Get(ctx context.Context, query ConfigGetParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "api/where/config.json"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, nil, opts...)
	return
}

type ConfigGetParams struct {
	Key param.Field[string] `query:"key"`
}

// URLQuery serializes [ConfigGetParams]'s query parameters as `url.Values`.
func (r ConfigGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
