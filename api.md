# Shared Response Types

- <a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk/shared">shared</a>.<a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk/shared#References">References</a>
- <a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk/shared">shared</a>.<a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk/shared#ResponseWrapper">ResponseWrapper</a>

# AgenciesWithCoverage

Response Types:

- <a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk">onebusaway</a>.<a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk#AgenciesWithCoverageListResponse">AgenciesWithCoverageListResponse</a>

Methods:

- <code title="get /api/where/agencies-with-coverage.json">client.AgenciesWithCoverage.<a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk#AgenciesWithCoverageService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk">onebusaway</a>.<a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk#AgenciesWithCoverageListResponse">AgenciesWithCoverageListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Agency

Response Types:

- <a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk">onebusaway</a>.<a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk#AgencyGetResponse">AgencyGetResponse</a>

Methods:

- <code title="get /api/where/agency/{agencyID}.json">client.Agency.<a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk#AgencyService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, agencyID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk">onebusaway</a>.<a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk#AgencyGetResponse">AgencyGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# VehiclesForAgency

Response Types:

- <a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk">onebusaway</a>.<a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk#VehiclesForAgencyListResponse">VehiclesForAgencyListResponse</a>

Methods:

- <code title="get /api/where/vehicles-for-agency/{agencyID}.json">client.VehiclesForAgency.<a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk#VehiclesForAgencyService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, agencyID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk">onebusaway</a>.<a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk#VehiclesForAgencyListParams">VehiclesForAgencyListParams</a>) (<a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk">onebusaway</a>.<a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk#VehiclesForAgencyListResponse">VehiclesForAgencyListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Config

Response Types:

- <a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk">onebusaway</a>.<a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk#ConfigGetResponse">ConfigGetResponse</a>

Methods:

- <code title="get /api/where/config.json">client.Config.<a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk#ConfigService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk">onebusaway</a>.<a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk#ConfigGetResponse">ConfigGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# CurrentTime

Response Types:

- <a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk">onebusaway</a>.<a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk#CurrentTimeGetResponse">CurrentTimeGetResponse</a>

Methods:

- <code title="get /api/where/current-time.json">client.CurrentTime.<a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk#CurrentTimeService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk">onebusaway</a>.<a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk#CurrentTimeGetResponse">CurrentTimeGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# StopsForLocation

Response Types:

- <a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk">onebusaway</a>.<a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk#StopsForLocationListResponse">StopsForLocationListResponse</a>

Methods:

- <code title="get /api/where/stops-for-location.json">client.StopsForLocation.<a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk#StopsForLocationService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk">onebusaway</a>.<a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk#StopsForLocationListParams">StopsForLocationListParams</a>) (<a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk">onebusaway</a>.<a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk#StopsForLocationListResponse">StopsForLocationListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# StopsForRoute

Response Types:

- <a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk">onebusaway</a>.<a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk#StopsForRouteListResponse">StopsForRouteListResponse</a>

Methods:

- <code title="get /api/where/stops-for-route/{routeID}.json">client.StopsForRoute.<a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk#StopsForRouteService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, routeID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk">onebusaway</a>.<a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk#StopsForRouteListParams">StopsForRouteListParams</a>) (<a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk">onebusaway</a>.<a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk#StopsForRouteListResponse">StopsForRouteListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Stop

Response Types:

- <a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk">onebusaway</a>.<a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk#StopGetResponse">StopGetResponse</a>

Methods:

- <code title="get /api/where/stop/{stopID}.json">client.Stop.<a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk#StopService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, stopID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk">onebusaway</a>.<a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk#StopGetResponse">StopGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# StopIDsForAgency

Response Types:

- <a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk">onebusaway</a>.<a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk#StopIDsForAgencyListResponse">StopIDsForAgencyListResponse</a>

Methods:

- <code title="get /api/where/stop-ids-for-agency/{agencyID}.json">client.StopIDsForAgency.<a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk#StopIDsForAgencyService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, agencyID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk">onebusaway</a>.<a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk#StopIDsForAgencyListResponse">StopIDsForAgencyListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# ScheduleForStop

Response Types:

- <a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk">onebusaway</a>.<a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk#ScheduleForStopGetResponse">ScheduleForStopGetResponse</a>

Methods:

