// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package onebusaway

import (
	"context"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/open-transit-go/internal/apijson"
	"github.com/stainless-sdks/open-transit-go/internal/apiquery"
	"github.com/stainless-sdks/open-transit-go/internal/param"
	"github.com/stainless-sdks/open-transit-go/internal/requestconfig"
	"github.com/stainless-sdks/open-transit-go/option"
	"github.com/stainless-sdks/open-transit-go/shared"
)

// SearchForRouteService contains methods and other services that help with
// interacting with the onebusaway-sdk API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSearchForRouteService] method instead.
type SearchForRouteService struct {
	Options []option.RequestOption
}

// NewSearchForRouteService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSearchForRouteService(opts ...option.RequestOption) (r *SearchForRouteService) {
	r = &SearchForRouteService{}
	r.Options = opts
	return
}

// Search for a route based on its name.
func (r *SearchForRouteService) List(ctx context.Context, query SearchForRouteListParams, opts ...option.RequestOption) (res *SearchForRouteListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/where/search/route.json"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type SearchForRouteListResponse struct {
	Data SearchForRouteListResponseData `json:"data"`
	JSON searchForRouteListResponseJSON `json:"-"`
	shared.ResponseWrapper
}

// searchForRouteListResponseJSON contains the JSON metadata for the struct
// [SearchForRouteListResponse]
type searchForRouteListResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SearchForRouteListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r searchForRouteListResponseJSON) RawJSON() string {
	return r.raw
}

type SearchForRouteListResponseData struct {
	LimitExceeded bool                                 `json:"limitExceeded,required"`
	List          []SearchForRouteListResponseDataList `json:"list,required"`
	OutOfRange    bool                                 `json:"outOfRange,required"`
	References    shared.References                    `json:"references,required"`
	JSON          searchForRouteListResponseDataJSON   `json:"-"`
}

// searchForRouteListResponseDataJSON contains the JSON metadata for the struct
// [SearchForRouteListResponseData]
type searchForRouteListResponseDataJSON struct {
	LimitExceeded apijson.Field
	List          apijson.Field
	OutOfRange    apijson.Field
	References    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *SearchForRouteListResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r searchForRouteListResponseDataJSON) RawJSON() string {
	return r.raw
}

type SearchForRouteListResponseDataList struct {
	ID                string                                 `json:"id,required"`
	AgencyID          string                                 `json:"agencyId,required"`
	Type              int64                                  `json:"type,required"`
	Color             string                                 `json:"color"`
	Description       string                                 `json:"description"`
	LongName          string                                 `json:"longName"`
	NullSafeShortName string                                 `json:"nullSafeShortName"`
	ShortName         string                                 `json:"shortName"`
	TextColor         string                                 `json:"textColor"`
	URL               string                                 `json:"url"`
	JSON              searchForRouteListResponseDataListJSON `json:"-"`
}

// searchForRouteListResponseDataListJSON contains the JSON metadata for the struct
// [SearchForRouteListResponseDataList]
type searchForRouteListResponseDataListJSON struct {
	ID                apijson.Field
	AgencyID          apijson.Field
	Type              apijson.Field
	Color             apijson.Field
	Description       apijson.Field
	LongName          apijson.Field
	NullSafeShortName apijson.Field
	ShortName         apijson.Field
	TextColor         apijson.Field
	URL               apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *SearchForRouteListResponseDataList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r searchForRouteListResponseDataListJSON) RawJSON() string {
	return r.raw
}

type SearchForRouteListParams struct {
	// The string to search for.
	Input param.Field[string] `query:"input,required"`
	// The max number of results to return. Defaults to 20.
	MaxCount param.Field[int64] `query:"maxCount"`
}

// URLQuery serializes [SearchForRouteListParams]'s query parameters as
// `url.Values`.
func (r SearchForRouteListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
