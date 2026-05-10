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
