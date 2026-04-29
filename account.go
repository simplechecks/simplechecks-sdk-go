// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package simplechecks

import (
	"context"
	"net/http"
	"slices"
	"time"

	"github.com/simplechecks/sdk-go/internal/apijson"
	"github.com/simplechecks/sdk-go/internal/requestconfig"
	"github.com/simplechecks/sdk-go/option"
	"github.com/simplechecks/sdk-go/packages/respjson"
)

// Account profile and balance.
//
// AccountService contains methods and other services that help with interacting
// with the simplechecks API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountService] method instead.
type AccountService struct {
	options []option.RequestOption
}

// NewAccountService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAccountService(opts ...option.RequestOption) (r AccountService) {
	r = AccountService{}
	r.options = opts
	return
}

// Returns the account row stitched together with the cached billing balance and
// the `paused` flag, so a single dashboard read fetches everything the customer's
// home page needs. Requires the `account:read` scope.
func (r *AccountService) Get(ctx context.Context, opts ...option.RequestOption) (res *AccountGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/account"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Account profile + cached billing balance. Returned by GET /v1/account.
type AccountGetResponse struct {
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
	Typeid string `json:"typeid" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Balance     respjson.Field
		CreatedAt   respjson.Field
		Name        respjson.Field
		Paused      respjson.Field
		Plan        respjson.Field
		Slug        respjson.Field
		Typeid      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccountGetResponse) RawJSON() string { return r.JSON.raw }
func (r *AccountGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
