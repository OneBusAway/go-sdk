# Shared Response Types

- <a href="https://pkg.go.dev/github.com/stainless-sdks/open-transit-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/open-transit-go/shared#ResponseWrapper">ResponseWrapper</a>

# API

## Where

### AgenciesWithCoverage

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/open-transit-go">opentransit</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/open-transit-go#APIWhereAgenciesWithCoverageListResponse">APIWhereAgenciesWithCoverageListResponse</a>

Methods:

- <code title="get /api/where/agencies-with-coverage.json">client.API.Where.AgenciesWithCoverage.<a href="https://pkg.go.dev/github.com/stainless-sdks/open-transit-go#APIWhereAgenciesWithCoverageService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/open-transit-go">opentransit</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/open-transit-go#APIWhereAgenciesWithCoverageListResponse">APIWhereAgenciesWithCoverageListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Config

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/open-transit-go">opentransit</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/open-transit-go#APIWhereConfigGetResponse">APIWhereConfigGetResponse</a>

Methods:

- <code title="get /api/where/config.json">client.API.Where.Config.<a href="https://pkg.go.dev/github.com/stainless-sdks/open-transit-go#APIWhereConfigService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/open-transit-go">opentransit</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/open-transit-go#APIWhereConfigGetResponse">APIWhereConfigGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### CurrentTime

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/open-transit-go">opentransit</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/open-transit-go#APIWhereCurrentTimeGetResponse">APIWhereCurrentTimeGetResponse</a>

Methods:

- <code title="get /api/where/current-time.json">client.API.Where.CurrentTime.<a href="https://pkg.go.dev/github.com/stainless-sdks/open-transit-go#APIWhereCurrentTimeService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/open-transit-go">opentransit</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/open-transit-go#APIWhereCurrentTimeGetResponse">APIWhereCurrentTimeGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### StopsForLocation

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/open-transit-go">opentransit</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/open-transit-go#APIWhereStopsForLocationListResponse">APIWhereStopsForLocationListResponse</a>

Methods:

- <code title="get /api/where/stops-for-location.json">client.API.Where.StopsForLocation.<a href="https://pkg.go.dev/github.com/stainless-sdks/open-transit-go#APIWhereStopsForLocationService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/open-transit-go">opentransit</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/open-transit-go#APIWhereStopsForLocationListParams">APIWhereStopsForLocationListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/open-transit-go">opentransit</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/open-transit-go#APIWhereStopsForLocationListResponse">APIWhereStopsForLocationListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### ArrivalAndDepartureForStop

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/open-transit-go">opentransit</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/open-transit-go#APIWhereArrivalAndDepartureForStopGetResponse">APIWhereArrivalAndDepartureForStopGetResponse</a>

Methods:

- <code title="get /api/where/arrival-and-departure-for-stop/{stopID}.json">client.API.Where.ArrivalAndDepartureForStop.<a href="https://pkg.go.dev/github.com/stainless-sdks/open-transit-go#APIWhereArrivalAndDepartureForStopService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, stopID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/open-transit-go">opentransit</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/open-transit-go#APIWhereArrivalAndDepartureForStopGetParams">APIWhereArrivalAndDepartureForStopGetParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/open-transit-go">opentransit</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/open-transit-go#APIWhereArrivalAndDepartureForStopGetResponse">APIWhereArrivalAndDepartureForStopGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### ArrivalsAndDeparturesForStop

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/open-transit-go">opentransit</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/open-transit-go#APIWhereArrivalsAndDeparturesForStopListResponse">APIWhereArrivalsAndDeparturesForStopListResponse</a>

Methods:

- <code title="get /api/where/arrivals-and-departures-for-stop/{stopID}.json">client.API.Where.ArrivalsAndDeparturesForStop.<a href="https://pkg.go.dev/github.com/stainless-sdks/open-transit-go#APIWhereArrivalsAndDeparturesForStopService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, stopID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/open-transit-go">opentransit</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/open-transit-go#APIWhereArrivalsAndDeparturesForStopListResponse">APIWhereArrivalsAndDeparturesForStopListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
