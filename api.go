// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package opentransit

import (
	"github.com/stainless-sdks/open-transit-go/option"
)

// APIService contains methods and other services that help with interacting with
// the OneBusAway API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAPIService] method instead.
type APIService struct {
	Options []option.RequestOption
	Where   *APIWhereService
}

// NewAPIService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAPIService(opts ...option.RequestOption) (r *APIService) {
	r = &APIService{}
	r.Options = opts
	r.Where = NewAPIWhereService(opts...)
	return
}
