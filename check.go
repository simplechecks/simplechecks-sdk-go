// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package simplechecksgo

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
	"github.com/simplechecks/simplechecks-sdk-go/internal/param"
	"github.com/simplechecks/simplechecks-sdk-go/internal/requestconfig"
	"github.com/simplechecks/simplechecks-sdk-go/option"
)

// CRUD for synthetic-monitoring checks.
//
// CheckService contains methods and other services that help with interacting with
// the simple-checks API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCheckService] method instead.
type CheckService struct {
	Options []option.RequestOption
	// Per-check alert configuration + test-fire endpoint (PR-Alerts/1).
	Alerts *CheckAlertService
}

// NewCheckService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewCheckService(opts ...option.RequestOption) (r *CheckService) {
	r = &CheckService{}
	r.Options = opts
	r.Alerts = NewCheckAlertService(opts...)
	return
}

// Creates a check bound to the resolved garrison for the given `provider` +
// `location`. Requires the `checks:write` scope.
func (r *CheckService) New(ctx context.Context, body CheckNewParams, opts ...option.RequestOption) (res *Check, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/checks"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Returns the check with the given id. 404 if no such check exists for the calling
// account. Requires the `checks:read` scope.
func (r *CheckService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *Check, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/checks/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// All fields in the body are optional; omitted fields are left unchanged. Requires
// the `checks:write` scope.
func (r *CheckService) Update(ctx context.Context, id string, body CheckUpdateParams, opts ...option.RequestOption) (res *Check, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/checks/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Returns the caller's checks with simple offset pagination. `next_offset` is set
// when a full page was returned and zero when there's no more data. Requires the
// `checks:read` scope.
func (r *CheckService) List(ctx context.Context, query CheckListParams, opts ...option.RequestOption) (res *CheckListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/checks"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Disables the check. Requires the `checks:write` scope.
func (r *CheckService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("v1/checks/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

type AlertChannel struct {
	// Channel-specific destination. URL for the webhook flavors
	// (slack/discord/teams/webhook), email address for `email`, integration key for
	// `pagerduty`, API key for `opsgenie`.
	Target string           `json:"target" api:"required"`
	Type   AlertChannelType `json:"type" api:"required"`
	// Type-specific options. Optional.
	Config map[string]interface{} `json:"config"`
	JSON   alertChannelJSON       `json:"-"`
}

// alertChannelJSON contains the JSON metadata for the struct [AlertChannel]
type alertChannelJSON struct {
	Target      apijson.Field
	Type        apijson.Field
	Config      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AlertChannel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r alertChannelJSON) RawJSON() string {
	return r.raw
}

type AlertChannelType string

const (
	AlertChannelTypeEmail     AlertChannelType = "email"
	AlertChannelTypeSlack     AlertChannelType = "slack"
	AlertChannelTypeDiscord   AlertChannelType = "discord"
	AlertChannelTypeTeams     AlertChannelType = "teams"
	AlertChannelTypeWebhook   AlertChannelType = "webhook"
	AlertChannelTypePagerduty AlertChannelType = "pagerduty"
	AlertChannelTypeOpsgenie  AlertChannelType = "opsgenie"
)

func (r AlertChannelType) IsKnown() bool {
	switch r {
	case AlertChannelTypeEmail, AlertChannelTypeSlack, AlertChannelTypeDiscord, AlertChannelTypeTeams, AlertChannelTypeWebhook, AlertChannelTypePagerduty, AlertChannelTypeOpsgenie:
		return true
	}
	return false
}

type AlertChannelParam struct {
	// Channel-specific destination. URL for the webhook flavors
	// (slack/discord/teams/webhook), email address for `email`, integration key for
	// `pagerduty`, API key for `opsgenie`.
	Target param.Field[string]           `json:"target" api:"required"`
	Type   param.Field[AlertChannelType] `json:"type" api:"required"`
	// Type-specific options. Optional.
	Config param.Field[map[string]interface{}] `json:"config"`
}

func (r AlertChannelParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AlertConfig struct {
	Channels []AlertChannel `json:"channels" api:"required"`
	// Number of consecutive globally-failing observations (after M-of-N consensus
	// collapses per-location status) required before an incident fires. Default = 1 =
	// "alert on first globally-failing observation."
	ConsecutiveFailuresThreshold int64 `json:"consecutive_failures_threshold" api:"required"`
	// M-of-N consensus rule denominator (expected total location count). When fewer
	// than `consensus_m` locations have observations, the evaluator falls back to "any
	// failing = failing" so brand-new checks don't miss outages.
	ConsensusM int64 `json:"consensus_m" api:"required"`
	// M-of-N consensus rule numerator. The evaluator considers the check
	// globally-failing only when at least this many locations are reporting fail
	// concurrently.
	ConsensusN int64 `json:"consensus_n" api:"required"`
	// When false, the evaluator skips this check entirely.
	Enabled bool `json:"enabled" api:"required"`
	// Server-set; ignored on write.
	AccountID string `json:"account_id" format:"uuid"`
	// Server-set; ignored on write.
	CheckID   string    `json:"check_id" format:"uuid"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Absolute-time windows during which the evaluator suppresses dispatch but still
	// updates state. Cron-style recurring windows are a future enhancement.
	MaintenanceWindows []MaintenanceWindow `json:"maintenance_windows"`
	UpdatedAt          time.Time           `json:"updated_at" format:"date-time"`
	JSON               alertConfigJSON     `json:"-"`
}

// alertConfigJSON contains the JSON metadata for the struct [AlertConfig]
type alertConfigJSON struct {
	Channels                     apijson.Field
	ConsecutiveFailuresThreshold apijson.Field
	ConsensusM                   apijson.Field
	ConsensusN                   apijson.Field
	Enabled                      apijson.Field
	AccountID                    apijson.Field
	CheckID                      apijson.Field
	CreatedAt                    apijson.Field
	MaintenanceWindows           apijson.Field
	UpdatedAt                    apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *AlertConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r alertConfigJSON) RawJSON() string {
	return r.raw
}

type AlertConfigParam struct {
	Channels param.Field[[]AlertChannelParam] `json:"channels" api:"required"`
	// Number of consecutive globally-failing observations (after M-of-N consensus
	// collapses per-location status) required before an incident fires. Default = 1 =
	// "alert on first globally-failing observation."
	ConsecutiveFailuresThreshold param.Field[int64] `json:"consecutive_failures_threshold" api:"required"`
	// M-of-N consensus rule denominator (expected total location count). When fewer
	// than `consensus_m` locations have observations, the evaluator falls back to "any
	// failing = failing" so brand-new checks don't miss outages.
	ConsensusM param.Field[int64] `json:"consensus_m" api:"required"`
	// M-of-N consensus rule numerator. The evaluator considers the check
	// globally-failing only when at least this many locations are reporting fail
	// concurrently.
	ConsensusN param.Field[int64] `json:"consensus_n" api:"required"`
	// When false, the evaluator skips this check entirely.
	Enabled param.Field[bool] `json:"enabled" api:"required"`
	// Server-set; ignored on write.
	AccountID param.Field[string] `json:"account_id" format:"uuid"`
	// Server-set; ignored on write.
	CheckID param.Field[string] `json:"check_id" format:"uuid"`
	// Absolute-time windows during which the evaluator suppresses dispatch but still
	// updates state. Cron-style recurring windows are a future enhancement.
	MaintenanceWindows param.Field[[]MaintenanceWindowParam] `json:"maintenance_windows"`
}

func (r AlertConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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
	Config map[string]interface{} `json:"config"`
	// Region/location on read responses is empty; populated on create requests only.
	Location string `json:"location"`
	// Cloud provider on read responses is empty; populated on create requests only.
	Provider string    `json:"provider"`
	JSON     checkJSON `json:"-"`
}

// checkJSON contains the JSON metadata for the struct [Check]
type checkJSON struct {
	ID            apijson.Field
	AccountTypeid apijson.Field
	CreatedAt     apijson.Field
	Enabled       apijson.Field
	GarrisonID    apijson.Field
	Name          apijson.Field
	Schedule      apijson.Field
	TargetURL     apijson.Field
	TimeoutMs     apijson.Field
	Type          apijson.Field
	UpdatedAt     apijson.Field
	ArtifactURL   apijson.Field
	Config        apijson.Field
	Location      apijson.Field
	Provider      apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *Check) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r checkJSON) RawJSON() string {
	return r.raw
}

type MaintenanceWindow struct {
	EndUnixMs   int64                 `json:"end_unix_ms" api:"required"`
	StartUnixMs int64                 `json:"start_unix_ms" api:"required"`
	JSON        maintenanceWindowJSON `json:"-"`
}

// maintenanceWindowJSON contains the JSON metadata for the struct
// [MaintenanceWindow]
type maintenanceWindowJSON struct {
	EndUnixMs   apijson.Field
	StartUnixMs apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MaintenanceWindow) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r maintenanceWindowJSON) RawJSON() string {
	return r.raw
}

type MaintenanceWindowParam struct {
	EndUnixMs   param.Field[int64] `json:"end_unix_ms" api:"required"`
	StartUnixMs param.Field[int64] `json:"start_unix_ms" api:"required"`
}

func (r MaintenanceWindowParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CheckListResponse struct {
	Checks []Check `json:"checks" api:"required"`
	// Offset to pass on the next request to continue pagination. Zero (or absent) when
	// there's no more data.
	NextOffset int64                 `json:"next_offset"`
	JSON       checkListResponseJSON `json:"-"`
}

// checkListResponseJSON contains the JSON metadata for the struct
// [CheckListResponse]
type checkListResponseJSON struct {
	Checks      apijson.Field
	NextOffset  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CheckListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r checkListResponseJSON) RawJSON() string {
	return r.raw
}

type CheckNewParams struct {
	Enabled param.Field[bool] `json:"enabled" api:"required"`
	// Provider-specific region/location.
	Location param.Field[string] `json:"location" api:"required"`
	Name     param.Field[string] `json:"name" api:"required"`
	// Cloud provider (`mock`, `ec2`, `ovh`, `azure`, `gcp`, `hetzner`).
	Provider    param.Field[string]                 `json:"provider" api:"required"`
	Schedule    param.Field[string]                 `json:"schedule" api:"required"`
	TargetURL   param.Field[string]                 `json:"target_url" api:"required" format:"uri"`
	Type        param.Field[string]                 `json:"type" api:"required"`
	ArtifactURL param.Field[string]                 `json:"artifact_url"`
	Config      param.Field[map[string]interface{}] `json:"config"`
	TimeoutMs   param.Field[int64]                  `json:"timeout_ms"`
}

func (r CheckNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CheckUpdateParams struct {
	ArtifactURL param.Field[string]                 `json:"artifact_url"`
	Config      param.Field[map[string]interface{}] `json:"config"`
	Enabled     param.Field[bool]                   `json:"enabled"`
	Name        param.Field[string]                 `json:"name"`
	Schedule    param.Field[string]                 `json:"schedule"`
	TargetURL   param.Field[string]                 `json:"target_url" format:"uri"`
	TimeoutMs   param.Field[int64]                  `json:"timeout_ms"`
	Type        param.Field[string]                 `json:"type"`
}

func (r CheckUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CheckListParams struct {
	// Max number of checks to return. Defaults to 100; the server caps further.
	Limit param.Field[int64] `query:"limit"`
	// Number of checks to skip. Pass the `next_offset` from the previous page.
	Offset param.Field[int64] `query:"offset"`
}

// URLQuery serializes [CheckListParams]'s query parameters as `url.Values`.
func (r CheckListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
