// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package opentransit

import (
	"github.com/stainless-sdks/open-transit-go/option"
)

// APIWhereService contains methods and other services that help with interacting
// with the OneBusAway API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAPIWhereService] method instead.
type APIWhereService struct {
	Options                      []option.RequestOption
	AgenciesWithCoverage         *APIWhereAgenciesWithCoverageService
	Config                       *APIWhereConfigService
	CurrentTime                  *APIWhereCurrentTimeService
	StopsForLocation             *APIWhereStopsForLocationService
	ArrivalAndDepartureForStop   *APIWhereArrivalAndDepartureForStopService
	ArrivalsAndDeparturesForStop *APIWhereArrivalsAndDeparturesForStopService
}

// NewAPIWhereService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAPIWhereService(opts ...option.RequestOption) (r *APIWhereService) {
	r = &APIWhereService{}
	r.Options = opts
	r.AgenciesWithCoverage = NewAPIWhereAgenciesWithCoverageService(opts...)
	r.Config = NewAPIWhereConfigService(opts...)
	r.CurrentTime = NewAPIWhereCurrentTimeService(opts...)
	r.StopsForLocation = NewAPIWhereStopsForLocationService(opts...)
	r.ArrivalAndDepartureForStop = NewAPIWhereArrivalAndDepartureForStopService(opts...)
	r.ArrivalsAndDeparturesForStop = NewAPIWhereArrivalsAndDeparturesForStopService(opts...)
	return
}