- <code title="get /api/where/schedule-for-stop/{stopID}.json">client.ScheduleForStop.<a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk#ScheduleForStopService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, stopID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk">onebusaway</a>.<a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk#ScheduleForStopGetParams">ScheduleForStopGetParams</a>) (<a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk">onebusaway</a>.<a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk#ScheduleForStopGetResponse">ScheduleForStopGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Route

Response Types:

- <a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk">onebusaway</a>.<a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk#RouteGetResponse">RouteGetResponse</a>

Methods:

- <code title="get /api/where/route/{routeID}.json">client.Route.<a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk#RouteService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, routeID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk">onebusaway</a>.<a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk#RouteGetResponse">RouteGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# RouteIDsForAgency

Response Types:

- <a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk">onebusaway</a>.<a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk#RouteIDsForAgencyListResponse">RouteIDsForAgencyListResponse</a>

Methods:

- <code title="get /api/where/route-ids-for-agency/{agencyID}.json">client.RouteIDsForAgency.<a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk#RouteIDsForAgencyService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, agencyID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk">onebusaway</a>.<a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk#RouteIDsForAgencyListResponse">RouteIDsForAgencyListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# RoutesForLocation

Response Types:

- <a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk">onebusaway</a>.<a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk#RoutesForLocationListResponse">RoutesForLocationListResponse</a>

Methods:

- <code title="get /api/where/routes-for-location.json">client.RoutesForLocation.<a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk#RoutesForLocationService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk">onebusaway</a>.<a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk#RoutesForLocationListParams">RoutesForLocationListParams</a>) (<a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk">onebusaway</a>.<a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk#RoutesForLocationListResponse">RoutesForLocationListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# RoutesForAgency

Response Types:

- <a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk">onebusaway</a>.<a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk#RoutesForAgencyListResponse">RoutesForAgencyListResponse</a>

Methods:

- <code title="get /api/where/routes-for-agency/{agencyID}.json">client.RoutesForAgency.<a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk#RoutesForAgencyService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, agencyID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk">onebusaway</a>.<a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk#RoutesForAgencyListResponse">RoutesForAgencyListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# ScheduleForRoute

Response Types:

- <a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk">onebusaway</a>.<a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk#ScheduleForRouteGetResponse">ScheduleForRouteGetResponse</a>

Methods:

- <code title="get /api/where/schedule-for-route/{routeID}.json">client.ScheduleForRoute.<a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk#ScheduleForRouteService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, routeID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk">onebusaway</a>.<a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk#ScheduleForRouteGetParams">ScheduleForRouteGetParams</a>) (<a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk">onebusaway</a>.<a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk#ScheduleForRouteGetResponse">ScheduleForRouteGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# ArrivalAndDeparture

Response Types:

- <a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk">onebusaway</a>.<a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk#ArrivalAndDepartureGetResponse">ArrivalAndDepartureGetResponse</a>
- <a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk">onebusaway</a>.<a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk#ArrivalAndDepartureListResponse">ArrivalAndDepartureListResponse</a>

Methods:

- <code title="get /api/where/arrival-and-departure-for-stop/{stopID}.json">client.ArrivalAndDeparture.<a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk#ArrivalAndDepartureService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, stopID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk">onebusaway</a>.<a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk#ArrivalAndDepartureGetParams">ArrivalAndDepartureGetParams</a>) (<a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk">onebusaway</a>.<a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk#ArrivalAndDepartureGetResponse">ArrivalAndDepartureGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/where/arrivals-and-departures-for-stop/{stopID}.json">client.ArrivalAndDeparture.<a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk#ArrivalAndDepartureService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, stopID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk">onebusaway</a>.<a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk#ArrivalAndDepartureListParams">ArrivalAndDepartureListParams</a>) (<a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk">onebusaway</a>.<a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk#ArrivalAndDepartureListResponse">ArrivalAndDepartureListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Trip

Response Types:

- <a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk">onebusaway</a>.<a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk#TripGetResponse">TripGetResponse</a>

Methods:

- <code title="get /api/where/trip/{tripID}.json">client.Trip.<a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk#TripService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, tripID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk">onebusaway</a>.<a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk#TripGetResponse">TripGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# TripsForLocation

Response Types:

- <a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk">onebusaway</a>.<a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk#TripsForLocationListResponse">TripsForLocationListResponse</a>

Methods:

