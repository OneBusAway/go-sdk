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

// StopIDsForAgencyService contains methods and other services that help with
// interacting with the onebusaway-sdk API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewStopIDsForAgencyService] method instead.
type StopIDsForAgencyService struct {
	Options []option.RequestOption
}

// NewStopIDsForAgencyService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewStopIDsForAgencyService(opts ...option.RequestOption) (r *StopIDsForAgencyService) {
	r = &StopIDsForAgencyService{}
	r.Options = opts
	return
}

// Get stop IDs for a specific agency
func (r *StopIDsForAgencyService) List(ctx context.Context, agencyID string, opts ...option.RequestOption) (res *StopIDsForAgencyListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if agencyID == "" {
		err = errors.New("missing required agencyID parameter")
		return
	}
	path := fmt.Sprintf("api/where/stop-ids-for-agency/agencyID.json")
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type StopIDsForAgencyListResponse struct {
	Data StopIDsForAgencyListResponseData `json:"data,required"`
	JSON stopIDsForAgencyListResponseJSON `json:"-"`
	shared.ResponseWrapper
}

// stopIDsForAgencyListResponseJSON contains the JSON metadata for the struct
// [StopIDsForAgencyListResponse]
type stopIDsForAgencyListResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StopIDsForAgencyListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stopIDsForAgencyListResponseJSON) RawJSON() string {
	return r.raw
}

type StopIDsForAgencyListResponseData struct {
	LimitExceeded bool                                 `json:"limitExceeded,required"`
	List          []string                             `json:"list,required"`
	References    shared.References                    `json:"references,required"`
	JSON          stopIDsForAgencyListResponseDataJSON `json:"-"`
}

// stopIDsForAgencyListResponseDataJSON contains the JSON metadata for the struct
// [StopIDsForAgencyListResponseData]
type stopIDsForAgencyListResponseDataJSON struct {
	LimitExceeded apijson.Field
	List          apijson.Field
	References    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *StopIDsForAgencyListResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stopIDsForAgencyListResponseDataJSON) RawJSON() string {
	return r.raw
}
