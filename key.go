// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package simplechecksgo

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/simplechecks/simplechecks-sdk-go/internal/apijson"
	"github.com/simplechecks/simplechecks-sdk-go/internal/param"
	"github.com/simplechecks/simplechecks-sdk-go/internal/requestconfig"
	"github.com/simplechecks/simplechecks-sdk-go/option"
)

// Manage personal access tokens (PATs).
//
// KeyService contains methods and other services that help with interacting with
// the simple-checks API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewKeyService] method instead.
type KeyService struct {
	Options []option.RequestOption
}

// NewKeyService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewKeyService(opts ...option.RequestOption) (r *KeyService) {
	r = &KeyService{}
	r.Options = opts
	return
}

// Mints a fresh PAT for the caller's account. The plaintext token is returned
// **once**; clients must persist it before discarding the response. Empty `scopes`
// means the server applies its default scope set. Requires the `keys:write` scope.
func (r *KeyService) New(ctx context.Context, body KeyNewParams, opts ...option.RequestOption) (res *KeyNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/keys"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Returns every API key for the caller's account, including revoked ones. The
// plaintext token is never returned by this endpoint — only by POST /v1/keys at
// mint time. Requires the `keys:read` scope.
func (r *KeyService) List(ctx context.Context, opts ...option.RequestOption) (res *KeyListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/keys"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Marks the key revoked. Subsequent ext_authz checks reject requests authenticated
// with this key. The row stays for audit. Requires the `keys:write` scope.
func (r *KeyService) Revoke(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("v1/keys/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// An account-scoped personal access token (PAT). The plaintext token never appears
// here; it's only returned by POST /v1/keys at mint time.
type APIKey struct {
	// Server-side key id (used for revoke).
	ID            string    `json:"id" api:"required"`
	AccountTypeid string    `json:"account_typeid" api:"required"`
	CreatedAt     time.Time `json:"created_at" api:"required" format:"date-time"`
	// Operator/customer-facing label.
	Name string `json:"name" api:"required"`
	// Logging-safe visible portion (e.g. `sc_live_xxx`).
	Prefix     string     `json:"prefix" api:"required"`
	Scopes     []string   `json:"scopes" api:"required"`
	LastUsedAt time.Time  `json:"last_used_at" format:"date-time"`
	RevokedAt  time.Time  `json:"revoked_at" format:"date-time"`
	JSON       apiKeyJSON `json:"-"`
}

// apiKeyJSON contains the JSON metadata for the struct [APIKey]
type apiKeyJSON struct {
	ID            apijson.Field
	AccountTypeid apijson.Field
	CreatedAt     apijson.Field
	Name          apijson.Field
	Prefix        apijson.Field
	Scopes        apijson.Field
	LastUsedAt    apijson.Field
	RevokedAt     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *APIKey) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiKeyJSON) RawJSON() string {
	return r.raw
}

type KeyNewResponse struct {
	KeyID string `json:"key_id" api:"required"`
	// Full `sc_live_…` token. Returned once; not retrievable later. Clients MUST
	// persist this before discarding the response.
	PlaintextToken string             `json:"plaintext_token" api:"required"`
	Prefix         string             `json:"prefix" api:"required"`
	JSON           keyNewResponseJSON `json:"-"`
}

// keyNewResponseJSON contains the JSON metadata for the struct [KeyNewResponse]
type keyNewResponseJSON struct {
	KeyID          apijson.Field
	PlaintextToken apijson.Field
	Prefix         apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *KeyNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r keyNewResponseJSON) RawJSON() string {
	return r.raw
}

type KeyListResponse struct {
	Keys []APIKey            `json:"keys" api:"required"`
	JSON keyListResponseJSON `json:"-"`
}

// keyListResponseJSON contains the JSON metadata for the struct [KeyListResponse]
type keyListResponseJSON struct {
	Keys        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *KeyListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r keyListResponseJSON) RawJSON() string {
	return r.raw
}

type KeyNewParams struct {
	// Operator/customer-facing label.
	Name param.Field[string] `json:"name" api:"required"`
	// Scope strings (e.g. `checks:read`). Empty = server applies its default set.
	// Unknown scopes return InvalidArgument.
	Scopes param.Field[[]string] `json:"scopes"`
}

func (r KeyNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
