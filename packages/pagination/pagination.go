// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pagination

import (
	"net/http"
	"strconv"

	"github.com/simplechecks/simplechecks-sdk-go/internal/apijson"
	"github.com/simplechecks/simplechecks-sdk-go/internal/requestconfig"
	"github.com/simplechecks/simplechecks-sdk-go/option"
)

type Offset[T any] struct {
	Checks     []T        `json:"checks"`
	NextOffset int64      `json:"next_offset" api:"nullable"`
	JSON       offsetJSON `json:"-"`
	cfg        *requestconfig.RequestConfig
	res        *http.Response
}

// offsetJSON contains the JSON metadata for the struct [Offset[T]]
type offsetJSON struct {
	Checks      apijson.Field
	NextOffset  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Offset[T]) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r offsetJSON) RawJSON() string {
	return r.raw
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *Offset[T]) GetNextPage() (res *Offset[T], err error) {
	if len(r.Checks) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)

	q := cfg.Request.URL.Query()
	offset, err := strconv.ParseInt(q.Get("offset"), 10, 64)
	if err != nil {
		offset = 0
	}
	length := int64(len(r.Checks))
	next := offset + length

	if length > 0 && next != 0 {
		err = cfg.Apply(option.WithQuery("offset", strconv.FormatInt(next, 10)))
		if err != nil {
			return nil, err
		}
	} else {
		return nil, nil
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *Offset[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &Offset[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type OffsetAutoPager[T any] struct {
	page *Offset[T]
	cur  T
	idx  int
	run  int
	err  error
}

func NewOffsetAutoPager[T any](page *Offset[T], err error) *OffsetAutoPager[T] {
	return &OffsetAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *OffsetAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Checks) == 0 {
		return false
	}
	if r.idx >= len(r.page.Checks) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Checks) == 0 {
			return false
		}
	}
	r.cur = r.page.Checks[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *OffsetAutoPager[T]) Current() T {
	return r.cur
}

func (r *OffsetAutoPager[T]) Err() error {
	return r.err
}

func (r *OffsetAutoPager[T]) Index() int {
	return r.run
}

type IncidentsOffset[T any] struct {
	Incidents  []T                 `json:"incidents"`
	NextOffset int64               `json:"next_offset" api:"nullable"`
	JSON       incidentsOffsetJSON `json:"-"`
	cfg        *requestconfig.RequestConfig
	res        *http.Response
}

// incidentsOffsetJSON contains the JSON metadata for the struct
// [IncidentsOffset[T]]
type incidentsOffsetJSON struct {
	Incidents   apijson.Field
	NextOffset  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IncidentsOffset[T]) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r incidentsOffsetJSON) RawJSON() string {
	return r.raw
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *IncidentsOffset[T]) GetNextPage() (res *IncidentsOffset[T], err error) {
	if len(r.Incidents) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)

	q := cfg.Request.URL.Query()
	offset, err := strconv.ParseInt(q.Get("offset"), 10, 64)
	if err != nil {
		offset = 0
	}
	length := int64(len(r.Incidents))
	next := offset + length

	if length > 0 && next != 0 {
		err = cfg.Apply(option.WithQuery("offset", strconv.FormatInt(next, 10)))
		if err != nil {
			return nil, err
		}
	} else {
		return nil, nil
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *IncidentsOffset[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &IncidentsOffset[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type IncidentsOffsetAutoPager[T any] struct {
	page *IncidentsOffset[T]
	cur  T
	idx  int
	run  int
	err  error
}

func NewIncidentsOffsetAutoPager[T any](page *IncidentsOffset[T], err error) *IncidentsOffsetAutoPager[T] {
	return &IncidentsOffsetAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *IncidentsOffsetAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Incidents) == 0 {
		return false
	}
	if r.idx >= len(r.page.Incidents) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Incidents) == 0 {
			return false
		}
	}
	r.cur = r.page.Incidents[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *IncidentsOffsetAutoPager[T]) Current() T {
	return r.cur
}

func (r *IncidentsOffsetAutoPager[T]) Err() error {
	return r.err
}

func (r *IncidentsOffsetAutoPager[T]) Index() int {
	return r.run
}

type RunsCursor[T any] struct {
	Runs       []T            `json:"runs"`
	NextCursor string         `json:"next_cursor" api:"nullable"`
	JSON       runsCursorJSON `json:"-"`
	cfg        *requestconfig.RequestConfig
	res        *http.Response
}

// runsCursorJSON contains the JSON metadata for the struct [RunsCursor[T]]
type runsCursorJSON struct {
	Runs        apijson.Field
	NextCursor  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RunsCursor[T]) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runsCursorJSON) RawJSON() string {
	return r.raw
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *RunsCursor[T]) GetNextPage() (res *RunsCursor[T], err error) {
	if len(r.Runs) == 0 {
		return nil, nil
	}
	next := r.NextCursor
	if len(next) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	err = cfg.Apply(option.WithQuery("cursor", next))
	if err != nil {
		return nil, err
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *RunsCursor[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &RunsCursor[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type RunsCursorAutoPager[T any] struct {
	page *RunsCursor[T]
	cur  T
	idx  int
	run  int
	err  error
}

func NewRunsCursorAutoPager[T any](page *RunsCursor[T], err error) *RunsCursorAutoPager[T] {
	return &RunsCursorAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *RunsCursorAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Runs) == 0 {
		return false
	}
	if r.idx >= len(r.page.Runs) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Runs) == 0 {
			return false
		}
	}
	r.cur = r.page.Runs[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *RunsCursorAutoPager[T]) Current() T {
	return r.cur
}

func (r *RunsCursorAutoPager[T]) Err() error {
	return r.err
}

func (r *RunsCursorAutoPager[T]) Index() int {
	return r.run
}

type AlertChannelsCursor[T any] struct {
	AlertChannels []T                     `json:"alert_channels"`
	NextCursor    string                  `json:"next_cursor" api:"nullable"`
	JSON          alertChannelsCursorJSON `json:"-"`
	cfg           *requestconfig.RequestConfig
	res           *http.Response
}

// alertChannelsCursorJSON contains the JSON metadata for the struct
// [AlertChannelsCursor[T]]
type alertChannelsCursorJSON struct {
	AlertChannels apijson.Field
	NextCursor    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AlertChannelsCursor[T]) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r alertChannelsCursorJSON) RawJSON() string {
	return r.raw
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *AlertChannelsCursor[T]) GetNextPage() (res *AlertChannelsCursor[T], err error) {
	if len(r.AlertChannels) == 0 {
		return nil, nil
	}
	next := r.NextCursor
	if len(next) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	err = cfg.Apply(option.WithQuery("cursor", next))
	if err != nil {
		return nil, err
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *AlertChannelsCursor[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &AlertChannelsCursor[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type AlertChannelsCursorAutoPager[T any] struct {
	page *AlertChannelsCursor[T]
	cur  T
	idx  int
	run  int
	err  error
}

func NewAlertChannelsCursorAutoPager[T any](page *AlertChannelsCursor[T], err error) *AlertChannelsCursorAutoPager[T] {
	return &AlertChannelsCursorAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *AlertChannelsCursorAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.AlertChannels) == 0 {
		return false
	}
	if r.idx >= len(r.page.AlertChannels) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.AlertChannels) == 0 {
			return false
		}
	}
	r.cur = r.page.AlertChannels[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *AlertChannelsCursorAutoPager[T]) Current() T {
	return r.cur
}

func (r *AlertChannelsCursorAutoPager[T]) Err() error {
	return r.err
}

func (r *AlertChannelsCursorAutoPager[T]) Index() int {
	return r.run
}

type AlertSubscriptionsCursor[T any] struct {
	AlertSubscriptions []T                          `json:"alert_subscriptions"`
	NextCursor         string                       `json:"next_cursor" api:"nullable"`
	JSON               alertSubscriptionsCursorJSON `json:"-"`
	cfg                *requestconfig.RequestConfig
	res                *http.Response
}

// alertSubscriptionsCursorJSON contains the JSON metadata for the struct
// [AlertSubscriptionsCursor[T]]
type alertSubscriptionsCursorJSON struct {
	AlertSubscriptions apijson.Field
	NextCursor         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AlertSubscriptionsCursor[T]) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r alertSubscriptionsCursorJSON) RawJSON() string {
	return r.raw
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *AlertSubscriptionsCursor[T]) GetNextPage() (res *AlertSubscriptionsCursor[T], err error) {
	if len(r.AlertSubscriptions) == 0 {
		return nil, nil
	}
	next := r.NextCursor
	if len(next) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	err = cfg.Apply(option.WithQuery("cursor", next))
	if err != nil {
		return nil, err
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *AlertSubscriptionsCursor[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &AlertSubscriptionsCursor[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type AlertSubscriptionsCursorAutoPager[T any] struct {
	page *AlertSubscriptionsCursor[T]
	cur  T
	idx  int
	run  int
	err  error
}

func NewAlertSubscriptionsCursorAutoPager[T any](page *AlertSubscriptionsCursor[T], err error) *AlertSubscriptionsCursorAutoPager[T] {
	return &AlertSubscriptionsCursorAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *AlertSubscriptionsCursorAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.AlertSubscriptions) == 0 {
		return false
	}
	if r.idx >= len(r.page.AlertSubscriptions) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.AlertSubscriptions) == 0 {
			return false
		}
	}
	r.cur = r.page.AlertSubscriptions[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *AlertSubscriptionsCursorAutoPager[T]) Current() T {
	return r.cur
}

func (r *AlertSubscriptionsCursorAutoPager[T]) Err() error {
	return r.err
}

func (r *AlertSubscriptionsCursorAutoPager[T]) Index() int {
	return r.run
}

type MaintenanceWindowsCursor[T any] struct {
	MaintenanceWindows []T                          `json:"maintenance_windows"`
	NextCursor         string                       `json:"next_cursor" api:"nullable"`
	JSON               maintenanceWindowsCursorJSON `json:"-"`
	cfg                *requestconfig.RequestConfig
	res                *http.Response
}

// maintenanceWindowsCursorJSON contains the JSON metadata for the struct
// [MaintenanceWindowsCursor[T]]
type maintenanceWindowsCursorJSON struct {
	MaintenanceWindows apijson.Field
	NextCursor         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *MaintenanceWindowsCursor[T]) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r maintenanceWindowsCursorJSON) RawJSON() string {
	return r.raw
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *MaintenanceWindowsCursor[T]) GetNextPage() (res *MaintenanceWindowsCursor[T], err error) {
	if len(r.MaintenanceWindows) == 0 {
		return nil, nil
	}
	next := r.NextCursor
	if len(next) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	err = cfg.Apply(option.WithQuery("cursor", next))
	if err != nil {
		return nil, err
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *MaintenanceWindowsCursor[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &MaintenanceWindowsCursor[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type MaintenanceWindowsCursorAutoPager[T any] struct {
	page *MaintenanceWindowsCursor[T]
	cur  T
	idx  int
	run  int
	err  error
}

func NewMaintenanceWindowsCursorAutoPager[T any](page *MaintenanceWindowsCursor[T], err error) *MaintenanceWindowsCursorAutoPager[T] {
	return &MaintenanceWindowsCursorAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *MaintenanceWindowsCursorAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.MaintenanceWindows) == 0 {
		return false
	}
	if r.idx >= len(r.page.MaintenanceWindows) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.MaintenanceWindows) == 0 {
			return false
		}
	}
	r.cur = r.page.MaintenanceWindows[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *MaintenanceWindowsCursorAutoPager[T]) Current() T {
	return r.cur
}

func (r *MaintenanceWindowsCursorAutoPager[T]) Err() error {
	return r.err
}

func (r *MaintenanceWindowsCursorAutoPager[T]) Index() int {
	return r.run
}
