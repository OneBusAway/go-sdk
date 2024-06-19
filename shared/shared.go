// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"github.com/stainless-sdks/open-transit-go/internal/apijson"
)

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
