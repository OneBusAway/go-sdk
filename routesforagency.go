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

// RoutesForAgencyService contains methods and other services that help with
// interacting with the onebusaway-sdk API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRoutesForAgencyService] method instead.
type RoutesForAgencyService struct {
	Options []option.RequestOption
}

// NewRoutesForAgencyService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRoutesForAgencyService(opts ...option.RequestOption) (r *RoutesForAgencyService) {
	r = &RoutesForAgencyService{}
	r.Options = opts
	return
}

// Retrieve the list of all routes for a particular agency by id
func (r *RoutesForAgencyService) List(ctx context.Context, agencyID string, opts ...option.RequestOption) (res *RoutesForAgencyListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if agencyID == "" {
		err = errors.New("missing required agencyID parameter")
		return
	}
	path := fmt.Sprintf("api/where/routes-for-agency/%s.json", agencyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type RoutesForAgencyListResponse struct {
	Data RoutesForAgencyListResponseData `json:"data,required"`
	JSON routesForAgencyListResponseJSON `json:"-"`
	shared.ResponseWrapper
}

// routesForAgencyListResponseJSON contains the JSON metadata for the struct
// [RoutesForAgencyListResponse]
type routesForAgencyListResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RoutesForAgencyListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routesForAgencyListResponseJSON) RawJSON() string {
	return r.raw
}

type RoutesForAgencyListResponseData struct {
	LimitExceeded bool                                  `json:"limitExceeded,required"`
	List          []RoutesForAgencyListResponseDataList `json:"list,required"`
	References    shared.References                     `json:"references,required"`
	JSON          routesForAgencyListResponseDataJSON   `json:"-"`
}

// routesForAgencyListResponseDataJSON contains the JSON metadata for the struct
// [RoutesForAgencyListResponseData]
type routesForAgencyListResponseDataJSON struct {
	LimitExceeded apijson.Field
	List          apijson.Field
	References    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *RoutesForAgencyListResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routesForAgencyListResponseDataJSON) RawJSON() string {
	return r.raw
}

type RoutesForAgencyListResponseDataList struct {
	ID                string                                  `json:"id,required"`
	AgencyID          string                                  `json:"agencyId,required"`
	Type              int64                                   `json:"type,required"`
	Color             string                                  `json:"color"`
	Description       string                                  `json:"description"`
	LongName          string                                  `json:"longName"`
	NullSafeShortName string                                  `json:"nullSafeShortName"`
	ShortName         string                                  `json:"shortName"`
	TextColor         string                                  `json:"textColor"`
	URL               string                                  `json:"url"`
	JSON              routesForAgencyListResponseDataListJSON `json:"-"`
}

// routesForAgencyListResponseDataListJSON contains the JSON metadata for the
// struct [RoutesForAgencyListResponseDataList]
type routesForAgencyListResponseDataListJSON struct {
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

func (r *RoutesForAgencyListResponseDataList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routesForAgencyListResponseDataListJSON) RawJSON() string {
	return r.raw
}