- <code title="get /api/where/trips-for-location.json">client.TripsForLocation.<a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk#TripsForLocationService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk">onebusaway</a>.<a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk#TripsForLocationListParams">TripsForLocationListParams</a>) (<a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk">onebusaway</a>.<a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk#TripsForLocationListResponse">TripsForLocationListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# TripDetails

Response Types:

- <a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk">onebusaway</a>.<a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk#TripDetailGetResponse">TripDetailGetResponse</a>

Methods:

- <code title="get /api/where/trip-details/{tripID}.json">client.TripDetails.<a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk#TripDetailService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, tripID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk">onebusaway</a>.<a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk#TripDetailGetParams">TripDetailGetParams</a>) (<a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk">onebusaway</a>.<a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk#TripDetailGetResponse">TripDetailGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# TripForVehicle

Response Types:

- <a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk">onebusaway</a>.<a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk#TripForVehicleGetResponse">TripForVehicleGetResponse</a>

Methods:

- <code title="get /api/where/trip-for-vehicle/{vehicleID}.json">client.TripForVehicle.<a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk#TripForVehicleService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, vehicleID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk">onebusaway</a>.<a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk#TripForVehicleGetParams">TripForVehicleGetParams</a>) (<a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk">onebusaway</a>.<a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk#TripForVehicleGetResponse">TripForVehicleGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# TripsForRoute

Response Types:

- <a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk">onebusaway</a>.<a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk#TripsForRouteListResponse">TripsForRouteListResponse</a>

Methods:

- <code title="get /api/where/trips-for-route/{routeID}.json">client.TripsForRoute.<a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk#TripsForRouteService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, routeID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk">onebusaway</a>.<a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk#TripsForRouteListParams">TripsForRouteListParams</a>) (<a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk">onebusaway</a>.<a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk#TripsForRouteListResponse">TripsForRouteListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# ReportProblemWithStop

Methods:

- <code title="get /api/where/report-problem-with-stop/{stopID}.json">client.ReportProblemWithStop.<a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk#ReportProblemWithStopService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, stopID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk">onebusaway</a>.<a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk#ReportProblemWithStopGetParams">ReportProblemWithStopGetParams</a>) (<a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk/shared">shared</a>.<a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk/shared#ResponseWrapper">ResponseWrapper</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# ReportProblemWithTrip

Methods:

- <code title="get /api/where/report-problem-with-trip/{tripID}.json">client.ReportProblemWithTrip.<a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk#ReportProblemWithTripService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, tripID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk">onebusaway</a>.<a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk#ReportProblemWithTripGetParams">ReportProblemWithTripGetParams</a>) (<a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk/shared">shared</a>.<a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk/shared#ResponseWrapper">ResponseWrapper</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# SearchForStop

Response Types:

- <a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk">onebusaway</a>.<a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk#SearchForStopListResponse">SearchForStopListResponse</a>

Methods:

- <code title="get /api/where/search/stop.json">client.SearchForStop.<a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk#SearchForStopService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk">onebusaway</a>.<a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk#SearchForStopListParams">SearchForStopListParams</a>) (<a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk">onebusaway</a>.<a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk#SearchForStopListResponse">SearchForStopListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# SearchForRoute

Response Types:

- <a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk">onebusaway</a>.<a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk#SearchForRouteListResponse">SearchForRouteListResponse</a>

Methods:

- <code title="get /api/where/search/route.json">client.SearchForRoute.<a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk#SearchForRouteService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk">onebusaway</a>.<a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk#SearchForRouteListParams">SearchForRouteListParams</a>) (<a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk">onebusaway</a>.<a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk#SearchForRouteListResponse">SearchForRouteListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Block

Response Types:

- <a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk">onebusaway</a>.<a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk#BlockGetResponse">BlockGetResponse</a>

Methods:

- <code title="get /api/where/block/{blockID}.json">client.Block.<a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk#BlockService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, blockID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk">onebusaway</a>.<a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk#BlockGetResponse">BlockGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Shape

Response Types:

- <a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk">onebusaway</a>.<a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk#ShapeGetResponse">ShapeGetResponse</a>

Methods:

- <code title="get /api/where/shape/{shapeID}.json">client.Shape.<a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk#ShapeService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, shapeID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk">onebusaway</a>.<a href="https://pkg.go.dev/github.com/OneBusAway/go-sdk#ShapeGetResponse">ShapeGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
