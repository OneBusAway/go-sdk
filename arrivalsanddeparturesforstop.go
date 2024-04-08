// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package tempopentransit

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/TEMP_open-transit-go/internal/apiquery"
	"github.com/stainless-sdks/TEMP_open-transit-go/internal/param"
	"github.com/stainless-sdks/TEMP_open-transit-go/internal/requestconfig"
	"github.com/stainless-sdks/TEMP_open-transit-go/option"
)

// ArrivalsAndDeparturesForStopService contains methods and other services that
// help with interacting with the open-transit API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewArrivalsAndDeparturesForStopService] method instead.
type ArrivalsAndDeparturesForStopService struct {
	Options []option.RequestOption
}

// NewArrivalsAndDeparturesForStopService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewArrivalsAndDeparturesForStopService(opts ...option.RequestOption) (r *ArrivalsAndDeparturesForStopService) {
	r = &ArrivalsAndDeparturesForStopService{}
	r.Options = opts
	return
}

// arrivals-and-departures-for-stop
func (r *ArrivalsAndDeparturesForStopService) Get(ctx context.Context, stopIDJson string, query ArrivalsAndDeparturesForStopGetParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("api/where/arrivals-and-departures-for-stop/%s", stopIDJson)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, nil, opts...)
	return
}

type ArrivalsAndDeparturesForStopGetParams struct {
	Key param.Field[string] `query:"key"`
}

// URLQuery serializes [ArrivalsAndDeparturesForStopGetParams]'s query parameters
// as `url.Values`.
func (r ArrivalsAndDeparturesForStopGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
