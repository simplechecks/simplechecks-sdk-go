// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package simplechecksgo

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

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
