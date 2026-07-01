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
	"github.com/simplechecks/simplechecks-sdk-go/packages/pagination"
)

// Reusable, account-scoped notification destinations (webhook, Slack, Discord,
// Teams, PagerDuty, Opsgenie, email). One channel can serve many checks. Includes
// a test-fire endpoint.
//
// AlertChannelService contains methods and other services that help with
// interacting with the simple-checks API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAlertChannelService] method instead.
type AlertChannelService struct {
	Options []option.RequestOption
}

// NewAlertChannelService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAlertChannelService(opts ...option.RequestOption) (r *AlertChannelService) {
	r = &AlertChannelService{}
	r.Options = opts
	return
}

// Creates a reusable notification destination. URL-bearing types (`webhook`,
// `slack`, `discord`, `teams`) are SSRF-filtered: targets resolving to private,
// loopback, or link-local addresses are rejected. The `target` is write-only —
// it's masked on every read. Requires the `alerts:write` scope (owner/admin only).
func (r *AlertChannelService) New(ctx context.Context, body AlertChannelNewParams, opts ...option.RequestOption) (res *AlertChannel, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/alert-channels"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Returns the alert channel. The `target` secret is masked. 404 if no such channel
// exists for the calling account. Requires the `alerts:read` scope.
func (r *AlertChannelService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *AlertChannel, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/alert-channels/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// All fields are optional; omitted fields are unchanged. A `target` equal to the
// masked value (`***<last4>`) is a no-op — only a fresh, non-masked secret updates
// the stored target. Requires the `alerts:write` scope (owner/admin only).
func (r *AlertChannelService) Update(ctx context.Context, id string, body AlertChannelUpdateParams, opts ...option.RequestOption) (res *AlertChannel, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/alert-channels/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Returns the caller's reusable alert channels with cursor pagination.
// `next_cursor` is set when a full page was returned and null on the final page.
// The `target` secret is always masked (`***<last4>`); the raw value is never
// returned. Requires the `alerts:read` scope.
func (r *AlertChannelService) List(ctx context.Context, query AlertChannelListParams, opts ...option.RequestOption) (res *pagination.AlertChannelsCursor[AlertChannel], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/alert-channels"
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

// Returns the caller's reusable alert channels with cursor pagination.
// `next_cursor` is set when a full page was returned and null on the final page.
// The `target` secret is always masked (`***<last4>`); the raw value is never
// returned. Requires the `alerts:read` scope.
func (r *AlertChannelService) ListAutoPaging(ctx context.Context, query AlertChannelListParams, opts ...option.RequestOption) *pagination.AlertChannelsCursorAutoPager[AlertChannel] {
	return pagination.NewAlertChannelsCursorAutoPager(r.List(ctx, query, opts...))
}

// Deletes the channel and cascades its subscriptions (the bound checks simply stop
// notifying it). Requires the `alerts:write` scope (owner/admin only).
func (r *AlertChannelService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("v1/alert-channels/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Enqueues a single `test_fire` dispatch through the channel so a customer can
// verify the destination works. Idempotent on the channel id (repeated clicks
// dedup). Requires the `alerts:write` scope (owner/admin only).
func (r *AlertChannelService) TestFire(ctx context.Context, id string, opts ...option.RequestOption) (res *AlertChannelTestFireResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/alert-channels/%s:test", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// A first-class, reusable alert channel. Referenced by many checks through alert
// subscriptions. The `target` secret is always returned masked (`***<last4>`).
type AlertChannel struct {
	// Channel id in `chan_<typeid>` form.
	ID string `json:"id" api:"required"`
	// Owning account's `acct_<typeid>`. Read-only.
	AccountTypeid string    `json:"account_typeid" api:"required"`
	CreatedAt     time.Time `json:"created_at" api:"required" format:"date-time"`
	// Account-unique display name.
	Name string `json:"name" api:"required"`
	// Masked destination secret (`***<last4>`). The raw value is write-only and never
	// returned.
	Target    string           `json:"target" api:"required"`
	Type      AlertChannelType `json:"type" api:"required"`
	UpdatedAt time.Time        `json:"updated_at" api:"required" format:"date-time"`
	// Type-specific options. Optional.
	Config map[string]interface{} `json:"config"`
	JSON   alertChannelJSON       `json:"-"`
}

// alertChannelJSON contains the JSON metadata for the struct [AlertChannel]
type alertChannelJSON struct {
	ID            apijson.Field
	AccountTypeid apijson.Field
	CreatedAt     apijson.Field
	Name          apijson.Field
	Target        apijson.Field
	Type          apijson.Field
	UpdatedAt     apijson.Field
	Config        apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AlertChannel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r alertChannelJSON) RawJSON() string {
	return r.raw
}

type AlertChannelType string

const (
	AlertChannelTypeSlack     AlertChannelType = "slack"
	AlertChannelTypeDiscord   AlertChannelType = "discord"
	AlertChannelTypeTeams     AlertChannelType = "teams"
	AlertChannelTypeWebhook   AlertChannelType = "webhook"
	AlertChannelTypePagerduty AlertChannelType = "pagerduty"
	AlertChannelTypeOpsgenie  AlertChannelType = "opsgenie"
	AlertChannelTypeEmail     AlertChannelType = "email"
)

func (r AlertChannelType) IsKnown() bool {
	switch r {
	case AlertChannelTypeSlack, AlertChannelTypeDiscord, AlertChannelTypeTeams, AlertChannelTypeWebhook, AlertChannelTypePagerduty, AlertChannelTypeOpsgenie, AlertChannelTypeEmail:
		return true
	}
	return false
}

type AlertChannelTestFireResponse struct {
	// 1 if a new dispatch was enqueued, 0 if it deduped.
	Enqueued int64 `json:"enqueued" api:"required"`
	// Synthetic incident id for the test dispatch.
	IncidentID string                           `json:"incident_id" api:"required" format:"uuid"`
	JSON       alertChannelTestFireResponseJSON `json:"-"`
}

// alertChannelTestFireResponseJSON contains the JSON metadata for the struct
// [AlertChannelTestFireResponse]
type alertChannelTestFireResponseJSON struct {
	Enqueued    apijson.Field
	IncidentID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AlertChannelTestFireResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r alertChannelTestFireResponseJSON) RawJSON() string {
	return r.raw
}

type AlertChannelNewParams struct {
	Name param.Field[string] `json:"name" api:"required"`
	// Destination secret. URL for the webhook flavors (slack/discord/teams/webhook),
	// email address for `email`, integration key for `pagerduty`, API key for
	// `opsgenie`. URL-bearing types are SSRF-filtered.
	Target param.Field[string]                    `json:"target" api:"required"`
	Type   param.Field[AlertChannelNewParamsType] `json:"type" api:"required"`
	Config param.Field[map[string]interface{}]    `json:"config"`
}

func (r AlertChannelNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AlertChannelNewParamsType string

const (
	AlertChannelNewParamsTypeSlack     AlertChannelNewParamsType = "slack"
	AlertChannelNewParamsTypeDiscord   AlertChannelNewParamsType = "discord"
	AlertChannelNewParamsTypeTeams     AlertChannelNewParamsType = "teams"
	AlertChannelNewParamsTypeWebhook   AlertChannelNewParamsType = "webhook"
	AlertChannelNewParamsTypePagerduty AlertChannelNewParamsType = "pagerduty"
	AlertChannelNewParamsTypeOpsgenie  AlertChannelNewParamsType = "opsgenie"
	AlertChannelNewParamsTypeEmail     AlertChannelNewParamsType = "email"
)

func (r AlertChannelNewParamsType) IsKnown() bool {
	switch r {
	case AlertChannelNewParamsTypeSlack, AlertChannelNewParamsTypeDiscord, AlertChannelNewParamsTypeTeams, AlertChannelNewParamsTypeWebhook, AlertChannelNewParamsTypePagerduty, AlertChannelNewParamsTypeOpsgenie, AlertChannelNewParamsTypeEmail:
		return true
	}
	return false
}

type AlertChannelUpdateParams struct {
	Config param.Field[map[string]interface{}]       `json:"config"`
	Name   param.Field[string]                       `json:"name"`
	Target param.Field[string]                       `json:"target"`
	Type   param.Field[AlertChannelUpdateParamsType] `json:"type"`
}

func (r AlertChannelUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AlertChannelUpdateParamsType string

const (
	AlertChannelUpdateParamsTypeSlack     AlertChannelUpdateParamsType = "slack"
	AlertChannelUpdateParamsTypeDiscord   AlertChannelUpdateParamsType = "discord"
	AlertChannelUpdateParamsTypeTeams     AlertChannelUpdateParamsType = "teams"
	AlertChannelUpdateParamsTypeWebhook   AlertChannelUpdateParamsType = "webhook"
	AlertChannelUpdateParamsTypePagerduty AlertChannelUpdateParamsType = "pagerduty"
	AlertChannelUpdateParamsTypeOpsgenie  AlertChannelUpdateParamsType = "opsgenie"
	AlertChannelUpdateParamsTypeEmail     AlertChannelUpdateParamsType = "email"
)

func (r AlertChannelUpdateParamsType) IsKnown() bool {
	switch r {
	case AlertChannelUpdateParamsTypeSlack, AlertChannelUpdateParamsTypeDiscord, AlertChannelUpdateParamsTypeTeams, AlertChannelUpdateParamsTypeWebhook, AlertChannelUpdateParamsTypePagerduty, AlertChannelUpdateParamsTypeOpsgenie, AlertChannelUpdateParamsTypeEmail:
		return true
	}
	return false
}

type AlertChannelListParams struct {
	// Opaque pagination token from the previous page's `next_cursor`.
	Cursor param.Field[string] `query:"cursor"`
	Limit  param.Field[int64]  `query:"limit"`
}

// URLQuery serializes [AlertChannelListParams]'s query parameters as `url.Values`.
func (r AlertChannelListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
