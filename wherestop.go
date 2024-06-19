// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package opentransit

import (
	"github.com/stainless-sdks/open-transit-go/option"
)

// WhereStopService contains methods and other services that help with interacting
// with the open-transit API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWhereStopService] method instead.
type WhereStopService struct {
	Options               []option.RequestOption
	ArrivalAndDeparture   *WhereStopArrivalAndDepartureService
	ArrivalsAndDepartures *WhereStopArrivalsAndDepartureService
}

// NewWhereStopService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWhereStopService(opts ...option.RequestOption) (r *WhereStopService) {
	r = &WhereStopService{}
	r.Options = opts
	r.ArrivalAndDeparture = NewWhereStopArrivalAndDepartureService(opts...)
	r.ArrivalsAndDepartures = NewWhereStopArrivalsAndDepartureService(opts...)
	return
}
