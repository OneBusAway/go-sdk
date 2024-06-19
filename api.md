# Where

## AgenciesWithCoverage

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/open-transit-go">opentransit</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/open-transit-go#WhereAgenciesWithCoverageListResponse">WhereAgenciesWithCoverageListResponse</a>

Methods:

- <code title="get /api/where/agencies-with-coverage.json">client.Where.AgenciesWithCoverage.<a href="https://pkg.go.dev/github.com/stainless-sdks/open-transit-go#WhereAgenciesWithCoverageService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/open-transit-go">opentransit</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/open-transit-go#WhereAgenciesWithCoverageListResponse">WhereAgenciesWithCoverageListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Config

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/open-transit-go">opentransit</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/open-transit-go#WhereConfigGetResponse">WhereConfigGetResponse</a>

Methods:

- <code title="get /api/where/config.json">client.Where.Config.<a href="https://pkg.go.dev/github.com/stainless-sdks/open-transit-go#WhereConfigService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/open-transit-go">opentransit</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/open-transit-go#WhereConfigGetResponse">WhereConfigGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## CurrentTime

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/open-transit-go">opentransit</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/open-transit-go#WhereCurrentTimeGetResponse">WhereCurrentTimeGetResponse</a>

Methods:

- <code title="get /api/where/current-time.json">client.Where.CurrentTime.<a href="https://pkg.go.dev/github.com/stainless-sdks/open-transit-go#WhereCurrentTimeService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/open-transit-go">opentransit</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/open-transit-go#WhereCurrentTimeGetResponse">WhereCurrentTimeGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## StopsForLocation

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/open-transit-go">opentransit</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/open-transit-go#WhereStopsForLocationListResponse">WhereStopsForLocationListResponse</a>

Methods:

- <code title="get /api/where/stops-for-location.json">client.Where.StopsForLocation.<a href="https://pkg.go.dev/github.com/stainless-sdks/open-transit-go#WhereStopsForLocationService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/open-transit-go">opentransit</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/open-transit-go#WhereStopsForLocationListParams">WhereStopsForLocationListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/open-transit-go">opentransit</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/open-transit-go#WhereStopsForLocationListResponse">WhereStopsForLocationListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Stop

### ArrivalAndDeparture

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/open-transit-go">opentransit</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/open-transit-go#WhereStopArrivalAndDepartureGetResponse">WhereStopArrivalAndDepartureGetResponse</a>

Methods:

- <code title="get /api/where/arrival-and-departure-for-stop/{stopID}.json">client.Where.Stop.ArrivalAndDeparture.<a href="https://pkg.go.dev/github.com/stainless-sdks/open-transit-go#WhereStopArrivalAndDepartureService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, stopID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/open-transit-go">opentransit</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/open-transit-go#WhereStopArrivalAndDepartureGetParams">WhereStopArrivalAndDepartureGetParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/open-transit-go">opentransit</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/open-transit-go#WhereStopArrivalAndDepartureGetResponse">WhereStopArrivalAndDepartureGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### ArrivalsAndDepartures

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/open-transit-go">opentransit</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/open-transit-go#WhereStopArrivalsAndDepartureListResponse">WhereStopArrivalsAndDepartureListResponse</a>

Methods:

- <code title="get /api/where/arrivals-and-departures-for-stop/{stopID}.json">client.Where.Stop.ArrivalsAndDepartures.<a href="https://pkg.go.dev/github.com/stainless-sdks/open-transit-go#WhereStopArrivalsAndDepartureService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, stopID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/open-transit-go">opentransit</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/open-transit-go#WhereStopArrivalsAndDepartureListResponse">WhereStopArrivalsAndDepartureListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
