// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package simplechecksgo

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/simplechecks/simplechecks-sdk-go/internal/apijson"
	"github.com/simplechecks/simplechecks-sdk-go/internal/apiquery"
	"github.com/simplechecks/simplechecks-sdk-go/internal/param"
	"github.com/simplechecks/simplechecks-sdk-go/internal/requestconfig"
	"github.com/simplechecks/simplechecks-sdk-go/option"
)

// Read-only access to past check executions.
//
// RunService contains methods and other services that help with interacting with
// the simple-checks API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRunService] method instead.
type RunService struct {
	Options []option.RequestOption
}

// NewRunService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewRunService(opts ...option.RequestOption) (r *RunService) {
	r = &RunService{}
	r.Options = opts
	return
}

// Returns the run matching `id`. The id's embedded UUIDv7 timestamp scopes the
// server-side scan to one day. Requires the `runs:read` scope.
func (r *RunService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *Run, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/runs/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Returns runs ordered by start time descending. Filter with `check_id`, `status`,
// `since` (unix-millis lower bound). `limit` defaults to 50 (max 200); `offset`
// paginates within the filtered set. Requires the `runs:read` scope.
//
// Run records come from the parquet result files garrisons write to S3; this
// endpoint scans up to the last 7 days by default. Older runs are not retained.
func (r *RunService) List(ctx context.Context, query RunListParams, opts ...option.RequestOption) (res *RunListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/runs"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// A single check execution. Runs are written by the garrison that executed the
// check; CC reads them from S3-resident parquet files for read-only public
// exposure here.
type Run struct {
	// Run typeid (`run_<26-char base32 UUIDv7>`).
	ID string `json:"id" api:"required"`
	// UUID of the parent check (matches `Check.id`).
	CheckID    string `json:"check_id" api:"required"`
	CheckName  string `json:"check_name" api:"required"`
	DurationMs int64  `json:"duration_ms" api:"required"`
	GarrisonID string `json:"garrison_id" api:"required"`
	InstanceID string `json:"instance_id" api:"required"`
	NodeName   string `json:"node_name" api:"required"`
	// Execution start time in unix milliseconds (UTC).
	StartedAtUnixMs int64     `json:"started_at_unix_ms" api:"required"`
	Status          RunStatus `json:"status" api:"required"`
	// Check type (`http`, `tcp`, `dns`, ...).
	Type         string `json:"type" api:"required"`
	ErrorMessage string `json:"error_message"`
	// Per-check-type metadata blob, JSON-encoded as a string.
	Metadata string  `json:"metadata"`
	JSON     runJSON `json:"-"`
}

// runJSON contains the JSON metadata for the struct [Run]
type runJSON struct {
	ID              apijson.Field
	CheckID         apijson.Field
	CheckName       apijson.Field
	DurationMs      apijson.Field
	GarrisonID      apijson.Field
	InstanceID      apijson.Field
	NodeName        apijson.Field
	StartedAtUnixMs apijson.Field
	Status          apijson.Field
	Type            apijson.Field
	ErrorMessage    apijson.Field
	Metadata        apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *Run) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runJSON) RawJSON() string {
	return r.raw
}

type RunStatus string

const (
	RunStatusPass    RunStatus = "PASS"
	RunStatusFail    RunStatus = "FAIL"
	RunStatusError   RunStatus = "ERROR"
	RunStatusTimeout RunStatus = "TIMEOUT"
)

func (r RunStatus) IsKnown() bool {
	switch r {
	case RunStatusPass, RunStatusFail, RunStatusError, RunStatusTimeout:
		return true
	}
	return false
}

type RunListResponse struct {
	Runs []Run               `json:"runs" api:"required"`
	JSON runListResponseJSON `json:"-"`
}

// runListResponseJSON contains the JSON metadata for the struct [RunListResponse]
type runListResponseJSON struct {
	Runs        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runListResponseJSON) RawJSON() string {
	return r.raw
}

type RunListParams struct {
	// Filter to a single check (UUID; matches `Check.id`).
	CheckID param.Field[string] `query:"check_id"`
	Limit   param.Field[int64]  `query:"limit"`
	Offset  param.Field[int64]  `query:"offset"`
	// Lower bound on `started_at_unix_ms`. Server clamps to a 7-day window.
	Since param.Field[int64] `query:"since"`
	// Filter to a single execution status.
	Status param.Field[RunListParamsStatus] `query:"status"`
}

// URLQuery serializes [RunListParams]'s query parameters as `url.Values`.
func (r RunListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter to a single execution status.
type RunListParamsStatus string

const (
	RunListParamsStatusPass    RunListParamsStatus = "PASS"
	RunListParamsStatusFail    RunListParamsStatus = "FAIL"
	RunListParamsStatusError   RunListParamsStatus = "ERROR"
	RunListParamsStatusTimeout RunListParamsStatus = "TIMEOUT"
)

func (r RunListParamsStatus) IsKnown() bool {
	switch r {
	case RunListParamsStatusPass, RunListParamsStatusFail, RunListParamsStatusError, RunListParamsStatusTimeout:
		return true
	}
	return false
}
