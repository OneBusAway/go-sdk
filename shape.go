// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package onebusaway

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/open-transit-go/internal/apijson"
	"github.com/stainless-sdks/open-transit-go/internal/requestconfig"
	"github.com/stainless-sdks/open-transit-go/option"
	"github.com/stainless-sdks/open-transit-go/shared"
)

// ShapeService contains methods and other services that help with interacting with
// the onebusaway-sdk API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewShapeService] method instead.
type ShapeService struct {
	Options []option.RequestOption
}

// NewShapeService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewShapeService(opts ...option.RequestOption) (r *ShapeService) {
	r = &ShapeService{}
	r.Options = opts
	return
}

// Retrieve a shape (the path traveled by a transit vehicle) by ID.
func (r *ShapeService) Get(ctx context.Context, shapeID string, opts ...option.RequestOption) (res *ShapeGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if shapeID == "" {
		err = errors.New("missing required shapeID parameter")
		return
	}
	path := fmt.Sprintf("api/where/shape/shapeID.json")
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ShapeGetResponse struct {
	Data ShapeGetResponseData `json:"data,required"`
	JSON shapeGetResponseJSON `json:"-"`
	shared.ResponseWrapper
}

// shapeGetResponseJSON contains the JSON metadata for the struct
// [ShapeGetResponse]
type shapeGetResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ShapeGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r shapeGetResponseJSON) RawJSON() string {
	return r.raw
}

type ShapeGetResponseData struct {
	Entry      ShapeGetResponseDataEntry `json:"entry,required"`
	References shared.References         `json:"references,required"`
	JSON       shapeGetResponseDataJSON  `json:"-"`
}

// shapeGetResponseDataJSON contains the JSON metadata for the struct
// [ShapeGetResponseData]
type shapeGetResponseDataJSON struct {
	Entry       apijson.Field
	References  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ShapeGetResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r shapeGetResponseDataJSON) RawJSON() string {
	return r.raw
}

type ShapeGetResponseDataEntry struct {
	Length int64 `json:"length,required"`
	// Encoded polyline format representing the shape of the path
	Points string                        `json:"points,required"`
	Levels string                        `json:"levels"`
	JSON   shapeGetResponseDataEntryJSON `json:"-"`
}

// shapeGetResponseDataEntryJSON contains the JSON metadata for the struct
// [ShapeGetResponseDataEntry]
type shapeGetResponseDataEntryJSON struct {
	Length      apijson.Field
	Points      apijson.Field
	Levels      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ShapeGetResponseDataEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r shapeGetResponseDataEntryJSON) RawJSON() string {
	return r.raw
}
