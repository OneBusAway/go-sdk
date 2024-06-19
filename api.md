# Shared Response Types

- <a href="https://pkg.go.dev/github.com/stainless-sdks/open-transit-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/open-transit-go/shared#ResponseWrapper">ResponseWrapper</a>

# AgenciesWithCoverage

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/open-transit-go">onebusaway</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/open-transit-go#AgenciesWithCoverageGetResponse">AgenciesWithCoverageGetResponse</a>

Methods:

- <code title="get /api/where/agencies-with-coverage.json">client.AgenciesWithCoverage.<a href="https://pkg.go.dev/github.com/stainless-sdks/open-transit-go#AgenciesWithCoverageService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/open-transit-go">onebusaway</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/open-transit-go#AgenciesWithCoverageGetResponse">AgenciesWithCoverageGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Config

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/open-transit-go">onebusaway</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/open-transit-go#ConfigGetResponse">ConfigGetResponse</a>

Methods:

- <code title="get /api/where/config.json">client.Config.<a href="https://pkg.go.dev/github.com/stainless-sdks/open-transit-go#ConfigService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/open-transit-go">onebusaway</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/open-transit-go#ConfigGetResponse">ConfigGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# CurrentTime

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/open-transit-go">onebusaway</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/open-transit-go#CurrentTimeGetResponse">CurrentTimeGetResponse</a>

Methods:

- <code title="get /api/where/current-time.json">client.CurrentTime.<a href="https://pkg.go.dev/github.com/stainless-sdks/open-transit-go#CurrentTimeService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/open-transit-go">onebusaway</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/open-transit-go#CurrentTimeGetResponse">CurrentTimeGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# StopsForLocation

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/open-transit-go">onebusaway</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/open-transit-go#StopsForLocationGetResponse">StopsForLocationGetResponse</a>

Methods:

- <code title="get /api/where/stops-for-location.json">client.StopsForLocation.<a href="https://pkg.go.dev/github.com/stainless-sdks/open-transit-go#StopsForLocationService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/open-transit-go">onebusaway</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/open-transit-go#StopsForLocationGetParams">StopsForLocationGetParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/open-transit-go">onebusaway</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/open-transit-go#StopsForLocationGetResponse">StopsForLocationGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# ArrivalAndDepartureForStop

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/open-transit-go">onebusaway</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/open-transit-go#ArrivalAndDepartureForStopGetResponse">ArrivalAndDepartureForStopGetResponse</a>

Methods:

- <code title="get /api/where/arrival-and-departure-for-stop/{stopID}.json">client.ArrivalAndDepartureForStop.<a href="https://pkg.go.dev/github.com/stainless-sdks/open-transit-go#ArrivalAndDepartureForStopService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, stopID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/open-transit-go">onebusaway</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/open-transit-go#ArrivalAndDepartureForStopGetParams">ArrivalAndDepartureForStopGetParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/open-transit-go">onebusaway</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/open-transit-go#ArrivalAndDepartureForStopGetResponse">ArrivalAndDepartureForStopGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# ArrivalsAndDeparturesForStop

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/open-transit-go">onebusaway</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/open-transit-go#ArrivalsAndDeparturesForStopGetResponse">ArrivalsAndDeparturesForStopGetResponse</a>

Methods:

- <code title="get /api/where/arrivals-and-departures-for-stop/{stopID}">client.ArrivalsAndDeparturesForStop.<a href="https://pkg.go.dev/github.com/stainless-sdks/open-transit-go#ArrivalsAndDeparturesForStopService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, stopID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/open-transit-go">onebusaway</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/open-transit-go#ArrivalsAndDeparturesForStopGetResponse">ArrivalsAndDeparturesForStopGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
