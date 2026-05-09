// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package simplechecksgo

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"slices"

	"github.com/simplechecks/simplechecks-sdk-go/internal/apijson"
	"github.com/simplechecks/simplechecks-sdk-go/internal/apiquery"
	"github.com/simplechecks/simplechecks-sdk-go/internal/param"
	"github.com/simplechecks/simplechecks-sdk-go/internal/requestconfig"
	"github.com/simplechecks/simplechecks-sdk-go/option"
	"github.com/simplechecks/simplechecks-sdk-go/packages/jsonl"
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

// Returns per-(check, location, minute-bucket) aggregate rows for the calling
// account, optionally filtered by check_id, location, and time range. Powers the
// customer dashboard ("uptime %", "pass rate", "average latency over period") and
// the public status page; you wouldn't typically render per-run rows from this
// endpoint at typical zoom levels.
//
// **Resolution.** Buckets are minute-aligned to UTC; the only accepted `bucket`
// value at MVP is `minute`. The param exists so future per-15s or per-hour rollups
// can slot in additively.
//
// **Eventual-consistency contract.** A bucket may continue to receive
// contributions after `now()` crosses its end boundary — late-arriving Garrison
// batches (network blip, scaling) feed the bucket they truncate to, which can be
// in the past. Treat any returned counts as a lower bound; dashboards refreshing
// the same window may see counts increase. The push cadence (15s) bounds how stale
// the aggregate is in steady state.
//
// **Latency stats.** `duration_avg_ms` is computed server-side from the underlying
// sum/count. `duration_min_ms` and `duration_max_ms` reflect the extremes seen in
// the bucket. Percentiles (p50/p95/p99) require online-mergeable sketches and are
// deferred to a follow-up.
//
// Requires the `runs:read` scope.
func (r *RunService) Aggregates(ctx context.Context, query RunAggregatesParams, opts ...option.RequestOption) (res *RunAggregatesResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/runs/aggregates"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns the gzipped JSONL execution log for the run. Each line is a structured
// event (timestamp, severity, source, message, optional kv map) tagged with the
// run id; producers include the executor (`run_start`, `run_end`) and per-type
// emitters (currently `http_request`, `http_response` for HTTP checks).
//
// The response always carries `Content-Encoding: gzip` and the bytes on the wire
// are the gzipped form; standards-compliant HTTP clients (browsers, curl,
// Go/Python/JS SDKs) decompress transparently. `sc logs <id>` (PR-Logs/2) consumes
// this endpoint.
//
// Tenancy is enforced before any byte fetch — a run id that doesn't belong to the
// calling account returns 404, not 403, so callers can't probe for the existence
// of other tenants' runs. Requires the `runs:read` scope.
func (r *RunService) LogsStreaming(ctx context.Context, id string, opts ...option.RequestOption) (stream *jsonl.Stream[io.Reader]) {
	var (
		raw *http.Response
		err error
	)
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/x-ndjson")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return jsonl.NewStream[io.Reader](nil, err)
	}
	path := fmt.Sprintf("v1/runs/%s/logs", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &raw, opts...)
	return jsonl.NewStream[io.Reader](raw, err)
}

type Aggregate struct {
	AccountID string `json:"account_id" api:"required" format:"uuid"`
	// Exclusive bucket end (= start + 60 000 ms today).
	BucketEndUnixMs int64 `json:"bucket_end_unix_ms" api:"required"`
	// Inclusive bucket start, unix-millis, minute-aligned to UTC.
	BucketStartUnixMs int64  `json:"bucket_start_unix_ms" api:"required"`
	CheckID           string `json:"check_id" api:"required" format:"uuid"`
	// Server-computed average from sum/count. Zero when the bucket has no runs.
	DurationAvgMs int64 `json:"duration_avg_ms" api:"required"`
	ErrorCount    int64 `json:"error_count" api:"required"`
	FailCount     int64 `json:"fail_count" api:"required"`
	// Garrison cloud / region label (e.g. `hetzner`, `ovh`, `aws`).
	Location     string `json:"location" api:"required"`
	PassCount    int64  `json:"pass_count" api:"required"`
	TimeoutCount int64  `json:"timeout_count" api:"required"`
	// Sum of all four status counts. Convenience for clients that compute uptime as
	// `pass_count / total_count`.
	TotalCount    int64         `json:"total_count" api:"required"`
	DurationMaxMs int64         `json:"duration_max_ms" api:"nullable"`
	DurationMinMs int64         `json:"duration_min_ms" api:"nullable"`
	JSON          aggregateJSON `json:"-"`
}

// aggregateJSON contains the JSON metadata for the struct [Aggregate]
type aggregateJSON struct {
	AccountID         apijson.Field
	BucketEndUnixMs   apijson.Field
	BucketStartUnixMs apijson.Field
	CheckID           apijson.Field
	DurationAvgMs     apijson.Field
	ErrorCount        apijson.Field
	FailCount         apijson.Field
	Location          apijson.Field
	PassCount         apijson.Field
	TimeoutCount      apijson.Field
	TotalCount        apijson.Field
	DurationMaxMs     apijson.Field
	DurationMinMs     apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *Aggregate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aggregateJSON) RawJSON() string {
	return r.raw
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

type RunAggregatesResponse struct {
	Aggregates []Aggregate               `json:"aggregates" api:"required"`
	JSON       runAggregatesResponseJSON `json:"-"`
}

// runAggregatesResponseJSON contains the JSON metadata for the struct
// [RunAggregatesResponse]
type runAggregatesResponseJSON struct {
	Aggregates  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunAggregatesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runAggregatesResponseJSON) RawJSON() string {
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

type RunAggregatesParams struct {
	// Bucket size. Only `minute` accepted today.
	Bucket param.Field[RunAggregatesParamsBucket] `query:"bucket"`
	// Filter to one check.
	CheckID param.Field[string] `query:"check_id" format:"uuid"`
	// Inclusive lower bound, unix-millis. Defaults to `now() - 1h`.
	From param.Field[int64] `query:"from"`
	// Maximum number of rows. Default 1000; hard cap 5000.
	Limit param.Field[int64] `query:"limit"`
	// Filter to one location (e.g. `hetzner`, `ovh`).
	Location param.Field[string] `query:"location"`
	// Exclusive upper bound, unix-millis. Defaults to `now() + 1m`.
	To param.Field[int64] `query:"to"`
}

// URLQuery serializes [RunAggregatesParams]'s query parameters as `url.Values`.
func (r RunAggregatesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Bucket size. Only `minute` accepted today.
type RunAggregatesParamsBucket string

const (
	RunAggregatesParamsBucketMinute RunAggregatesParamsBucket = "minute"
)

func (r RunAggregatesParamsBucket) IsKnown() bool {
	switch r {
	case RunAggregatesParamsBucketMinute:
		return true
	}
	return false
}
