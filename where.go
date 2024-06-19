// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package opentransit

import (
	"github.com/stainless-sdks/open-transit-go/option"
)

// WhereService contains methods and other services that help with interacting with
// the open-transit API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWhereService] method instead.
type WhereService struct {
	Options              []option.RequestOption
	AgenciesWithCoverage *WhereAgenciesWithCoverageService
	Config               *WhereConfigService
	CurrentTime          *WhereCurrentTimeService
	StopsForLocation     *WhereStopsForLocationService
	Stop                 *WhereStopService
}

// NewWhereService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewWhereService(opts ...option.RequestOption) (r *WhereService) {
	r = &WhereService{}
	r.Options = opts
	r.AgenciesWithCoverage = NewWhereAgenciesWithCoverageService(opts...)
	r.Config = NewWhereConfigService(opts...)
	r.CurrentTime = NewWhereCurrentTimeService(opts...)
	r.StopsForLocation = NewWhereStopsForLocationService(opts...)
	r.Stop = NewWhereStopService(opts...)
	return
}
