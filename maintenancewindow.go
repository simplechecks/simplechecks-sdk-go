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

// Account-scoped windows that pause execution of their targeted checks for the
// scheduled interval(s); paused runs are not recorded and never count against
// uptime.
//
// MaintenanceWindowService contains methods and other services that help with
// interacting with the simple-checks API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMaintenanceWindowService] method instead.
type MaintenanceWindowService struct {
	Options []option.RequestOption
}

// NewMaintenanceWindowService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewMaintenanceWindowService(opts ...option.RequestOption) (r *MaintenanceWindowService) {
	r = &MaintenanceWindowService{}
	r.Options = opts
	return
}

// Creates a maintenance window that pauses execution of its targeted checks for
// the scheduled interval(s). `schedule_kind` is `one_time` or `recurring`;
// recurrence fields (`repeat_unit`, `repeat_interval`, `repeat_ends_unix_ms`) are
// valid only for a recurring window. `timezone` is an IANA name. `check_ids` are
// raw check UUIDs and must belong to your account; a check id that doesn't
// returns 404. Requires the `alerts:write` scope (owner/admin only).
func (r *MaintenanceWindowService) New(ctx context.Context, body MaintenanceWindowNewParams, opts ...option.RequestOption) (res *MaintenanceWindow, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/maintenance-windows"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Returns the window with its targeting. 404 if no such window exists for the
// calling account. Requires the `alerts:read` scope.
func (r *MaintenanceWindowService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *MaintenanceWindow, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/maintenance-windows/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Updates the supplied fields. A non-null `check_ids` replaces the targeting set;
// a check id that isn't your account's returns 404. The effective schedule is
// re-validated. Omitted fields are unchanged. Requires the `alerts:write` scope
// (owner/admin only).
func (r *MaintenanceWindowService) Update(ctx context.Context, id string, body MaintenanceWindowUpdateParams, opts ...option.RequestOption) (res *MaintenanceWindow, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/maintenance-windows/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Returns the caller's maintenance windows with cursor pagination. Each window
// carries its explicit check targeting (`check_ids`). `next_cursor` is set when a
// full page was returned and null on the final page. Requires the `alerts:read`
// scope.
func (r *MaintenanceWindowService) List(ctx context.Context, query MaintenanceWindowListParams, opts ...option.RequestOption) (res *pagination.MaintenanceWindowsCursor[MaintenanceWindow], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/maintenance-windows"
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

// Returns the caller's maintenance windows with cursor pagination. Each window
// carries its explicit check targeting (`check_ids`). `next_cursor` is set when a
// full page was returned and null on the final page. Requires the `alerts:read`
// scope.
func (r *MaintenanceWindowService) ListAutoPaging(ctx context.Context, query MaintenanceWindowListParams, opts ...option.RequestOption) *pagination.MaintenanceWindowsCursorAutoPager[MaintenanceWindow] {
	return pagination.NewMaintenanceWindowsCursorAutoPager(r.List(ctx, query, opts...))
}

// Removes the window and its targeting; affected checks resume normal execution.
// Requires the `alerts:write` scope (owner/admin only).
func (r *MaintenanceWindowService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("v1/maintenance-windows/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// A maintenance window that pauses execution of its targeted checks for the
// scheduled interval(s). The DST-correct occurrence expansion is performed by the
// control plane; this resource carries the stored schedule shape plus its explicit
// targeting.
type MaintenanceWindow struct {
	// Window id in `mwin_<typeid>` form.
	ID string `json:"id" api:"required"`
	// Owning account's `acct_<typeid>`. Read-only.
	AccountTypeid string `json:"account_typeid" api:"required"`
	// Raw UUIDs of the targeted checks.
	CheckIDs []string `json:"check_ids" api:"required" format:"uuid"`
	// Reserved for tag-based targeting; accepted but not yet consumed.
	CheckTags []string  `json:"check_tags" api:"required"`
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Window duration in milliseconds (> 0).
	DurationMs   int64                         `json:"duration_ms" api:"required"`
	Name         string                        `json:"name" api:"required"`
	ScheduleKind MaintenanceWindowScheduleKind `json:"schedule_kind" api:"required"`
	// First occurrence start, Unix epoch milliseconds.
	StartUnixMs int64 `json:"start_unix_ms" api:"required"`
	// IANA timezone name (e.g. "America/Chicago"). Defaults to UTC.
	Timezone  string    `json:"timezone" api:"required"`
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Recurrence end bound, Unix epoch ms; recurring only.
	RepeatEndsUnixMs int64 `json:"repeat_ends_unix_ms"`
	// Recurrence interval (e.g. every N units); recurring only.
	RepeatInterval int64 `json:"repeat_interval"`
	// Recurrence unit; present only for a recurring window.
	RepeatUnit MaintenanceWindowRepeatUnit `json:"repeat_unit"`
	JSON       maintenanceWindowJSON       `json:"-"`
}

// maintenanceWindowJSON contains the JSON metadata for the struct
// [MaintenanceWindow]
type maintenanceWindowJSON struct {
	ID               apijson.Field
	AccountTypeid    apijson.Field
	CheckIDs         apijson.Field
	CheckTags        apijson.Field
	CreatedAt        apijson.Field
	DurationMs       apijson.Field
	Name             apijson.Field
	ScheduleKind     apijson.Field
	StartUnixMs      apijson.Field
	Timezone         apijson.Field
	UpdatedAt        apijson.Field
	RepeatEndsUnixMs apijson.Field
	RepeatInterval   apijson.Field
	RepeatUnit       apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *MaintenanceWindow) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r maintenanceWindowJSON) RawJSON() string {
	return r.raw
}

type MaintenanceWindowScheduleKind string

const (
	MaintenanceWindowScheduleKindOneTime   MaintenanceWindowScheduleKind = "one_time"
	MaintenanceWindowScheduleKindRecurring MaintenanceWindowScheduleKind = "recurring"
)

func (r MaintenanceWindowScheduleKind) IsKnown() bool {
	switch r {
	case MaintenanceWindowScheduleKindOneTime, MaintenanceWindowScheduleKindRecurring:
		return true
	}
	return false
}

// Recurrence unit; present only for a recurring window.
type MaintenanceWindowRepeatUnit string

const (
	MaintenanceWindowRepeatUnitDay   MaintenanceWindowRepeatUnit = "DAY"
	MaintenanceWindowRepeatUnitWeek  MaintenanceWindowRepeatUnit = "WEEK"
	MaintenanceWindowRepeatUnitMonth MaintenanceWindowRepeatUnit = "MONTH"
)

func (r MaintenanceWindowRepeatUnit) IsKnown() bool {
	switch r {
	case MaintenanceWindowRepeatUnitDay, MaintenanceWindowRepeatUnitWeek, MaintenanceWindowRepeatUnitMonth:
		return true
	}
	return false
}

type MaintenanceWindowNewParams struct {
	// Window duration in milliseconds; must be positive.
	DurationMs   param.Field[int64]                                  `json:"duration_ms" api:"required"`
	Name         param.Field[string]                                 `json:"name" api:"required"`
	ScheduleKind param.Field[MaintenanceWindowNewParamsScheduleKind] `json:"schedule_kind" api:"required"`
	StartUnixMs  param.Field[int64]                                  `json:"start_unix_ms" api:"required"`
	// Raw check UUIDs to target (must belong to your account).
	CheckIDs  param.Field[[]string] `json:"check_ids" format:"uuid"`
	CheckTags param.Field[[]string] `json:"check_tags"`
	// Valid only for a recurring window.
	RepeatEndsUnixMs param.Field[int64] `json:"repeat_ends_unix_ms"`
	// Valid only for a recurring window; must be positive.
	RepeatInterval param.Field[int64] `json:"repeat_interval"`
	// Valid only for a recurring window.
	RepeatUnit param.Field[MaintenanceWindowNewParamsRepeatUnit] `json:"repeat_unit"`
	// IANA timezone name. Defaults to UTC when omitted.
	Timezone param.Field[string] `json:"timezone"`
}

func (r MaintenanceWindowNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MaintenanceWindowNewParamsScheduleKind string

const (
	MaintenanceWindowNewParamsScheduleKindOneTime   MaintenanceWindowNewParamsScheduleKind = "one_time"
	MaintenanceWindowNewParamsScheduleKindRecurring MaintenanceWindowNewParamsScheduleKind = "recurring"
)

func (r MaintenanceWindowNewParamsScheduleKind) IsKnown() bool {
	switch r {
	case MaintenanceWindowNewParamsScheduleKindOneTime, MaintenanceWindowNewParamsScheduleKindRecurring:
		return true
	}
	return false
}

// Valid only for a recurring window.
type MaintenanceWindowNewParamsRepeatUnit string

const (
	MaintenanceWindowNewParamsRepeatUnitDay   MaintenanceWindowNewParamsRepeatUnit = "DAY"
	MaintenanceWindowNewParamsRepeatUnitWeek  MaintenanceWindowNewParamsRepeatUnit = "WEEK"
	MaintenanceWindowNewParamsRepeatUnitMonth MaintenanceWindowNewParamsRepeatUnit = "MONTH"
)

func (r MaintenanceWindowNewParamsRepeatUnit) IsKnown() bool {
	switch r {
	case MaintenanceWindowNewParamsRepeatUnitDay, MaintenanceWindowNewParamsRepeatUnitWeek, MaintenanceWindowNewParamsRepeatUnitMonth:
		return true
	}
	return false
}

type MaintenanceWindowUpdateParams struct {
	CheckIDs         param.Field[[]string]                                  `json:"check_ids" format:"uuid"`
	CheckTags        param.Field[[]string]                                  `json:"check_tags"`
	DurationMs       param.Field[int64]                                     `json:"duration_ms"`
	Name             param.Field[string]                                    `json:"name"`
	RepeatEndsUnixMs param.Field[int64]                                     `json:"repeat_ends_unix_ms"`
	RepeatInterval   param.Field[int64]                                     `json:"repeat_interval"`
	RepeatUnit       param.Field[MaintenanceWindowUpdateParamsRepeatUnit]   `json:"repeat_unit"`
	ScheduleKind     param.Field[MaintenanceWindowUpdateParamsScheduleKind] `json:"schedule_kind"`
	StartUnixMs      param.Field[int64]                                     `json:"start_unix_ms"`
	Timezone         param.Field[string]                                    `json:"timezone"`
}

func (r MaintenanceWindowUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MaintenanceWindowUpdateParamsRepeatUnit string

const (
	MaintenanceWindowUpdateParamsRepeatUnitDay   MaintenanceWindowUpdateParamsRepeatUnit = "DAY"
	MaintenanceWindowUpdateParamsRepeatUnitWeek  MaintenanceWindowUpdateParamsRepeatUnit = "WEEK"
	MaintenanceWindowUpdateParamsRepeatUnitMonth MaintenanceWindowUpdateParamsRepeatUnit = "MONTH"
)

func (r MaintenanceWindowUpdateParamsRepeatUnit) IsKnown() bool {
	switch r {
	case MaintenanceWindowUpdateParamsRepeatUnitDay, MaintenanceWindowUpdateParamsRepeatUnitWeek, MaintenanceWindowUpdateParamsRepeatUnitMonth:
		return true
	}
	return false
}

type MaintenanceWindowUpdateParamsScheduleKind string

const (
	MaintenanceWindowUpdateParamsScheduleKindOneTime   MaintenanceWindowUpdateParamsScheduleKind = "one_time"
	MaintenanceWindowUpdateParamsScheduleKindRecurring MaintenanceWindowUpdateParamsScheduleKind = "recurring"
)

func (r MaintenanceWindowUpdateParamsScheduleKind) IsKnown() bool {
	switch r {
	case MaintenanceWindowUpdateParamsScheduleKindOneTime, MaintenanceWindowUpdateParamsScheduleKindRecurring:
		return true
	}
	return false
}

type MaintenanceWindowListParams struct {
	// Opaque pagination token from the previous page's `next_cursor`.
	Cursor param.Field[string] `query:"cursor"`
	Limit  param.Field[int64]  `query:"limit"`
}

// URLQuery serializes [MaintenanceWindowListParams]'s query parameters as
// `url.Values`.
func (r MaintenanceWindowListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
