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

// Run-credit balance + Stripe Checkout for top-ups.
//
// BalanceService contains methods and other services that help with interacting
// with the simple-checks API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBalanceService] method instead.
type BalanceService struct {
	Options []option.RequestOption
}

// NewBalanceService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewBalanceService(opts ...option.RequestOption) (r *BalanceService) {
	r = &BalanceService{}
	r.Options = opts
	return
}

// Thin sibling of GET /v1/account that returns just the balance and paused flag.
// The CLI's `sc balance` command pulls this so it doesn't have to fetch the full
// account row each time. Requires the `account:read` scope.
func (r *BalanceService) Get(ctx context.Context, opts ...option.RequestOption) (res *Balance, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/balance"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type Balance struct {
	// Cached run-credit balance.
	Balance int64 `json:"balance" api:"required"`
	// True when execution is paused (e.g. balance exhausted).
	Paused bool        `json:"paused" api:"required"`
	JSON   balanceJSON `json:"-"`
}

// balanceJSON contains the JSON metadata for the struct [Balance]
type balanceJSON struct {
	Balance     apijson.Field
	Paused      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Balance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r balanceJSON) RawJSON() string {
	return r.raw
}
