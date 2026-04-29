// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package simplechecks

import (
	"context"
	"net/http"
	"slices"

	"github.com/simplechecks/sdk-go/internal/apijson"
	"github.com/simplechecks/sdk-go/internal/requestconfig"
	"github.com/simplechecks/sdk-go/option"
	"github.com/simplechecks/sdk-go/packages/respjson"
)

// Liveness + readiness.
//
// HealthzService contains methods and other services that help with interacting
// with the simplechecks API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewHealthzService] method instead.
type HealthzService struct {
	options []option.RequestOption
}

// NewHealthzService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewHealthzService(opts ...option.RequestOption) (r HealthzService) {
	r = HealthzService{}
	r.options = opts
	return
}

// Returns 200 when the process is up. webapp/api is stateless, so "the process is
// up" is the entire health story; Kubernetes uses this for both liveness and
// readiness probes. Public — no auth.
func (r *HealthzService) Check(ctx context.Context, opts ...option.RequestOption) (res *HealthzCheckResponse, err error) {
	var preClientOpts = []option.RequestOption{requestconfig.WithSecurity(requestconfig.Security{})}
	opts = slices.Concat(preClientOpts, r.options, opts)
	path := "healthz"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type HealthzCheckResponse struct {
	Status string `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r HealthzCheckResponse) RawJSON() string { return r.JSON.raw }
func (r *HealthzCheckResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
