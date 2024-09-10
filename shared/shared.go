// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"github.com/stainless-sdks/open-transit-go/internal/apijson"
)

type References struct {
	Agencies   []ReferencesAgency    `json:"agencies,required"`
	Routes     []ReferencesRoute     `json:"routes,required"`
	Situations []ReferencesSituation `json:"situations,required"`
	Stops      []ReferencesStop      `json:"stops,required"`
	StopTimes  []ReferencesStopTime  `json:"stopTimes,required"`
	Trips      []ReferencesTrip      `json:"trips,required"`
	JSON       referencesJSON        `json:"-"`
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
	ID                string              `json:"id,required"`
	AgencyID          string              `json:"agencyId,required"`
	Type              int64               `json:"type,required"`
	Color             string              `json:"color"`
	Description       string              `json:"description"`
	LongName          string              `json:"longName"`
	NullSafeShortName string              `json:"nullSafeShortName"`
	ShortName         string              `json:"shortName"`
	TextColor         string              `json:"textColor"`
	URL               string              `json:"url"`
	JSON              referencesRouteJSON `json:"-"`
}

// referencesRouteJSON contains the JSON metadata for the struct [ReferencesRoute]
type referencesRouteJSON struct {
	ID                apijson.Field
	AgencyID          apijson.Field
	Type              apijson.Field
	Color             apijson.Field
	Description       apijson.Field
	LongName          apijson.Field
	NullSafeShortName apijson.Field
	ShortName         apijson.Field
	TextColor         apijson.Field
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

type ReferencesSituation struct {
	// Unique identifier for the situation.
	ID string `json:"id,required"`
	// Unix timestamp of when this situation was created.
	CreationTime  int64                              `json:"creationTime,required"`
	ActiveWindows []ReferencesSituationsActiveWindow `json:"activeWindows"`
	AllAffects    []ReferencesSituationsAllAffect    `json:"allAffects"`
	// Message regarding the consequence of the situation.
	ConsequenceMessage string                                  `json:"consequenceMessage"`
	Consequences       []ReferencesSituationsConsequence       `json:"consequences"`
	Description        ReferencesSituationsDescription         `json:"description"`
	PublicationWindows []ReferencesSituationsPublicationWindow `json:"publicationWindows"`
	// Reason for the service alert, taken from TPEG codes.
	Reason ReferencesSituationsReason `json:"reason"`
	// Severity of the situation.
	Severity string                      `json:"severity"`
	Summary  ReferencesSituationsSummary `json:"summary"`
	URL      ReferencesSituationsURL     `json:"url"`
	JSON     referencesSituationJSON     `json:"-"`
}

// referencesSituationJSON contains the JSON metadata for the struct
// [ReferencesSituation]
type referencesSituationJSON struct {
	ID                 apijson.Field
	CreationTime       apijson.Field
	ActiveWindows      apijson.Field
	AllAffects         apijson.Field
	ConsequenceMessage apijson.Field
	Consequences       apijson.Field
	Description        apijson.Field
	PublicationWindows apijson.Field
	Reason             apijson.Field
	Severity           apijson.Field
	Summary            apijson.Field
	URL                apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ReferencesSituation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r referencesSituationJSON) RawJSON() string {
	return r.raw
}

type ReferencesSituationsActiveWindow struct {
	// Start time of the active window as a Unix timestamp.
	From int64 `json:"from"`
	// End time of the active window as a Unix timestamp.
	To   int64                                `json:"to"`
	JSON referencesSituationsActiveWindowJSON `json:"-"`
}

// referencesSituationsActiveWindowJSON contains the JSON metadata for the struct
// [ReferencesSituationsActiveWindow]
type referencesSituationsActiveWindowJSON struct {
	From        apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ReferencesSituationsActiveWindow) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r referencesSituationsActiveWindowJSON) RawJSON() string {
	return r.raw
}

