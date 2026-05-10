// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package simplechecksgo

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/simplechecks/simplechecks-sdk-go/internal/apijson"
	"github.com/simplechecks/simplechecks-sdk-go/internal/requestconfig"
	"github.com/simplechecks/simplechecks-sdk-go/option"
)

// Per-check alert configuration + test-fire endpoint (PR-Alerts/1).
//
// CheckAlertService contains methods and other services that help with interacting
// with the simple-checks API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCheckAlertService] method instead.
type CheckAlertService struct {
	Options []option.RequestOption
}

// NewCheckAlertService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCheckAlertService(opts ...option.RequestOption) (r *CheckAlertService) {
	r = &CheckAlertService{}
	r.Options = opts
	return
}

// Returns the per-check alert configuration: enabled flag, thresholds, M-of-N
// consensus, maintenance windows, channels. Requires the `checks:read` scope.
func (r *CheckAlertService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *AlertConfig, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/checks/%s/alerts", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Subsequent runs will not be evaluated for alerts. State rows in `alert_state`
// and `alert_location_state` cascade with the underlying check; deleting just the
// config leaves them behind harmlessly. Requires the `checks:write` scope.
func (r *CheckAlertService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("v1/checks/%s/alerts", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Idempotent upsert. The same body shape is returned by GET. Channels supported:
// email, slack, discord, teams, webhook, pagerduty, opsgenie. The PR-Alerts/1
// evaluator runs M-of-N consensus before incident-firing; if fewer than
// `consensus_m` locations have observations, the rule falls back to "any failing =
// failing" so brand-new checks don't miss outages.
//
// Eventual-consistency contract: after a config write, the evaluator picks up the
// new thresholds on the next ingest cycle (15s push cadence).
//
// Requires the `checks:write` scope.
func (r *CheckAlertService) Replace(ctx context.Context, id string, body CheckAlertReplaceParams, opts ...option.RequestOption) (res *AlertConfig, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/checks/%s/alerts", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

// Synthesizes a `test_fire` dispatch per channel and enqueues them for the async
// dispatcher. Useful for verifying that a Slack webhook URL or PagerDuty
// integration key actually works without waiting for a real failure. The test
// dispatches do not affect alert state or incident lifecycle. Requires the
// `checks:write` scope.
func (r *CheckAlertService) TestFire(ctx context.Context, id string, opts ...option.RequestOption) (res *CheckAlertTestFireResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/checks/%s/alerts:test", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

type CheckAlertTestFireResponse struct {
	// Total channels configured on the check.
	ChannelCount int64 `json:"channel_count" api:"required"`
	// Number of dispatches accepted (un-deduped).
	Enqueued int64 `json:"enqueued" api:"required"`
	// Synthetic incident id used to dedupe the test dispatches against accidental
	// double-clicks.
	IncidentID string                         `json:"incident_id" api:"required" format:"uuid"`
	JSON       checkAlertTestFireResponseJSON `json:"-"`
}

// checkAlertTestFireResponseJSON contains the JSON metadata for the struct
// [CheckAlertTestFireResponse]
type checkAlertTestFireResponseJSON struct {
	ChannelCount apijson.Field
	Enqueued     apijson.Field
	IncidentID   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *CheckAlertTestFireResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r checkAlertTestFireResponseJSON) RawJSON() string {
	return r.raw
}

type CheckAlertReplaceParams struct {
	AlertConfig AlertConfigParam `json:"alert_config" api:"required"`
}

func (r CheckAlertReplaceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.AlertConfig)
}
