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

// AgenciesWithCoverageService contains methods and other services that help with
// interacting with the open-transit API. Note, unlike clients, this service does
// not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewAgenciesWithCoverageService] method instead.
type AgenciesWithCoverageService struct {
	Options []option.RequestOption
}

// NewAgenciesWithCoverageService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAgenciesWithCoverageService(opts ...option.RequestOption) (r *AgenciesWithCoverageService) {
	r = &AgenciesWithCoverageService{}
	r.Options = opts
	return
}

// agencies-with-coverage
func (r *AgenciesWithCoverageService) List(ctx context.Context, query AgenciesWithCoverageListParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "api/where/agencies-with-coverage.json"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, nil, opts...)
	return
}

type AgenciesWithCoverageListParams struct {
	Key param.Field[string] `query:"key"`
}

// URLQuery serializes [AgenciesWithCoverageListParams]'s query parameters as
// `url.Values`.
func (r AgenciesWithCoverageListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