type ReferencesSituationsAllAffect struct {
	// Identifier for the agency.
	AgencyID string `json:"agencyId"`
	// Identifier for the application.
	ApplicationID string `json:"applicationId"`
	// Identifier for the direction.
	DirectionID string `json:"directionId"`
	// Identifier for the route.
	RouteID string `json:"routeId"`
	// Identifier for the stop.
	StopID string `json:"stopId"`
	// Identifier for the trip.
	TripID string                            `json:"tripId"`
	JSON   referencesSituationsAllAffectJSON `json:"-"`
}

// referencesSituationsAllAffectJSON contains the JSON metadata for the struct
// [ReferencesSituationsAllAffect]
type referencesSituationsAllAffectJSON struct {
	AgencyID      apijson.Field
	ApplicationID apijson.Field
	DirectionID   apijson.Field
	RouteID       apijson.Field
	StopID        apijson.Field
	TripID        apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ReferencesSituationsAllAffect) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r referencesSituationsAllAffectJSON) RawJSON() string {
	return r.raw
}

type ReferencesSituationsConsequence struct {
	// Condition of the consequence.
	Condition        string                                           `json:"condition"`
	ConditionDetails ReferencesSituationsConsequencesConditionDetails `json:"conditionDetails"`
	JSON             referencesSituationsConsequenceJSON              `json:"-"`
}

// referencesSituationsConsequenceJSON contains the JSON metadata for the struct
// [ReferencesSituationsConsequence]
type referencesSituationsConsequenceJSON struct {
	Condition        apijson.Field
	ConditionDetails apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ReferencesSituationsConsequence) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r referencesSituationsConsequenceJSON) RawJSON() string {
	return r.raw
}

type ReferencesSituationsConsequencesConditionDetails struct {
	DiversionPath    ReferencesSituationsConsequencesConditionDetailsDiversionPath `json:"diversionPath"`
	DiversionStopIDs []string                                                      `json:"diversionStopIds"`
	JSON             referencesSituationsConsequencesConditionDetailsJSON          `json:"-"`
}

// referencesSituationsConsequencesConditionDetailsJSON contains the JSON metadata
// for the struct [ReferencesSituationsConsequencesConditionDetails]
type referencesSituationsConsequencesConditionDetailsJSON struct {
	DiversionPath    apijson.Field
	DiversionStopIDs apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ReferencesSituationsConsequencesConditionDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r referencesSituationsConsequencesConditionDetailsJSON) RawJSON() string {
	return r.raw
}

type ReferencesSituationsConsequencesConditionDetailsDiversionPath struct {
	// Length of the diversion path.
	Length int64 `json:"length"`
	// Levels of the diversion path.
	Levels string `json:"levels"`
	// Points of the diversion path.
	Points string                                                            `json:"points"`
	JSON   referencesSituationsConsequencesConditionDetailsDiversionPathJSON `json:"-"`
}

// referencesSituationsConsequencesConditionDetailsDiversionPathJSON contains the
// JSON metadata for the struct
// [ReferencesSituationsConsequencesConditionDetailsDiversionPath]
type referencesSituationsConsequencesConditionDetailsDiversionPathJSON struct {
	Length      apijson.Field
	Levels      apijson.Field
	Points      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ReferencesSituationsConsequencesConditionDetailsDiversionPath) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r referencesSituationsConsequencesConditionDetailsDiversionPathJSON) RawJSON() string {
	return r.raw
}

type ReferencesSituationsDescription struct {
	// Language of the description.
	Lang string `json:"lang"`
	// Longer description of the situation.
	Value string                              `json:"value"`
	JSON  referencesSituationsDescriptionJSON `json:"-"`
}

// referencesSituationsDescriptionJSON contains the JSON metadata for the struct
// [ReferencesSituationsDescription]
type referencesSituationsDescriptionJSON struct {
	Lang        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ReferencesSituationsDescription) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r referencesSituationsDescriptionJSON) RawJSON() string {
	return r.raw
}

type ReferencesSituationsPublicationWindow struct {
	// Start time of the time window as a Unix timestamp.
	From int64 `json:"from,required"`
	// End time of the time window as a Unix timestamp.
	To   int64                                     `json:"to,required"`
	JSON referencesSituationsPublicationWindowJSON `json:"-"`
}

