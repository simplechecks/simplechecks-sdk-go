// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package simplechecksgo

import (
	"context"
	"net/http"
	"slices"
	"time"

	"github.com/simplechecks/simplechecks-sdk-go/internal/apijson"
	"github.com/simplechecks/simplechecks-sdk-go/internal/param"
	"github.com/simplechecks/simplechecks-sdk-go/internal/requestconfig"
	"github.com/simplechecks/simplechecks-sdk-go/option"
)

// Run-credit balance, Stripe Checkout top-ups, and purchase history.
//
// CheckoutSessionService contains methods and other services that help with
// interacting with the simple-checks API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCheckoutSessionService] method instead.
type CheckoutSessionService struct {
	Options []option.RequestOption
}

// NewCheckoutSessionService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCheckoutSessionService(opts ...option.RequestOption) (r *CheckoutSessionService) {
	r = &CheckoutSessionService{}
	r.Options = opts
	return
}

// Returns a Stripe-hosted checkout URL the customer pays on. The webhook fulfils
// the purchase asynchronously after the customer completes payment. Requires the
// `billing:write` scope (opt-in; not in the default scope set, since spending
// money should be a deliberate choice).
func (r *CheckoutSessionService) New(ctx context.Context, body CheckoutSessionNewParams, opts ...option.RequestOption) (res *CheckoutSession, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/checkout-session"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type CheckoutSession struct {
	// Stripe-hosted page the customer pays on.
	CheckoutURL     string              `json:"checkout_url" api:"required" format:"uri"`
	StripeSessionID string              `json:"stripe_session_id" api:"required"`
	ExpiresAt       time.Time           `json:"expires_at" format:"date-time"`
	JSON            checkoutSessionJSON `json:"-"`
}

// checkoutSessionJSON contains the JSON metadata for the struct [CheckoutSession]
type checkoutSessionJSON struct {
	CheckoutURL     apijson.Field
	StripeSessionID apijson.Field
	ExpiresAt       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *CheckoutSession) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r checkoutSessionJSON) RawJSON() string {
	return r.raw
}

type CheckoutSessionNewParams struct {
	BundleSKU param.Field[CheckoutSessionNewParamsBundleSKU] `json:"bundle_sku" api:"required"`
}

func (r CheckoutSessionNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CheckoutSessionNewParamsBundleSKU string

const (
	CheckoutSessionNewParamsBundleSKUStarter CheckoutSessionNewParamsBundleSKU = "starter"
	CheckoutSessionNewParamsBundleSKUGrowth  CheckoutSessionNewParamsBundleSKU = "growth"
	CheckoutSessionNewParamsBundleSKUScale   CheckoutSessionNewParamsBundleSKU = "scale"
	CheckoutSessionNewParamsBundleSKUTeam    CheckoutSessionNewParamsBundleSKU = "team"
)

func (r CheckoutSessionNewParamsBundleSKU) IsKnown() bool {
	switch r {
	case CheckoutSessionNewParamsBundleSKUStarter, CheckoutSessionNewParamsBundleSKUGrowth, CheckoutSessionNewParamsBundleSKUScale, CheckoutSessionNewParamsBundleSKUTeam:
		return true
	}
	return false
}
