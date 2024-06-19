// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"github.com/stainless-sdks/open-transit-go/internal/apijson"
)

type References struct {
	Agencies   []ReferencesAgency `json:"agencies"`
	Routes     []ReferencesRoute  `json:"routes"`
	Situations []interface{}      `json:"situations"`
	Stops      []ReferencesStop   `json:"stops"`
	StopTimes  []interface{}      `json:"stopTimes"`
	Trips      []ReferencesTrip   `json:"trips"`
	JSON       referencesJSON     `json:"-"`
}

// referencesJSON contains the JSON metadata for the struct [References]
type referencesJSON struct {
	Agencies    apijson.Field
	Routes      apijson.Field
	Situations  apijson.Field
	Stops       apijson.Field
	StopTimes   apijson.Field
	Trips       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *References) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r referencesJSON) RawJSON() string {
	return r.raw
}

type ReferencesAgency struct {
	ID             string               `json:"id,required"`
	Name           string               `json:"name,required"`
	Timezone       string               `json:"timezone,required"`
	URL            string               `json:"url,required"`
	Disclaimer     string               `json:"disclaimer"`
	Email          string               `json:"email"`
	FareURL        string               `json:"fareUrl"`
	Lang           string               `json:"lang"`
	Phone          string               `json:"phone"`
	PrivateService bool                 `json:"privateService"`
	JSON           referencesAgencyJSON `json:"-"`
}

// referencesAgencyJSON contains the JSON metadata for the struct
// [ReferencesAgency]
type referencesAgencyJSON struct {
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

func (r *ReferencesAgency) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r referencesAgencyJSON) RawJSON() string {
	return r.raw
}

type ReferencesRoute struct {
	ID                string              `json:"id"`
	AgencyID          string              `json:"agencyId"`
	Color             string              `json:"color"`
	Description       string              `json:"description"`
	LongName          string              `json:"longName"`
	NullSafeShortName string              `json:"nullSafeShortName"`
	ShortName         string              `json:"shortName"`
	TextColor         string              `json:"textColor"`
	Type              int64               `json:"type"`
	URL               string              `json:"url"`
	JSON              referencesRouteJSON `json:"-"`
}

// referencesRouteJSON contains the JSON metadata for the struct [ReferencesRoute]
type referencesRouteJSON struct {
	ID                apijson.Field
	AgencyID          apijson.Field
	Color             apijson.Field
	Description       apijson.Field
	LongName          apijson.Field
	NullSafeShortName apijson.Field
	ShortName         apijson.Field
	TextColor         apijson.Field
	Type              apijson.Field
	URL               apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ReferencesRoute) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r referencesRouteJSON) RawJSON() string {
	return r.raw
}

type ReferencesStop struct {
	ID                 string             `json:"id,required"`
	Code               string             `json:"code,required"`
	Lat                float64            `json:"lat,required"`
	Lon                float64            `json:"lon,required"`
	Name               string             `json:"name,required"`
	Direction          string             `json:"direction"`
	LocationType       int64              `json:"locationType"`
	Parent             string             `json:"parent"`
	RouteIDs           []string           `json:"routeIds"`
	StaticRouteIDs     []string           `json:"staticRouteIds"`
	WheelchairBoarding string             `json:"wheelchairBoarding"`
	JSON               referencesStopJSON `json:"-"`
}

// referencesStopJSON contains the JSON metadata for the struct [ReferencesStop]
type referencesStopJSON struct {
	ID                 apijson.Field
	Code               apijson.Field
	Lat                apijson.Field
	Lon                apijson.Field
	Name               apijson.Field
	Direction          apijson.Field
	LocationType       apijson.Field
	Parent             apijson.Field
	RouteIDs           apijson.Field
	StaticRouteIDs     apijson.Field
	WheelchairBoarding apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ReferencesStop) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r referencesStopJSON) RawJSON() string {
	return r.raw
}

type ReferencesTrip struct {
	ID             string             `json:"id,required"`
	RouteID        string             `json:"routeId,required"`
	BlockID        string             `json:"blockId"`
	DirectionID    string             `json:"directionId"`
	PeakOffpeak    int64              `json:"peakOffpeak"`
	RouteShortName string             `json:"routeShortName"`
	ServiceID      string             `json:"serviceId"`
	ShapeID        string             `json:"shapeId"`
	TimeZone       string             `json:"timeZone"`
	TripHeadsign   string             `json:"tripHeadsign"`
	TripShortName  string             `json:"tripShortName"`
	JSON           referencesTripJSON `json:"-"`
}

// referencesTripJSON contains the JSON metadata for the struct [ReferencesTrip]
type referencesTripJSON struct {
	ID             apijson.Field
	RouteID        apijson.Field
	BlockID        apijson.Field
	DirectionID    apijson.Field
	PeakOffpeak    apijson.Field
	RouteShortName apijson.Field
	ServiceID      apijson.Field
	ShapeID        apijson.Field
	TimeZone       apijson.Field
	TripHeadsign   apijson.Field
	TripShortName  apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ReferencesTrip) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r referencesTripJSON) RawJSON() string {
	return r.raw
}

type ResponseWrapper struct {
	Code        int64               `json:"code,required"`
	CurrentTime int64               `json:"currentTime,required"`
	Text        string              `json:"text,required"`
	Version     int64               `json:"version,required"`
	JSON        responseWrapperJSON `json:"-"`
}

// responseWrapperJSON contains the JSON metadata for the struct [ResponseWrapper]
type responseWrapperJSON struct {
	Code        apijson.Field
	CurrentTime apijson.Field
	Text        apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseWrapper) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r responseWrapperJSON) RawJSON() string {
	return r.raw
}
