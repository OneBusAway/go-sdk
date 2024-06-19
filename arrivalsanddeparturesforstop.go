// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package onebusaway

import (
	"github.com/stainless-sdks/open-transit-go/option"
)

// ArrivalsAndDeparturesForStopService contains methods and other services that
// help with interacting with the OneBusAway API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewArrivalsAndDeparturesForStopService] method instead.
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
