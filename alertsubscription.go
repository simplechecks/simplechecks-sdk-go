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

// Bindings of a check to an alert channel, each carrying its own notify-on-failure
// / notify-on-recovery flags.
//
// AlertSubscriptionService contains methods and other services that help with
// interacting with the simple-checks API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAlertSubscriptionService] method instead.
type AlertSubscriptionService struct {
	Options []option.RequestOption
}

// NewAlertSubscriptionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAlertSubscriptionService(opts ...option.RequestOption) (r *AlertSubscriptionService) {
	r = &AlertSubscriptionService{}
	r.Options = opts
	return
}

// Binds a check to a channel and carries the per-binding notify flags
// (`notify_on_failure`, `notify_on_recovery`, both default true). The binding is
// account-scoped: a check or channel that isn't yours yields 404. A duplicate
// `(check_id, channel_id)` binding yields 409. Requires the `alerts:write` scope
// (owner/admin only).
func (r *AlertSubscriptionService) New(ctx context.Context, body AlertSubscriptionNewParams, opts ...option.RequestOption) (res *AlertSubscription, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/alert-subscriptions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Returns the subscription. 404 if no such subscription exists for the calling
// account. Requires the `alerts:read` scope.
func (r *AlertSubscriptionService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *AlertSubscription, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/alert-subscriptions/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Updates only the notify flags (`notify_on_failure`, `notify_on_recovery`); the
// check and channel bindings are immutable. Omitted flags are unchanged. Requires
// the `alerts:write` scope (owner/admin only).
func (r *AlertSubscriptionService) Update(ctx context.Context, id string, body AlertSubscriptionUpdateParams, opts ...option.RequestOption) (res *AlertSubscription, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/alert-subscriptions/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Returns the caller's check↔channel subscriptions with cursor pagination.
// Optionally filter by `check_id` and/or `channel_id`. `next_cursor` is set when a
// full page was returned and null on the final page. Requires the `alerts:read`
// scope.
func (r *AlertSubscriptionService) List(ctx context.Context, query AlertSubscriptionListParams, opts ...option.RequestOption) (res *pagination.AlertSubscriptionsCursor[AlertSubscription], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/alert-subscriptions"
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

// Returns the caller's check↔channel subscriptions with cursor pagination.
// Optionally filter by `check_id` and/or `channel_id`. `next_cursor` is set when a
// full page was returned and null on the final page. Requires the `alerts:read`
// scope.
func (r *AlertSubscriptionService) ListAutoPaging(ctx context.Context, query AlertSubscriptionListParams, opts ...option.RequestOption) *pagination.AlertSubscriptionsCursorAutoPager[AlertSubscription] {
	return pagination.NewAlertSubscriptionsCursorAutoPager(r.List(ctx, query, opts...))
}

// Removes the binding; the check stops notifying that channel. Requires the
// `alerts:write` scope (owner/admin only).
func (r *AlertSubscriptionService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("v1/alert-subscriptions/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// A binding between one check and one alert channel, carrying the per-binding
// notify flags. The same channel can be subscribed by many checks, each with its
// own flags.
type AlertSubscription struct {
	// Subscription id in `asub_<typeid>` form.
	ID string `json:"id" api:"required"`
	// Owning account's `acct_<typeid>`. Read-only.
	AccountTypeid string `json:"account_typeid" api:"required"`
	// The bound channel's id in `chan_<typeid>` form.
	ChannelID string `json:"channel_id" api:"required"`
	// The subscribed check's id.
	CheckID   string    `json:"check_id" api:"required" format:"uuid"`
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// When true, an incident-started event dispatches to this channel.
	NotifyOnFailure bool `json:"notify_on_failure" api:"required"`
	// When true, an incident-recovered event dispatches to this channel.
	NotifyOnRecovery bool                  `json:"notify_on_recovery" api:"required"`
	UpdatedAt        time.Time             `json:"updated_at" api:"required" format:"date-time"`
	JSON             alertSubscriptionJSON `json:"-"`
}

// alertSubscriptionJSON contains the JSON metadata for the struct
// [AlertSubscription]
type alertSubscriptionJSON struct {
	ID               apijson.Field
	AccountTypeid    apijson.Field
	ChannelID        apijson.Field
	CheckID          apijson.Field
	CreatedAt        apijson.Field
	NotifyOnFailure  apijson.Field
	NotifyOnRecovery apijson.Field
	UpdatedAt        apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AlertSubscription) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r alertSubscriptionJSON) RawJSON() string {
	return r.raw
}

type AlertSubscriptionNewParams struct {
	// The channel to bind in `chan_<typeid>` form (must belong to your account).
	ChannelID param.Field[string] `json:"channel_id" api:"required"`
	// The check to subscribe (must belong to your account).
	CheckID param.Field[string] `json:"check_id" api:"required" format:"uuid"`
	// Defaults to true when omitted.
	NotifyOnFailure param.Field[bool] `json:"notify_on_failure"`
	// Defaults to true when omitted.
	NotifyOnRecovery param.Field[bool] `json:"notify_on_recovery"`
}

func (r AlertSubscriptionNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AlertSubscriptionUpdateParams struct {
	NotifyOnFailure  param.Field[bool] `json:"notify_on_failure"`
	NotifyOnRecovery param.Field[bool] `json:"notify_on_recovery"`
}

func (r AlertSubscriptionUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AlertSubscriptionListParams struct {
	// Filter to subscriptions for this channel (`chan_<typeid>`).
	ChannelID param.Field[string] `query:"channel_id"`
	// Filter to subscriptions for this check (raw check UUID).
	CheckID param.Field[string] `query:"check_id" format:"uuid"`
	// Opaque pagination token from the previous page's `next_cursor`.
	Cursor param.Field[string] `query:"cursor"`
	Limit  param.Field[int64]  `query:"limit"`
}

// URLQuery serializes [AlertSubscriptionListParams]'s query parameters as
// `url.Values`.
func (r AlertSubscriptionListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
