// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package simplechecksgo

import (
	"context"
	"net/http"
	"slices"
	"time"

	"github.com/simplechecks/simplechecks-sdk-go/internal/apijson"
	"github.com/simplechecks/simplechecks-sdk-go/internal/requestconfig"
	"github.com/simplechecks/simplechecks-sdk-go/option"
)

// Account profile and balance.
//
// AccountService contains methods and other services that help with interacting
// with the simple-checks API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountService] method instead.
type AccountService struct {
	Options []option.RequestOption
}

// NewAccountService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAccountService(opts ...option.RequestOption) (r *AccountService) {
	r = &AccountService{}
	r.Options = opts
	return
}

// Returns the account row stitched together with the cached billing balance and
// the `paused` flag, so a single dashboard read fetches everything the customer's
// home page needs. Requires the `account:read` scope.
func (r *AccountService) Get(ctx context.Context, opts ...option.RequestOption) (res *Account, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/account"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Account profile + cached billing balance. Returned by GET /v1/account.
type Account struct {
	// Cached run-credit balance, in run-credit units.
	Balance   int64     `json:"balance" api:"required"`
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	Name      string    `json:"name" api:"required"`
	// True when execution is paused (e.g. balance exhausted).
	Paused bool `json:"paused" api:"required"`
	// Billing plan identifier.
	Plan string `json:"plan" api:"required"`
	// Renameable URL-friendly handle. Display only — never use as a system identifier.
	Slug string `json:"slug" api:"required"`
	// Stable account identifier (`acct_<typeid>`). Used in API responses and audit
	// logs.
	Typeid string      `json:"typeid" api:"required"`
	JSON   accountJSON `json:"-"`
}

// accountJSON contains the JSON metadata for the struct [Account]
type accountJSON struct {
	Balance     apijson.Field
	CreatedAt   apijson.Field
	Name        apijson.Field
	Paused      apijson.Field
	Plan        apijson.Field
	Slug        apijson.Field
	Typeid      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Account) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountJSON) RawJSON() string {
	return r.raw
}
