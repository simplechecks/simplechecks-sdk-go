// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package simplechecksgo

import (
	"context"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/simplechecks/simplechecks-sdk-go/internal/apijson"
	"github.com/simplechecks/simplechecks-sdk-go/internal/apiquery"
	"github.com/simplechecks/simplechecks-sdk-go/internal/param"
	"github.com/simplechecks/simplechecks-sdk-go/internal/requestconfig"
	"github.com/simplechecks/simplechecks-sdk-go/option"
)

// Run-credit balance, Stripe Checkout top-ups, and purchase history.
//
// PurchaseService contains methods and other services that help with interacting
// with the simple-checks API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPurchaseService] method instead.
type PurchaseService struct {
	Options []option.RequestOption
}

// NewPurchaseService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPurchaseService(opts ...option.RequestOption) (r *PurchaseService) {
	r = &PurchaseService{}
	r.Options = opts
	return
}

// Returns every Stripe Checkout bundle purchase for the caller's account, newest
// first. Powers the "Invoices" section of Settings → Billing in the webapp. The
// `receipt_url`, when present, links to the Stripe-hosted receipt PDF. Reading
// purchase history requires only the default-scope `account:read` — spending money
// on a new purchase requires the opt-in `billing:write` scope (POST
// /v1/checkout-session).
func (r *PurchaseService) List(ctx context.Context, query PurchaseListParams, opts ...option.RequestOption) (res *PurchaseListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/purchases"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// One row of the customer's Stripe Checkout bundle purchase history. Tokens are
// credited at fulfillment time; pending and failed rows reflect Checkout sessions
// that did not complete.
type Purchase struct {
	// Server-side purchase id.
	ID string `json:"id" api:"required"`
	// Customer-paid amount in the smallest currency unit (e.g., USD cents).
	AmountCents int64 `json:"amount_cents" api:"required"`
	// Bundle identifier (e.g., `starter`, `growth`, `scale`, `team`).
	BundleSKU string `json:"bundle_sku" api:"required"`
	// When the Checkout session was minted.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// ISO 4217 currency code (e.g., `usd`).
	Currency string         `json:"currency" api:"required"`
	Status   PurchaseStatus `json:"status" api:"required"`
	// Stripe Checkout session that originated this purchase.
	StripeSessionID string `json:"stripe_session_id" api:"required"`
	// Total tokens credited on fulfillment (includes any bonus).
	Tokens int64 `json:"tokens" api:"required"`
	// When the payment landed and tokens were credited; absent for non-fulfilled rows.
	FulfilledAt time.Time `json:"fulfilled_at" format:"date-time"`
	// Stripe-hosted receipt PDF URL. Absent for in-flight purchases and for fulfilled
	// purchases whose payment event did not surface a receipt (e.g., asynchronous
	// payment methods).
	ReceiptURL string       `json:"receipt_url" format:"uri"`
	JSON       purchaseJSON `json:"-"`
}

// purchaseJSON contains the JSON metadata for the struct [Purchase]
type purchaseJSON struct {
	ID              apijson.Field
	AmountCents     apijson.Field
	BundleSKU       apijson.Field
	CreatedAt       apijson.Field
	Currency        apijson.Field
	Status          apijson.Field
	StripeSessionID apijson.Field
	Tokens          apijson.Field
	FulfilledAt     apijson.Field
	ReceiptURL      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *Purchase) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r purchaseJSON) RawJSON() string {
	return r.raw
}

type PurchaseStatus string

const (
	PurchaseStatusPending   PurchaseStatus = "pending"
	PurchaseStatusFulfilled PurchaseStatus = "fulfilled"
	PurchaseStatusFailed    PurchaseStatus = "failed"
)

func (r PurchaseStatus) IsKnown() bool {
	switch r {
	case PurchaseStatusPending, PurchaseStatusFulfilled, PurchaseStatusFailed:
		return true
	}
	return false
}

type PurchaseListResponse struct {
	Purchases []Purchase               `json:"purchases" api:"required"`
	JSON      purchaseListResponseJSON `json:"-"`
}

// purchaseListResponseJSON contains the JSON metadata for the struct
// [PurchaseListResponse]
type purchaseListResponseJSON struct {
	Purchases   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PurchaseListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r purchaseListResponseJSON) RawJSON() string {
	return r.raw
}

type PurchaseListParams struct {
	// Page size. Server applies a default of 100 when omitted or when set to 0; values
	// above the server cap are clamped.
	Limit param.Field[int64] `query:"limit"`
	// Pagination offset within the newest-first list.
	Offset param.Field[int64] `query:"offset"`
}

// URLQuery serializes [PurchaseListParams]'s query parameters as `url.Values`.
func (r PurchaseListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
