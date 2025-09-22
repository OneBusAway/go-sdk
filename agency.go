// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package onebusaway

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/OneBusAway/go-sdk/internal/apijson"
	"github.com/OneBusAway/go-sdk/internal/requestconfig"
	"github.com/OneBusAway/go-sdk/option"
	"github.com/OneBusAway/go-sdk/shared"
)

// AgencyService contains methods and other services that help with interacting
// with the onebusaway-sdk API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAgencyService] method instead.
type AgencyService struct {
	Options []option.RequestOption
}

// NewAgencyService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAgencyService(opts ...option.RequestOption) (r *AgencyService) {
	r = &AgencyService{}
	r.Options = opts
	return
}

// Retrieve information for a specific transit agency identified by its unique ID.
func (r *AgencyService) Get(ctx context.Context, agencyID string, opts ...option.RequestOption) (res *AgencyGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if agencyID == "" {
		err = errors.New("missing required agencyID parameter")
		return
	}
	path := fmt.Sprintf("api/where/agency/%s.json", agencyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AgencyGetResponse struct {
	Data AgencyGetResponseData `json:"data,required"`
	JSON agencyGetResponseJSON `json:"-"`
	shared.ResponseWrapper
}

// agencyGetResponseJSON contains the JSON metadata for the struct
// [AgencyGetResponse]
type agencyGetResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AgencyGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r agencyGetResponseJSON) RawJSON() string {
	return r.raw
}

type AgencyGetResponseData struct {
	Entry         AgencyGetResponseDataEntry `json:"entry,required"`
	LimitExceeded bool                       `json:"limitExceeded,required"`
	References    shared.References          `json:"references,required"`
	JSON          agencyGetResponseDataJSON  `json:"-"`
}

// agencyGetResponseDataJSON contains the JSON metadata for the struct
// [AgencyGetResponseData]
type agencyGetResponseDataJSON struct {
	Entry         apijson.Field
	LimitExceeded apijson.Field
	References    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AgencyGetResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r agencyGetResponseDataJSON) RawJSON() string {
	return r.raw
}

type AgencyGetResponseDataEntry struct {
	ID             string                         `json:"id,required"`
	Name           string                         `json:"name,required"`
	Timezone       string                         `json:"timezone,required"`
	URL            string                         `json:"url,required"`
	Disclaimer     string                         `json:"disclaimer"`
	Email          string                         `json:"email"`
	FareURL        string                         `json:"fareUrl"`
	Lang           string                         `json:"lang"`
	Phone          string                         `json:"phone"`
	PrivateService bool                           `json:"privateService"`
	JSON           agencyGetResponseDataEntryJSON `json:"-"`
}

// agencyGetResponseDataEntryJSON contains the JSON metadata for the struct
// [AgencyGetResponseDataEntry]
type agencyGetResponseDataEntryJSON struct {
	ID             apijson.Field
	Name           apijson.Field
	Timezone       apijson.Field
	URL            apijson.Field
	Disclaimer     apijson.Field
	Email          apijson.Field
	FareURL        apijson.Field
	Lang           apijson.Field
	Phone          apijson.Field
	PrivateService apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AgencyGetResponseDataEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r agencyGetResponseDataEntryJSON) RawJSON() string {
	return r.raw
}