// referencesSituationsPublicationWindowJSON contains the JSON metadata for the
// struct [ReferencesSituationsPublicationWindow]
type referencesSituationsPublicationWindowJSON struct {
	From        apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ReferencesSituationsPublicationWindow) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r referencesSituationsPublicationWindowJSON) RawJSON() string {
	return r.raw
}

// Reason for the service alert, taken from TPEG codes.
type ReferencesSituationsReason string

const (
	ReferencesSituationsReasonEquipmentReason     ReferencesSituationsReason = "equipmentReason"
	ReferencesSituationsReasonEnvironmentReason   ReferencesSituationsReason = "environmentReason"
	ReferencesSituationsReasonPersonnelReason     ReferencesSituationsReason = "personnelReason"
	ReferencesSituationsReasonMiscellaneousReason ReferencesSituationsReason = "miscellaneousReason"
	ReferencesSituationsReasonSecurityAlert       ReferencesSituationsReason = "securityAlert"
)

func (r ReferencesSituationsReason) IsKnown() bool {
	switch r {
	case ReferencesSituationsReasonEquipmentReason, ReferencesSituationsReasonEnvironmentReason, ReferencesSituationsReasonPersonnelReason, ReferencesSituationsReasonMiscellaneousReason, ReferencesSituationsReasonSecurityAlert:
		return true
	}
	return false
}

type ReferencesSituationsSummary struct {
	// Language of the summary.
	Lang string `json:"lang"`
	// Short summary of the situation.
	Value string                          `json:"value"`
	JSON  referencesSituationsSummaryJSON `json:"-"`
}

// referencesSituationsSummaryJSON contains the JSON metadata for the struct
// [ReferencesSituationsSummary]
type referencesSituationsSummaryJSON struct {
	Lang        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ReferencesSituationsSummary) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r referencesSituationsSummaryJSON) RawJSON() string {
	return r.raw
}

type ReferencesSituationsURL struct {
	// Language of the URL.
	Lang string `json:"lang"`
	// URL for more information about the situation.
	Value string                      `json:"value"`
	JSON  referencesSituationsURLJSON `json:"-"`
}

// referencesSituationsURLJSON contains the JSON metadata for the struct
// [ReferencesSituationsURL]
type referencesSituationsURLJSON struct {
	Lang        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ReferencesSituationsURL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r referencesSituationsURLJSON) RawJSON() string {
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

type ReferencesStopTime struct {
	ArrivalTime         int64                  `json:"arrivalTime"`
	DepartureTime       int64                  `json:"departureTime"`
	DistanceAlongTrip   float64                `json:"distanceAlongTrip"`
	HistoricalOccupancy string                 `json:"historicalOccupancy"`
	StopHeadsign        string                 `json:"stopHeadsign"`
	StopID              string                 `json:"stopId"`
	JSON                referencesStopTimeJSON `json:"-"`
}

// referencesStopTimeJSON contains the JSON metadata for the struct
// [ReferencesStopTime]
type referencesStopTimeJSON struct {
	ArrivalTime         apijson.Field
	DepartureTime       apijson.Field
	DistanceAlongTrip   apijson.Field
	HistoricalOccupancy apijson.Field
	StopHeadsign        apijson.Field
	StopID              apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *ReferencesStopTime) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r referencesStopTimeJSON) RawJSON() string {
	return r.raw
}

type ReferencesTrip struct {
	ID             string             `json:"id,required"`
	RouteID        string             `json:"routeId,required"`
	ServiceID      string             `json:"serviceId,required"`
	BlockID        string             `json:"blockId"`
	DirectionID    string             `json:"directionId"`
	PeakOffpeak    int64              `json:"peakOffpeak"`
	RouteShortName string             `json:"routeShortName"`
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
	ServiceID      apijson.Field
	BlockID        apijson.Field
	DirectionID    apijson.Field
	PeakOffpeak    apijson.Field
	RouteShortName apijson.Field
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
