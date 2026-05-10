// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package simplechecksgo

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/simplechecks/simplechecks-sdk-go/internal/apijson"
	"github.com/simplechecks/simplechecks-sdk-go/internal/apiquery"
	"github.com/simplechecks/simplechecks-sdk-go/internal/param"
	"github.com/simplechecks/simplechecks-sdk-go/internal/requestconfig"
	"github.com/simplechecks/simplechecks-sdk-go/option"
)

// Read-only incident timeline derived from alert state.
//
// IncidentService contains methods and other services that help with interacting
// with the simple-checks API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewIncidentService] method instead.
type IncidentService struct {
	Options []option.RequestOption
}

// NewIncidentService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewIncidentService(opts ...option.RequestOption) (r *IncidentService) {
	r = &IncidentService{}
	r.Options = opts
	return
}

// Returns incidents derived on read from `alert_state` (ongoing) and
// `alert_dispatches` (resolved). Ordered ongoing-first, then most-recent-resolved
// first. Pagination is offset-based; pass `next_offset` back to continue.
//
// Status semantics:
//
//   - `ongoing` — `alert_state.current_incident_id` is set; `resolved_at_unix_ms` is
//     omitted.
//   - `resolved` — a recovery dispatch has been enqueued; both timestamps are
//     populated.
//
// Incidents that fired entirely inside a maintenance window won't appear here —
// the dispatcher doesn't ledger suppressed dispatches. That matches the customer
// expectation that maintenance windows mean "don't notify, don't surface as
// urgent."
//
// Requires the `checks:read` scope (incidents are per-check; we reuse the existing
// scope rather than minting a new one).
func (r *IncidentService) List(ctx context.Context, query IncidentListParams, opts ...option.RequestOption) (res *IncidentListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/incidents"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// One alert-state lifecycle entry. Derived on read from `alert_state` +
// `alert_dispatches` — there's no separate incidents table because the data is
// fully reconstructable from the rows the evaluator already writes.
type Incident struct {
	// Incident id (UUID; from `alert_state.current_incident_id`).
	ID        string `json:"id" api:"required" format:"uuid"`
	CheckID   string `json:"check_id" api:"required" format:"uuid"`
	CheckName string `json:"check_name" api:"required"`
	// When the evaluator fired the incident (unix-millis).
	StartedAtUnixMs int64          `json:"started_at_unix_ms" api:"required"`
	Status          IncidentStatus `json:"status" api:"required"`
	// Unix-millis of the recovery dispatch. Absent on ongoing incidents.
	ResolvedAtUnixMs int64        `json:"resolved_at_unix_ms" api:"nullable"`
	JSON             incidentJSON `json:"-"`
}

// incidentJSON contains the JSON metadata for the struct [Incident]
type incidentJSON struct {
	ID               apijson.Field
	CheckID          apijson.Field
	CheckName        apijson.Field
	StartedAtUnixMs  apijson.Field
	Status           apijson.Field
	ResolvedAtUnixMs apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *Incident) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r incidentJSON) RawJSON() string {
	return r.raw
}

type IncidentStatus string

const (
	IncidentStatusOngoing  IncidentStatus = "ongoing"
	IncidentStatusResolved IncidentStatus = "resolved"
)

func (r IncidentStatus) IsKnown() bool {
	switch r {
	case IncidentStatusOngoing, IncidentStatusResolved:
		return true
	}
	return false
}

type IncidentListResponse struct {
	Incidents []Incident `json:"incidents" api:"required"`
	// Offset to pass on the next request. Zero (or absent) when there's no more data.
	NextOffset int64                    `json:"next_offset"`
	JSON       incidentListResponseJSON `json:"-"`
}

// incidentListResponseJSON contains the JSON metadata for the struct
// [IncidentListResponse]
type incidentListResponseJSON struct {
	Incidents   apijson.Field
	NextOffset  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IncidentListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r incidentListResponseJSON) RawJSON() string {
	return r.raw
}

type IncidentListParams struct {
	// Max number of incidents to return. Defaults to 50; server caps at 500.
	Limit param.Field[int64] `query:"limit"`
	// Number of incidents to skip. Pass the `next_offset` from the previous page.
	Offset param.Field[int64] `query:"offset"`
}

// URLQuery serializes [IncidentListParams]'s query parameters as `url.Values`.
func (r IncidentListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
