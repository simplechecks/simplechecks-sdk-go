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
	"github.com/simplechecks/simplechecks-sdk-go/packages/pagination"
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

// Returns the full record for the run matching `id` — the slim list fields plus
// the run's `metadata` (a JSON object) and a list of downloadable `artifacts`
// (each an opaque URL). Runs are retained for 30 days; an aged-out or unknown id
// returns 404. Requires the `runs:read` scope.
func (r *RunService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *RunDetail, err error) {
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
// `location`, and a `since`/`until` unix-millis window. `limit` defaults to 50
// (max 200). Pages are cursor-based: when more rows remain, the response carries a
// `next_cursor` — pass it back as `cursor` to fetch the next page. Requires the
// `runs:read` scope.
//
// Run records are served from the central runs table; runs are retained for 30
// days. Each record carries structured `provider`/`region`/`location` fields and a
// short `error_summary` rather than infrastructure internals.
func (r *RunService) List(ctx context.Context, query RunListParams, opts ...option.RequestOption) (res *pagination.RunsCursor[RunListItem], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/runs"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Returns runs ordered by start time descending. Filter with `check_id`, `status`,
// `location`, and a `since`/`until` unix-millis window. `limit` defaults to 50
// (max 200). Pages are cursor-based: when more rows remain, the response carries a
// `next_cursor` — pass it back as `cursor` to fetch the next page. Requires the
// `runs:read` scope.
//
// Run records are served from the central runs table; runs are retained for 30
// days. Each record carries structured `provider`/`region`/`location` fields and a
// short `error_summary` rather than infrastructure internals.
func (r *RunService) ListAutoPaging(ctx context.Context, query RunListParams, opts ...option.RequestOption) *pagination.RunsCursorAutoPager[RunListItem] {
	return pagination.NewRunsCursorAutoPager(r.List(ctx, query, opts...))
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

// The full record for one check execution: the list fields plus the run's
// `metadata` (a JSON object) and the set of downloadable `artifacts`. Location is
// structured; no infrastructure identifiers are exposed.
type RunDetail struct {
	// Run typeid (`run_<26-char base32 UUIDv7>`).
	ID string `json:"id" api:"required"`
	// Downloadable artifacts for this run (empty when none).
	Artifacts []RunDetailArtifact `json:"artifacts" api:"required"`
	// UUID of the parent check (matches `Check.id`).
	CheckID     string `json:"check_id" api:"required"`
	CheckName   string `json:"check_name" api:"required"`
	DurationMs  int64  `json:"duration_ms" api:"required"`
	HasErrors   bool   `json:"has_errors" api:"required"`
	HasFailures bool   `json:"has_failures" api:"required"`
	// Execution start time in unix milliseconds (UTC).
	StartedAtUnixMs int64           `json:"started_at_unix_ms" api:"required"`
	Status          RunDetailStatus `json:"status" api:"required"`
	// Check type (`http`, `tcp`, `dns`, ...).
	Type string `json:"type" api:"required"`
	// Reserved; always null at this version.
	Degraded bool `json:"degraded" api:"nullable"`
	// Full failure message; null on a passing run.
	ErrorMessage string `json:"error_message" api:"nullable"`
	// Human-readable location label. Null when unresolved.
	Location string `json:"location" api:"nullable"`
	// Per-check-type metadata as a JSON object; null when absent.
	Metadata map[string]interface{} `json:"metadata" api:"nullable"`
	// Cloud provider that ran the check. Null when unresolved.
	Provider string `json:"provider" api:"nullable"`
	// Provider-native region id. Null when unresolved.
	Region string        `json:"region" api:"nullable"`
	JSON   runDetailJSON `json:"-"`
}

// runDetailJSON contains the JSON metadata for the struct [RunDetail]
type runDetailJSON struct {
	ID              apijson.Field
	Artifacts       apijson.Field
	CheckID         apijson.Field
	CheckName       apijson.Field
	DurationMs      apijson.Field
	HasErrors       apijson.Field
	HasFailures     apijson.Field
	StartedAtUnixMs apijson.Field
	Status          apijson.Field
	Type            apijson.Field
	Degraded        apijson.Field
	ErrorMessage    apijson.Field
	Location        apijson.Field
	Metadata        apijson.Field
	Provider        apijson.Field
	Region          apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *RunDetail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runDetailJSON) RawJSON() string {
	return r.raw
}

// One downloadable artifact for a run.
type RunDetailArtifact struct {
	// Artifact kind (closed set).
	Kind RunDetailArtifactsKind `json:"kind" api:"required"`
	// Opaque, webapp-relative download path (`/v1/runs/{id}/artifacts/{kind}`).
	URL  string                `json:"url" api:"required"`
	JSON runDetailArtifactJSON `json:"-"`
}

// runDetailArtifactJSON contains the JSON metadata for the struct
// [RunDetailArtifact]
type runDetailArtifactJSON struct {
	Kind        apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunDetailArtifact) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runDetailArtifactJSON) RawJSON() string {
	return r.raw
}

// Artifact kind (closed set).
type RunDetailArtifactsKind string

const (
	RunDetailArtifactsKindScreenshot RunDetailArtifactsKind = "screenshot"
	RunDetailArtifactsKindTrace      RunDetailArtifactsKind = "trace"
	RunDetailArtifactsKindHar        RunDetailArtifactsKind = "har"
)

func (r RunDetailArtifactsKind) IsKnown() bool {
	switch r {
	case RunDetailArtifactsKindScreenshot, RunDetailArtifactsKindTrace, RunDetailArtifactsKindHar:
		return true
	}
	return false
}

type RunDetailStatus string

const (
	RunDetailStatusPass    RunDetailStatus = "PASS"
	RunDetailStatusFail    RunDetailStatus = "FAIL"
	RunDetailStatusError   RunDetailStatus = "ERROR"
	RunDetailStatusTimeout RunDetailStatus = "TIMEOUT"
)

func (r RunDetailStatus) IsKnown() bool {
	switch r {
	case RunDetailStatusPass, RunDetailStatusFail, RunDetailStatusError, RunDetailStatusTimeout:
		return true
	}
	return false
}

// One check execution in a list. Location is exposed as structured
// `provider`/`region`/`location` fields rather than infrastructure internals; the
// row carries cheap boolean flags and a short `error_summary` (null on a passing
// run). For the full record (metadata + downloadable artifacts), fetch
// `GET /v1/runs/{id}`.
type RunListItem struct {
	// Run typeid (`run_<26-char base32 UUIDv7>`).
	ID string `json:"id" api:"required"`
	// UUID of the parent check (matches `Check.id`).
	CheckID     string `json:"check_id" api:"required"`
	CheckName   string `json:"check_name" api:"required"`
	DurationMs  int64  `json:"duration_ms" api:"required"`
	HasErrors   bool   `json:"has_errors" api:"required"`
	HasFailures bool   `json:"has_failures" api:"required"`
	// Execution start time in unix milliseconds (UTC).
	StartedAtUnixMs int64             `json:"started_at_unix_ms" api:"required"`
	Status          RunListItemStatus `json:"status" api:"required"`
	// Check type (`http`, `tcp`, `dns`, ...).
	Type string `json:"type" api:"required"`
	// Reserved; always null at this version.
	Degraded bool `json:"degraded" api:"nullable"`
	// Short failure summary; null on a passing run.
	ErrorSummary string `json:"error_summary" api:"nullable"`
	// Human-readable location label (e.g. `Falkenstein, DE`). Null when unresolved.
	Location string `json:"location" api:"nullable"`
	// Cloud provider that ran the check (e.g. `hetzner`, `ovh`). Null when unresolved.
	Provider string `json:"provider" api:"nullable"`
	// Provider-native region id (e.g. `fsn1`, `gra7`). Null when unresolved.
	Region string          `json:"region" api:"nullable"`
	JSON   runListItemJSON `json:"-"`
}

// runListItemJSON contains the JSON metadata for the struct [RunListItem]
type runListItemJSON struct {
	ID              apijson.Field
	CheckID         apijson.Field
	CheckName       apijson.Field
	DurationMs      apijson.Field
	HasErrors       apijson.Field
	HasFailures     apijson.Field
	StartedAtUnixMs apijson.Field
	Status          apijson.Field
	Type            apijson.Field
	Degraded        apijson.Field
	ErrorSummary    apijson.Field
	Location        apijson.Field
	Provider        apijson.Field
	Region          apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *RunListItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runListItemJSON) RawJSON() string {
	return r.raw
}

type RunListItemStatus string

const (
	RunListItemStatusPass    RunListItemStatus = "PASS"
	RunListItemStatusFail    RunListItemStatus = "FAIL"
	RunListItemStatusError   RunListItemStatus = "ERROR"
	RunListItemStatusTimeout RunListItemStatus = "TIMEOUT"
)

func (r RunListItemStatus) IsKnown() bool {
	switch r {
	case RunListItemStatusPass, RunListItemStatusFail, RunListItemStatusError, RunListItemStatusTimeout:
		return true
	}
	return false
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
	// Opaque pagination token from the previous page's `next_cursor`.
	Cursor param.Field[string] `query:"cursor"`
	// Page size; defaults to 50, max 200.
	Limit param.Field[int64] `query:"limit"`
	// Filter to a single provider-native region id (e.g. `fsn1`).
	Location param.Field[string] `query:"location"`
	// Lower bound on `started_at_unix_ms` (inclusive).
	Since param.Field[int64] `query:"since"`
	// Filter to a single execution status.
	Status param.Field[RunListParamsStatus] `query:"status"`
	// Upper bound on `started_at_unix_ms` (inclusive).
	Until param.Field[int64] `query:"until"`
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
