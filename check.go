// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package simplechecks

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/simplechecks/simplechecks-sdk-go/internal/apijson"
	"github.com/simplechecks/simplechecks-sdk-go/internal/apiquery"
	"github.com/simplechecks/simplechecks-sdk-go/internal/requestconfig"
	"github.com/simplechecks/simplechecks-sdk-go/option"
	"github.com/simplechecks/simplechecks-sdk-go/packages/param"
	"github.com/simplechecks/simplechecks-sdk-go/packages/respjson"
)

// CRUD for synthetic-monitoring checks.
//
// CheckService contains methods and other services that help with interacting with
// the simplechecks API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCheckService] method instead.
type CheckService struct {
	options []option.RequestOption
}

// NewCheckService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewCheckService(opts ...option.RequestOption) (r CheckService) {
	r = CheckService{}
	r.options = opts
	return
}

// Creates a check bound to the resolved garrison for the given `provider` +
// `location`. Requires the `checks:write` scope.
func (r *CheckService) New(ctx context.Context, body CheckNewParams, opts ...option.RequestOption) (res *Check, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/checks"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Returns the check with the given id. 404 if no such check exists for the calling
// account. Requires the `checks:read` scope.
func (r *CheckService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *Check, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/checks/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// All fields in the body are optional; omitted fields are left unchanged. Requires
// the `checks:write` scope.
func (r *CheckService) Update(ctx context.Context, id string, body CheckUpdateParams, opts ...option.RequestOption) (res *Check, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/checks/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Returns the caller's checks with simple offset pagination. `next_offset` is set
// when a full page was returned and zero when there's no more data. Requires the
// `checks:read` scope.
func (r *CheckService) List(ctx context.Context, query CheckListParams, opts ...option.RequestOption) (res *CheckListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/checks"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Disables the check. Requires the `checks:write` scope.
func (r *CheckService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("v1/checks/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

type Check struct {
	ID string `json:"id" api:"required"`
	// Owning account's `acct_<typeid>`. Read-only.
	AccountTypeid string    `json:"account_typeid" api:"required"`
	CreatedAt     time.Time `json:"created_at" api:"required" format:"date-time"`
	Enabled       bool      `json:"enabled" api:"required"`
	// Garrison the check is bound to. Server-assigned.
	GarrisonID string `json:"garrison_id" api:"required"`
	Name       string `json:"name" api:"required"`
	// Cron expression; minute granularity.
	Schedule  string `json:"schedule" api:"required"`
	TargetURL string `json:"target_url" api:"required" format:"uri"`
	TimeoutMs int64  `json:"timeout_ms" api:"required"`
	// Check type. Currently only `http` is publicly documented.
	Type      string    `json:"type" api:"required"`
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Optional artifact reference (e.g. uploaded Playwright bundle).
	ArtifactURL string `json:"artifact_url"`
	// Per-check-type configuration blob. Opaque on the wire.
	Config map[string]any `json:"config"`
	// Region/location on read responses is empty; populated on create requests only.
	Location string `json:"location"`
	// Cloud provider on read responses is empty; populated on create requests only.
	Provider string `json:"provider"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		AccountTypeid respjson.Field
		CreatedAt     respjson.Field
		Enabled       respjson.Field
		GarrisonID    respjson.Field
		Name          respjson.Field
		Schedule      respjson.Field
		TargetURL     respjson.Field
		TimeoutMs     respjson.Field
		Type          respjson.Field
		UpdatedAt     respjson.Field
		ArtifactURL   respjson.Field
		Config        respjson.Field
		Location      respjson.Field
		Provider      respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Check) RawJSON() string { return r.JSON.raw }
func (r *Check) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CheckListResponse struct {
	Checks []Check `json:"checks" api:"required"`
	// Offset to pass on the next request to continue pagination. Zero (or absent) when
	// there's no more data.
	NextOffset int64 `json:"next_offset"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Checks      respjson.Field
		NextOffset  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CheckListResponse) RawJSON() string { return r.JSON.raw }
func (r *CheckListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CheckNewParams struct {
	Enabled bool `json:"enabled" api:"required"`
	// Provider-specific region/location.
	Location string `json:"location" api:"required"`
	Name     string `json:"name" api:"required"`
	// Cloud provider (`mock`, `ec2`, `ovh`, `azure`, `gcp`, `hetzner`).
	Provider    string            `json:"provider" api:"required"`
	Schedule    string            `json:"schedule" api:"required"`
	TargetURL   string            `json:"target_url" api:"required" format:"uri"`
	Type        string            `json:"type" api:"required"`
	ArtifactURL param.Opt[string] `json:"artifact_url,omitzero"`
	TimeoutMs   param.Opt[int64]  `json:"timeout_ms,omitzero"`
	Config      map[string]any    `json:"config,omitzero"`
	paramObj
}

func (r CheckNewParams) MarshalJSON() (data []byte, err error) {
	type shadow CheckNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CheckNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CheckUpdateParams struct {
	ArtifactURL param.Opt[string] `json:"artifact_url,omitzero"`
	Enabled     param.Opt[bool]   `json:"enabled,omitzero"`
	Name        param.Opt[string] `json:"name,omitzero"`
	Schedule    param.Opt[string] `json:"schedule,omitzero"`
	TargetURL   param.Opt[string] `json:"target_url,omitzero" format:"uri"`
	TimeoutMs   param.Opt[int64]  `json:"timeout_ms,omitzero"`
	Type        param.Opt[string] `json:"type,omitzero"`
	Config      map[string]any    `json:"config,omitzero"`
	paramObj
}

func (r CheckUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow CheckUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CheckUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CheckListParams struct {
	// Max number of checks to return. Defaults to 100; the server caps further.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Number of checks to skip. Pass the `next_offset` from the previous page.
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CheckListParams]'s query parameters as `url.Values`.
func (r CheckListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
