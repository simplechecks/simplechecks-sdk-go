// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package simplechecksgo

import (
	"context"
	"net/http"
	"slices"

	"github.com/simplechecks/simplechecks-sdk-go/internal/apijson"
	"github.com/simplechecks/simplechecks-sdk-go/internal/requestconfig"
	"github.com/simplechecks/simplechecks-sdk-go/option"
)

// PricingService contains methods and other services that help with interacting
// with the simple-checks API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPricingService] method instead.
type PricingService struct {
	Options []option.RequestOption
}

// NewPricingService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewPricingService(opts ...option.RequestOption) (r *PricingService) {
	r = &PricingService{}
	r.Options = opts
	return
}

// Returns the active token-pricing table so a client can show the per-provider
// cost of a check at configuration time. The cost of one run is
// `floor(weight × multiplier_milli / 1000)`, where `weight` is the check type's
// compute weight plus its artifact-egress component, and the multiplier resolves
// `(provider, location)` → `(provider, "")` → `1.0` (returned as `1000` milli).
// The result equals what metering debits, so a UI preview is exact.
//
// The provider multiplier is the customer-facing cost lever: cheaper providers
// (e.g. OVH, Hetzner) carry a multiplier below 1.0. Reads of this table are free.
//
// Requires the `account:read` scope — pricing is incidental to account/check
// configuration, not a per-check write.
func (r *PricingService) Get(ctx context.Context, opts ...option.RequestOption) (res *Pricing, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/pricing"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// The active token-pricing table. cost(run) =
// `floor(weight × multiplier_milli / 1000)`, multiplier resolving
// `(provider, location)` → `(provider, "")` → `1.0` (1000 milli).
type Pricing struct {
	CheckTypes  []PricingCheckType  `json:"check_types" api:"required"`
	Multipliers []PricingMultiplier `json:"multipliers" api:"required"`
	JSON        pricingJSON         `json:"-"`
}

// pricingJSON contains the JSON metadata for the struct [Pricing]
type pricingJSON struct {
	CheckTypes  apijson.Field
	Multipliers apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Pricing) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pricingJSON) RawJSON() string {
	return r.raw
}

// One check type's per-run weight (compute + artifact egress), pre-multiplier.
type PricingCheckType struct {
	// Check type identifier.
	CheckType string `json:"check_type" api:"required"`
	// The artifact-egress portion of `weight` (0 for non-artifact types). Surfaced so
	// a UI can label the artifact-retrieval cost of a browser/playwright run.
	EgressWeight int64 `json:"egress_weight" api:"required"`
	// Per-run weight, compute plus artifact egress (pre-multiplier).
	Weight int64                `json:"weight" api:"required"`
	JSON   pricingCheckTypeJSON `json:"-"`
}

// pricingCheckTypeJSON contains the JSON metadata for the struct
// [PricingCheckType]
type pricingCheckTypeJSON struct {
	CheckType    apijson.Field
	EgressWeight apijson.Field
	Weight       apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *PricingCheckType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pricingCheckTypeJSON) RawJSON() string {
	return r.raw
}

// One (provider, location) cost multiplier.
type PricingMultiplier struct {
	// Provider-native location id; empty for a provider-wide default.
	Location string `json:"location" api:"required"`
	// Multiplier × 1000 (e.g. 500 = 0.5×, the cheap-provider wedge).
	MultiplierMilli int64                 `json:"multiplier_milli" api:"required"`
	Provider        string                `json:"provider" api:"required"`
	JSON            pricingMultiplierJSON `json:"-"`
}

// pricingMultiplierJSON contains the JSON metadata for the struct
// [PricingMultiplier]
type pricingMultiplierJSON struct {
	Location        apijson.Field
	MultiplierMilli apijson.Field
	Provider        apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *PricingMultiplier) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pricingMultiplierJSON) RawJSON() string {
	return r.raw
}
